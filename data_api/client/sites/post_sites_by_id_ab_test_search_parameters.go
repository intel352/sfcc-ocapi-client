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

// NewPostSitesByIDAbTestSearchParams creates a new PostSitesByIDAbTestSearchParams object
// with the default values initialized.
func NewPostSitesByIDAbTestSearchParams() *PostSitesByIDAbTestSearchParams {
	var ()
	return &PostSitesByIDAbTestSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSitesByIDAbTestSearchParamsWithTimeout creates a new PostSitesByIDAbTestSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSitesByIDAbTestSearchParamsWithTimeout(timeout time.Duration) *PostSitesByIDAbTestSearchParams {
	var ()
	return &PostSitesByIDAbTestSearchParams{

		timeout: timeout,
	}
}

// NewPostSitesByIDAbTestSearchParamsWithContext creates a new PostSitesByIDAbTestSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSitesByIDAbTestSearchParamsWithContext(ctx context.Context) *PostSitesByIDAbTestSearchParams {
	var ()
	return &PostSitesByIDAbTestSearchParams{

		Context: ctx,
	}
}

// NewPostSitesByIDAbTestSearchParamsWithHTTPClient creates a new PostSitesByIDAbTestSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSitesByIDAbTestSearchParamsWithHTTPClient(client *http.Client) *PostSitesByIDAbTestSearchParams {
	var ()
	return &PostSitesByIDAbTestSearchParams{
		HTTPClient: client,
	}
}

/*PostSitesByIDAbTestSearchParams contains all the parameters to send to the API endpoint
for the post sites by ID ab test search operation typically these are written to a http.Request
*/
type PostSitesByIDAbTestSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*SiteID
	  ID of the site that the A/B tests are contained within.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) WithTimeout(timeout time.Duration) *PostSitesByIDAbTestSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) WithContext(ctx context.Context) *PostSitesByIDAbTestSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) WithHTTPClient(client *http.Client) *PostSitesByIDAbTestSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) WithBody(body *models.SearchRequest) *PostSitesByIDAbTestSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithSiteID adds the siteID to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) WithSiteID(siteID string) *PostSitesByIDAbTestSearchParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post sites by ID ab test search params
func (o *PostSitesByIDAbTestSearchParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSitesByIDAbTestSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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