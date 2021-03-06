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

// NewPutSitesByIDCouponsByIDParams creates a new PutSitesByIDCouponsByIDParams object
// with the default values initialized.
func NewPutSitesByIDCouponsByIDParams() *PutSitesByIDCouponsByIDParams {
	var ()
	return &PutSitesByIDCouponsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDCouponsByIDParamsWithTimeout creates a new PutSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDCouponsByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDCouponsByIDParams {
	var ()
	return &PutSitesByIDCouponsByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDCouponsByIDParamsWithContext creates a new PutSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDCouponsByIDParamsWithContext(ctx context.Context) *PutSitesByIDCouponsByIDParams {
	var ()
	return &PutSitesByIDCouponsByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDCouponsByIDParamsWithHTTPClient creates a new PutSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDCouponsByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDCouponsByIDParams {
	var ()
	return &PutSitesByIDCouponsByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDCouponsByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID coupons by ID operation typically these are written to a http.Request
*/
type PutSitesByIDCouponsByIDParams struct {

	/*Body*/
	Body *models.Coupon
	/*CouponID
	  The id of the coupon to create.

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

// WithTimeout adds the timeout to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDCouponsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithContext(ctx context.Context) *PutSitesByIDCouponsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDCouponsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithBody(body *models.Coupon) *PutSitesByIDCouponsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetBody(body *models.Coupon) {
	o.Body = body
}

// WithCouponID adds the couponID to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithCouponID(couponID string) *PutSitesByIDCouponsByIDParams {
	o.SetCouponID(couponID)
	return o
}

// SetCouponID adds the couponId to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetCouponID(couponID string) {
	o.CouponID = couponID
}

// WithSiteID adds the siteID to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) WithSiteID(siteID string) *PutSitesByIDCouponsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID coupons by ID params
func (o *PutSitesByIDCouponsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDCouponsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
