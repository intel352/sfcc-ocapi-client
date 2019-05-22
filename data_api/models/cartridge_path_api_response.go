// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CartridgePathAPIResponse Response of cartridge path related operation
// swagger:model cartridge_path_api_response
type CartridgePathAPIResponse struct {

	// Updated cartridge path
	Cartridges string `json:"cartridges,omitempty"`

	// Site id
	SiteID string `json:"site_id,omitempty"`
}

// Validate validates this cartridge path api response
func (m *CartridgePathAPIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CartridgePathAPIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CartridgePathAPIResponse) UnmarshalBinary(b []byte) error {
	var res CartridgePathAPIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
