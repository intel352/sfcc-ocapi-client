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

// NewPostSitesByIDSourceCodeGroupSearchParams creates a new PostSitesByIDSourceCodeGroupSearchParams object
// with the default values initialized.
func NewPostSitesByIDSourceCodeGroupSearchParams() *PostSitesByIDSourceCodeGroupSearchParams {
	var ()
	return &PostSitesByIDSourceCodeGroupSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesByIDSourceCodeGroupSearchParamsWithTimeout creates a new PostSitesByIDSourceCodeGroupSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesByIDSourceCodeGroupSearchParamsWithTimeout(timeout time.Duration) *PostSitesByIDSourceCodeGroupSearchParams {
	var ()
	return &PostSitesByIDSourceCodeGroupSearchParams{

		timeout: timeout,
	}
}

// NewPostSitesByIDSourceCodeGroupSearchParamsWithContext creates a new PostSitesByIDSourceCodeGroupSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesByIDSourceCodeGroupSearchParamsWithContext(ctx context.Context) *PostSitesByIDSourceCodeGroupSearchParams {
	var ()
	return &PostSitesByIDSourceCodeGroupSearchParams{

		Context: ctx,
	}
}

// NewPostSitesByIDSourceCodeGroupSearchParamsWithHTTPClient creates a new PostSitesByIDSourceCodeGroupSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesByIDSourceCodeGroupSearchParamsWithHTTPClient(client *http.Client) *PostSitesByIDSourceCodeGroupSearchParams {
	var ()
	return &PostSitesByIDSourceCodeGroupSearchParams{
		HTTPClient: client,
	}
}

/*PostSitesByIDSourceCodeGroupSearchParams contains all the parameters to send to the API endpoint
for the post sites by ID source code group search operation typically these are written to a http.Request
*/
type PostSitesByIDSourceCodeGroupSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*SiteID
	  The id of the site.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) WithTimeout(timeout time.Duration) *PostSitesByIDSourceCodeGroupSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) WithContext(ctx context.Context) *PostSitesByIDSourceCodeGroupSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) WithHTTPClient(client *http.Client) *PostSitesByIDSourceCodeGroupSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) WithBody(body *models.SearchRequest) *PostSitesByIDSourceCodeGroupSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) WithSiteID(siteID string) *PostSitesByIDSourceCodeGroupSearchParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites by ID source code group search params
func (o *PostSitesByIDSourceCodeGroupSearchParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesByIDSourceCodeGroupSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
