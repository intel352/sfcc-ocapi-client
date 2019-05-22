// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteJobsByIDExecutionsByID Deletes job execution information using the specified ID for the job having the specified job ID.
*/
func (a *Client) DeleteJobsByIDExecutionsByID(params *DeleteJobsByIDExecutionsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteJobsByIDExecutionsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteJobsByIDExecutionsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteJobsByIDExecutionsByID",
		Method:             "DELETE",
		PathPattern:        "/jobs/{job_id}/executions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteJobsByIDExecutionsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJobsByIDExecutionsByIDNoContent), nil

}

/*
GetJobsByIDExecutionsByID Returns job execution information using the specified ID for the job having the specified job ID.
*/
func (a *Client) GetJobsByIDExecutionsByID(params *GetJobsByIDExecutionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsByIDExecutionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJobsByIDExecutionsByID",
		Method:             "GET",
		PathPattern:        "/jobs/{job_id}/executions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJobsByIDExecutionsByIDReader{formats: a.formats},
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
PostJobsByIDExecutions Executes the job with the given job ID by creating and returning a job execution for it. The job might still be
 executed when the job execution is returned. Note that this resource is also intended for running system jobs.
*/
func (a *Client) PostJobsByIDExecutions(params *PostJobsByIDExecutionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostJobsByIDExecutionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postJobsByIDExecutions",
		Method:             "POST",
		PathPattern:        "/jobs/{job_id}/executions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostJobsByIDExecutionsReader{formats: a.formats},
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