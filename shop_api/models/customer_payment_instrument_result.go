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

// CustomerPaymentInstrumentResult Document representing a customer payment instrument result. The payment data contained is masked where needed for security purposes.
// swagger:model customer_payment_instrument_result
type CustomerPaymentInstrumentResult struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// The customer payment instruments list.
	Data []*CustomerPaymentInstrument `json:"data"`

	// The total number of documents.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this customer payment instrument result
func (m *CustomerPaymentInstrumentResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerPaymentInstrumentResult) validateData(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CustomerPaymentInstrumentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerPaymentInstrumentResult) UnmarshalBinary(b []byte) error {
	var res CustomerPaymentInstrumentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}