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

// SlotSearchResult Document representing a slot search result.
// swagger:model slot_search_result
type SlotSearchResult struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// data
	Data []interface{} `json:"data"`

	// The zero-based index of the record that we want to start with, used to optimize special handling
	// Minimum: 0
	DbStartRecord *int32 `json:"db_start_record_,omitempty"`

	// List of expansions to be applied to each search results. Expands are optional
	Expand []string `json:"expand"`

	// The sorted array of search hits. Can be empty.
	Hits []*Slot `json:"hits"`

	// The URL of the next result page.
	Next *ResultPage `json:"next,omitempty"`

	// The URL of the previous result page.
	Previous *ResultPage `json:"previous,omitempty"`

	// The query passed into the search
	Query Query `json:"query,omitempty"`

	// The fields that you want to select.
	Select string `json:"select,omitempty"`

	// The list of sort clauses configured for the search request. Sort clauses are optional.
	Sorts []*Sort `json:"sorts"`

	// The zero-based index of the first search hit to include in the result.
	// Minimum: 0
	Start *int32 `json:"start,omitempty"`

	// The total number of documents.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this slot search result
func (m *SlotSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDbStartRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SlotSearchResult) validateDbStartRecord(formats strfmt.Registry) error {

	if swag.IsZero(m.DbStartRecord) { // not required
		return nil
	}

	if err := validate.MinimumInt("db_start_record_", "body", int64(*m.DbStartRecord), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SlotSearchResult) validateHits(formats strfmt.Registry) error {

	if swag.IsZero(m.Hits) { // not required
		return nil
	}

	for i := 0; i < len(m.Hits); i++ {
		if swag.IsZero(m.Hits[i]) { // not required
			continue
		}

		if m.Hits[i] != nil {
			if err := m.Hits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlotSearchResult) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *SlotSearchResult) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if m.Previous != nil {
		if err := m.Previous.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("previous")
			}
			return err
		}
	}

	return nil
}

func (m *SlotSearchResult) validateSorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Sorts) { // not required
		return nil
	}

	for i := 0; i < len(m.Sorts); i++ {
		if swag.IsZero(m.Sorts[i]) { // not required
			continue
		}

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SlotSearchResult) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.MinimumInt("start", "body", int64(*m.Start), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlotSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlotSearchResult) UnmarshalBinary(b []byte) error {
	var res SlotSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}