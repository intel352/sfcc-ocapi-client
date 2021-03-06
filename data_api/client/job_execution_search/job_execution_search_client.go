// Code generated by go-swagger; DO NOT EDIT.

package job_execution_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new job execution search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job execution search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostJobExecutionSearch Searches for job executions.

 The query attribute specifies a complex query that can be used to narrow down the search. This is the list of
 searchable attributes:

 id - String
 job_id - String
 start_time - Date
 end_time - Date
 status - String

 This is the list of sortable attributes:

 job_id - String
 start_time - Date
 end_time - Date
 status - String

*/
func (a *Client) PostJobExecutionSearch(params *PostJobExecutionSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostJobExecutionSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJobExecutionSearch",
		Method:             "POST",
		PathPattern:        "/job_execution_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostJobExecutionSearchReader{formats: a.formats},
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
