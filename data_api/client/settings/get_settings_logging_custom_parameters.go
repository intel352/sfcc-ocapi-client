// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// NewGetSettingsLoggingCustomParams creates a new GetSettingsLoggingCustomParams object
// with the default values initialized.
func NewGetSettingsLoggingCustomParams() *GetSettingsLoggingCustomParams {

	return &GetSettingsLoggingCustomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsLoggingCustomParamsWithTimeout creates a new GetSettingsLoggingCustomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsLoggingCustomParamsWithTimeout(timeout time.Duration) *GetSettingsLoggingCustomParams {

	return &GetSettingsLoggingCustomParams{

		timeout: timeout,
	}
}

// NewGetSettingsLoggingCustomParamsWithContext creates a new GetSettingsLoggingCustomParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsLoggingCustomParamsWithContext(ctx context.Context) *GetSettingsLoggingCustomParams {

	return &GetSettingsLoggingCustomParams{

		Context: ctx,
	}
}

// NewGetSettingsLoggingCustomParamsWithHTTPClient creates a new GetSettingsLoggingCustomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsLoggingCustomParamsWithHTTPClient(client *http.Client) *GetSettingsLoggingCustomParams {

	return &GetSettingsLoggingCustomParams{
		HTTPClient: client,
	}
}

/*GetSettingsLoggingCustomParams contains all the parameters to send to the API endpoint
for the get settings logging custom operation typically these are written to a http.Request
*/
type GetSettingsLoggingCustomParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) WithTimeout(timeout time.Duration) *GetSettingsLoggingCustomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) WithContext(ctx context.Context) *GetSettingsLoggingCustomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) WithHTTPClient(client *http.Client) *GetSettingsLoggingCustomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings logging custom params
func (o *GetSettingsLoggingCustomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsLoggingCustomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}