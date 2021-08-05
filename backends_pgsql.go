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
	Forms     map[string]Form
	Responses map[string][]Response

	Conn *pgxpool.Pool
}

func (b PgsqlBackend) Type() string {
	return "Pgsql"
}

// NewPgsqlBackend initializes a new empty PgsqlBackend
func NewPgsqlBackend(databaseurl string) PgsqlBackend {
	conn, err := pgxpool.Connect(context.Background(), databaseurl)

	if err != nil {
		panic(err)
	}

	if _, err := conn.Exec(context.Background(), "create table if not exists __config(id varchar(50) primary key, type varchar(20), schema json)"); err != nil {
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
	if _, err := tx.Exec(context.Background(), "drop table form_"+formId); err != nil {
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
			switch {
			case reflect.TypeOf(responsemap[k]).Kind() == reflect.String:
				if _, err := tx.Exec(context.Background(), fmt.Sprintf("alter table %s add column if not exists %s varchar(250)", "form_"+form.ID(), k)); err != nil {
					fmt.Println("dfsd")
					return err
				}
				break
			case reflect.TypeOf(responsemap[k]).Kind() == reflect.Bool:
				if _, err := tx.Exec(context.Background(), fmt.Sprintf("alter table %s add column if not exists %s boolean", "form_"+form.ID(), k)); err != nil {
					fmt.Println("aaa")
					return err
				}
				break
			case reflect.TypeOf(responsemap[k]).Kind() == reflect.Map:
				if _, err := tx.Exec(context.Background(), fmt.Sprintf("alter table %s add column if not exists %s json", "form_"+form.ID(), k)); err != nil {
					fmt.Println("vvv")
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
	rows, err := b.Conn.Query(context.Background(), fmt.Sprintf("select * from form_%s", formId))
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
			create += fmt.Sprintf("%s varchar(250)", k)
			break
		case v.Type == "boolean":
			create += fmt.Sprintf("%s boolean", k)
			break
		case v.Type == "float64" || v.Type == "float32":
			create += fmt.Sprintf("%s numeric", k)
			break
		case v.Type == "int8" || v.Type == "uint8" || v.Type == "int16" || v.Type == "uint16" || v.Type == "int32" || v.Type == "uint32":
			create += fmt.Sprintf("%s integer", k)
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
