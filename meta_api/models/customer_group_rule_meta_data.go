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

// CustomerGroupRuleMetaData Document representing meta data for customer group rules.
// swagger:model customer_group_rule_meta_data
type CustomerGroupRuleMetaData struct {

	// The list of attribute groups.
	Groups []*RuleAttributeGroup `json:"groups"`

	// The quota limits for customer group rules.
	Quotas []*Quota `json:"quotas"`
}

// Validate validates this customer group rule meta data
func (m *CustomerGroupRuleMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerGroupRuleMetaData) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerGroupRuleMetaData) validateQuotas(formats strfmt.Registry) error {

	if swag.IsZero(m.Quotas) { // not required
		return nil
	}

	for i := 0; i < len(m.Quotas); i++ {
		if swag.IsZero(m.Quotas[i]) { // not required
			continue
		}

		if m.Quotas[i] != nil {
			if err := m.Quotas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("quotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerGroupRuleMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerGroupRuleMetaData) UnmarshalBinary(b []byte) error {
	var res CustomerGroupRuleMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
