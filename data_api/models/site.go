// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Site Document representing a site.
// swagger:model site
type Site struct {

	// The cartridge Path of the site
	Cartridges string `json:"cartridges,omitempty"`

	// The link to the customer list.
	CustomerListLink *CustomerListLink `json:"customer_list_link,omitempty"`

	// The description of this site.
	Description map[string]string `json:"description,omitempty"`

	// The display name entered by the user.
	DisplayName map[string]string `json:"display_name,omitempty"`

	// The id of this site.
	// Required: true
	// Min Length: 1
	ID *string `json:"id"`

	// The deletion status of this site, true if in deletion
	InDeletion bool `json:"in_deletion,omitempty"`

	// A link directly to the site
	Link string `json:"link,omitempty"`

	// storefront status
	// Enum: [online maintenance to_be_deleted protected]
	StorefrontStatus string `json:"storefront_status,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerListLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorefrontStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateCustomerListLink(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerListLink) { // not required
		return nil
	}

	if m.CustomerListLink != nil {
		if err := m.CustomerListLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer_list_link")
			}
			return err
		}
	}

	return nil
}

func (m *Site) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	return nil
}

var siteTypeStorefrontStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","maintenance","to_be_deleted","protected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTypeStorefrontStatusPropEnum = append(siteTypeStorefrontStatusPropEnum, v)
	}
}

const (

	// SiteStorefrontStatusOnline captures enum value "online"
	SiteStorefrontStatusOnline string = "online"

	// SiteStorefrontStatusMaintenance captures enum value "maintenance"
	SiteStorefrontStatusMaintenance string = "maintenance"

	// SiteStorefrontStatusToBeDeleted captures enum value "to_be_deleted"
	SiteStorefrontStatusToBeDeleted string = "to_be_deleted"

	// SiteStorefrontStatusProtected captures enum value "protected"
	SiteStorefrontStatusProtected string = "protected"
)

// prop value enum
func (m *Site) validateStorefrontStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteTypeStorefrontStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Site) validateStorefrontStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.StorefrontStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateStorefrontStatusEnum("storefront_status", "body", m.StorefrontStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
