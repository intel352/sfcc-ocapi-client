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

// OptionValue Document representing an option value.
// swagger:model option_value
type OptionValue struct {

	// A flag indicating whether this option value is the default one.
	Default bool `json:"default,omitempty"`

	// The id of the option value.
	// Required: true
	// Max Length: 100
	// Min Length: 1
	ID *string `json:"id"`

	// The localized name of the option value.
	Name string `json:"name,omitempty"`

	// The effective price of this option value.
	Price float64 `json:"price,omitempty"`
}

// Validate validates this option value
func (m *OptionValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OptionValue) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OptionValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OptionValue) UnmarshalBinary(b []byte) error {
	var res OptionValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
