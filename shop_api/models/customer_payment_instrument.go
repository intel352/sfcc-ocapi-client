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

// CustomerPaymentInstrument Document representing a customer payment instrument.
// swagger:model customer_payment_instrument
type CustomerPaymentInstrument struct {

	// The bank routing number.
	// Max Length: 256
	BankRoutingNumber string `json:"bank_routing_number,omitempty"`

	// Returns the value of attribute 'creationDate'.
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// Returns the value of attribute 'lastModified'.
	// Read Only: true
	// Format: date-time
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// The masked gift certificate code.
	MaskedGiftCertificateCode string `json:"masked_gift_certificate_code,omitempty"`

	// The payment bank account.
	PaymentBankAccount *PaymentBankAccount `json:"payment_bank_account,omitempty"`

	// The payment card.
	PaymentCard *PaymentCard `json:"payment_card,omitempty"`

	// The payment instrument ID.
	PaymentInstrumentID string `json:"payment_instrument_id,omitempty"`

	// The payment method id. Optional if a customer payment instrument id is specified.
	// Max Length: 256
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

// Validate validates this customer payment instrument
func (m *CustomerPaymentInstrument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankRoutingNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentBankAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerPaymentInstrument) validateBankRoutingNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.BankRoutingNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("bank_routing_number", "body", string(m.BankRoutingNumber), 256); err != nil {
		return err
	}

	return nil
}

func (m *CustomerPaymentInstrument) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_date", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerPaymentInstrument) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerPaymentInstrument) validatePaymentBankAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentBankAccount) { // not required
		return nil
	}

	if m.PaymentBankAccount != nil {
		if err := m.PaymentBankAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_bank_account")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerPaymentInstrument) validatePaymentCard(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentCard) { // not required
		return nil
	}

	if m.PaymentCard != nil {
		if err := m.PaymentCard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_card")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerPaymentInstrument) validatePaymentMethodID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.MaxLength("payment_method_id", "body", string(m.PaymentMethodID), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerPaymentInstrument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerPaymentInstrument) UnmarshalBinary(b []byte) error {
	var res CustomerPaymentInstrument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}