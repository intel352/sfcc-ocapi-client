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

// NewDeleteSitesByIDCouponsByIDParams creates a new DeleteSitesByIDCouponsByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDCouponsByIDParams() *DeleteSitesByIDCouponsByIDParams {
	var ()
	return &DeleteSitesByIDCouponsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDCouponsByIDParamsWithTimeout creates a new DeleteSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDCouponsByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDCouponsByIDParams {
	var ()
	return &DeleteSitesByIDCouponsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDCouponsByIDParamsWithContext creates a new DeleteSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDCouponsByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDCouponsByIDParams {
	var ()
	return &DeleteSitesByIDCouponsByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDCouponsByIDParamsWithHTTPClient creates a new DeleteSitesByIDCouponsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDCouponsByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDCouponsByIDParams {
	var ()
	return &DeleteSitesByIDCouponsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDCouponsByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID coupons by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDCouponsByIDParams struct {

	/*CouponID
	  Id of the coupon to delete from the site.

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

// WithTimeout adds the timeout to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDCouponsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDCouponsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDCouponsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCouponID adds the couponID to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) WithCouponID(couponID string) *DeleteSitesByIDCouponsByIDParams {
	o.SetCouponID(couponID)
	return o
}

// SetCouponID adds the couponId to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) SetCouponID(couponID string) {
	o.CouponID = couponID
}

// WithSiteID adds the siteID to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) WithSiteID(siteID string) *DeleteSitesByIDCouponsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID coupons by ID params
func (o *DeleteSitesByIDCouponsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDCouponsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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