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

// RoleModulePermission Document representing a module permission.
// swagger:model role_module_permission
type RoleModulePermission struct {

	// The permission application (e.g. "bm", "csc").
	// Required: true
	// Min Length: 1
	Application *string `json:"application"`

	// The related menu action name of the module permission.
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Flag to indicate a system menu action. This is false for custom menu actions.
	System bool `json:"system,omitempty"`

	// The permission type ("module").
	// Required: true
	// Min Length: 1
	Type *string `json:"type"`

	// The non domain specific value for the module permission, e.g. ACCESS or READONLY.
	Value string `json:"value,omitempty"`

	// The map of value per domain for the module permission, e.g. ACCESS or READONLY per domain name.
	Values map[string]string `json:"values,omitempty"`
}

// Validate validates this role module permission
func (m *RoleModulePermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleModulePermission) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	if err := validate.MinLength("application", "body", string(*m.Application), 1); err != nil {
		return err
	}

	return nil
}

func (m *RoleModulePermission) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *RoleModulePermission) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleModulePermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleModulePermission) UnmarshalBinary(b []byte) error {
	var res RoleModulePermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}