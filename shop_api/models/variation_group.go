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

// VariationGroup Document representing a variation group.
// swagger:model variation_group
type VariationGroup struct {

	// The URL addressing the product.
	Link string `json:"link,omitempty"`

	// A flag indicating whether the variation group is orderable.
	Orderable bool `json:"orderable,omitempty"`

	// The sales price of the variation group.
	Price float64 `json:"price,omitempty"`

	// The id (SKU) of the variation group.
	// Max Length: 100
	// Min Length: 1
	ProductID string `json:"product_id,omitempty"`

	// The actual variation attribute id - value pairs.
	VariationValues map[string]string `json:"variation_values,omitempty"`
}

// Validate validates this variation group
func (m *VariationGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariationGroup) validateProductID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductID) { // not required
		return nil
	}

	if err := validate.MinLength("product_id", "body", string(m.ProductID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("product_id", "body", string(m.ProductID), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariationGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariationGroup) UnmarshalBinary(b []byte) error {
	var res VariationGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}