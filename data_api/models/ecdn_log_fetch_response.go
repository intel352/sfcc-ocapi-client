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

// EcdnLogFetchResponse Response object, providing the status of the current log fetch request.
// swagger:model ecdn_log_fetch_response
type EcdnLogFetchResponse struct {

	// ID of the log fetch request
	ID string `json:"id,omitempty"`

	// HTTPS Download link to the fetched log file, which has a lifetime of 30 minutes. This link will only appear, if the current status of the log fetching is 'finished'.
	Link string `json:"link,omitempty"`

	// Current status of the log fetch request
	// Enum: [pending running finished]
	Status string `json:"status,omitempty"`
}

// Validate validates this ecdn log fetch response
func (m *EcdnLogFetchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ecdnLogFetchResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","running","finished"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ecdnLogFetchResponseTypeStatusPropEnum = append(ecdnLogFetchResponseTypeStatusPropEnum, v)
	}
}

const (

	// EcdnLogFetchResponseStatusPending captures enum value "pending"
	EcdnLogFetchResponseStatusPending string = "pending"

	// EcdnLogFetchResponseStatusRunning captures enum value "running"
	EcdnLogFetchResponseStatusRunning string = "running"

	// EcdnLogFetchResponseStatusFinished captures enum value "finished"
	EcdnLogFetchResponseStatusFinished string = "finished"
)

// prop value enum
func (m *EcdnLogFetchResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, ecdnLogFetchResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EcdnLogFetchResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EcdnLogFetchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EcdnLogFetchResponse) UnmarshalBinary(b []byte) error {
	var res EcdnLogFetchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
