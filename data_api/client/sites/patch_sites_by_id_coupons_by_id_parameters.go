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

// NewPatchSitesByIDCouponsByIDParams creates a new PatchSitesByIDCouponsByIDParams object
// with the default values initialized.
func NewPatchSitesByIDCouponsByIDParams() *PatchSitesByIDCouponsByIDParams {
	var ()
	return &PatchSitesByIDCouponsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSitesByIDCouponsByIDParamsWithTimeout creates a new PatchSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSitesByIDCouponsByIDParamsWithTimeout(timeout time.Duration) *PatchSitesByIDCouponsByIDParams {
	var ()
	return &PatchSitesByIDCouponsByIDParams{

		timeout: timeout,
	}
}

// NewPatchSitesByIDCouponsByIDParamsWithContext creates a new PatchSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSitesByIDCouponsByIDParamsWithContext(ctx context.Context) *PatchSitesByIDCouponsByIDParams {
	var ()
	return &PatchSitesByIDCouponsByIDParams{

		Context: ctx,
	}
}

// NewPatchSitesByIDCouponsByIDParamsWithHTTPClient creates a new PatchSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSitesByIDCouponsByIDParamsWithHTTPClient(client *http.Client) *PatchSitesByIDCouponsByIDParams {
	var ()
	return &PatchSitesByIDCouponsByIDParams{
		HTTPClient: client,
	}
}

/*PatchSitesByIDCouponsByIDParams contains all the parameters to send to the API endpoint
for the patch sites by ID coupons by ID operation typically these are written to a http.Request
*/
type PatchSitesByIDCouponsByIDParams struct {

	/*Body*/
	Body *models.Coupon
	/*CouponID
	  The id of the requested coupon.

	*/
	CouponID string
	/*SiteID
	  The site context.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithTimeout(timeout time.Duration) *PatchSitesByIDCouponsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithContext(ctx context.Context) *PatchSitesByIDCouponsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithHTTPClient(client *http.Client) *PatchSitesByIDCouponsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithBody(body *models.Coupon) *PatchSitesByIDCouponsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetBody(body *models.Coupon) {
	o.Body = body
}

// WithCouponID adds the couponID to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithCouponID(couponID string) *PatchSitesByIDCouponsByIDParams {
	o.SetCouponID(couponID)
	return o
}

// SetCouponID adds the couponId to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetCouponID(couponID string) {
	o.CouponID = couponID
}

// WithSiteID adds the siteID to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) WithSiteID(siteID string) *PatchSitesByIDCouponsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the patch sites by ID coupons by ID params
func (o *PatchSitesByIDCouponsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSitesByIDCouponsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param coupon_id
	if err := r.SetPathParam("coupon_id", o.CouponID); err != nil {
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