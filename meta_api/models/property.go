// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Property Document property description.
// swagger:model property
type Property struct {

	// Either HTTP- or document-internal-link to a document as data type.
	NrDollarRef string `json:"$ref,omitempty"`

	// additional properties
	AdditionalProperties *Property `json:"additionalProperties,omitempty"`

	// Single element map containing definition of a document. Obsolete if
	//  '$ref' is set.
	Definition map[string]Document `json:"definition,omitempty"`

	// Property description.
	Description string `json:"description,omitempty"`

	// Array of all valid property values.
	Enum []interface{} `json:"enum"`

	// Output format information.
	// Enum: [int32 int64 float double byte date date-time time text email html password money site-specific localized]
	Format string `json:"format,omitempty"`

	// Property type description for type 'array'.
	Items *Property `json:"items,omitempty"`

	// max items
	MaxItems int32 `json:"maxItems,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// Maximum property integer- or number-value.
	Maximum float64 `json:"maximum,omitempty"`

	// min items
	MinItems int32 `json:"minItems,omitempty"`

	// min length
	MinLength int32 `json:"minLength,omitempty"`

	// Minimum property integer- or number-value.
	Minimum float64 `json:"minimum,omitempty"`

	// Regular expression to limit string value.
	Pattern string `json:"pattern,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// Data type information.
	// Enum: [integer number string boolean array object]
	Type string `json:"type,omitempty"`

	// x enum labels
	XEnumLabels []interface{} `json:"x-enum-labels"`

	// x label
	XLabel map[string]string `json:"x-label,omitempty"`

	// x sub type definitions
	XSubTypeDefinitions map[string]Document `json:"x-sub_type_definitions,omitempty"`

	// x sub types
	XSubTypes map[string]string `json:"x-sub_types,omitempty"`
}

// Validate validates this property
func (m *Property) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXSubTypeDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Property) validateAdditionalProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalProperties) { // not required
		return nil
	}

	if m.AdditionalProperties != nil {
		if err := m.AdditionalProperties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additionalProperties")
			}
			return err
		}
	}

	return nil
}

func (m *Property) validateDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.Definition) { // not required
		return nil
	}

	for k := range m.Definition {

		if err := validate.Required("definition"+"."+k, "body", m.Definition[k]); err != nil {
			return err
		}
		if val, ok := m.Definition[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var propertyTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["int32","int64","float","double","byte","date","date-time","time","text","email","html","password","money","site-specific","localized"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyTypeFormatPropEnum = append(propertyTypeFormatPropEnum, v)
	}
}

const (

	// PropertyFormatInt32 captures enum value "int32"
	PropertyFormatInt32 string = "int32"

	// PropertyFormatInt64 captures enum value "int64"
	PropertyFormatInt64 string = "int64"

	// PropertyFormatFloat captures enum value "float"
	PropertyFormatFloat string = "float"

	// PropertyFormatDouble captures enum value "double"
	PropertyFormatDouble string = "double"

	// PropertyFormatByte captures enum value "byte"
	PropertyFormatByte string = "byte"

	// PropertyFormatDate captures enum value "date"
	PropertyFormatDate string = "date"

	// PropertyFormatDateTime captures enum value "date-time"
	PropertyFormatDateTime string = "date-time"

	// PropertyFormatTime captures enum value "time"
	PropertyFormatTime string = "time"

	// PropertyFormatText captures enum value "text"
	PropertyFormatText string = "text"

	// PropertyFormatEmail captures enum value "email"
	PropertyFormatEmail string = "email"

	// PropertyFormatHTML captures enum value "html"
	PropertyFormatHTML string = "html"

	// PropertyFormatPassword captures enum value "password"
	PropertyFormatPassword string = "password"

	// PropertyFormatMoney captures enum value "money"
	PropertyFormatMoney string = "money"

	// PropertyFormatSiteSpecific captures enum value "site-specific"
	PropertyFormatSiteSpecific string = "site-specific"

	// PropertyFormatLocalized captures enum value "localized"
	PropertyFormatLocalized string = "localized"
)

// prop value enum
func (m *Property) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, propertyTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Property) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *Property) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	if m.Items != nil {
		if err := m.Items.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

var propertyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["integer","number","string","boolean","array","object"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		propertyTypeTypePropEnum = append(propertyTypeTypePropEnum, v)
	}
}

const (

	// PropertyTypeInteger captures enum value "integer"
	PropertyTypeInteger string = "integer"

	// PropertyTypeNumber captures enum value "number"
	PropertyTypeNumber string = "number"

	// PropertyTypeString captures enum value "string"
	PropertyTypeString string = "string"

	// PropertyTypeBoolean captures enum value "boolean"
	PropertyTypeBoolean string = "boolean"

	// PropertyTypeArray captures enum value "array"
	PropertyTypeArray string = "array"

	// PropertyTypeObject captures enum value "object"
	PropertyTypeObject string = "object"
)

// prop value enum
func (m *Property) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, propertyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Property) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Property) validateXSubTypeDefinitions(formats strfmt.Registry) error {

	if swag.IsZero(m.XSubTypeDefinitions) { // not required
		return nil
	}

	for k := range m.XSubTypeDefinitions {

		if err := validate.Required("x-sub_type_definitions"+"."+k, "body", m.XSubTypeDefinitions[k]); err != nil {
			return err
		}
		if val, ok := m.XSubTypeDefinitions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Property) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Property) UnmarshalBinary(b []byte) error {
	var res Property
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}