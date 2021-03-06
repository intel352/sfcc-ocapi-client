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

// NewPostSitesByIDGiftCertificatesParams creates a new PostSitesByIDGiftCertificatesParams object
// with the default values initialized.
func NewPostSitesByIDGiftCertificatesParams() *PostSitesByIDGiftCertificatesParams {
	var ()
	return &PostSitesByIDGiftCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesByIDGiftCertificatesParamsWithTimeout creates a new PostSitesByIDGiftCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesByIDGiftCertificatesParamsWithTimeout(timeout time.Duration) *PostSitesByIDGiftCertificatesParams {
	var ()
	return &PostSitesByIDGiftCertificatesParams{

		timeout: timeout,
	}
}

// NewPostSitesByIDGiftCertificatesParamsWithContext creates a new PostSitesByIDGiftCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesByIDGiftCertificatesParamsWithContext(ctx context.Context) *PostSitesByIDGiftCertificatesParams {
	var ()
	return &PostSitesByIDGiftCertificatesParams{

		Context: ctx,
	}
}

// NewPostSitesByIDGiftCertificatesParamsWithHTTPClient creates a new PostSitesByIDGiftCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesByIDGiftCertificatesParamsWithHTTPClient(client *http.Client) *PostSitesByIDGiftCertificatesParams {
	var ()
	return &PostSitesByIDGiftCertificatesParams{
		HTTPClient: client,
	}
}

/*PostSitesByIDGiftCertificatesParams contains all the parameters to send to the API endpoint
for the post sites by ID gift certificates operation typically these are written to a http.Request
*/
type PostSitesByIDGiftCertificatesParams struct {

	/*Body*/
	Body *models.GiftCertificate
	/*SiteID
	  The id of the site.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) WithTimeout(timeout time.Duration) *PostSitesByIDGiftCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) WithContext(ctx context.Context) *PostSitesByIDGiftCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) WithHTTPClient(client *http.Client) *PostSitesByIDGiftCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) WithBody(body *models.GiftCertificate) *PostSitesByIDGiftCertificatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) SetBody(body *models.GiftCertificate) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) WithSiteID(siteID string) *PostSitesByIDGiftCertificatesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites by ID gift certificates params
func (o *PostSitesByIDGiftCertificatesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesByIDGiftCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
