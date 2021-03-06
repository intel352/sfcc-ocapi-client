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

// NewPostSitesByIDStoreSearchParams creates a new PostSitesByIDStoreSearchParams object
// with the default values initialized.
func NewPostSitesByIDStoreSearchParams() *PostSitesByIDStoreSearchParams {
	var ()
	return &PostSitesByIDStoreSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesByIDStoreSearchParamsWithTimeout creates a new PostSitesByIDStoreSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesByIDStoreSearchParamsWithTimeout(timeout time.Duration) *PostSitesByIDStoreSearchParams {
	var ()
	return &PostSitesByIDStoreSearchParams{

		timeout: timeout,
	}
}

// NewPostSitesByIDStoreSearchParamsWithContext creates a new PostSitesByIDStoreSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesByIDStoreSearchParamsWithContext(ctx context.Context) *PostSitesByIDStoreSearchParams {
	var ()
	return &PostSitesByIDStoreSearchParams{

		Context: ctx,
	}
}

// NewPostSitesByIDStoreSearchParamsWithHTTPClient creates a new PostSitesByIDStoreSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesByIDStoreSearchParamsWithHTTPClient(client *http.Client) *PostSitesByIDStoreSearchParams {
	var ()
	return &PostSitesByIDStoreSearchParams{
		HTTPClient: client,
	}
}

/*PostSitesByIDStoreSearchParams contains all the parameters to send to the API endpoint
for the post sites by ID store search operation typically these are written to a http.Request
*/
type PostSitesByIDStoreSearchParams struct {

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

// WithTimeout adds the timeout to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) WithTimeout(timeout time.Duration) *PostSitesByIDStoreSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) WithContext(ctx context.Context) *PostSitesByIDStoreSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) WithHTTPClient(client *http.Client) *PostSitesByIDStoreSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) WithBody(body *models.SearchRequest) *PostSitesByIDStoreSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) WithSiteID(siteID string) *PostSitesByIDStoreSearchParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites by ID store search params
func (o *PostSitesByIDStoreSearchParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesByIDStoreSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
