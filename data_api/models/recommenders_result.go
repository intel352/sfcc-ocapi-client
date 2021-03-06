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

// RecommendersResult A list of recommenders available for use in recommendation requests.
// swagger:model recommenders_result
type RecommendersResult struct {

	// The recommender objects
	Recommenders []*Recommender `json:"recommenders"`
}

// Validate validates this recommenders result
func (m *RecommendersResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecommenders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecommendersResult) validateRecommenders(formats strfmt.Registry) error {

	if swag.IsZero(m.Recommenders) { // not required
		return nil
	}

	for i := 0; i < len(m.Recommenders); i++ {
		if swag.IsZero(m.Recommenders[i]) { // not required
			continue
		}

		if m.Recommenders[i] != nil {
			if err := m.Recommenders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recommenders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecommendersResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecommendersResult) UnmarshalBinary(b []byte) error {
	var res RecommendersResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
