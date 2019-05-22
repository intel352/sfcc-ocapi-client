// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPermissionsBmParams creates a new GetPermissionsBmParams object
// with the default values initialized.
func NewGetPermissionsBmParams() *GetPermissionsBmParams {
	var ()
	return &GetPermissionsBmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPermissionsBmParamsWithTimeout creates a new GetPermissionsBmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPermissionsBmParamsWithTimeout(timeout time.Duration) *GetPermissionsBmParams {
	var ()
	return &GetPermissionsBmParams{

		timeout: timeout,
	}
}

// NewGetPermissionsBmParamsWithContext creates a new GetPermissionsBmParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPermissionsBmParamsWithContext(ctx context.Context) *GetPermissionsBmParams {
	var ()
	return &GetPermissionsBmParams{

		Context: ctx,
	}
}

// NewGetPermissionsBmParamsWithHTTPClient creates a new GetPermissionsBmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPermissionsBmParamsWithHTTPClient(client *http.Client) *GetPermissionsBmParams {
	var ()
	return &GetPermissionsBmParams{
		HTTPClient: client,
	}
}

/*GetPermissionsBmParams contains all the parameters to send to the API endpoint
for the get permissions bm operation typically these are written to a http.Request
*/
type GetPermissionsBmParams struct {

	/*Expand
	 The permission expand. A comma separated list with the allowed values
	(module, functional, webdav, locale).

	*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get permissions bm params
func (o *GetPermissionsBmParams) WithTimeout(timeout time.Duration) *GetPermissionsBmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get permissions bm params
func (o *GetPermissionsBmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get permissions bm params
func (o *GetPermissionsBmParams) WithContext(ctx context.Context) *GetPermissionsBmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get permissions bm params
func (o *GetPermissionsBmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get permissions bm params
func (o *GetPermissionsBmParams) WithHTTPClient(client *http.Client) *GetPermissionsBmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get permissions bm params
func (o *GetPermissionsBmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get permissions bm params
func (o *GetPermissionsBmParams) WithExpand(expand []string) *GetPermissionsBmParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get permissions bm params
func (o *GetPermissionsBmParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetPermissionsBmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}