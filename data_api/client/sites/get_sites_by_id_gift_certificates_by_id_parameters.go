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

// NewGetSitesByIDGiftCertificatesByIDParams creates a new GetSitesByIDGiftCertificatesByIDParams object
// with the default values initialized.
func NewGetSitesByIDGiftCertificatesByIDParams() *GetSitesByIDGiftCertificatesByIDParams {
	var ()
	return &GetSitesByIDGiftCertificatesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDGiftCertificatesByIDParamsWithTimeout creates a new GetSitesByIDGiftCertificatesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDGiftCertificatesByIDParamsWithTimeout(timeout time.Duration) *GetSitesByIDGiftCertificatesByIDParams {
	var ()
	return &GetSitesByIDGiftCertificatesByIDParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDGiftCertificatesByIDParamsWithContext creates a new GetSitesByIDGiftCertificatesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDGiftCertificatesByIDParamsWithContext(ctx context.Context) *GetSitesByIDGiftCertificatesByIDParams {
	var ()
	return &GetSitesByIDGiftCertificatesByIDParams{

		Context: ctx,
	}
}

// NewGetSitesByIDGiftCertificatesByIDParamsWithHTTPClient creates a new GetSitesByIDGiftCertificatesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDGiftCertificatesByIDParamsWithHTTPClient(client *http.Client) *GetSitesByIDGiftCertificatesByIDParams {
	var ()
	return &GetSitesByIDGiftCertificatesByIDParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDGiftCertificatesByIDParams contains all the parameters to send to the API endpoint
for the get sites by ID gift certificates by ID operation typically these are written to a http.Request
*/
type GetSitesByIDGiftCertificatesByIDParams struct {

	/*MerchantID
	  The merchant id of the requested gift certificate.

	*/
	MerchantID string
	/*SiteID
	  The id of the site.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) WithTimeout(timeout time.Duration) *GetSitesByIDGiftCertificatesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) WithContext(ctx context.Context) *GetSitesByIDGiftCertificatesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) WithHTTPClient(client *http.Client) *GetSitesByIDGiftCertificatesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMerchantID adds the merchantID to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) WithMerchantID(merchantID string) *GetSitesByIDGiftCertificatesByIDParams {
	o.SetMerchantID(merchantID)
	return o
}

// SetMerchantID adds the merchantId to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) SetMerchantID(merchantID string) {
	o.MerchantID = merchantID
}

// WithSiteID adds the siteID to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) WithSiteID(siteID string) *GetSitesByIDGiftCertificatesByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID gift certificates by ID params
func (o *GetSitesByIDGiftCertificatesByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDGiftCertificatesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param merchant_id
	if err := r.SetPathParam("merchant_id", o.MerchantID); err != nil {
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