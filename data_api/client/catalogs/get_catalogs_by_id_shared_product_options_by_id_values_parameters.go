// Code generated by go-swagger; DO NOT EDIT.

package catalogs

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

// NewGetCatalogsByIDSharedProductOptionsByIDValuesParams creates a new GetCatalogsByIDSharedProductOptionsByIDValuesParams object
// with the default values initialized.
func NewGetCatalogsByIDSharedProductOptionsByIDValuesParams() *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	var ()
	return &GetCatalogsByIDSharedProductOptionsByIDValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithTimeout creates a new GetCatalogsByIDSharedProductOptionsByIDValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithTimeout(timeout time.Duration) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	var ()
	return &GetCatalogsByIDSharedProductOptionsByIDValuesParams{

		timeout: timeout,
	}
}

// NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithContext creates a new GetCatalogsByIDSharedProductOptionsByIDValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithContext(ctx context.Context) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	var ()
	return &GetCatalogsByIDSharedProductOptionsByIDValuesParams{

		Context: ctx,
	}
}

// NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithHTTPClient creates a new GetCatalogsByIDSharedProductOptionsByIDValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogsByIDSharedProductOptionsByIDValuesParamsWithHTTPClient(client *http.Client) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	var ()
	return &GetCatalogsByIDSharedProductOptionsByIDValuesParams{
		HTTPClient: client,
	}
}

/*GetCatalogsByIDSharedProductOptionsByIDValuesParams contains all the parameters to send to the API endpoint
for the get catalogs by ID shared product options by ID values operation typically these are written to a http.Request
*/
type GetCatalogsByIDSharedProductOptionsByIDValuesParams struct {

	/*CatalogID
	  The id of the catalog.

	*/
	CatalogID string
	/*Count*/
	Count *int32
	/*OptionID
	  The id of the shared product option.

	*/
	OptionID string
	/*Select*/
	Select *string
	/*Start*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithTimeout(timeout time.Duration) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithContext(ctx context.Context) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithHTTPClient(client *http.Client) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithCatalogID(catalogID string) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithCount adds the count to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithCount(count *int32) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetCount(count *int32) {
	o.Count = count
}

// WithOptionID adds the optionID to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithOptionID(optionID string) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetOptionID(optionID)
	return o
}

// SetOptionID adds the optionId to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetOptionID(optionID string) {
	o.OptionID = optionID
}

// WithSelect adds the selectVar to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithSelect(selectVar *string) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithStart adds the start to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WithStart(start *int32) *GetCatalogsByIDSharedProductOptionsByIDValuesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get catalogs by ID shared product options by ID values params
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogsByIDSharedProductOptionsByIDValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

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

	// path param option_id
	if err := r.SetPathParam("option_id", o.OptionID); err != nil {
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