package forms

import (
	"errors"
	"fmt"
)

// Backend is the interface for backends (storage, data handling)
type Backend interface {
	Type() string
	AddForm(Form) error
	DeleteForm(string) error
	GetForms() ([]Form, error)
	GetForm(string) (Form, error)
	SubmitResponse(Form, Response) error
	GetFormResponses(string) ([]Response, error)
}

// NoBackend is a fake backend implementation
type NoBackend struct {
}

// TempBackend is an in-memory temporary backend primarily used for testing and development purposes
type TempBackend struct {
	Forms     map[string]Form
	Responses map[string][]Response
}

func (b TempBackend) Type() string {
	return "Temp"
}

// NewTempBackend initializes a new empty TempBackend
func NewTempBackend() TempBackend {
	return TempBackend{
		Forms:     map[string]Form{},
		Responses: map[string][]Response{},
	}
}

// AddForm creates a new form in the backend
func (b TempBackend) AddForm(f Form) error {
	if b.Forms[f.ID()] != nil {
		return errors.New("a form already exists with this id")
	}
	b.Forms[f.ID()] = f
	fmt.Println(b)
	return nil
}

// DeleteForm removes the form from the backend using formIs
func (b TempBackend) DeleteForm(formId string) error {
	if b.Forms[formId] == nil {
		return errors.New("form does not exist")
	}
	delete(b.Forms, formId)
	return nil
}

// GetForms retrieves all forms from this backend
func (b TempBackend) GetForms() ([]Form, error) {
	result := []Form{}
	for _, f := range b.Forms {
		result = append(result, f)
	}
	return result, nil
}

// GetForm retrieves the form from the backend given a formId
func (b TempBackend) GetForm(formId string) (Form, error) {
	for _, f := range b.Forms {
		if f.ID() == formId {
			return f, nil
		}
	}
	return nil, errors.New("form not found")
}

func (b TempBackend) SubmitResponse(form Form, response Response) error {
	if b.Forms[form.ID()] == nil {
		return errors.New("form not found")
	}

	if b.Responses[form.ID()] == nil {
		b.Responses[form.ID()] = []Response{response}
		return nil
	}

	b.Responses[form.ID()] = append(b.Responses[form.ID()], response)
	return nil
}

func (b TempBackend) GetFormResponses(formId string) ([]Response, error) {
	if b.Forms[formId] == nil {
		return []Response{}, errors.New("form not found")
	}

	if b.Responses[formId] == nil {
		return []Response{}, nil
	}

	return b.Responses[formId], nil
}
