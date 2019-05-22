// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SourceCodeGroup Document representing a source code group
// swagger:model source_code_group
type SourceCodeGroup struct {

	// The active flag, a computed value based on start and end time
	Active bool `json:"active,omitempty"`

	// The active redirect information
	ActiveRedirect *SourceCodeRedirectInfo `json:"active_redirect,omitempty"`

	// The cookie duration in days
	// Maximum: 999
	// Minimum: 0
	CookieDuration *int32 `json:"cookie_duration,omitempty"`

	// Returns the value of attribute 'creationDate'.
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// The description
	Description string `json:"description,omitempty"`

	// The enabled flag for storefront to consider the source code group, default to false.
	Enabled bool `json:"enabled,omitempty"`

	// The end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// The id of source code group
	// Required: true
	// Max Length: 28
	// Min Length: 1
	ID *string `json:"id"`

	// The inactive redirect information
	InactiveRedirect *SourceCodeRedirectInfo `json:"inactive_redirect,omitempty"`

	// Returns the value of attribute 'lastModified'.
	// Read Only: true
	// Format: date-time
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// URL that is used to get this instance, read only
	Link string `json:"link,omitempty"`

	// Source Code specifications
	Specifications []*SourceCodeSpecification `json:"specifications"`

	// The start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this source code group
func (m *SourceCodeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveRedirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookieDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInactiveRedirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpecifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceCodeGroup) validateActiveRedirect(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveRedirect) { // not required
		return nil
	}

	if m.ActiveRedirect != nil {
		if err := m.ActiveRedirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active_redirect")
			}
			return err
		}
	}

	return nil
}

func (m *SourceCodeGroup) validateCookieDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.CookieDuration) { // not required
		return nil
	}

	if err := validate.MinimumInt("cookie_duration", "body", int64(*m.CookieDuration), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cookie_duration", "body", int64(*m.CookieDuration), 999, false); err != nil {
		return err
	}

	return nil
}

func (m *SourceCodeGroup) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_date", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceCodeGroup) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceCodeGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 28); err != nil {
		return err
	}

	return nil
}

func (m *SourceCodeGroup) validateInactiveRedirect(formats strfmt.Registry) error {

	if swag.IsZero(m.InactiveRedirect) { // not required
		return nil
	}

	if m.InactiveRedirect != nil {
		if err := m.InactiveRedirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inactive_redirect")
			}
			return err
		}
	}

	return nil
}

func (m *SourceCodeGroup) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceCodeGroup) validateSpecifications(formats strfmt.Registry) error {

	if swag.IsZero(m.Specifications) { // not required
		return nil
	}

	for i := 0; i < len(m.Specifications); i++ {
		if swag.IsZero(m.Specifications[i]) { // not required
			continue
		}

		if m.Specifications[i] != nil {
			if err := m.Specifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("specifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SourceCodeGroup) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceCodeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceCodeGroup) UnmarshalBinary(b []byte) error {
	var res SourceCodeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}