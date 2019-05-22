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

// WebdavPermissions Document representing the available WebDAV permissions.
// swagger:model webdav_permissions
type WebdavPermissions struct {

	// The available WebDAV permission scopes (e.g. unscoped).
	Scopes []string `json:"scopes"`

	// The collection of available unscoped WebDAV permissions.
	Unscoped []*WebdavPermission `json:"unscoped"`
}

// Validate validates this webdav permissions
func (m *WebdavPermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnscoped(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebdavPermissions) validateUnscoped(formats strfmt.Registry) error {

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
func (m *WebdavPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebdavPermissions) UnmarshalBinary(b []byte) error {
	var res WebdavPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
