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

// NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParams creates a new PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams object
// with the default values initialized.
func NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParams() *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithTimeout creates a new PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithContext creates a new PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithContext(ctx context.Context) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithHTTPClient creates a new PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDCampaignsByIDSortingRulesByIDByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID campaigns by ID sorting rules by ID by ID operation typically these are written to a http.Request
*/
type PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams struct {

	/*CampaignID
	  The ID of the campaign to which the sorting rule is to be assigned.

	*/
	CampaignID string
	/*CategoryID
	  The ID of the category that is associated with the sorting rule.

	*/
	CategoryID string
	/*RuleContext*/
	RuleContext *string
	/*SiteID
	  The ID of the site that contains the campaign, sorting rule and category.

	*/
	SiteID string
	/*SortingRuleID
	  The ID of sorting rule that is to be assigned to the campaign.

	*/
	SortingRuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithContext(ctx context.Context) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithCampaignID(campaignID string) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WithCategoryID adds the categoryID to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithCategoryID(categoryID string) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithRuleContext adds the ruleContext to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithRuleContext(ruleContext *string) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetRuleContext(ruleContext)
	return o
}

// SetRuleContext adds the ruleContext to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetRuleContext(ruleContext *string) {
	o.RuleContext = ruleContext
}

// WithSiteID adds the siteID to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithSiteID(siteID string) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSortingRuleID adds the sortingRuleID to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WithSortingRuleID(sortingRuleID string) *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams {
	o.SetSortingRuleID(sortingRuleID)
	return o
}

// SetSortingRuleID adds the sortingRuleId to the put sites by ID campaigns by ID sorting rules by ID by ID params
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) SetSortingRuleID(sortingRuleID string) {
	o.SortingRuleID = sortingRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDCampaignsByIDSortingRulesByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaign_id
	if err := r.SetPathParam("campaign_id", o.CampaignID); err != nil {
		return err
	}

	// path param category_id
	if err := r.SetPathParam("category_id", o.CategoryID); err != nil {
		return err
	}

	if o.RuleContext != nil {

		// query param rule_context
		var qrRuleContext string
		if o.RuleContext != nil {
			qrRuleContext = *o.RuleContext
		}
		qRuleContext := qrRuleContext
		if qRuleContext != "" {
			if err := r.SetQueryParam("rule_context", qRuleContext); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param sorting_rule_id
	if err := r.SetPathParam("sorting_rule_id", o.SortingRuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
