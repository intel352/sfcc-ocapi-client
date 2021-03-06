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

// ApproachingDiscountResult A result of a approaching discount request.
// swagger:model approaching_discount_result
type ApproachingDiscountResult struct {

	// Lists approaching discounts.
	ApproachingDiscounts []*ApproachingDiscount `json:"approaching_discounts"`
}

// Validate validates this approaching discount result
func (m *ApproachingDiscountResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApproachingDiscounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApproachingDiscountResult) validateApproachingDiscounts(formats strfmt.Registry) error {

	if swag.IsZero(m.ApproachingDiscounts) { // not required
		return nil
	}

	for i := 0; i < len(m.ApproachingDiscounts); i++ {
		if swag.IsZero(m.ApproachingDiscounts[i]) { // not required
			continue
		}

		if m.ApproachingDiscounts[i] != nil {
			if err := m.ApproachingDiscounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("approaching_discounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApproachingDiscountResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApproachingDiscountResult) UnmarshalBinary(b []byte) error {
	var res ApproachingDiscountResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
