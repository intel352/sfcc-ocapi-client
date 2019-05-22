// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewPostProductsByIDVariantSearchParams creates a new PostProductsByIDVariantSearchParams object
// with the default values initialized.
func NewPostProductsByIDVariantSearchParams() *PostProductsByIDVariantSearchParams {
	var ()
	return &PostProductsByIDVariantSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProductsByIDVariantSearchParamsWithTimeout creates a new PostProductsByIDVariantSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProductsByIDVariantSearchParamsWithTimeout(timeout time.Duration) *PostProductsByIDVariantSearchParams {
	var ()
	return &PostProductsByIDVariantSearchParams{

		timeout: timeout,
	}
}

// NewPostProductsByIDVariantSearchParamsWithContext creates a new PostProductsByIDVariantSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProductsByIDVariantSearchParamsWithContext(ctx context.Context) *PostProductsByIDVariantSearchParams {
	var ()
	return &PostProductsByIDVariantSearchParams{

		Context: ctx,
	}
}

// NewPostProductsByIDVariantSearchParamsWithHTTPClient creates a new PostProductsByIDVariantSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProductsByIDVariantSearchParamsWithHTTPClient(client *http.Client) *PostProductsByIDVariantSearchParams {
	var ()
	return &PostProductsByIDVariantSearchParams{
		HTTPClient: client,
	}
}

/*PostProductsByIDVariantSearchParams contains all the parameters to send to the API endpoint
for the post products by ID variant search operation typically these are written to a http.Request
*/
type PostProductsByIDVariantSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*ID
	  The product id of master product or variation group product

	*/
	ID string
	/*SiteID*/
	SiteID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithTimeout(timeout time.Duration) *PostProductsByIDVariantSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithContext(ctx context.Context) *PostProductsByIDVariantSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithHTTPClient(client *http.Client) *PostProductsByIDVariantSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithBody(body *models.SearchRequest) *PostProductsByIDVariantSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithID adds the id to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithID(id string) *PostProductsByIDVariantSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) WithSiteID(siteID *string) *PostProductsByIDVariantSearchParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the post products by ID variant search params
func (o *PostProductsByIDVariantSearchParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PostProductsByIDVariantSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
