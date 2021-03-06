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

// SlotContent Document representing the content type for a slot.
// swagger:model slot_content
type SlotContent struct {

	// The HTML body (valid only for type 'html').
	Body map[string]MarkupText `json:"body,omitempty"`

	// The category ids (valid only for type 'categories').
	CategoryIds []string `json:"category_ids"`

	// The content asset ids (valid only for type 'content_assets').
	ContentAssetIds []string `json:"content_asset_ids"`

	// The product ids (valid only for type 'products').
	ProductIds []string `json:"product_ids"`

	// The type of content in the slot.
	// Required: true
	// Enum: [products categories content_assets html recommended_products]
	Type *string `json:"type"`
}

// Validate validates this slot content
func (m *SlotContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
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

func (m *SlotContent) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	for k := range m.Body {

		if err := validate.Required("body"+"."+k, "body", m.Body[k]); err != nil {
			return err
		}
		if val, ok := m.Body[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var slotContentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["products","categories","content_assets","html","recommended_products"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		slotContentTypeTypePropEnum = append(slotContentTypeTypePropEnum, v)
	}
}

const (

	// SlotContentTypeProducts captures enum value "products"
	SlotContentTypeProducts string = "products"

	// SlotContentTypeCategories captures enum value "categories"
	SlotContentTypeCategories string = "categories"

	// SlotContentTypeContentAssets captures enum value "content_assets"
	SlotContentTypeContentAssets string = "content_assets"

	// SlotContentTypeHTML captures enum value "html"
	SlotContentTypeHTML string = "html"

	// SlotContentTypeRecommendedProducts captures enum value "recommended_products"
	SlotContentTypeRecommendedProducts string = "recommended_products"
)

// prop value enum
func (m *SlotContent) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, slotContentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SlotContent) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlotContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlotContent) UnmarshalBinary(b []byte) error {
	var res SlotContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
