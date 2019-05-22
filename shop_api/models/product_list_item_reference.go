// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductListItemReference product list item reference
// swagger:model product_list_item_reference
type ProductListItemReference struct {

	// The id of the product list item.
	// Required: true
	ID *string `json:"id"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// product details link
	ProductDetailsLink *ProductDetailsLink `json:"product_details_link,omitempty"`

	// The link of the product list, the item is assigned
	ProductList *ProductListLink `json:"product_list,omitempty"`

	// public
	Public bool `json:"public,omitempty"`

	// purchased quantity
	PurchasedQuantity float64 `json:"purchased_quantity,omitempty"`

	// quantity
	// Minimum: 0
	Quantity *float64 `json:"quantity,omitempty"`

	// type
	// Enum: [product gift_certificate]
	Type string `json:"type,omitempty"`
}

// Validate validates this product list item reference
func (m *ProductListItemReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductDetailsLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductListItemReference) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProductListItemReference) validateProductDetailsLink(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductDetailsLink) { // not required
		return nil
	}

	if m.ProductDetailsLink != nil {
		if err := m.ProductDetailsLink.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product_details_link")
			}
			return err
		}
	}

	return nil
}

func (m *ProductListItemReference) validateProductList(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductList) { // not required
		return nil
	}

	if m.ProductList != nil {
		if err := m.ProductList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product_list")
			}
			return err
		}
	}

	return nil
}

func (m *ProductListItemReference) validateQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	if err := validate.Minimum("quantity", "body", float64(*m.Quantity), 0, false); err != nil {
		return err
	}

	return nil
}

var productListItemReferenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["product","gift_certificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productListItemReferenceTypeTypePropEnum = append(productListItemReferenceTypeTypePropEnum, v)
	}
}

const (

	// ProductListItemReferenceTypeProduct captures enum value "product"
	ProductListItemReferenceTypeProduct string = "product"

	// ProductListItemReferenceTypeGiftCertificate captures enum value "gift_certificate"
	ProductListItemReferenceTypeGiftCertificate string = "gift_certificate"
)

// prop value enum
func (m *ProductListItemReference) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productListItemReferenceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProductListItemReference) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductListItemReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductListItemReference) UnmarshalBinary(b []byte) error {
	var res ProductListItemReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}