// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PromotionAbtestGroupAssignment promotion abtest group assignment
// swagger:model promotion_abtest_group_assignment
type PromotionAbtestGroupAssignment struct {

	// abtest description
	AbtestDescription string `json:"abtest_description,omitempty"`

	// abtest id
	AbtestID string `json:"abtest_id,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// schedule
	Schedule *Schedule `json:"schedule,omitempty"`

	// segment description
	SegmentDescription string `json:"segment_description,omitempty"`

	// segment id
	SegmentID string `json:"segment_id,omitempty"`
}

// Validate validates this promotion abtest group assignment
func (m *PromotionAbtestGroupAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PromotionAbtestGroupAssignment) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PromotionAbtestGroupAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PromotionAbtestGroupAssignment) UnmarshalBinary(b []byte) error {
	var res PromotionAbtestGroupAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
