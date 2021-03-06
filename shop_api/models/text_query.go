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

// TextQuery A text query is used to match some text (i.e. a search phrase possibly consisting of multiple terms) against one or
//  multiple fields. In case multiple fields are provided, the phrase conceptually forms a logical OR over the fields. In
//  this case, the terms of the phrase basically have to match within the text, that would result in concatenating all
//  given fields.
//
// swagger:model text_query
type TextQuery struct {

	// The document fields the search phrase has to match against.
	// Required: true
	// Min Items: 1
	Fields []string `json:"fields"`

	// A search phrase, which may consist of multiple terms.
	// Required: true
	SearchPhrase *string `json:"search_phrase"`
}

// Validate validates this text query
func (m *TextQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchPhrase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextQuery) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	iFieldsSize := int64(len(m.Fields))

	if err := validate.MinItems("fields", "body", iFieldsSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *TextQuery) validateSearchPhrase(formats strfmt.Registry) error {

	if err := validate.Required("search_phrase", "body", m.SearchPhrase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextQuery) UnmarshalBinary(b []byte) error {
	var res TextQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
