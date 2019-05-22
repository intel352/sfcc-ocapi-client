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

// ContentSearchRefinement Document representing a search refinement attribute.
// swagger:model content_search_refinement
type ContentSearchRefinement struct {

	// The id of the search refinement attribute. In the case of an attribute refinement, this is the
	//  attribute id. Custom attributes are marked by the prefix "c_".
	// Required: true
	AttributeID *string `json:"attribute_id"`

	// The localized label of the refinement.
	Label string `json:"label,omitempty"`

	// The sorted array of refinement values. The array can be empty.
	Values []*ContentSearchRefinementValue `json:"values"`
}

// Validate validates this content search refinement
func (m *ContentSearchRefinement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeID(formats); err != nil {
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

func (m *ContentSearchRefinement) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", m.AttributeID); err != nil {
		return err
	}

	return nil
}

func (m *ContentSearchRefinement) validateValues(formats strfmt.Registry) error {

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
func (m *ContentSearchRefinement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentSearchRefinement) UnmarshalBinary(b []byte) error {
	var res ContentSearchRefinement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
