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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSitesByIDCouponsByIDCampaignsParams creates a new GetSitesByIDCouponsByIDCampaignsParams object
// with the default values initialized.
func NewGetSitesByIDCouponsByIDCampaignsParams() *GetSitesByIDCouponsByIDCampaignsParams {
	var ()
	return &GetSitesByIDCouponsByIDCampaignsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDCouponsByIDCampaignsParamsWithTimeout creates a new GetSitesByIDCouponsByIDCampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDCouponsByIDCampaignsParamsWithTimeout(timeout time.Duration) *GetSitesByIDCouponsByIDCampaignsParams {
	var ()
	return &GetSitesByIDCouponsByIDCampaignsParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDCouponsByIDCampaignsParamsWithContext creates a new GetSitesByIDCouponsByIDCampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDCouponsByIDCampaignsParamsWithContext(ctx context.Context) *GetSitesByIDCouponsByIDCampaignsParams {
	var ()
	return &GetSitesByIDCouponsByIDCampaignsParams{

		Context: ctx,
	}
}

// NewGetSitesByIDCouponsByIDCampaignsParamsWithHTTPClient creates a new GetSitesByIDCouponsByIDCampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDCouponsByIDCampaignsParamsWithHTTPClient(client *http.Client) *GetSitesByIDCouponsByIDCampaignsParams {
	var ()
	return &GetSitesByIDCouponsByIDCampaignsParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDCouponsByIDCampaignsParams contains all the parameters to send to the API endpoint
for the get sites by ID coupons by ID campaigns operation typically these are written to a http.Request
*/
type GetSitesByIDCouponsByIDCampaignsParams struct {

	/*Count*/
	Count *int32
	/*CouponID
	  The id of the coupon that is assigned to campaigns directly or through promotions.

	*/
	CouponID string
	/*Select*/
	Select *string
	/*SiteID
	  The site context.

	*/
	SiteID string
	/*Start*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithTimeout(timeout time.Duration) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithContext(ctx context.Context) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithHTTPClient(client *http.Client) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithCount(count *int32) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetCount(count *int32) {
	o.Count = count
}

// WithCouponID adds the couponID to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithCouponID(couponID string) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetCouponID(couponID)
	return o
}

// SetCouponID adds the couponId to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetCouponID(couponID string) {
	o.CouponID = couponID
}

// WithSelect adds the selectVar to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithSelect(selectVar *string) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithSiteID adds the siteID to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithSiteID(siteID string) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStart adds the start to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) WithStart(start *int32) *GetSitesByIDCouponsByIDCampaignsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get sites by ID coupons by ID campaigns params
func (o *GetSitesByIDCouponsByIDCampaignsParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDCouponsByIDCampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	// path param coupon_id
	if err := r.SetPathParam("coupon_id", o.CouponID); err != nil {
		return err
	}

	if o.Select != nil {

		// query param select
		var qrSelect string
		if o.Select != nil {
			qrSelect = *o.Select
		}
		qSelect := qrSelect
		if qSelect != "" {
			if err := r.SetQueryParam("select", qSelect); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}