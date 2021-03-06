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

// PaymentMethodResult Result document of payment methods applicable for a basket.
// swagger:model payment_method_result
type PaymentMethodResult struct {

	// The applicable payment methods.
	ApplicablePaymentMethods []*PaymentMethod `json:"applicable_payment_methods"`
}

// Validate validates this payment method result
func (m *PaymentMethodResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicablePaymentMethods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentMethodResult) validateApplicablePaymentMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicablePaymentMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplicablePaymentMethods); i++ {
		if swag.IsZero(m.ApplicablePaymentMethods[i]) { // not required
			continue
		}

		if m.ApplicablePaymentMethods[i] != nil {
			if err := m.ApplicablePaymentMethods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applicable_payment_methods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentMethodResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentMethodResult) UnmarshalBinary(b []byte) error {
	var res PaymentMethodResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
