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

// PasswordChangeRequest Document representing a password change request.
// swagger:model password_change_request
type PasswordChangeRequest struct {

	// The customer's current password.
	// Required: true
	// Max Length: 4096
	CurrentPassword *string `json:"current_password"`

	// The customer's new password.
	// Required: true
	// Max Length: 4096
	Password *string `json:"password"`
}

// Validate validates this password change request
func (m *PasswordChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PasswordChangeRequest) validateCurrentPassword(formats strfmt.Registry) error {

	if err := validate.Required("current_password", "body", m.CurrentPassword); err != nil {
		return err
	}

	if err := validate.MaxLength("current_password", "body", string(*m.CurrentPassword), 4096); err != nil {
		return err
	}

	return nil
}

func (m *PasswordChangeRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 4096); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PasswordChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordChangeRequest) UnmarshalBinary(b []byte) error {
	var res PasswordChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}