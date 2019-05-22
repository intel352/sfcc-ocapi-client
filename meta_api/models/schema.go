// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Schema Object, containing either a link to a document or the definition of it.
// swagger:model schema
type Schema struct {

	// HTTP- or document-internal link to the schema document.
	NrDollarRef string `json:"$ref,omitempty"`

	// Single element map for document definition.
	Definition map[string]Document `json:"definition,omitempty"`
}

// Validate validates this schema
func (m *Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Schema) validateDefinition(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema) UnmarshalBinary(b []byte) error {
	var res Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
