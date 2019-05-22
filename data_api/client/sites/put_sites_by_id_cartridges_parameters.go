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

// NewPutSitesByIDCartridgesParams creates a new PutSitesByIDCartridgesParams object
// with the default values initialized.
func NewPutSitesByIDCartridgesParams() *PutSitesByIDCartridgesParams {
	var ()
	return &PutSitesByIDCartridgesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDCartridgesParamsWithTimeout creates a new PutSitesByIDCartridgesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDCartridgesParamsWithTimeout(timeout time.Duration) *PutSitesByIDCartridgesParams {
	var ()
	return &PutSitesByIDCartridgesParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDCartridgesParamsWithContext creates a new PutSitesByIDCartridgesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDCartridgesParamsWithContext(ctx context.Context) *PutSitesByIDCartridgesParams {
	var ()
	return &PutSitesByIDCartridgesParams{

		Context: ctx,
	}
}

// NewPutSitesByIDCartridgesParamsWithHTTPClient creates a new PutSitesByIDCartridgesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDCartridgesParamsWithHTTPClient(client *http.Client) *PutSitesByIDCartridgesParams {
	var ()
	return &PutSitesByIDCartridgesParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDCartridgesParams contains all the parameters to send to the API endpoint
for the put sites by ID cartridges operation typically these are written to a http.Request
*/
type PutSitesByIDCartridgesParams struct {

	/*Body*/
	Body *models.CartridgePathCreateRequest
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) WithTimeout(timeout time.Duration) *PutSitesByIDCartridgesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) WithContext(ctx context.Context) *PutSitesByIDCartridgesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) WithHTTPClient(client *http.Client) *PutSitesByIDCartridgesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) WithBody(body *models.CartridgePathCreateRequest) *PutSitesByIDCartridgesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) SetBody(body *models.CartridgePathCreateRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) WithSiteID(siteID string) *PutSitesByIDCartridgesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID cartridges params
func (o *PutSitesByIDCartridgesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDCartridgesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
