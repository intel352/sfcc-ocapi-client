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

// CustomerGroups Document representing an unfiltered list of customer groups.
// swagger:model customer_groups
type CustomerGroups struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// data
	Data []*CustomerGroup `json:"data"`

	// The list of expands set for the search request. Expands are optional.
	Expand []string `json:"expand"`

	// The URL of the next result page.
	Next string `json:"next,omitempty"`

	// The URL of the previous result page.
	Previous string `json:"previous,omitempty"`

	// The fields that you want to select.
	Select string `json:"select,omitempty"`

	// The zero-based index of the first search hit to include in the result.
	// Minimum: 0
	Start *int32 `json:"start,omitempty"`

	// The total number of documents.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this customer groups
func (m *CustomerGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerGroups) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerGroups) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.MinimumInt("start", "body", int64(*m.Start), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerGroups) UnmarshalBinary(b []byte) error {
	var res CustomerGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
