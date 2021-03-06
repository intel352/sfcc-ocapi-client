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

// NestedQuery Nested query allows to query upon nested documents that are part of a larger document. The classical example is a
//  product master with variants (in one big document) where you want to constraint a search to masters that have
//  variants that match multiple constraints (like color = blue AND size = M).
//
//
// swagger:model nested_query
type NestedQuery struct {

	// path
	// Required: true
	Path *string `json:"path"`

	// query
	// Required: true
	Query Query `json:"query"`

	// score mode
	// Enum: [avg total max none]
	ScoreMode string `json:"score_mode,omitempty"`
}

// Validate validates this nested query
func (m *NestedQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoreMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedQuery) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *NestedQuery) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

var nestedQueryTypeScoreModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["avg","total","max","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedQueryTypeScoreModePropEnum = append(nestedQueryTypeScoreModePropEnum, v)
	}
}

const (

	// NestedQueryScoreModeAvg captures enum value "avg"
	NestedQueryScoreModeAvg string = "avg"

	// NestedQueryScoreModeTotal captures enum value "total"
	NestedQueryScoreModeTotal string = "total"

	// NestedQueryScoreModeMax captures enum value "max"
	NestedQueryScoreModeMax string = "max"

	// NestedQueryScoreModeNone captures enum value "none"
	NestedQueryScoreModeNone string = "none"
)

// prop value enum
func (m *NestedQuery) validateScoreModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nestedQueryTypeScoreModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NestedQuery) validateScoreMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ScoreMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateScoreModeEnum("score_mode", "body", m.ScoreMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedQuery) UnmarshalBinary(b []byte) error {
	var res NestedQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
