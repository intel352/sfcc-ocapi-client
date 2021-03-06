// Code generated by go-swagger; DO NOT EDIT.

package job_execution_search

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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPostJobExecutionSearchParams creates a new PostJobExecutionSearchParams object
// with the default values initialized.
func NewPostJobExecutionSearchParams() *PostJobExecutionSearchParams {
	var ()
	return &PostJobExecutionSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostJobExecutionSearchParamsWithTimeout creates a new PostJobExecutionSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostJobExecutionSearchParamsWithTimeout(timeout time.Duration) *PostJobExecutionSearchParams {
	var ()
	return &PostJobExecutionSearchParams{

		timeout: timeout,
	}
}

// NewPostJobExecutionSearchParamsWithContext creates a new PostJobExecutionSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostJobExecutionSearchParamsWithContext(ctx context.Context) *PostJobExecutionSearchParams {
	var ()
	return &PostJobExecutionSearchParams{

		Context: ctx,
	}
}

// NewPostJobExecutionSearchParamsWithHTTPClient creates a new PostJobExecutionSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostJobExecutionSearchParamsWithHTTPClient(client *http.Client) *PostJobExecutionSearchParams {
	var ()
	return &PostJobExecutionSearchParams{
		HTTPClient: client,
	}
}

/*PostJobExecutionSearchParams contains all the parameters to send to the API endpoint
for the post job execution search operation typically these are written to a http.Request
*/
type PostJobExecutionSearchParams struct {

	/*Body*/
	Body *models.SearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post job execution search params
func (o *PostJobExecutionSearchParams) WithTimeout(timeout time.Duration) *PostJobExecutionSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post job execution search params
func (o *PostJobExecutionSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post job execution search params
func (o *PostJobExecutionSearchParams) WithContext(ctx context.Context) *PostJobExecutionSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post job execution search params
func (o *PostJobExecutionSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post job execution search params
func (o *PostJobExecutionSearchParams) WithHTTPClient(client *http.Client) *PostJobExecutionSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post job execution search params
func (o *PostJobExecutionSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post job execution search params
func (o *PostJobExecutionSearchParams) WithBody(body *models.SearchRequest) *PostJobExecutionSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post job execution search params
func (o *PostJobExecutionSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostJobExecutionSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
