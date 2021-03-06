// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TimeOfDay Document representing a time schedule within a single day.
// swagger:model time_of_day
type TimeOfDay struct {

	// The time to start from. Time format: HH:mm:ss. Seconds
	//  are ignored and set to 0.
	TimeFrom string `json:"time_from,omitempty"`

	// The time to end on. Time format: HH:mm:ss. Seconds
	//  are ignored and set to 0.
	TimeTo string `json:"time_to,omitempty"`
}

// Validate validates this time of day
func (m *TimeOfDay) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeOfDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOfDay) UnmarshalBinary(b []byte) error {
	var res TimeOfDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
