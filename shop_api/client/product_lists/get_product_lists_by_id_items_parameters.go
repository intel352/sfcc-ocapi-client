// Code generated by go-swagger; DO NOT EDIT.

package product_lists

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

// NewGetProductListsByIDItemsParams creates a new GetProductListsByIDItemsParams object
// with the default values initialized.
func NewGetProductListsByIDItemsParams() *GetProductListsByIDItemsParams {
	var ()
	return &GetProductListsByIDItemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductListsByIDItemsParamsWithTimeout creates a new GetProductListsByIDItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductListsByIDItemsParamsWithTimeout(timeout time.Duration) *GetProductListsByIDItemsParams {
	var ()
	return &GetProductListsByIDItemsParams{

		timeout: timeout,
	}
}

// NewGetProductListsByIDItemsParamsWithContext creates a new GetProductListsByIDItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductListsByIDItemsParamsWithContext(ctx context.Context) *GetProductListsByIDItemsParams {
	var ()
	return &GetProductListsByIDItemsParams{

		Context: ctx,
	}
}

// NewGetProductListsByIDItemsParamsWithHTTPClient creates a new GetProductListsByIDItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductListsByIDItemsParamsWithHTTPClient(client *http.Client) *GetProductListsByIDItemsParams {
	var ()
	return &GetProductListsByIDItemsParams{
		HTTPClient: client,
	}
}

/*GetProductListsByIDItemsParams contains all the parameters to send to the API endpoint
for the get product lists by ID items operation typically these are written to a http.Request
*/
type GetProductListsByIDItemsParams struct {

	/*Expand*/
	Expand []string
	/*ListID
	  The id of the list.

	*/
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) WithTimeout(timeout time.Duration) *GetProductListsByIDItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) WithContext(ctx context.Context) *GetProductListsByIDItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) WithHTTPClient(client *http.Client) *GetProductListsByIDItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) WithExpand(expand []string) *GetProductListsByIDItemsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithListID adds the listID to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) WithListID(listID string) *GetProductListsByIDItemsParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the get product lists by ID items params
func (o *GetProductListsByIDItemsParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductListsByIDItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param list_id
	if err := r.SetPathParam("list_id", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
