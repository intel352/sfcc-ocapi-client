// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetJobsByIDExecutionsByIDParams creates a new GetJobsByIDExecutionsByIDParams object
// with the default values initialized.
func NewGetJobsByIDExecutionsByIDParams() *GetJobsByIDExecutionsByIDParams {
	var ()
	return &GetJobsByIDExecutionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobsByIDExecutionsByIDParamsWithTimeout creates a new GetJobsByIDExecutionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobsByIDExecutionsByIDParamsWithTimeout(timeout time.Duration) *GetJobsByIDExecutionsByIDParams {
	var ()
	return &GetJobsByIDExecutionsByIDParams{

		timeout: timeout,
	}
}

// NewGetJobsByIDExecutionsByIDParamsWithContext creates a new GetJobsByIDExecutionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobsByIDExecutionsByIDParamsWithContext(ctx context.Context) *GetJobsByIDExecutionsByIDParams {
	var ()
	return &GetJobsByIDExecutionsByIDParams{

		Context: ctx,
	}
}

// NewGetJobsByIDExecutionsByIDParamsWithHTTPClient creates a new GetJobsByIDExecutionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobsByIDExecutionsByIDParamsWithHTTPClient(client *http.Client) *GetJobsByIDExecutionsByIDParams {
	var ()
	return &GetJobsByIDExecutionsByIDParams{
		HTTPClient: client,
	}
}

/*GetJobsByIDExecutionsByIDParams contains all the parameters to send to the API endpoint
for the get jobs by ID executions by ID operation typically these are written to a http.Request
*/
type GetJobsByIDExecutionsByIDParams struct {

	/*ID
	  the ID of the job execution

	*/
	ID string
	/*JobID
	  the ID of the job.

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) WithTimeout(timeout time.Duration) *GetJobsByIDExecutionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) WithContext(ctx context.Context) *GetJobsByIDExecutionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) WithHTTPClient(client *http.Client) *GetJobsByIDExecutionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) WithID(id string) *GetJobsByIDExecutionsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) SetID(id string) {
	o.ID = id
}

// WithJobID adds the jobID to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) WithJobID(jobID string) *GetJobsByIDExecutionsByIDParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the get jobs by ID executions by ID params
func (o *GetJobsByIDExecutionsByIDParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobsByIDExecutionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}