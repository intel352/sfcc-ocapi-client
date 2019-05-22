// Code generated by go-swagger; DO NOT EDIT.

package ocapi_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ocapi configs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ocapi configs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteOcapiConfigsByID Delete a client.
*/
func (a *Client) DeleteOcapiConfigsByID(params *DeleteOcapiConfigsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOcapiConfigsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOcapiConfigsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOcapiConfigsByID",
		Method:             "DELETE",
		PathPattern:        "/ocapi_configs/{clientId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOcapiConfigsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOcapiConfigsByIDNoContent), nil

}

/*
GetOcapiConfigsByID Get all allowed resources for the client.
*/
func (a *Client) GetOcapiConfigsByID(params *GetOcapiConfigsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOcapiConfigsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOcapiConfigsByID",
		Method:             "GET",
		PathPattern:        "/ocapi_configs/{clientId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOcapiConfigsByIDReader{formats: a.formats},
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
PostOcapiConfigsByID Add a client to existing OCAPI configurations. Return errors when client already exists.
*/
func (a *Client) PostOcapiConfigsByID(params *PostOcapiConfigsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOcapiConfigsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOcapiConfigsByID",
		Method:             "POST",
		PathPattern:        "/ocapi_configs/{clientId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOcapiConfigsByIDReader{formats: a.formats},
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
PutOcapiConfigsByID Add a client to existing OCAPI configurations. Overwrite config if the client already exists.
*/
func (a *Client) PutOcapiConfigsByID(params *PutOcapiConfigsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOcapiConfigsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putOcapiConfigsByID",
		Method:             "PUT",
		PathPattern:        "/ocapi_configs/{clientId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutOcapiConfigsByIDReader{formats: a.formats},
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