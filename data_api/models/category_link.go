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

// CategoryLink Document representing a category link
// swagger:model category_link
type CategoryLink struct {

	// The date the link was last modified
	// Format: date-time
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// URL that is used to get this instance
	Link string `json:"link,omitempty"`

	// The position in the source catalog / category for this link relative to the other links in the same category.
	// Minimum: 0
	Position *float64 `json:"position,omitempty"`

	// The source catalog for the link
	SourceCatalogID string `json:"source_catalog_id,omitempty"`

	// The name of the source catalog
	SourceCatalogName map[string]string `json:"source_catalog_name,omitempty"`

	// The source category for the link
	SourceCategoryID string `json:"source_category_id,omitempty"`

	// The name of the source category
	SourceCategoryName map[string]string `json:"source_category_name,omitempty"`

	// The target category for the link
	TargetCatalogID string `json:"target_catalog_id,omitempty"`

	// The name of the target catalog
	TargetCatalogName map[string]string `json:"target_catalog_name,omitempty"`

	// The target category for the link
	TargetCategoryID string `json:"target_category_id,omitempty"`

	// The name of the target category
	TargetCategoryName map[string]string `json:"target_category_name,omitempty"`

	// The link type
	// Enum: [other accessories cross_selling up_selling spare_parts]
	Type string `json:"type,omitempty"`
}

// Validate validates this category link
func (m *CategoryLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
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

func (m *CategoryLink) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CategoryLink) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.Minimum("position", "body", float64(*m.Position), 0, false); err != nil {
		return err
	}

	return nil
}

var categoryLinkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["other","accessories","cross_selling","up_selling","spare_parts"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		categoryLinkTypeTypePropEnum = append(categoryLinkTypeTypePropEnum, v)
	}
}

const (

	// CategoryLinkTypeOther captures enum value "other"
	CategoryLinkTypeOther string = "other"

	// CategoryLinkTypeAccessories captures enum value "accessories"
	CategoryLinkTypeAccessories string = "accessories"

	// CategoryLinkTypeCrossSelling captures enum value "cross_selling"
	CategoryLinkTypeCrossSelling string = "cross_selling"

	// CategoryLinkTypeUpSelling captures enum value "up_selling"
	CategoryLinkTypeUpSelling string = "up_selling"

	// CategoryLinkTypeSpareParts captures enum value "spare_parts"
	CategoryLinkTypeSpareParts string = "spare_parts"
)

// prop value enum
func (m *CategoryLink) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, categoryLinkTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CategoryLink) validateType(formats strfmt.Registry) error {

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
func (m *CategoryLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryLink) UnmarshalBinary(b []byte) error {
	var res CategoryLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
