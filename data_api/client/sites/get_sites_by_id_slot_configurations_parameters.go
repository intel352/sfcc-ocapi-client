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

// NewGetSitesByIDSlotConfigurationsParams creates a new GetSitesByIDSlotConfigurationsParams object
// with the default values initialized.
func NewGetSitesByIDSlotConfigurationsParams() *GetSitesByIDSlotConfigurationsParams {
	var ()
	return &GetSitesByIDSlotConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDSlotConfigurationsParamsWithTimeout creates a new GetSitesByIDSlotConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDSlotConfigurationsParamsWithTimeout(timeout time.Duration) *GetSitesByIDSlotConfigurationsParams {
	var ()
	return &GetSitesByIDSlotConfigurationsParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDSlotConfigurationsParamsWithContext creates a new GetSitesByIDSlotConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDSlotConfigurationsParamsWithContext(ctx context.Context) *GetSitesByIDSlotConfigurationsParams {
	var ()
	return &GetSitesByIDSlotConfigurationsParams{

		Context: ctx,
	}
}

// NewGetSitesByIDSlotConfigurationsParamsWithHTTPClient creates a new GetSitesByIDSlotConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDSlotConfigurationsParamsWithHTTPClient(client *http.Client) *GetSitesByIDSlotConfigurationsParams {
	var ()
	return &GetSitesByIDSlotConfigurationsParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDSlotConfigurationsParams contains all the parameters to send to the API endpoint
for the get sites by ID slot configurations operation typically these are written to a http.Request
*/
type GetSitesByIDSlotConfigurationsParams struct {

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

// WithTimeout adds the timeout to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithTimeout(timeout time.Duration) *GetSitesByIDSlotConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithContext(ctx context.Context) *GetSitesByIDSlotConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithHTTPClient(client *http.Client) *GetSitesByIDSlotConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithCount(count *int32) *GetSitesByIDSlotConfigurationsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetCount(count *int32) {
	o.Count = count
}

// WithSelect adds the selectVar to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithSelect(selectVar *string) *GetSitesByIDSlotConfigurationsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithSiteID adds the siteID to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithSiteID(siteID string) *GetSitesByIDSlotConfigurationsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStart adds the start to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) WithStart(start *int32) *GetSitesByIDSlotConfigurationsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get sites by ID slot configurations params
func (o *GetSitesByIDSlotConfigurationsParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDSlotConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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