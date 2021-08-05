package forms

import (
	"bytes"
	"encoding/json"
)

// DataType is the type to describe the forms data types (Text, Email, Boolean, Number, Location, ...)
type DataType int

const (
	TextType DataType = iota
	EmailType
	BooleanType
	NumberType
	LocationType
	NestedType
)

func (t DataType) String() string {
	return toString[t]
}

var toString = map[DataType]string{
	TextType:     "Text",
	EmailType:    "Email",
	BooleanType:  "Boolean",
	NumberType:   "Number",
	LocationType: "Location",
	NestedType:   "Nested",
}

var toID = map[string]DataType{
	"Text":     TextType,
	"Email":    EmailType,
	"Boolean":  BooleanType,
	"Number":   NumberType,
	"Location": LocationType,
	"Nested":   NestedType,
	"Parent":   NestedType, // to be removed, it's just because I still have tests with this
}

// MarshalJSON marshals the enum as a quoted json string
func (s DataType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *DataType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Text' in this case.
	*s = toID[j]
	return nil
}
