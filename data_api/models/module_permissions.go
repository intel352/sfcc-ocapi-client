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

// ModulePermissions Document representing the available module permissions in shape of menu items and menu actions.
// swagger:model module_permissions
type ModulePermissions struct {

	// The collection of available organization menu items.
	Organization []*MenuItem `json:"organization"`

	// The available menu item scopes (e.g. organization, site).
	Scopes []string `json:"scopes"`

	// The list of available site menu items.
	Site []*MenuItem `json:"site"`
}

// Validate validates this module permissions
func (m *ModulePermissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModulePermissions) validateOrganization(formats strfmt.Registry) error {

	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	for i := 0; i < len(m.Organization); i++ {
		if swag.IsZero(m.Organization[i]) { // not required
			continue
		}

		if m.Organization[i] != nil {
			if err := m.Organization[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organization" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModulePermissions) validateSite(formats strfmt.Registry) error {

	if swag.IsZero(m.Site) { // not required
		return nil
	}

	for i := 0; i < len(m.Site); i++ {
		if swag.IsZero(m.Site[i]) { // not required
			continue
		}

		if m.Site[i] != nil {
			if err := m.Site[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("site" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModulePermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModulePermissions) UnmarshalBinary(b []byte) error {
	var res ModulePermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
