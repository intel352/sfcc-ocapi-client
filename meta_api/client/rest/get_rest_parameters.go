// Code generated by go-swagger; DO NOT EDIT.

package rest

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

// NewGetRestParams creates a new GetRestParams object
// with the default values initialized.
func NewGetRestParams() *GetRestParams {

	return &GetRestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestParamsWithTimeout creates a new GetRestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestParamsWithTimeout(timeout time.Duration) *GetRestParams {

	return &GetRestParams{

		timeout: timeout,
	}
}

// NewGetRestParamsWithContext creates a new GetRestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestParamsWithContext(ctx context.Context) *GetRestParams {

	return &GetRestParams{

		Context: ctx,
	}
}

// NewGetRestParamsWithHTTPClient creates a new GetRestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestParamsWithHTTPClient(client *http.Client) *GetRestParams {

	return &GetRestParams{
		HTTPClient: client,
	}
}

/*GetRestParams contains all the parameters to send to the API endpoint
for the get rest operation typically these are written to a http.Request
*/
type GetRestParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rest params
func (o *GetRestParams) WithTimeout(timeout time.Duration) *GetRestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rest params
func (o *GetRestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rest params
func (o *GetRestParams) WithContext(ctx context.Context) *GetRestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rest params
func (o *GetRestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rest params
func (o *GetRestParams) WithHTTPClient(client *http.Client) *GetRestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rest params
func (o *GetRestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
