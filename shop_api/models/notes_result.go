// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NotesResult A result of a note request.
//
//  Contains notes for an object - for example, for a basket.
// swagger:model notes_result
type NotesResult struct {

	// The notes for an object.
	Notes []*Note `json:"notes"`
}

// Validate validates this notes result
func (m *NotesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotesResult) validateNotes(formats strfmt.Registry) error {

	if swag.IsZero(m.Notes) { // not required
		return nil
	}

	for i := 0; i < len(m.Notes); i++ {
		if swag.IsZero(m.Notes[i]) { // not required
			continue
		}

		if m.Notes[i] != nil {
			if err := m.Notes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotesResult) UnmarshalBinary(b []byte) error {
	var res NotesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}