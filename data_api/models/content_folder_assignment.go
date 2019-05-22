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

// ContentFolderAssignment Document representing a content folder assignment.
// swagger:model content_folder_assignment
type ContentFolderAssignment struct {

	// The content id.
	// Max Length: 256
	ContentID string `json:"content_id,omitempty"`

	// The content link.
	ContentLink string `json:"content_link,omitempty"`

	// A flag indicating whether the assignment is the default one.
	Default bool `json:"default,omitempty"`

	// The folder id.
	// Max Length: 256
	FolderID string `json:"folder_id,omitempty"`

	// The folder link.
	FolderLink string `json:"folder_link,omitempty"`

	// The position of the content asset in the folder.
	// Minimum: 0
	Position *float64 `json:"position,omitempty"`
}

// Validate validates this content folder assignment
func (m *ContentFolderAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentFolderAssignment) validateContentID(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentID) { // not required
		return nil
	}

	if err := validate.MaxLength("content_id", "body", string(m.ContentID), 256); err != nil {
		return err
	}

	return nil
}

func (m *ContentFolderAssignment) validateFolderID(formats strfmt.Registry) error {

	if swag.IsZero(m.FolderID) { // not required
		return nil
	}

	if err := validate.MaxLength("folder_id", "body", string(m.FolderID), 256); err != nil {
		return err
	}

	return nil
}

func (m *ContentFolderAssignment) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.Minimum("position", "body", float64(*m.Position), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentFolderAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentFolderAssignment) UnmarshalBinary(b []byte) error {
	var res ContentFolderAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
