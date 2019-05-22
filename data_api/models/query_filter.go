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

// QueryFilter Document representing a query filter. A query filter wraps any query and allows it to be used as a filter.
//
// swagger:model query_filter
type QueryFilter struct {

	// The query, which should be used as a filter.
	// Required: true
	Query Query `json:"query"`
}

// Validate validates this query filter
func (m *QueryFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryFilter) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryFilter) UnmarshalBinary(b []byte) error {
	var res QueryFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}