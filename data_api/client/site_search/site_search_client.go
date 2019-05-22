// Code generated by go-swagger; DO NOT EDIT.

package site_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new site search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for site search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostSiteSearch Searches for sites.

 The query attribute specifies a complex query that can be used to narrow down the search. These are the list
 of searchable attributes:

 id - String
 display_name - Localized <String>
 description - Localized <String>
 in_deletion - Boolean


 The output of the query can also be sorted. These are the list of sortable attributes:

 id - String
 display_name - Localized <String>
 description - Localized <String>
 in_deletion - Boolean

*/
func (a *Client) PostSiteSearch(params *PostSiteSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSiteSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSiteSearch",
		Method:             "POST",
		PathPattern:        "/site_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSiteSearchReader{formats: a.formats},
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
