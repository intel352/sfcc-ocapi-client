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

// NewPatchSitesByIDCustomerGroupsByIDParams creates a new PatchSitesByIDCustomerGroupsByIDParams object
// with the default values initialized.
func NewPatchSitesByIDCustomerGroupsByIDParams() *PatchSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PatchSitesByIDCustomerGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSitesByIDCustomerGroupsByIDParamsWithTimeout creates a new PatchSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSitesByIDCustomerGroupsByIDParamsWithTimeout(timeout time.Duration) *PatchSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PatchSitesByIDCustomerGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPatchSitesByIDCustomerGroupsByIDParamsWithContext creates a new PatchSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSitesByIDCustomerGroupsByIDParamsWithContext(ctx context.Context) *PatchSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PatchSitesByIDCustomerGroupsByIDParams{

		Context: ctx,
	}
}

// NewPatchSitesByIDCustomerGroupsByIDParamsWithHTTPClient creates a new PatchSitesByIDCustomerGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSitesByIDCustomerGroupsByIDParamsWithHTTPClient(client *http.Client) *PatchSitesByIDCustomerGroupsByIDParams {
	var ()
	return &PatchSitesByIDCustomerGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PatchSitesByIDCustomerGroupsByIDParams contains all the parameters to send to the API endpoint
for the patch sites by ID customer groups by ID operation typically these are written to a http.Request
*/
type PatchSitesByIDCustomerGroupsByIDParams struct {

	/*Body*/
	Body *models.CustomerGroup
	/*ID
	  The id of the requested customer group.

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

// WithTimeout adds the timeout to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithTimeout(timeout time.Duration) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithContext(ctx context.Context) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithHTTPClient(client *http.Client) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithBody(body *models.CustomerGroup) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetBody(body *models.CustomerGroup) {
	o.Body = body
}

// WithID adds the id to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithID(id string) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) WithSiteID(siteID string) *PatchSitesByIDCustomerGroupsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the patch sites by ID customer groups by ID params
func (o *PatchSitesByIDCustomerGroupsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSitesByIDCustomerGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
