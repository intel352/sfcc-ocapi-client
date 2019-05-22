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

// NewGetRestByIDByIDDocumentsParams creates a new GetRestByIDByIDDocumentsParams object
// with the default values initialized.
func NewGetRestByIDByIDDocumentsParams() *GetRestByIDByIDDocumentsParams {
	var ()
	return &GetRestByIDByIDDocumentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRestByIDByIDDocumentsParamsWithTimeout creates a new GetRestByIDByIDDocumentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRestByIDByIDDocumentsParamsWithTimeout(timeout time.Duration) *GetRestByIDByIDDocumentsParams {
	var ()
	return &GetRestByIDByIDDocumentsParams{

		timeout: timeout,
	}
}

// NewGetRestByIDByIDDocumentsParamsWithContext creates a new GetRestByIDByIDDocumentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRestByIDByIDDocumentsParamsWithContext(ctx context.Context) *GetRestByIDByIDDocumentsParams {
	var ()
	return &GetRestByIDByIDDocumentsParams{

		Context: ctx,
	}
}

// NewGetRestByIDByIDDocumentsParamsWithHTTPClient creates a new GetRestByIDByIDDocumentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRestByIDByIDDocumentsParamsWithHTTPClient(client *http.Client) *GetRestByIDByIDDocumentsParams {
	var ()
	return &GetRestByIDByIDDocumentsParams{
		HTTPClient: client,
	}
}

/*GetRestByIDByIDDocumentsParams contains all the parameters to send to the API endpoint
for the get rest by ID by ID documents operation typically these are written to a http.Request
*/
type GetRestByIDByIDDocumentsParams struct {

	/*APIName
	  the name of the API

	*/
	APIName string
	/*Version
	  the version to get the description for

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) WithTimeout(timeout time.Duration) *GetRestByIDByIDDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) WithContext(ctx context.Context) *GetRestByIDByIDDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) WithHTTPClient(client *http.Client) *GetRestByIDByIDDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIName adds the aPIName to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) WithAPIName(aPIName string) *GetRestByIDByIDDocumentsParams {
	o.SetAPIName(aPIName)
	return o
}

// SetAPIName adds the apiName to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) SetAPIName(aPIName string) {
	o.APIName = aPIName
}

// WithVersion adds the version to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) WithVersion(version string) *GetRestByIDByIDDocumentsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get rest by ID by ID documents params
func (o *GetRestByIDByIDDocumentsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetRestByIDByIDDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api_name
	if err := r.SetPathParam("api_name", o.APIName); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
