// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SuggestedProduct Document representing a product search hit.
// swagger:model suggested_product
type SuggestedProduct struct {

	// The ISO 4217 mnemonic code of the currency.
	Currency string `json:"currency,omitempty"`

	// The first image of the product hit for the configured viewtype.
	Image *Image `json:"image,omitempty"`

	// The URL addressing the product.
	Link string `json:"link,omitempty"`

	// The sales price of the product. In the case of complex products like a master or a set, this is the minimum price of
	//  related child products.
	Price float64 `json:"price,omitempty"`

	// The id (SKU) of the product.
	ProductID string `json:"product_id,omitempty"`

	// The localized name of the product.
	ProductName string `json:"product_name,omitempty"`
}

// Validate validates this suggested product
func (m *SuggestedProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuggestedProduct) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SuggestedProduct) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuggestedProduct) UnmarshalBinary(b []byte) error {
	var res SuggestedProduct
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
