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

// NewGetCatalogsByIDParams creates a new GetCatalogsByIDParams object
// with the default values initialized.
func NewGetCatalogsByIDParams() *GetCatalogsByIDParams {
	var ()
	return &GetCatalogsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogsByIDParamsWithTimeout creates a new GetCatalogsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogsByIDParamsWithTimeout(timeout time.Duration) *GetCatalogsByIDParams {
	var ()
	return &GetCatalogsByIDParams{

		timeout: timeout,
	}
}

// NewGetCatalogsByIDParamsWithContext creates a new GetCatalogsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogsByIDParamsWithContext(ctx context.Context) *GetCatalogsByIDParams {
	var ()
	return &GetCatalogsByIDParams{

		Context: ctx,
	}
}

// NewGetCatalogsByIDParamsWithHTTPClient creates a new GetCatalogsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogsByIDParamsWithHTTPClient(client *http.Client) *GetCatalogsByIDParams {
	var ()
	return &GetCatalogsByIDParams{
		HTTPClient: client,
	}
}

/*GetCatalogsByIDParams contains all the parameters to send to the API endpoint
for the get catalogs by ID operation typically these are written to a http.Request
*/
type GetCatalogsByIDParams struct {

	/*CatalogID
	  The id of the requested catalog.

	*/
	CatalogID string
	/*Expand*/
	Expand []string
	/*Select*/
	Select *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithTimeout(timeout time.Duration) *GetCatalogsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithContext(ctx context.Context) *GetCatalogsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithHTTPClient(client *http.Client) *GetCatalogsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithCatalogID(catalogID string) *GetCatalogsByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithExpand adds the expand to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithExpand(expand []string) *GetCatalogsByIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithSelect adds the selectVar to the get catalogs by ID params
func (o *GetCatalogsByIDParams) WithSelect(selectVar *string) *GetCatalogsByIDParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get catalogs by ID params
func (o *GetCatalogsByIDParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
