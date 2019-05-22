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

// NewGetSitesByIDAiRecommenderNamesParams creates a new GetSitesByIDAiRecommenderNamesParams object
// with the default values initialized.
func NewGetSitesByIDAiRecommenderNamesParams() *GetSitesByIDAiRecommenderNamesParams {
	var ()
	return &GetSitesByIDAiRecommenderNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDAiRecommenderNamesParamsWithTimeout creates a new GetSitesByIDAiRecommenderNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDAiRecommenderNamesParamsWithTimeout(timeout time.Duration) *GetSitesByIDAiRecommenderNamesParams {
	var ()
	return &GetSitesByIDAiRecommenderNamesParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDAiRecommenderNamesParamsWithContext creates a new GetSitesByIDAiRecommenderNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDAiRecommenderNamesParamsWithContext(ctx context.Context) *GetSitesByIDAiRecommenderNamesParams {
	var ()
	return &GetSitesByIDAiRecommenderNamesParams{

		Context: ctx,
	}
}

// NewGetSitesByIDAiRecommenderNamesParamsWithHTTPClient creates a new GetSitesByIDAiRecommenderNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDAiRecommenderNamesParamsWithHTTPClient(client *http.Client) *GetSitesByIDAiRecommenderNamesParams {
	var ()
	return &GetSitesByIDAiRecommenderNamesParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDAiRecommenderNamesParams contains all the parameters to send to the API endpoint
for the get sites by ID ai recommender names operation typically these are written to a http.Request
*/
type GetSitesByIDAiRecommenderNamesParams struct {

	/*SiteID
	  Site ID to get a list of recommenders for.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) WithTimeout(timeout time.Duration) *GetSitesByIDAiRecommenderNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) WithContext(ctx context.Context) *GetSitesByIDAiRecommenderNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) WithHTTPClient(client *http.Client) *GetSitesByIDAiRecommenderNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) WithSiteID(siteID string) *GetSitesByIDAiRecommenderNamesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID ai recommender names params
func (o *GetSitesByIDAiRecommenderNamesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDAiRecommenderNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
