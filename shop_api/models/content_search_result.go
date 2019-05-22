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

// ContentSearchResult Document representing a content search result.
// swagger:model content_search_result
type ContentSearchResult struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// data
	Data []interface{} `json:"data"`

	// The sorted array of search hits. Can be empty.
	Hits []*Content `json:"hits"`

	// The URL of the next result page.
	Next string `json:"next,omitempty"`

	// The URL of the previous result page.
	Previous string `json:"previous,omitempty"`

	// The query String that was searched for.
	Query string `json:"query,omitempty"`

	// The sorted array of search refinements. Can be empty.
	Refinements []*ContentSearchRefinement `json:"refinements"`

	// Map of selected refinement attribute id/value(s) pairs. The sorting order is the same like in request URL.
	SelectedRefinements map[string]string `json:"selected_refinements,omitempty"`

	// The zero-based index of the first search hit to include in the result.
	// Minimum: 0
	Start *int32 `json:"start,omitempty"`

	// The total number of documents.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this content search result
func (m *ContentSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefinements(formats); err != nil {
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

func (m *ContentSearchResult) validateHits(formats strfmt.Registry) error {

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

func (m *ContentSearchResult) validateRefinements(formats strfmt.Registry) error {

	if swag.IsZero(m.Refinements) { // not required
		return nil
	}

	for i := 0; i < len(m.Refinements); i++ {
		if swag.IsZero(m.Refinements[i]) { // not required
			continue
		}

		if m.Refinements[i] != nil {
			if err := m.Refinements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("refinements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentSearchResult) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.MinimumInt("start", "body", int64(*m.Start), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentSearchResult) UnmarshalBinary(b []byte) error {
	var res ContentSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
