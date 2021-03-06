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

// CodeVersionResult Result document containing an array of code versions.
// swagger:model code_version_result
type CodeVersionResult struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// The array of code versions
	Data []*CodeVersion `json:"data"`

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

// Validate validates this code version result
func (m *CodeVersionResult) Validate(formats strfmt.Registry) error {
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

func (m *CodeVersionResult) validateData(formats strfmt.Registry) error {

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

func (m *CodeVersionResult) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.MinimumInt("start", "body", int64(*m.Start), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CodeVersionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CodeVersionResult) UnmarshalBinary(b []byte) error {
	var res CodeVersionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
