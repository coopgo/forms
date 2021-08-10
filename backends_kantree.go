package forms

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

// KantreeBackend is a backend using Kantree
type KantreeBackend struct {
	BaseUrl    string
	APIKey     string
	ProjectId  string
	TitleField string
}

// Type returns the backend type (Kantree)
func (b KantreeBackend) Type() string {
	return "Kantree"
}

func (b KantreeBackend) Configuration() map[string]interface{} {
	return map[string]interface{}{
		"type": b.Type(),
		"configuration": map[string]string{
			"api_key":     b.APIKey,
			"project_id":  b.ProjectId,
			"title_field": b.TitleField,
		},
	}
}

// NewKantreeBackend initializes a new empty KantreeBackend
func NewKantreeBackend(configuration map[string]interface{}) KantreeBackend {
	return KantreeBackend{
		BaseUrl:    "https://kantree.io/api/1.0/",
		APIKey:     configuration["api_key"].(string),
		ProjectId:  configuration["project_id"].(string),
		TitleField: configuration["title_field"].(string),
	}
}

// AddForm creates a new form in the backend
func (b KantreeBackend) AddForm(f Form) error {
	return errors.New("not implemented")
}

// DeleteForm removes the form from the backend using formId
func (b KantreeBackend) DeleteForm(formId string) error {

	return errors.New("not implemented")
}

// GetForms retrieves all forms from this backend
func (b KantreeBackend) GetForms() ([]Form, error) {
	return nil, errors.New("not implemented")
}

// GetForm retrieves the form from the backend given a formId
func (b KantreeBackend) GetForm(id string) (Form, error) {
	return NewUnstructuredForm(""), errors.New("not implemented")
}

// Submitresponse registers a response to the form given it's formid
func (b KantreeBackend) SubmitResponse(form Form, response Response) error {
	//TODO submit response
	req, err := http.NewRequest("GET", fmt.Sprintf("%sprojects/%s", b.BaseUrl, b.ProjectId), nil)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(b.APIKey+":")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err)
		return err
	}

	if reflect.ValueOf(data["top_level_card_id"]).Kind() != reflect.Float64 {
		return errors.New("top_level_card_id type is " + reflect.ValueOf(data["top_level_card_id"]).Kind().String() + " instead of float64")
	}

	toplevelcardid := data["top_level_card_id"].(float64)

	kantreecard := map[string]interface{}{
		"title":      response.(map[string]interface{})[b.TitleField],
		"attributes": response,
	}

	fmt.Println(kantreecard)

	kcdata, err := json.Marshal(kantreecard)
	if err != nil {
		return err
	}

	fmt.Println(string(kcdata))

	fmt.Println(fmt.Sprintf("%scards/%.0f/children", b.BaseUrl, toplevelcardid))
	req2, err := http.NewRequest("POST", fmt.Sprintf("%scards/%.0f/children", b.BaseUrl, toplevelcardid), bytes.NewBuffer(kcdata))

	if err != nil {
		fmt.Println("error 3")
		fmt.Println(err)
		return err
	}

	req2.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(b.APIKey+":")))

	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println("error 4")
		fmt.Println(err)
		return err
	}
	defer resp2.Body.Close()

	var newdata interface{}

	err = json.NewDecoder(resp2.Body).Decode(&newdata)
	if err != nil {
		fmt.Println("error 1")
		fmt.Println(err)
		return err
	}

	fmt.Println(newdata)
	fmt.Println(resp2.Status)
	fmt.Println(err)

	return nil
}

// GetFormResponses retrieves the responses for a given form (with formId)
func (b KantreeBackend) GetFormResponses(formId string) ([]Response, error) {
	return nil, errors.New("not implemented")

}

//AddFormBackend adds a new backend : unavailable for Kantree backends
func (b KantreeBackend) AddFormBackend(string, Backend) error {
	return errors.New("not implemented")
}

//GetFormBackends retrieves form backends : unavailable for Kantree backends
func (b KantreeBackend) GetFormBackends(string) ([]Backend, error) {
	return nil, errors.New("not implemented")
}
