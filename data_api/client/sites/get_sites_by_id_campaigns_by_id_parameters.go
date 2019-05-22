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

// NewGetSitesByIDCampaignsByIDParams creates a new GetSitesByIDCampaignsByIDParams object
// with the default values initialized.
func NewGetSitesByIDCampaignsByIDParams() *GetSitesByIDCampaignsByIDParams {
	var ()
	return &GetSitesByIDCampaignsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDCampaignsByIDParamsWithTimeout creates a new GetSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDCampaignsByIDParamsWithTimeout(timeout time.Duration) *GetSitesByIDCampaignsByIDParams {
	var ()
	return &GetSitesByIDCampaignsByIDParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDCampaignsByIDParamsWithContext creates a new GetSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDCampaignsByIDParamsWithContext(ctx context.Context) *GetSitesByIDCampaignsByIDParams {
	var ()
	return &GetSitesByIDCampaignsByIDParams{

		Context: ctx,
	}
}

// NewGetSitesByIDCampaignsByIDParamsWithHTTPClient creates a new GetSitesByIDCampaignsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDCampaignsByIDParamsWithHTTPClient(client *http.Client) *GetSitesByIDCampaignsByIDParams {
	var ()
	return &GetSitesByIDCampaignsByIDParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDCampaignsByIDParams contains all the parameters to send to the API endpoint
for the get sites by ID campaigns by ID operation typically these are written to a http.Request
*/
type GetSitesByIDCampaignsByIDParams struct {

	/*CampaignID
	  The id of the requested campaign.

	*/
	CampaignID string
	/*SiteID
	  The site the requested campaign belongs to.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) WithTimeout(timeout time.Duration) *GetSitesByIDCampaignsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) WithContext(ctx context.Context) *GetSitesByIDCampaignsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) WithHTTPClient(client *http.Client) *GetSitesByIDCampaignsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) WithCampaignID(campaignID string) *GetSitesByIDCampaignsByIDParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WithSiteID adds the siteID to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) WithSiteID(siteID string) *GetSitesByIDCampaignsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID campaigns by ID params
func (o *GetSitesByIDCampaignsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDCampaignsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
