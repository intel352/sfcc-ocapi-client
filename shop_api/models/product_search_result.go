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

// ProductSearchResult Document representing a product search result.
// swagger:model product_search_result
type ProductSearchResult struct {

	// The number of returned documents.
	Count int32 `json:"count,omitempty"`

	// data
	Data []interface{} `json:"data"`

	// fetch date
	FetchDate int32 `json:"fetch_date,omitempty"`

	// The sorted array of search hits. This array can be empty.
	Hits []*ProductSearchHit `json:"hits"`

	// The URL of the next result page.
	Next string `json:"next,omitempty"`

	// The URL of the previous result page.
	Previous string `json:"previous,omitempty"`

	// The query String that was searched for.
	Query string `json:"query,omitempty"`

	// The sorted array of search refinements. This array can be empty.
	Refinements []*ProductSearchRefinement `json:"refinements"`

	// The suggestion given by the system for the submitted search phrase.
	SearchPhraseSuggestions *Suggestion `json:"search_phrase_suggestions,omitempty"`

	// A map of selected refinement attribute id/value(s) pairs. The sorting order is the same as in request URL.
	SelectedRefinements map[string]string `json:"selected_refinements,omitempty"`

	// The id of the applied sorting option.
	SelectedSortingOption string `json:"selected_sorting_option,omitempty"`

	// The sorted array of search sorting options. This array can be empty.
	SortingOptions []*ProductSearchSortingOption `json:"sorting_options"`

	// The zero-based index of the first search hit to include in the result.
	// Minimum: 0
	Start *int32 `json:"start,omitempty"`

	// The total number of documents.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this product search result
func (m *ProductSearchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefinements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchPhraseSuggestions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortingOptions(formats); err != nil {
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

func (m *ProductSearchResult) validateHits(formats strfmt.Registry) error {

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

func (m *ProductSearchResult) validateRefinements(formats strfmt.Registry) error {

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

func (m *ProductSearchResult) validateSearchPhraseSuggestions(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchPhraseSuggestions) { // not required
		return nil
	}

	if m.SearchPhraseSuggestions != nil {
		if err := m.SearchPhraseSuggestions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("search_phrase_suggestions")
			}
			return err
		}
	}

	return nil
}

func (m *ProductSearchResult) validateSortingOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.SortingOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.SortingOptions); i++ {
		if swag.IsZero(m.SortingOptions[i]) { // not required
			continue
		}

		if m.SortingOptions[i] != nil {
			if err := m.SortingOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorting_options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductSearchResult) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.MinimumInt("start", "body", int64(*m.Start), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductSearchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductSearchResult) UnmarshalBinary(b []byte) error {
	var res ProductSearchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
