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

// PublicProductListLink Document representing a link to a public product list.
// swagger:model public_product_list_link
type PublicProductListLink struct {

	// The description of this product list.
	Description string `json:"description,omitempty"`

	// The target of the link.
	Link string `json:"link,omitempty"`

	// The name of this product list.
	Name string `json:"name,omitempty"`

	// The link title.
	Title string `json:"title,omitempty"`

	// The type of the product list.
	// Enum: [wish_list gift_registry shopping_list custom_1 custom_2 custom_3]
	Type string `json:"type,omitempty"`
}

// Validate validates this public product list link
func (m *PublicProductListLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var publicProductListLinkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wish_list","gift_registry","shopping_list","custom_1","custom_2","custom_3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		publicProductListLinkTypeTypePropEnum = append(publicProductListLinkTypeTypePropEnum, v)
	}
}

const (

	// PublicProductListLinkTypeWishList captures enum value "wish_list"
	PublicProductListLinkTypeWishList string = "wish_list"

	// PublicProductListLinkTypeGiftRegistry captures enum value "gift_registry"
	PublicProductListLinkTypeGiftRegistry string = "gift_registry"

	// PublicProductListLinkTypeShoppingList captures enum value "shopping_list"
	PublicProductListLinkTypeShoppingList string = "shopping_list"

	// PublicProductListLinkTypeCustom1 captures enum value "custom_1"
	PublicProductListLinkTypeCustom1 string = "custom_1"

	// PublicProductListLinkTypeCustom2 captures enum value "custom_2"
	PublicProductListLinkTypeCustom2 string = "custom_2"

	// PublicProductListLinkTypeCustom3 captures enum value "custom_3"
	PublicProductListLinkTypeCustom3 string = "custom_3"
)

// prop value enum
func (m *PublicProductListLink) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, publicProductListLinkTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PublicProductListLink) validateType(formats strfmt.Registry) error {

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
func (m *PublicProductListLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicProductListLink) UnmarshalBinary(b []byte) error {
	var res PublicProductListLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
