// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Uicustomization uicustomization
// swagger:model uicustomization
type Uicustomization struct {

	// data
	Data Uidata `json:"data,omitempty"`

	// observables
	Observables []string `json:"observables"`

	// restricted
	Restricted bool `json:"restricted,omitempty"`

	// template
	Template string `json:"template,omitempty"`

	// validators
	Validators Validators `json:"validators,omitempty"`
}

// Validate validates this uicustomization
func (m *Uicustomization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Uicustomization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Uicustomization) UnmarshalBinary(b []byte) error {
	var res Uicustomization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}