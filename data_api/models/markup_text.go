// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MarkupText markup text
// swagger:model markup_text
type MarkupText struct {

	// The rendered HTML
	Markup string `json:"markup,omitempty"`

	// The raw markup text
	Source string `json:"source,omitempty"`
}

// Validate validates this markup text
func (m *MarkupText) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarkupText) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarkupText) UnmarshalBinary(b []byte) error {
	var res MarkupText
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}