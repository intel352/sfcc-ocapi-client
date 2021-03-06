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

// RoleLocalePermissions Document listing the locale permissions assigned to a certain role.
// swagger:model role_locale_permissions
type RoleLocalePermissions struct {

	// The list of unscoped locale permissions.
	Unscoped []*RoleLocalePermission `json:"unscoped"`
}

// Validate validates this role locale permissions
func (m *RoleLocalePermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnscoped(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleLocalePermissions) validateUnscoped(formats strfmt.Registry) error {

	if swag.IsZero(m.Unscoped) { // not required
		return nil
	}

	for i := 0; i < len(m.Unscoped); i++ {
		if swag.IsZero(m.Unscoped[i]) { // not required
			continue
		}

		if m.Unscoped[i] != nil {
			if err := m.Unscoped[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unscoped" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleLocalePermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleLocalePermissions) UnmarshalBinary(b []byte) error {
	var res RoleLocalePermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
