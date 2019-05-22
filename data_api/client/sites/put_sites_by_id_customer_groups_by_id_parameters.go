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

// NewPutSitesByIDCustomerGroupsByIDParams creates a new PutSitesByIDCustomerGroupsByIDParams object
// with the default values initialized.
func NewPutSitesByIDCustomerGroupsByIDParams() *PutSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCustomerGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDCustomerGroupsByIDParamsWithTimeout creates a new PutSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDCustomerGroupsByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCustomerGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDCustomerGroupsByIDParamsWithContext creates a new PutSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDCustomerGroupsByIDParamsWithContext(ctx context.Context) *PutSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCustomerGroupsByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDCustomerGroupsByIDParamsWithHTTPClient creates a new PutSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDCustomerGroupsByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PutSitesByIDCustomerGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDCustomerGroupsByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID customer groups by ID operation typically these are written to a http.Request
*/
type PutSitesByIDCustomerGroupsByIDParams struct {

	/*Body*/
	Body *models.CustomerGroup
	/*ID
	  The id of the customer group to create.

	*/
	ID string
	/*SiteID
	  The id of the site.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithContext(ctx context.Context) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithBody(body *models.CustomerGroup) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetBody(body *models.CustomerGroup) {
	o.Body = body
}

// WithID adds the id to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithID(id string) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) WithSiteID(siteID string) *PutSitesByIDCustomerGroupsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID customer groups by ID params
func (o *PutSitesByIDCustomerGroupsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDCustomerGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
