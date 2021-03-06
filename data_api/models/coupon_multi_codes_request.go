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

// CouponMultiCodesRequest A request object to add and remove coupon codes from a document
// swagger:model coupon_multi_codes_request
type CouponMultiCodesRequest struct {

	// The list of coupon codes to add or delete
	// Required: true
	Codes []string `json:"codes"`
}

// Validate validates this coupon multi codes request
func (m *CouponMultiCodesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CouponMultiCodesRequest) validateCodes(formats strfmt.Registry) error {

	if err := validate.Required("codes", "body", m.Codes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CouponMultiCodesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CouponMultiCodesRequest) UnmarshalBinary(b []byte) error {
	var res CouponMultiCodesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
