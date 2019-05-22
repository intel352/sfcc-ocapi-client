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

// NewGetSitesByIDStoresParams creates a new GetSitesByIDStoresParams object
// with the default values initialized.
func NewGetSitesByIDStoresParams() *GetSitesByIDStoresParams {
	var ()
	return &GetSitesByIDStoresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDStoresParamsWithTimeout creates a new GetSitesByIDStoresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDStoresParamsWithTimeout(timeout time.Duration) *GetSitesByIDStoresParams {
	var ()
	return &GetSitesByIDStoresParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDStoresParamsWithContext creates a new GetSitesByIDStoresParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDStoresParamsWithContext(ctx context.Context) *GetSitesByIDStoresParams {
	var ()
	return &GetSitesByIDStoresParams{

		Context: ctx,
	}
}

// NewGetSitesByIDStoresParamsWithHTTPClient creates a new GetSitesByIDStoresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDStoresParamsWithHTTPClient(client *http.Client) *GetSitesByIDStoresParams {
	var ()
	return &GetSitesByIDStoresParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDStoresParams contains all the parameters to send to the API endpoint
for the get sites by ID stores operation typically these are written to a http.Request
*/
type GetSitesByIDStoresParams struct {

	/*Count*/
	Count *int32
	/*Select*/
	Select *string
	/*SiteID*/
	SiteID string
	/*Start*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithTimeout(timeout time.Duration) *GetSitesByIDStoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithContext(ctx context.Context) *GetSitesByIDStoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithHTTPClient(client *http.Client) *GetSitesByIDStoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithCount(count *int32) *GetSitesByIDStoresParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetCount(count *int32) {
	o.Count = count
}

// WithSelect adds the selectVar to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithSelect(selectVar *string) *GetSitesByIDStoresParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithSiteID adds the siteID to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithSiteID(siteID string) *GetSitesByIDStoresParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStart adds the start to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) WithStart(start *int32) *GetSitesByIDStoresParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get sites by ID stores params
func (o *GetSitesByIDStoresParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDStoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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