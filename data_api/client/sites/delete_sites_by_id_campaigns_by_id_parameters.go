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

// NewDeleteSitesByIDCampaignsByIDParams creates a new DeleteSitesByIDCampaignsByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDCampaignsByIDParams() *DeleteSitesByIDCampaignsByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDCampaignsByIDParamsWithTimeout creates a new DeleteSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDCampaignsByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDCampaignsByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDCampaignsByIDParamsWithContext creates a new DeleteSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDCampaignsByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDCampaignsByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDCampaignsByIDParamsWithHTTPClient creates a new DeleteSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDCampaignsByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDCampaignsByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDCampaignsByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID campaigns by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDCampaignsByIDParams struct {

	/*CampaignID
	  A campaign id to remove

	*/
	CampaignID string
	/*SiteID
	  The site context.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDCampaignsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDCampaignsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDCampaignsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) WithCampaignID(campaignID string) *DeleteSitesByIDCampaignsByIDParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WithSiteID adds the siteID to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) WithSiteID(siteID string) *DeleteSitesByIDCampaignsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID campaigns by ID params
func (o *DeleteSitesByIDCampaignsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDCampaignsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaign_id
	if err := r.SetPathParam("campaign_id", o.CampaignID); err != nil {
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
