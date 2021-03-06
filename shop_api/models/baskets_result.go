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

// BasketsResult Result document containing an array of baskets.
// swagger:model baskets_result
type BasketsResult struct {

	// The list of baskets for a customer.
	Baskets []*Basket `json:"baskets"`

	// The total number of baskets.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this baskets result
func (m *BasketsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaskets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasketsResult) validateBaskets(formats strfmt.Registry) error {

	if swag.IsZero(m.Baskets) { // not required
		return nil
	}

	for i := 0; i < len(m.Baskets); i++ {
		if swag.IsZero(m.Baskets[i]) { // not required
			continue
		}

		if m.Baskets[i] != nil {
			if err := m.Baskets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("baskets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BasketsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasketsResult) UnmarshalBinary(b []byte) error {
	var res BasketsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
