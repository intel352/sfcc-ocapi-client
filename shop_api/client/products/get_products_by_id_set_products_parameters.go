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
)

// NewGetProductsByIDSetProductsParams creates a new GetProductsByIDSetProductsParams object
// with the default values initialized.
func NewGetProductsByIDSetProductsParams() *GetProductsByIDSetProductsParams {
	var ()
	return &GetProductsByIDSetProductsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsByIDSetProductsParamsWithTimeout creates a new GetProductsByIDSetProductsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsByIDSetProductsParamsWithTimeout(timeout time.Duration) *GetProductsByIDSetProductsParams {
	var ()
	return &GetProductsByIDSetProductsParams{

		timeout: timeout,
	}
}

// NewGetProductsByIDSetProductsParamsWithContext creates a new GetProductsByIDSetProductsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsByIDSetProductsParamsWithContext(ctx context.Context) *GetProductsByIDSetProductsParams {
	var ()
	return &GetProductsByIDSetProductsParams{

		Context: ctx,
	}
}

// NewGetProductsByIDSetProductsParamsWithHTTPClient creates a new GetProductsByIDSetProductsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsByIDSetProductsParamsWithHTTPClient(client *http.Client) *GetProductsByIDSetProductsParams {
	var ()
	return &GetProductsByIDSetProductsParams{
		HTTPClient: client,
	}
}

/*GetProductsByIDSetProductsParams contains all the parameters to send to the API endpoint
for the get products by ID set products operation typically these are written to a http.Request
*/
type GetProductsByIDSetProductsParams struct {

	/*ID
	  The requested product id.

	*/
	ID string
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) WithTimeout(timeout time.Duration) *GetProductsByIDSetProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) WithContext(ctx context.Context) *GetProductsByIDSetProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) WithHTTPClient(client *http.Client) *GetProductsByIDSetProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) WithID(id string) *GetProductsByIDSetProductsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) SetID(id string) {
	o.ID = id
}

// WithLocale adds the locale to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) WithLocale(locale *string) *GetProductsByIDSetProductsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get products by ID set products params
func (o *GetProductsByIDSetProductsParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsByIDSetProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Locale != nil {

		// query param locale
		var qrLocale string
		if o.Locale != nil {
			qrLocale = *o.Locale
		}
		qLocale := qrLocale
		if qLocale != "" {
			if err := r.SetQueryParam("locale", qLocale); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
