// Code generated by go-swagger; DO NOT EDIT.

package code_versions

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

// NewGetCodeVersionsParams creates a new GetCodeVersionsParams object
// with the default values initialized.
func NewGetCodeVersionsParams() *GetCodeVersionsParams {

	return &GetCodeVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCodeVersionsParamsWithTimeout creates a new GetCodeVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCodeVersionsParamsWithTimeout(timeout time.Duration) *GetCodeVersionsParams {

	return &GetCodeVersionsParams{

		timeout: timeout,
	}
}

// NewGetCodeVersionsParamsWithContext creates a new GetCodeVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCodeVersionsParamsWithContext(ctx context.Context) *GetCodeVersionsParams {

	return &GetCodeVersionsParams{

		Context: ctx,
	}
}

// NewGetCodeVersionsParamsWithHTTPClient creates a new GetCodeVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCodeVersionsParamsWithHTTPClient(client *http.Client) *GetCodeVersionsParams {

	return &GetCodeVersionsParams{
		HTTPClient: client,
	}
}

/*GetCodeVersionsParams contains all the parameters to send to the API endpoint
for the get code versions operation typically these are written to a http.Request
*/
type GetCodeVersionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get code versions params
func (o *GetCodeVersionsParams) WithTimeout(timeout time.Duration) *GetCodeVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get code versions params
func (o *GetCodeVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get code versions params
func (o *GetCodeVersionsParams) WithContext(ctx context.Context) *GetCodeVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get code versions params
func (o *GetCodeVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get code versions params
func (o *GetCodeVersionsParams) WithHTTPClient(client *http.Client) *GetCodeVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get code versions params
func (o *GetCodeVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCodeVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
