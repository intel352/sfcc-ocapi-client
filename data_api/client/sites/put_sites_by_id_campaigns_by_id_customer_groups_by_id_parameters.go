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

// NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParams creates a new PutSitesByIDCampaignsByIDCustomerGroupsByIDParams object
// with the default values initialized.
func NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParams() *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDCustomerGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithTimeout creates a new PutSitesByIDCampaignsByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDCustomerGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithContext creates a new PutSitesByIDCampaignsByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithContext(ctx context.Context) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDCustomerGroupsByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithHTTPClient creates a new PutSitesByIDCampaignsByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDCampaignsByIDCustomerGroupsByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDCustomerGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDCampaignsByIDCustomerGroupsByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID campaigns by ID customer groups by ID operation typically these are written to a http.Request
*/
type PutSitesByIDCampaignsByIDCustomerGroupsByIDParams struct {

	/*CampaignID
	  The campaign ID that coupons are to be bound to

	*/
	CampaignID string
	/*CustomerGroupID
	  The customer group ID to bind to a campaign

	*/
	CustomerGroupID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithContext(ctx context.Context) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithCampaignID(campaignID string) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WithCustomerGroupID adds the customerGroupID to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithCustomerGroupID(customerGroupID string) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetCustomerGroupID(customerGroupID)
	return o
}

// SetCustomerGroupID adds the customerGroupId to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetCustomerGroupID(customerGroupID string) {
	o.CustomerGroupID = customerGroupID
}

// WithSiteID adds the siteID to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WithSiteID(siteID string) *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID campaigns by ID customer groups by ID params
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDCampaignsByIDCustomerGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaign_id
	if err := r.SetPathParam("campaign_id", o.CampaignID); err != nil {
		return err
	}

	// path param customer_group_id
	if err := r.SetPathParam("customer_group_id", o.CustomerGroupID); err != nil {
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
