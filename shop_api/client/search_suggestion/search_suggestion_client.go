// Code generated by go-swagger; DO NOT EDIT.

package search_suggestion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search suggestion API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search suggestion API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSearchSuggestion Provides keyword search functionality for products, categories, content, brands and custom suggestions.
 Returns suggested products, suggested categories, suggested content, suggested brands and custom suggestions
 for the given search phrase.
*/
func (a *Client) GetSearchSuggestion(params *GetSearchSuggestionParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchSuggestionParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSearchSuggestion",
		Method:             "GET",
		PathPattern:        "/search_suggestion",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSearchSuggestionReader{formats: a.formats},
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
