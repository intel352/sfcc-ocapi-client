// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceObject Object to represent a resource in ocapi_config_api_request
// swagger:model resource_object
type ResourceObject struct {

	// Resource Cache time
	CacheTime int32 `json:"cache_time,omitempty"`

	// Configuration of the resource
	Config map[string]string `json:"config,omitempty"`

	// Allowed methods of the resource
	// Required: true
	// Min Items: 1
	Methods []string `json:"methods"`

	// Indicate if the personalized caching is enabled for the resource
	PersonalizedCachingEnabled bool `json:"personalized_caching_enabled,omitempty"`

	// Read attributes of the resource
	ReadAttributes string `json:"read_attributes,omitempty"`

	// Resource ID
	// Required: true
	// Min Length: 1
	ResourceID *string `json:"resource_id"`

	// Version range
	VersionRange *VersionRangeObject `json:"version_range,omitempty"`

	// Write attributes of the resource
	WriteAttributes string `json:"write_attributes,omitempty"`
}

// Validate validates this resource object
func (m *ResourceObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceObject) validateMethods(formats strfmt.Registry) error {

	if err := validate.Required("methods", "body", m.Methods); err != nil {
		return err
	}

	iMethodsSize := int64(len(m.Methods))

	if err := validate.MinItems("methods", "body", iMethodsSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *ResourceObject) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	if err := validate.MinLength("resource_id", "body", string(*m.ResourceID), 1); err != nil {
		return err
	}

	return nil
}

func (m *ResourceObject) validateVersionRange(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionRange) { // not required
		return nil
	}

	if m.VersionRange != nil {
		if err := m.VersionRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject) UnmarshalBinary(b []byte) error {
	var res ResourceObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
