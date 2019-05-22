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

// PaymentBankAccountRequest Document representing a payment bank account request.
// swagger:model payment_bank_account_request
type PaymentBankAccountRequest struct {

	// The drivers license.
	// Max Length: 256
	DriversLicense string `json:"drivers_license,omitempty"`

	// The driver license state code.
	// Max Length: 256
	DriversLicenseStateCode string `json:"drivers_license_state_code,omitempty"`

	// The holder of the bank account.
	// Max Length: 256
	Holder string `json:"holder,omitempty"`

	// The payment bank account number.
	// Max Length: 256
	Number string `json:"number,omitempty"`
}

// Validate validates this payment bank account request
func (m *PaymentBankAccountRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriversLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriversLicenseStateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentBankAccountRequest) validateDriversLicense(formats strfmt.Registry) error {

	if swag.IsZero(m.DriversLicense) { // not required
		return nil
	}

	if err := validate.MaxLength("drivers_license", "body", string(m.DriversLicense), 256); err != nil {
		return err
	}

	return nil
}

func (m *PaymentBankAccountRequest) validateDriversLicenseStateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.DriversLicenseStateCode) { // not required
		return nil
	}

	if err := validate.MaxLength("drivers_license_state_code", "body", string(m.DriversLicenseStateCode), 256); err != nil {
		return err
	}

	return nil
}

func (m *PaymentBankAccountRequest) validateHolder(formats strfmt.Registry) error {

	if swag.IsZero(m.Holder) { // not required
		return nil
	}

	if err := validate.MaxLength("holder", "body", string(m.Holder), 256); err != nil {
		return err
	}

	return nil
}

func (m *PaymentBankAccountRequest) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaxLength("number", "body", string(m.Number), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentBankAccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentBankAccountRequest) UnmarshalBinary(b []byte) error {
	var res PaymentBankAccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
