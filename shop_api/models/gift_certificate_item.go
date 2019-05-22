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

// GiftCertificateItem A gift certificate item.
// swagger:model gift_certificate_item
type GiftCertificateItem struct {

	// The certificate item amount.
	// Required: true
	Amount *float64 `json:"amount"`

	// Id used to identify this item
	GiftCertificateItemID string `json:"gift_certificate_item_id,omitempty"`

	// The certificate's message.
	// Max Length: 4000
	Message string `json:"message,omitempty"`

	// The recipient's email.
	// Required: true
	// Min Length: 1
	RecipientEmail *string `json:"recipient_email"`

	// The recipient's name.
	RecipientName string `json:"recipient_name,omitempty"`

	// The sender's name.
	SenderName string `json:"sender_name,omitempty"`

	// The shipment id.
	ShipmentID string `json:"shipment_id,omitempty"`
}

// Validate validates this gift certificate item
func (m *GiftCertificateItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GiftCertificateItem) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GiftCertificateItem) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if err := validate.MaxLength("message", "body", string(m.Message), 4000); err != nil {
		return err
	}

	return nil
}

func (m *GiftCertificateItem) validateRecipientEmail(formats strfmt.Registry) error {

	if err := validate.Required("recipient_email", "body", m.RecipientEmail); err != nil {
		return err
	}

	if err := validate.MinLength("recipient_email", "body", string(*m.RecipientEmail), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GiftCertificateItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GiftCertificateItem) UnmarshalBinary(b []byte) error {
	var res GiftCertificateItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}