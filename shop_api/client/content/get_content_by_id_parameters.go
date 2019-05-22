// Code generated by go-swagger; DO NOT EDIT.

package content

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

// NewGetContentByIDParams creates a new GetContentByIDParams object
// with the default values initialized.
func NewGetContentByIDParams() *GetContentByIDParams {
	var ()
	return &GetContentByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContentByIDParamsWithTimeout creates a new GetContentByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContentByIDParamsWithTimeout(timeout time.Duration) *GetContentByIDParams {
	var ()
	return &GetContentByIDParams{

		timeout: timeout,
	}
}

// NewGetContentByIDParamsWithContext creates a new GetContentByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContentByIDParamsWithContext(ctx context.Context) *GetContentByIDParams {
	var ()
	return &GetContentByIDParams{

		Context: ctx,
	}
}

// NewGetContentByIDParamsWithHTTPClient creates a new GetContentByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContentByIDParamsWithHTTPClient(client *http.Client) *GetContentByIDParams {
	var ()
	return &GetContentByIDParams{
		HTTPClient: client,
	}
}

/*GetContentByIDParams contains all the parameters to send to the API endpoint
for the get content by ID operation typically these are written to a http.Request
*/
type GetContentByIDParams struct {

	/*ID
	  The id of the requested content asset.

	*/
	ID string
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get content by ID params
func (o *GetContentByIDParams) WithTimeout(timeout time.Duration) *GetContentByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get content by ID params
func (o *GetContentByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get content by ID params
func (o *GetContentByIDParams) WithContext(ctx context.Context) *GetContentByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get content by ID params
func (o *GetContentByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get content by ID params
func (o *GetContentByIDParams) WithHTTPClient(client *http.Client) *GetContentByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get content by ID params
func (o *GetContentByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get content by ID params
func (o *GetContentByIDParams) WithID(id string) *GetContentByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get content by ID params
func (o *GetContentByIDParams) SetID(id string) {
	o.ID = id
}

// WithLocale adds the locale to the get content by ID params
func (o *GetContentByIDParams) WithLocale(locale *string) *GetContentByIDParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get content by ID params
func (o *GetContentByIDParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetContentByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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