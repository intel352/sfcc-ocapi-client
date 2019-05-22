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

// NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams creates a new DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams() *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithTimeout creates a new DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithContext creates a new DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithHTTPClient creates a new DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID campaigns by ID slot configurations by ID by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams struct {

	/*CampaignID
	  The id of the campaign

	*/
	CampaignID string
	/*Context*/
	Context *string
	/*SiteID
	  The id of the site

	*/
	SiteID string
	/*SlotConfigID
	  The id of the slot configuration

	*/
	SlotConfigID string
	/*SlotID
	  The of the slot

	*/
	SlotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithCampaignID(campaignID string) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WithContext adds the context to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithContext(context *string) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetContext(context)
	return o
}

// SetContext adds the context to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetContext(context *string) {
	o.Context = context
}

// WithSiteID adds the siteID to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithSiteID(siteID string) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSlotConfigID adds the slotConfigID to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithSlotConfigID(slotConfigID string) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetSlotConfigID(slotConfigID)
	return o
}

// SetSlotConfigID adds the slotConfigId to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetSlotConfigID(slotConfigID string) {
	o.SlotConfigID = slotConfigID
}

// WithSlotID adds the slotID to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WithSlotID(slotID string) *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams {
	o.SetSlotID(slotID)
	return o
}

// SetSlotID adds the slotId to the delete sites by ID campaigns by ID slot configurations by ID by ID params
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) SetSlotID(slotID string) {
	o.SlotID = slotID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDCampaignsByIDSlotConfigurationsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaign_id
	if err := r.SetPathParam("campaign_id", o.CampaignID); err != nil {
		return err
	}

	if o.Context != nil {

		// query param context
		var qrContext string
		if o.Context != nil {
			qrContext = *o.Context
		}
		qContext := qrContext
		if qContext != "" {
			if err := r.SetQueryParam("context", qContext); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param slot_config_id
	if err := r.SetPathParam("slot_config_id", o.SlotConfigID); err != nil {
		return err
	}

	// path param slot_id
	if err := r.SetPathParam("slot_id", o.SlotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
