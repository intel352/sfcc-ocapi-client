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

// NewPostSitesByIDCouponSearchParams creates a new PostSitesByIDCouponSearchParams object
// with the default values initialized.
func NewPostSitesByIDCouponSearchParams() *PostSitesByIDCouponSearchParams {
	var ()
	return &PostSitesByIDCouponSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesByIDCouponSearchParamsWithTimeout creates a new PostSitesByIDCouponSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesByIDCouponSearchParamsWithTimeout(timeout time.Duration) *PostSitesByIDCouponSearchParams {
	var ()
	return &PostSitesByIDCouponSearchParams{

		timeout: timeout,
	}
}

// NewPostSitesByIDCouponSearchParamsWithContext creates a new PostSitesByIDCouponSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesByIDCouponSearchParamsWithContext(ctx context.Context) *PostSitesByIDCouponSearchParams {
	var ()
	return &PostSitesByIDCouponSearchParams{

		Context: ctx,
	}
}

// NewPostSitesByIDCouponSearchParamsWithHTTPClient creates a new PostSitesByIDCouponSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesByIDCouponSearchParamsWithHTTPClient(client *http.Client) *PostSitesByIDCouponSearchParams {
	var ()
	return &PostSitesByIDCouponSearchParams{
		HTTPClient: client,
	}
}

/*PostSitesByIDCouponSearchParams contains all the parameters to send to the API endpoint
for the post sites by ID coupon search operation typically these are written to a http.Request
*/
type PostSitesByIDCouponSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*SiteID
	  The site context.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) WithTimeout(timeout time.Duration) *PostSitesByIDCouponSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) WithContext(ctx context.Context) *PostSitesByIDCouponSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) WithHTTPClient(client *http.Client) *PostSitesByIDCouponSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) WithBody(body *models.SearchRequest) *PostSitesByIDCouponSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) WithSiteID(siteID string) *PostSitesByIDCouponSearchParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites by ID coupon search params
func (o *PostSitesByIDCouponSearchParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesByIDCouponSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
