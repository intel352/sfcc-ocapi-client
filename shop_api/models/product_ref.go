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

// ProductRef Document representing a product reference.
// swagger:model product_ref
type ProductRef struct {

	// The ID of the product reference.
	// Required: true
	ID *string `json:"id"`

	// The link to the product reference.
	// Required: true
	// Max Length: 100
	Link *string `json:"link"`
}

// Validate validates this product ref
func (m *ProductRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProductRef) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	if err := validate.MaxLength("link", "body", string(*m.Link), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductRef) UnmarshalBinary(b []byte) error {
	var res ProductRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
