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

// NewDeleteSitesByIDCartridgesByIDParams creates a new DeleteSitesByIDCartridgesByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDCartridgesByIDParams() *DeleteSitesByIDCartridgesByIDParams {
	var ()
	return &DeleteSitesByIDCartridgesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDCartridgesByIDParamsWithTimeout creates a new DeleteSitesByIDCartridgesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDCartridgesByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDCartridgesByIDParams {
	var ()
	return &DeleteSitesByIDCartridgesByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDCartridgesByIDParamsWithContext creates a new DeleteSitesByIDCartridgesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDCartridgesByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDCartridgesByIDParams {
	var ()
	return &DeleteSitesByIDCartridgesByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDCartridgesByIDParamsWithHTTPClient creates a new DeleteSitesByIDCartridgesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDCartridgesByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDCartridgesByIDParams {
	var ()
	return &DeleteSitesByIDCartridgesByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDCartridgesByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID cartridges by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDCartridgesByIDParams struct {

	/*CartridgeName
	  request body

	*/
	CartridgeName string
	/*SiteID
	  ID of the site.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDCartridgesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDCartridgesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDCartridgesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCartridgeName adds the cartridgeName to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) WithCartridgeName(cartridgeName string) *DeleteSitesByIDCartridgesByIDParams {
	o.SetCartridgeName(cartridgeName)
	return o
}

// SetCartridgeName adds the cartridgeName to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) SetCartridgeName(cartridgeName string) {
	o.CartridgeName = cartridgeName
}

// WithSiteID adds the siteID to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) WithSiteID(siteID string) *DeleteSitesByIDCartridgesByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID cartridges by ID params
func (o *DeleteSitesByIDCartridgesByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDCartridgesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cartridge_name
	if err := r.SetPathParam("cartridge_name", o.CartridgeName); err != nil {
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