// Code generated by go-swagger; DO NOT EDIT.

package global_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new global preferences API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for global preferences API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetGlobalPreferencesPreferenceGroupsByIDByID For the specified instance, read the custom preferences in the preference group at the global(organization) level.
 Specify 'current' to retrieve the preferences for the instance on which this call is being made. The system will recognize its type.
*/
func (a *Client) GetGlobalPreferencesPreferenceGroupsByIDByID(params *GetGlobalPreferencesPreferenceGroupsByIDByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalPreferencesPreferenceGroupsByIDByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGlobalPreferencesPreferenceGroupsByIDByID",
		Method:             "GET",
		PathPattern:        "/global_preferences/preference_groups/{group_id}/{instance_type}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGlobalPreferencesPreferenceGroupsByIDByIDReader{formats: a.formats},
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
PatchGlobalPreferencesPreferenceGroupsByIDByID For the specified instance, update one or more custom preferences in the preference group  at the global(organization) level.
*/
func (a *Client) PatchGlobalPreferencesPreferenceGroupsByIDByID(params *PatchGlobalPreferencesPreferenceGroupsByIDByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchGlobalPreferencesPreferenceGroupsByIDByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchGlobalPreferencesPreferenceGroupsByIDByID",
		Method:             "PATCH",
		PathPattern:        "/global_preferences/preference_groups/{group_id}/{instance_type}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchGlobalPreferencesPreferenceGroupsByIDByIDReader{formats: a.formats},
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
