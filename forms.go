package forms

import (
	"fmt"

	jtd "github.com/jsontypedef/json-typedef-go"
)

// Form is the generic interface for forms.
type Form interface {
	ID() string
	Type() string
	GetSchema() *jtd.Schema
	GetAdditionalBackends() []Backend
	AddBackend(Backend) error
	Validate(Response) bool
}

// StructuredForm is a Form with a defined schema. Submitted data must match the given schema
type StructuredForm struct {
	id                 string
	Schema             *jtd.Schema
	AdditionalBackends []Backend
}

// NewStructuredForm creates a new structured form with a given ID and schema
func NewStructuredForm(id string, schema *jtd.Schema, backends ...Backend) *StructuredForm {
	return &StructuredForm{
		id:                 id,
		Schema:             schema,
		AdditionalBackends: backends,
	}
}

// ID returns the structured form ID
func (f *StructuredForm) ID() string {
	return f.id
}

// Type returns the type of the form (Structured)
func (f *StructuredForm) Type() string {
	return "Structured"
}

// GetSchema returns the data schema of the form
func (f *StructuredForm) GetSchema() *jtd.Schema {
	return f.Schema
}

// GetAdditionalBackends returns the additional backends for this form
func (f *StructuredForm) GetAdditionalBackends() []Backend {
	return f.AdditionalBackends
}

// AddBackend adds a new backend to the form
func (f *StructuredForm) AddBackend(b Backend) error {
	f.AdditionalBackends = append(f.AdditionalBackends, b)
	return nil
}

// Validate validates that the response follows the form's schema
func (f *StructuredForm) Validate(response Response) bool {
	if errs, _ := jtd.Validate(*f.Schema, response); errs != nil && len(errs) > 0 {
		fmt.Println(errs)
		return false
	}
	return true
}

// UnstructuredForm is a form without any given schema. It accepts unstructured data
type UnstructuredForm struct {
	id                 string
	AdditionalBackends []Backend
}

// NewUnstructuredForm creates a new unstructured form with a given id
func NewUnstructuredForm(id string, backends ...Backend) *UnstructuredForm {
	return &UnstructuredForm{
		id:                 id,
		AdditionalBackends: backends,
	}
}

// ID returns the ID of the unstructured form
func (f *UnstructuredForm) ID() string {
	return f.id
}

// Type returns the type of the form (Unstructured)
func (f *UnstructuredForm) Type() string {
	return "Unstructured"
}

// GetSchema returns nil because unstructured form do not have schemas
func (f *UnstructuredForm) GetSchema() *jtd.Schema {
	return nil
}

// GetAdditionalBackends returns the additional backends for this form
func (f *UnstructuredForm) GetAdditionalBackends() []Backend {
	return f.AdditionalBackends
}

// AddBackend adds a new backend to the form
func (f *UnstructuredForm) AddBackend(b Backend) error {
	f.AdditionalBackends = append(f.AdditionalBackends, b)
	return nil
}

// Validate validates that the response follows the form's schema : always true as unstructured form
func (f *UnstructuredForm) Validate(response Response) bool {
	return true
}
