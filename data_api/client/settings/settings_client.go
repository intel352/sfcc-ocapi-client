// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSettingsLoggingCustom Action to get custom log settings.
*/
func (a *Client) GetSettingsLoggingCustom(params *GetSettingsLoggingCustomParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsLoggingCustomParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSettingsLoggingCustom",
		Method:             "GET",
		PathPattern:        "/settings/logging/custom",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSettingsLoggingCustomReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PatchSettingsLoggingCustom Updates the custom log settings.
*/
func (a *Client) PatchSettingsLoggingCustom(params *PatchSettingsLoggingCustomParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchSettingsLoggingCustomParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchSettingsLoggingCustom",
		Method:             "PATCH",
		PathPattern:        "/settings/logging/custom",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchSettingsLoggingCustomReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
