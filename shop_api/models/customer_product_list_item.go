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

// CustomerProductListItem Document representing a customer product list item.
// swagger:model customer_product_list_item
type CustomerProductListItem struct {

	// The id of this product list item.
	ID string `json:"id,omitempty"`

	// The priority of the item.
	Priority int32 `json:"priority,omitempty"`

	// The product item
	Product *Product `json:"product,omitempty"`

	// A link to the product.
	ProductDetailsLink *ProductSimpleLink `json:"product_details_link,omitempty"`

	// The id of the product.
	ProductID string `json:"product_id,omitempty"`

	// Is this product list item available for access by other customers?
	Public bool `json:"public,omitempty"`

	// The quantity of products already purchased.
	PurchasedQuantity float64 `json:"purchased_quantity,omitempty"`

	// The quantity of this product list item.
	// Minimum: 0
	Quantity *float64 `json:"quantity,omitempty"`

	// The type of the item.
	// Enum: [product gift_certificate]
	Type string `json:"type,omitempty"`
}

// Validate validates this customer product list item
func (m *CustomerProductListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductDetailsLink(formats); err != nil {
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

func (m *CustomerProductListItem) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerProductListItem) validateProductDetailsLink(formats strfmt.Registry) error {

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

func (m *CustomerProductListItem) validateQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.Quantity) { // not required
		return nil
	}

	if err := validate.Minimum("quantity", "body", float64(*m.Quantity), 0, false); err != nil {
		return err
	}

	return nil
}

var customerProductListItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["product","gift_certificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customerProductListItemTypeTypePropEnum = append(customerProductListItemTypeTypePropEnum, v)
	}
}

const (

	// CustomerProductListItemTypeProduct captures enum value "product"
	CustomerProductListItemTypeProduct string = "product"

	// CustomerProductListItemTypeGiftCertificate captures enum value "gift_certificate"
	CustomerProductListItemTypeGiftCertificate string = "gift_certificate"
)

// prop value enum
func (m *CustomerProductListItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customerProductListItemTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomerProductListItem) validateType(formats strfmt.Registry) error {

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
func (m *CustomerProductListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerProductListItem) UnmarshalBinary(b []byte) error {
	var res CustomerProductListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
