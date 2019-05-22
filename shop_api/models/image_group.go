// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ImageGroup Document representing an image group containing a list of images for a particular view type and an optional variation value.
// swagger:model image_group
type ImageGroup struct {

	// The images of the image group.
	Images []*Image `json:"images"`

	// Returns a list of variation attributes applying to this image group.
	VariationAttributes []*VariationAttribute `json:"variation_attributes"`

	// The image view type.
	ViewType string `json:"view_type,omitempty"`
}

// Validate validates this image group
func (m *ImageGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariationAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageGroup) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImageGroup) validateVariationAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.VariationAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.VariationAttributes); i++ {
		if swag.IsZero(m.VariationAttributes[i]) { // not required
			continue
		}

		if m.VariationAttributes[i] != nil {
			if err := m.VariationAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variation_attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageGroup) UnmarshalBinary(b []byte) error {
	var res ImageGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}