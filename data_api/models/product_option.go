// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductOption Document representing a product option
// swagger:model product_option
type ProductOption struct {

	// The localized custom name of the product option.
	CustomName map[string]string `json:"custom_name,omitempty"`

	// The default product option value.
	DefaultProductOptionValue string `json:"default_product_option_value,omitempty"`

	// The localized description of the product option.
	Description map[string]string `json:"description,omitempty"`

	// The object attribute definition id which is also the identifier for the product option.
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// The image assigned to the product option.
	Image *MediaFile `json:"image,omitempty"`

	// The URL link to the product option.
	Link string `json:"link,omitempty"`

	// The name of the object attribute definition.
	Name string `json:"name,omitempty"`

	// The selected option value of the product option.
	SelectedOptionValue string `json:"selected_option_value,omitempty"`

	// The flag that indicates if the product option is shared or local.
	Shared bool `json:"shared,omitempty"`

	// The sorting mode for the product option values.
	// Enum: [byexplicitorder byoptionprice]
	SortingMode string `json:"sorting_mode,omitempty"`

	// The sorted array of values of the product option.
	Values []*ProductOptionValue `json:"values"`
}

// Validate validates this product option
func (m *ProductOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductOption) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *ProductOption) validateImage(formats strfmt.Registry) error {

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

var productOptionTypeSortingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["byexplicitorder","byoptionprice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productOptionTypeSortingModePropEnum = append(productOptionTypeSortingModePropEnum, v)
	}
}

const (

	// ProductOptionSortingModeByexplicitorder captures enum value "byexplicitorder"
	ProductOptionSortingModeByexplicitorder string = "byexplicitorder"

	// ProductOptionSortingModeByoptionprice captures enum value "byoptionprice"
	ProductOptionSortingModeByoptionprice string = "byoptionprice"
)

// prop value enum
func (m *ProductOption) validateSortingModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productOptionTypeSortingModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProductOption) validateSortingMode(formats strfmt.Registry) error {

	if swag.IsZero(m.SortingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateSortingModeEnum("sorting_mode", "body", m.SortingMode); err != nil {
		return err
	}

	return nil
}

func (m *ProductOption) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductOption) UnmarshalBinary(b []byte) error {
	var res ProductOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
