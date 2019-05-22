// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RuleOperator Document representing an operator for a customer group attribute.
// swagger:model rule_operator
type RuleOperator struct {

	// The display value of the operator.
	DisplayValue string `json:"display_value,omitempty"`

	// The operator of the attribute.
	Operator string `json:"operator,omitempty"`

	// Possible value restrictions for the operator.
	ValueRestriction string `json:"value_restriction,omitempty"`

	// The attribute type for which the operator is valid.
	//  This is only set if the operator is restricted to a specific attribute type.
	ValueType string `json:"value_type,omitempty"`
}

// Validate validates this rule operator
func (m *RuleOperator) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleOperator) UnmarshalBinary(b []byte) error {
	var res RuleOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}