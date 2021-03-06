// Code generated by go-swagger; DO NOT EDIT.

package ocapi_configs

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

// NewDeleteOcapiConfigsByIDParams creates a new DeleteOcapiConfigsByIDParams object
// with the default values initialized.
func NewDeleteOcapiConfigsByIDParams() *DeleteOcapiConfigsByIDParams {
	var ()
	return &DeleteOcapiConfigsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOcapiConfigsByIDParamsWithTimeout creates a new DeleteOcapiConfigsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOcapiConfigsByIDParamsWithTimeout(timeout time.Duration) *DeleteOcapiConfigsByIDParams {
	var ()
	return &DeleteOcapiConfigsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteOcapiConfigsByIDParamsWithContext creates a new DeleteOcapiConfigsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOcapiConfigsByIDParamsWithContext(ctx context.Context) *DeleteOcapiConfigsByIDParams {
	var ()
	return &DeleteOcapiConfigsByIDParams{

		Context: ctx,
	}
}

// NewDeleteOcapiConfigsByIDParamsWithHTTPClient creates a new DeleteOcapiConfigsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOcapiConfigsByIDParamsWithHTTPClient(client *http.Client) *DeleteOcapiConfigsByIDParams {
	var ()
	return &DeleteOcapiConfigsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteOcapiConfigsByIDParams contains all the parameters to send to the API endpoint
for the delete ocapi configs by ID operation typically these are written to a http.Request
*/
type DeleteOcapiConfigsByIDParams struct {

	/*ClientID
	  Id of the client to be deleted.

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) WithTimeout(timeout time.Duration) *DeleteOcapiConfigsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) WithContext(ctx context.Context) *DeleteOcapiConfigsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) WithHTTPClient(client *http.Client) *DeleteOcapiConfigsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) WithClientID(clientID string) *DeleteOcapiConfigsByIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the delete ocapi configs by ID params
func (o *DeleteOcapiConfigsByIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOcapiConfigsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientId
	if err := r.SetPathParam("clientId", o.ClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
