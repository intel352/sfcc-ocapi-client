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

// SlotConfigurationCampaignAssignment Document representing a slot_configuration_campaign_assignment
// swagger:model slot_configuration_campaign_assignment
type SlotConfigurationCampaignAssignment struct {

	// The campaign.
	Campaign *Campaign `json:"campaign,omitempty"`

	// The id of the campaign that has the slot configuration assigned to it.
	// Max Length: 256
	// Min Length: 1
	CampaignID string `json:"campaign_id,omitempty"`

	// The slot context.
	// Required: true
	// Enum: [global category folder]
	Context *string `json:"context"`

	// The list of customer groups.
	CustomerGroups []string `json:"customer_groups"`

	// The description of the slot configuration.
	// Max Length: 4000
	Description string `json:"description,omitempty"`

	// True if the assignment resource is enabled
	Enabled bool `json:"enabled,omitempty"`

	// The URL to the slot configuration-campaign assignment.
	Link string `json:"link,omitempty"`

	// The rank of the slot confiuration-campaign assignment.
	//  This is different than the rank of the slot configuration.
	// Minimum: 1
	Rank int32 `json:"rank,omitempty"`

	// The schedule of the assignment resource
	Schedule *Schedule `json:"schedule,omitempty"`

	// The slotConfiguration
	SlotConfiguration *SlotConfiguration `json:"slot_configuration,omitempty"`

	// The ID of the slot configuration.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	SlotConfigurationID *string `json:"slot_configuration_id"`

	// The UUID of the slot configuration.
	// Max Length: 28
	SlotConfigurationUUID string `json:"slot_configuration_uuid,omitempty"`

	// The ID of the slot's context, for example, the category ID for a slot with category context.
	// Max Length: 256
	SlotContextID string `json:"slot_context_id,omitempty"`

	// The ID of the slot.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	SlotID *string `json:"slot_id"`
}

// Validate validates this slot configuration campaign assignment
func (m *SlotConfigurationCampaignAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCampaignID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlotConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlotConfigurationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlotConfigurationUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlotContextID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateCampaign(formats strfmt.Registry) error {

	if swag.IsZero(m.Campaign) { // not required
		return nil
	}

	if m.Campaign != nil {
		if err := m.Campaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("campaign")
			}
			return err
		}
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateCampaignID(formats strfmt.Registry) error {

	if swag.IsZero(m.CampaignID) { // not required
		return nil
	}

	if err := validate.MinLength("campaign_id", "body", string(m.CampaignID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("campaign_id", "body", string(m.CampaignID), 256); err != nil {
		return err
	}

	return nil
}

var slotConfigurationCampaignAssignmentTypeContextPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["global","category","folder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		slotConfigurationCampaignAssignmentTypeContextPropEnum = append(slotConfigurationCampaignAssignmentTypeContextPropEnum, v)
	}
}

const (

	// SlotConfigurationCampaignAssignmentContextGlobal captures enum value "global"
	SlotConfigurationCampaignAssignmentContextGlobal string = "global"

	// SlotConfigurationCampaignAssignmentContextCategory captures enum value "category"
	SlotConfigurationCampaignAssignmentContextCategory string = "category"

	// SlotConfigurationCampaignAssignmentContextFolder captures enum value "folder"
	SlotConfigurationCampaignAssignmentContextFolder string = "folder"
)

// prop value enum
func (m *SlotConfigurationCampaignAssignment) validateContextEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, slotConfigurationCampaignAssignmentTypeContextPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateContext(formats strfmt.Registry) error {

	if err := validate.Required("context", "body", m.Context); err != nil {
		return err
	}

	// value enum
	if err := m.validateContextEnum("context", "body", *m.Context); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 4000); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateRank(formats strfmt.Registry) error {

	if swag.IsZero(m.Rank) { // not required
		return nil
	}

	if err := validate.MinimumInt("rank", "body", int64(m.Rank), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateSchedule(formats strfmt.Registry) error {

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

func (m *SlotConfigurationCampaignAssignment) validateSlotConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.SlotConfiguration) { // not required
		return nil
	}

	if m.SlotConfiguration != nil {
		if err := m.SlotConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slot_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateSlotConfigurationID(formats strfmt.Registry) error {

	if err := validate.Required("slot_configuration_id", "body", m.SlotConfigurationID); err != nil {
		return err
	}

	if err := validate.MinLength("slot_configuration_id", "body", string(*m.SlotConfigurationID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slot_configuration_id", "body", string(*m.SlotConfigurationID), 256); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateSlotConfigurationUUID(formats strfmt.Registry) error {

	if swag.IsZero(m.SlotConfigurationUUID) { // not required
		return nil
	}

	if err := validate.MaxLength("slot_configuration_uuid", "body", string(m.SlotConfigurationUUID), 28); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateSlotContextID(formats strfmt.Registry) error {

	if swag.IsZero(m.SlotContextID) { // not required
		return nil
	}

	if err := validate.MaxLength("slot_context_id", "body", string(m.SlotContextID), 256); err != nil {
		return err
	}

	return nil
}

func (m *SlotConfigurationCampaignAssignment) validateSlotID(formats strfmt.Registry) error {

	if err := validate.Required("slot_id", "body", m.SlotID); err != nil {
		return err
	}

	if err := validate.MinLength("slot_id", "body", string(*m.SlotID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slot_id", "body", string(*m.SlotID), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SlotConfigurationCampaignAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlotConfigurationCampaignAssignment) UnmarshalBinary(b []byte) error {
	var res SlotConfigurationCampaignAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
