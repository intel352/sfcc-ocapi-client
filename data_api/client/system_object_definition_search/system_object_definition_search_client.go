// Code generated by go-swagger; DO NOT EDIT.

package system_object_definition_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system object definition search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system object definition search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostSystemObjectDefinitionSearch Searches for system objects.

 The query attribute specifies a complex query that can be used to narrow down the search. These are the list
 of searchable attributes:

 object_type - String
 display_name - Localized<String>
 description - Localized<String>
 read_only - Boolean


 The output of the query can also be sorted. These are the list of sortable attributes:

 object_type - String
 display_name - Localized<String>
 description - Localized<String>
 read_only - Boolean

*/
func (a *Client) PostSystemObjectDefinitionSearch(params *PostSystemObjectDefinitionSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSystemObjectDefinitionSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSystemObjectDefinitionSearch",
		Method:             "POST",
		PathPattern:        "/system_object_definition_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSystemObjectDefinitionSearchReader{formats: a.formats},
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
