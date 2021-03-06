// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SuggestedContent suggested content
// swagger:model suggested_content
type SuggestedContent struct {

	// The id of the content.
	ID string `json:"id,omitempty"`

	// The URL addressing the content.
	Link string `json:"link,omitempty"`

	// The localized name of the content.
	Name string `json:"name,omitempty"`
}

// Validate validates this suggested content
func (m *SuggestedContent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuggestedContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestedContent) UnmarshalBinary(b []byte) error {
	var res SuggestedContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
