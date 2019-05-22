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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParams creates a new PutSitesByIDSlotsByIDSlotConfigurationsByIDParams object
// with the default values initialized.
func NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParams() *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	var ()
	return &PutSitesByIDSlotsByIDSlotConfigurationsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithTimeout creates a new PutSitesByIDSlotsByIDSlotConfigurationsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	var ()
	return &PutSitesByIDSlotsByIDSlotConfigurationsByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithContext creates a new PutSitesByIDSlotsByIDSlotConfigurationsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithContext(ctx context.Context) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	var ()
	return &PutSitesByIDSlotsByIDSlotConfigurationsByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithHTTPClient creates a new PutSitesByIDSlotsByIDSlotConfigurationsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDSlotsByIDSlotConfigurationsByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	var ()
	return &PutSitesByIDSlotsByIDSlotConfigurationsByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDSlotsByIDSlotConfigurationsByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID slots by ID slot configurations by ID operation typically these are written to a http.Request
*/
type PutSitesByIDSlotsByIDSlotConfigurationsByIDParams struct {

	/*Body*/
	Body *models.SlotConfiguration
	/*ConfigurationID
	  The id of the slot configuration.

	*/
	ConfigurationID string
	/*Context*/
	Context *string
	/*SiteID
	  The id of the site for which you want to create the slot configuration.

	*/
	SiteID string
	/*SlotID
	  The id of the slot.

	*/
	SlotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithContext(ctx context.Context) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithBody(body *models.SlotConfiguration) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetBody(body *models.SlotConfiguration) {
	o.Body = body
}

// WithConfigurationID adds the configurationID to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithConfigurationID(configurationID string) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetConfigurationID(configurationID)
	return o
}

// SetConfigurationID adds the configurationId to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetConfigurationID(configurationID string) {
	o.ConfigurationID = configurationID
}

// WithContext adds the context to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithContext(context *string) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetContext(context)
	return o
}

// SetContext adds the context to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetContext(context *string) {
	o.Context = context
}

// WithSiteID adds the siteID to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithSiteID(siteID string) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSlotID adds the slotID to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WithSlotID(slotID string) *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams {
	o.SetSlotID(slotID)
	return o
}

// SetSlotID adds the slotId to the put sites by ID slots by ID slot configurations by ID params
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) SetSlotID(slotID string) {
	o.SlotID = slotID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDSlotsByIDSlotConfigurationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param configuration_id
	if err := r.SetPathParam("configuration_id", o.ConfigurationID); err != nil {
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

	// path param slot_id
	if err := r.SetPathParam("slot_id", o.SlotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}