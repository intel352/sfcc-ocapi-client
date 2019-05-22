// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShippingMethod Document representing a shipping method.
// swagger:model shipping_method
type ShippingMethod struct {

	// c store pickup enabled
	// Required: true
	CStorePickupEnabled *bool `json:"c_storePickupEnabled"`

	// The localized description of the shipping method.
	Description string `json:"description,omitempty"`

	// The external shipping method.
	ExternalShippingMethod string `json:"external_shipping_method,omitempty"`

	// The shipping method id.
	// Required: true
	// Max Length: 256
	ID *string `json:"id"`

	// The localized name of the shipping method.
	Name string `json:"name,omitempty"`

	// The shipping cost total, including shipment level costs and
	//  product level fix and surcharge costs.
	Price float64 `json:"price,omitempty"`

	// The array of active customer shipping promotions for this shipping
	//  method. This array can be empty.
	ShippingPromotions []*ShippingPromotion `json:"shipping_promotions"`
}

// Validate validates this shipping method
func (m *ShippingMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCStorePickupEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingPromotions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingMethod) validateCStorePickupEnabled(formats strfmt.Registry) error {

	if err := validate.Required("c_storePickupEnabled", "body", m.CStorePickupEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ShippingMethod) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 256); err != nil {
		return err
	}

	return nil
}

func (m *ShippingMethod) validateShippingPromotions(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingPromotions) { // not required
		return nil
	}

	for i := 0; i < len(m.ShippingPromotions); i++ {
		if swag.IsZero(m.ShippingPromotions[i]) { // not required
			continue
		}

		if m.ShippingPromotions[i] != nil {
			if err := m.ShippingPromotions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipping_promotions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShippingMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShippingMethod) UnmarshalBinary(b []byte) error {
	var res ShippingMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}