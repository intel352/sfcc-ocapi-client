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

// OrderPaymentCardRequest Document representing an order payment card request.
// swagger:model order_payment_card_request
type OrderPaymentCardRequest struct {

	// card type
	// Max Length: 256
	CardType string `json:"card_type,omitempty"`

	// credit card token
	CreditCardToken string `json:"credit_card_token,omitempty"`

	// expiration month
	ExpirationMonth int32 `json:"expiration_month,omitempty"`

	// expiration year
	ExpirationYear int32 `json:"expiration_year,omitempty"`

	// holder
	// Max Length: 256
	Holder string `json:"holder,omitempty"`

	// issue number
	// Max Length: 256
	IssueNumber string `json:"issue_number,omitempty"`

	// number
	// Max Length: 4000
	Number string `json:"number,omitempty"`

	// The security code for the payment card.
	// Max Length: 4
	SecurityCode string `json:"security_code,omitempty"`

	// valid from month
	// Maximum: 12
	// Minimum: 1
	ValidFromMonth int32 `json:"valid_from_month,omitempty"`

	// valid from year
	ValidFromYear int32 `json:"valid_from_year,omitempty"`
}

// Validate validates this order payment card request
func (m *OrderPaymentCardRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCardType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHolder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFromMonth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderPaymentCardRequest) validateCardType(formats strfmt.Registry) error {

	if swag.IsZero(m.CardType) { // not required
		return nil
	}

	if err := validate.MaxLength("card_type", "body", string(m.CardType), 256); err != nil {
		return err
	}

	return nil
}

func (m *OrderPaymentCardRequest) validateHolder(formats strfmt.Registry) error {

	if swag.IsZero(m.Holder) { // not required
		return nil
	}

	if err := validate.MaxLength("holder", "body", string(m.Holder), 256); err != nil {
		return err
	}

	return nil
}

func (m *OrderPaymentCardRequest) validateIssueNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.IssueNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("issue_number", "body", string(m.IssueNumber), 256); err != nil {
		return err
	}

	return nil
}

func (m *OrderPaymentCardRequest) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaxLength("number", "body", string(m.Number), 4000); err != nil {
		return err
	}

	return nil
}

func (m *OrderPaymentCardRequest) validateSecurityCode(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityCode) { // not required
		return nil
	}

	if err := validate.MaxLength("security_code", "body", string(m.SecurityCode), 4); err != nil {
		return err
	}

	return nil
}

func (m *OrderPaymentCardRequest) validateValidFromMonth(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidFromMonth) { // not required
		return nil
	}

	if err := validate.MinimumInt("valid_from_month", "body", int64(m.ValidFromMonth), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("valid_from_month", "body", int64(m.ValidFromMonth), 12, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderPaymentCardRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderPaymentCardRequest) UnmarshalBinary(b []byte) error {
	var res OrderPaymentCardRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
