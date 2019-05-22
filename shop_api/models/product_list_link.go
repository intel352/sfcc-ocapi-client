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

// ProductListLink Document representing a link to a product list.
// swagger:model product_list_link
type ProductListLink struct {

	// The description of this product list.
	Description string `json:"description,omitempty"`

	// The target of the link.
	Link string `json:"link,omitempty"`

	// The name of this product list.
	Name string `json:"name,omitempty"`

	// A flag indicating whether the owner made this product list available for access
	//  by other customers.
	Public bool `json:"public,omitempty"`

	// The link title.
	Title string `json:"title,omitempty"`

	// The type of the product list.
	// Enum: [wish_list gift_registry shopping_list custom_1 custom_2 custom_3]
	Type string `json:"type,omitempty"`
}

// Validate validates this product list link
func (m *ProductListLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var productListLinkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wish_list","gift_registry","shopping_list","custom_1","custom_2","custom_3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productListLinkTypeTypePropEnum = append(productListLinkTypeTypePropEnum, v)
	}
}

const (

	// ProductListLinkTypeWishList captures enum value "wish_list"
	ProductListLinkTypeWishList string = "wish_list"

	// ProductListLinkTypeGiftRegistry captures enum value "gift_registry"
	ProductListLinkTypeGiftRegistry string = "gift_registry"

	// ProductListLinkTypeShoppingList captures enum value "shopping_list"
	ProductListLinkTypeShoppingList string = "shopping_list"

	// ProductListLinkTypeCustom1 captures enum value "custom_1"
	ProductListLinkTypeCustom1 string = "custom_1"

	// ProductListLinkTypeCustom2 captures enum value "custom_2"
	ProductListLinkTypeCustom2 string = "custom_2"

	// ProductListLinkTypeCustom3 captures enum value "custom_3"
	ProductListLinkTypeCustom3 string = "custom_3"
)

// prop value enum
func (m *ProductListLink) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productListLinkTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProductListLink) validateType(formats strfmt.Registry) error {

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
func (m *ProductListLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductListLink) UnmarshalBinary(b []byte) error {
	var res ProductListLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
