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

// NewDeleteSitesByIDPromotionsByIDParams creates a new DeleteSitesByIDPromotionsByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDPromotionsByIDParams() *DeleteSitesByIDPromotionsByIDParams {
	var ()
	return &DeleteSitesByIDPromotionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDPromotionsByIDParamsWithTimeout creates a new DeleteSitesByIDPromotionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDPromotionsByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDPromotionsByIDParams {
	var ()
	return &DeleteSitesByIDPromotionsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDPromotionsByIDParamsWithContext creates a new DeleteSitesByIDPromotionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDPromotionsByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDPromotionsByIDParams {
	var ()
	return &DeleteSitesByIDPromotionsByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDPromotionsByIDParamsWithHTTPClient creates a new DeleteSitesByIDPromotionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDPromotionsByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDPromotionsByIDParams {
	var ()
	return &DeleteSitesByIDPromotionsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDPromotionsByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID promotions by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDPromotionsByIDParams struct {

	/*ID
	  Promotion id to remove

	*/
	ID string
	/*SiteID
	  The site context.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDPromotionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDPromotionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDPromotionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) WithID(id string) *DeleteSitesByIDPromotionsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) WithSiteID(siteID string) *DeleteSitesByIDPromotionsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID promotions by ID params
func (o *DeleteSitesByIDPromotionsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDPromotionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
