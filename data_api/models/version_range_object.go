// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VersionRangeObject Object to represent the Version Range. Used in resources_object.
// swagger:model version_range_object
type VersionRangeObject struct {

	// Starting version
	From string `json:"from,omitempty"`

	// Ending version
	Until string `json:"until,omitempty"`
}

// Validate validates this version range object
func (m *VersionRangeObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VersionRangeObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VersionRangeObject) UnmarshalBinary(b []byte) error {
	var res VersionRangeObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
