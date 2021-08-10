package forms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	jtd "github.com/jsontypedef/json-typedef-go"
)

// PgsqlBackend is a backend using PostgreSQL
type PgsqlBackend struct {
	Conn *pgxpool.Pool
}

func (b PgsqlBackend) Type() string {
	return "Pgsql"
}

func (b PgsqlBackend) Configuration() map[string]interface{} {
	return map[string]interface{}{
		"type": b.Type(),
	}
}

// NewPgsqlBackend initializes a new empty PgsqlBackend
func NewPgsqlBackend(databaseurl string) PgsqlBackend {
	conn, err := pgxpool.Connect(context.Background(), databaseurl)

	if err != nil {
		panic(err)
	}

	if _, err := conn.Exec(context.Background(), "create table if not exists __config(id varchar(50) primary key, type varchar(20), schema json, additional_backends json)"); err != nil {
		panic(err)
	}

	return PgsqlBackend{
		Conn: conn,
	}
}

// AddForm creates a new form in the backend
func (b PgsqlBackend) AddForm(f Form) error {
	tx, err := b.Conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(
		context.Background(),
		"insert into __config (id, type, schema) values($1, $2, $3)", f.ID(), f.Type(), f.GetSchema(),
	); err != nil {
		return errors.New("a form already exists with this id")
	}

	sql := fmt.Sprintf("create table form_%s (", f.ID()) +
		schemaToSQL(f.GetSchema()) +
		")"

	if _, err := tx.Exec(
		context.Background(),
		sql,
	); err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// DeleteForm removes the form from the backend using formId
func (b PgsqlBackend) DeleteForm(formId string) error {
	tx, err := b.Conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())
	if _, err := tx.Exec(context.Background(), "delete from __config where id=$1", formId); err != nil {
		return err
	}
	if _, err := tx.Exec(context.Background(), fmt.Sprintf("drop table %s", pgx.Identifier{"form_" + formId}.Sanitize())); err != nil {
		return err
	}
	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}

// GetForms retrieves all forms from this backend
func (b PgsqlBackend) GetForms() ([]Form, error) {
	forms := []Form{}
	rows, err := b.Conn.Query(context.Background(), "select id, type, schema from __config")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var formid string
		var formtype string
		var formschema []byte
		err := rows.Scan(&formid, &formtype, &formschema)
		if err != nil {
			return nil, err
		}

		if formtype == "Structured" {
			var s *jtd.Schema
			json.Unmarshal(formschema, s)

			forms = append(forms, NewStructuredForm(formid, s))
		} else {
			forms = append(forms, NewUnstructuredForm(formid))
		}

	}
	return forms, nil
}

// GetForm retrieves the form from the backend given a formId
func (b PgsqlBackend) GetForm(id string) (Form, error) {
	var formid string
	var formtype string
	var formschema []byte
	if err := b.Conn.QueryRow(context.Background(), "select id, type, schema from __config where id=$1", id).Scan(&formid, &formtype, &formschema); err != nil {
		return nil, err
	}

	if formtype == "Structured" {
		var s jtd.Schema
		if err := json.Unmarshal(formschema, &s); err != nil {
			return nil, err
		}

		form := NewStructuredForm(formid, &s)

		return form, nil
	} else {
		form := NewUnstructuredForm(formid)

		return form, nil
	}
}

// Submitresponse registers a response to the form given it's formid
func (b PgsqlBackend) SubmitResponse(form Form, response Response) error {
	keys := []string{}
	values := []interface{}{}
	responsemap := response.(map[string]interface{})
	for k, v := range responsemap {
		keys = append(keys, k)
		values = append(values, v)
	}
	if form.Type() == "Structuted" {
		if _, err := b.Conn.CopyFrom(
			context.Background(),
			pgx.Identifier{"form_" + form.ID()},
			keys,
			pgx.CopyFromRows([][]interface{}{values}),
		); err != nil {
			return err
		}
	} else {
		tx, err := b.Conn.Begin(context.Background())
		defer tx.Rollback(context.Background())
		if err != nil {
			return err
		}
		for _, k := range keys {
			kind := reflect.TypeOf(responsemap[k]).Kind()
			switch {
			case kind == reflect.String:
				if _, err := tx.Exec(
					context.Background(),
					fmt.Sprintf("alter table %s add column if not exists %s varchar(250)",
						pgx.Identifier{"form_" + form.ID()}.Sanitize(),
						pgx.Identifier{k}.Sanitize(),
					),
				); err != nil {
					return err
				}
				break
			case kind == reflect.Bool:
				if _, err := tx.Exec(
					context.Background(),
					fmt.Sprintf(
						"alter table %s add column if not exists %s boolean",
						pgx.Identifier{"form_" + form.ID()}.Sanitize(),
						pgx.Identifier{k}.Sanitize(),
					),
				); err != nil {
					return err
				}
				break
			case kind == reflect.Map:
				if _, err := tx.Exec(
					context.Background(),
					fmt.Sprintf(
						"alter table %s add column if not exists %s json",
						pgx.Identifier{"form_" + form.ID()}.Sanitize(),
						pgx.Identifier{k}.Sanitize(),
					),
				); err != nil {
					return err
				}
				break
			case kind == reflect.Int || kind == reflect.Uint || kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Int32 || kind == reflect.Int64:
				if _, err := tx.Exec(
					context.Background(),
					fmt.Sprintf(
						"alter table %s add column if not exists %s integer",
						pgx.Identifier{"form_" + form.ID()}.Sanitize(),
						pgx.Identifier{k}.Sanitize(),
					),
				); err != nil {
					return err
				}
				break
			case kind == reflect.Float32 || kind == reflect.Float64:
				if _, err := tx.Exec(
					context.Background(),
					fmt.Sprintf(
						"alter table %s add column if not exists %s numeric",
						pgx.Identifier{"form_" + form.ID()}.Sanitize(),
						pgx.Identifier{k}.Sanitize(),
					),
				); err != nil {
					return err
				}
				break

				// TODO other types !!!
				// TODO recursive
			}
		}
		if _, err := tx.CopyFrom(
			context.Background(),
			pgx.Identifier{"form_" + form.ID()},
			keys,
			pgx.CopyFromRows([][]interface{}{values}),
		); err != nil {
			return err
		}

		err = tx.Commit(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

// GetFormResponses retrieves the responses for a given form (with formId)
func (b PgsqlBackend) GetFormResponses(formId string) ([]Response, error) {
	responses := []Response{}
	rows, err := b.Conn.Query(context.Background(), fmt.Sprintf("select * from %s", pgx.Identifier{"form_" + formId}.Sanitize()))
	if err != nil {
		fmt.Println("err 1")
		return nil, err
	}

	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		r := map[string]interface{}{}
		for i := 0; i < len(columns); i++ {
			r[columns[i]] = values[i]
		}
		responses = append(responses, reflect.ValueOf(r).Interface())
	}

	return responses, nil
}

// AddFormBackend adds an additional backend to a given form
func (b PgsqlBackend) AddFormBackend(id string, backend Backend) error {
	// var formid string
	// var formtype string
	// var formschema []byte
	var formbackends []byte
	if err := b.Conn.QueryRow(context.Background(), "select additional_backends from __config where id=$1", id).Scan(&formbackends); err != nil {
		fmt.Println("1 : ", err)
		return err
	}

	backends := []map[string]interface{}{}

	if formbackends != nil {
		if err := json.Unmarshal(formbackends, &backends); err != nil {
			fmt.Println("2 : ", err)
			return err
		}
	}

	backends = append(backends, backend.Configuration())

	if _, err := b.Conn.Exec(context.Background(), "update __config set additional_backends = $1 where id=$2", backends, id); err != nil {
		fmt.Println("3 : ", err)
		return err
	}
	return nil
}

func (b PgsqlBackend) GetFormBackends(id string) ([]Backend, error) {
	var formbackends []byte
	if err := b.Conn.QueryRow(context.Background(), "select additional_backends from __config where id=$1", id).Scan(&formbackends); err != nil {
		return nil, err
	}

	backends := []Backend{}

	var backendconfigs []map[string]interface{}

	if err := json.Unmarshal(formbackends, &backendconfigs); err != nil {
		return nil, err
	}

	for _, bc := range backendconfigs {
		if reflect.ValueOf(bc["type"]).Kind() == reflect.String && bc["type"].(string) == "Kantree" {
			backends = append(backends, NewKantreeBackend(bc["configuration"].(map[string]interface{})))
		}
	}

	return backends, nil
}

// schemaToSQL is a small function to generate table columns name and types from a JTD schema (used during table creation)
// Returns "" (empty string) in case of a nil schema
func schemaToSQL(s *jtd.Schema) string {
	if s == nil {
		return ""
	}
	create := ""
	first := true
	for k, v := range s.Properties {
		if !first {
			create += ", "
		}
		first = false
		switch {
		case v.Type == "string":
			create += fmt.Sprintf("%s varchar(250)", pgx.Identifier{k}.Sanitize())
			break
		case v.Type == "boolean":
			create += fmt.Sprintf("%s boolean", pgx.Identifier{k}.Sanitize())
			break
		case v.Type == "float64" || v.Type == "float32":
			create += fmt.Sprintf("%s numeric", pgx.Identifier{k}.Sanitize())
			break
		case v.Type == "int8" || v.Type == "uint8" || v.Type == "int16" || v.Type == "uint16" || v.Type == "int32" || v.Type == "uint32":
			create += fmt.Sprintf("%s integer", pgx.Identifier{k}.Sanitize())
			break
		}
	}
	for k, v := range s.OptionalProperties {
		if !first {
			create += ", "
		}
		first = false
		switch {
		case v.Type == "string":
			create += fmt.Sprintf("%s varchar(250) ", k)
			break
		case v.Type == "boolean":
			create += fmt.Sprintf("%s boolean ", k)
			break
		case v.Type == "float64" || v.Type == "float32":
			create += fmt.Sprintf("%s numeric ", k)
			break
		case v.Type == "int8" || v.Type == "uint8" || v.Type == "int16" || v.Type == "uint16" || v.Type == "int32" || v.Type == "uint32":
			create += fmt.Sprintf("%s integer ", k)
			break
		}
	}
	return create
}
