// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewDeleteSitesByIDAbTestsByIDParams creates a new DeleteSitesByIDAbTestsByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDAbTestsByIDParams() *DeleteSitesByIDAbTestsByIDParams {
	var ()
	return &DeleteSitesByIDAbTestsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDAbTestsByIDParamsWithTimeout creates a new DeleteSitesByIDAbTestsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDAbTestsByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDAbTestsByIDParams {
	var ()
	return &DeleteSitesByIDAbTestsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDAbTestsByIDParamsWithContext creates a new DeleteSitesByIDAbTestsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDAbTestsByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDAbTestsByIDParams {
	var ()
	return &DeleteSitesByIDAbTestsByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDAbTestsByIDParamsWithHTTPClient creates a new DeleteSitesByIDAbTestsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDAbTestsByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDAbTestsByIDParams {
	var ()
	return &DeleteSitesByIDAbTestsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDAbTestsByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID ab tests by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDAbTestsByIDParams struct {

	/*ID
	  The id of the requested A/B Test.

	*/
	ID string
	/*SiteID
	  ID of the site that the A/B tests are contained within.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDAbTestsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDAbTestsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDAbTestsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) WithID(id string) *DeleteSitesByIDAbTestsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) WithSiteID(siteID string) *DeleteSitesByIDAbTestsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID ab tests by ID params
func (o *DeleteSitesByIDAbTestsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDAbTestsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
