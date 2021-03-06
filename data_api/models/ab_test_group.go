// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AbTestGroup Document representing an A/B Test Group
// swagger:model ab_test_group
type AbTestGroup struct {

	// Test Group percentage allocation
	Allocation int32 `json:"allocation,omitempty"`

	// Flag to determine if this Test Group is a customer experience
	CustomExperience bool `json:"custom_experience,omitempty"`

	// Test Group description
	Description string `json:"description,omitempty"`

	// Test group id
	ID string `json:"id,omitempty"`
}

// Validate validates this ab test group
func (m *AbTestGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AbTestGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbTestGroup) UnmarshalBinary(b []byte) error {
	var res AbTestGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
