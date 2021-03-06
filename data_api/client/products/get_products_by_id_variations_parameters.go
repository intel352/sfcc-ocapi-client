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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProductsByIDVariationsParams creates a new GetProductsByIDVariationsParams object
// with the default values initialized.
func NewGetProductsByIDVariationsParams() *GetProductsByIDVariationsParams {
	var ()
	return &GetProductsByIDVariationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsByIDVariationsParamsWithTimeout creates a new GetProductsByIDVariationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsByIDVariationsParamsWithTimeout(timeout time.Duration) *GetProductsByIDVariationsParams {
	var ()
	return &GetProductsByIDVariationsParams{

		timeout: timeout,
	}
}

// NewGetProductsByIDVariationsParamsWithContext creates a new GetProductsByIDVariationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsByIDVariationsParamsWithContext(ctx context.Context) *GetProductsByIDVariationsParams {
	var ()
	return &GetProductsByIDVariationsParams{

		Context: ctx,
	}
}

// NewGetProductsByIDVariationsParamsWithHTTPClient creates a new GetProductsByIDVariationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsByIDVariationsParamsWithHTTPClient(client *http.Client) *GetProductsByIDVariationsParams {
	var ()
	return &GetProductsByIDVariationsParams{
		HTTPClient: client,
	}
}

/*GetProductsByIDVariationsParams contains all the parameters to send to the API endpoint
for the get products by ID variations operation typically these are written to a http.Request
*/
type GetProductsByIDVariationsParams struct {

	/*Count*/
	Count *int32
	/*Expand*/
	Expand []string
	/*MasterProductID
	  The id of the master product.

	*/
	MasterProductID string
	/*Select*/
	Select *string
	/*SiteID*/
	SiteID *string
	/*Start*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithTimeout(timeout time.Duration) *GetProductsByIDVariationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithContext(ctx context.Context) *GetProductsByIDVariationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithHTTPClient(client *http.Client) *GetProductsByIDVariationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithCount(count *int32) *GetProductsByIDVariationsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetCount(count *int32) {
	o.Count = count
}

// WithExpand adds the expand to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithExpand(expand []string) *GetProductsByIDVariationsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithMasterProductID adds the masterProductID to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithMasterProductID(masterProductID string) *GetProductsByIDVariationsParams {
	o.SetMasterProductID(masterProductID)
	return o
}

// SetMasterProductID adds the masterProductId to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetMasterProductID(masterProductID string) {
	o.MasterProductID = masterProductID
}

// WithSelect adds the selectVar to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithSelect(selectVar *string) *GetProductsByIDVariationsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithSiteID adds the siteID to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithSiteID(siteID *string) *GetProductsByIDVariationsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStart adds the start to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) WithStart(start *int32) *GetProductsByIDVariationsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get products by ID variations params
func (o *GetProductsByIDVariationsParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsByIDVariationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param master_product_id
	if err := r.SetPathParam("master_product_id", o.MasterProductID); err != nil {
		return err
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
