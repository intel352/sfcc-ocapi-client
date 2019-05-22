// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecommendationType Document representing a recommendation type.
// swagger:model recommendation_type
type RecommendationType struct {

	// The localized display value of the recommendation type.
	DisplayValue string `json:"display_value,omitempty"`

	// The value of the recommendation type.
	Value int32 `json:"value,omitempty"`
}

// Validate validates this recommendation type
func (m *RecommendationType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecommendationType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecommendationType) UnmarshalBinary(b []byte) error {
	var res RecommendationType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}