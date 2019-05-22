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

// Recurrence Document representing a schedule recurrence.
// swagger:model recurrence
type Recurrence struct {

	// The days of week for recurrence.
	// Enum: [monday tuesday wednesday thursday friday saturday sunday]
	DayOfWeek []string `json:"day_of_week"`

	// The time of the day for recurrence.
	TimeOfDay *TimeOfDay `json:"time_of_day,omitempty"`
}

// Validate validates this recurrence
func (m *Recurrence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOfDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recurrenceTypeDayOfWeekPropEnum []interface{}

func init() {
	var res [][]string
	if err := json.Unmarshal([]byte(`["monday","tuesday","wednesday","thursday","friday","saturday","sunday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurrenceTypeDayOfWeekPropEnum = append(recurrenceTypeDayOfWeekPropEnum, v)
	}
}

// prop value enum
func (m *Recurrence) validateDayOfWeekEnum(path, location string, value []string) error {
	if err := validate.Enum(path, location, value, recurrenceTypeDayOfWeekPropEnum); err != nil {
		return err
	}
	return nil
}

var recurrenceDayOfWeekItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["monday","tuesday","wednesday","thursday","friday","saturday","sunday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recurrenceDayOfWeekItemsEnum = append(recurrenceDayOfWeekItemsEnum, v)
	}
}

func (m *Recurrence) validateDayOfWeekItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recurrenceDayOfWeekItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Recurrence) validateDayOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.DayOfWeek) { // not required
		return nil
	}

	// for slice
	if err := m.validateDayOfWeekEnum("day_of_week", "body", m.DayOfWeek); err != nil {
		return err
	}

	for i := 0; i < len(m.DayOfWeek); i++ {

		// value enum
		if err := m.validateDayOfWeekItemsEnum("day_of_week"+"."+strconv.Itoa(i), "body", m.DayOfWeek[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Recurrence) validateTimeOfDay(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeOfDay) { // not required
		return nil
	}

	if m.TimeOfDay != nil {
		if err := m.TimeOfDay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_of_day")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recurrence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recurrence) UnmarshalBinary(b []byte) error {
	var res Recurrence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
