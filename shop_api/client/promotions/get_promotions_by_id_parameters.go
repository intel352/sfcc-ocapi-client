// Code generated by go-swagger; DO NOT EDIT.

package promotions

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

// NewGetPromotionsByIDParams creates a new GetPromotionsByIDParams object
// with the default values initialized.
func NewGetPromotionsByIDParams() *GetPromotionsByIDParams {
	var ()
	return &GetPromotionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPromotionsByIDParamsWithTimeout creates a new GetPromotionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPromotionsByIDParamsWithTimeout(timeout time.Duration) *GetPromotionsByIDParams {
	var ()
	return &GetPromotionsByIDParams{

		timeout: timeout,
	}
}

// NewGetPromotionsByIDParamsWithContext creates a new GetPromotionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPromotionsByIDParamsWithContext(ctx context.Context) *GetPromotionsByIDParams {
	var ()
	return &GetPromotionsByIDParams{

		Context: ctx,
	}
}

// NewGetPromotionsByIDParamsWithHTTPClient creates a new GetPromotionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPromotionsByIDParamsWithHTTPClient(client *http.Client) *GetPromotionsByIDParams {
	var ()
	return &GetPromotionsByIDParams{
		HTTPClient: client,
	}
}

/*GetPromotionsByIDParams contains all the parameters to send to the API endpoint
for the get promotions by ID operation typically these are written to a http.Request
*/
type GetPromotionsByIDParams struct {

	/*ID
	  The id of the requested promotion.

	*/
	ID string
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get promotions by ID params
func (o *GetPromotionsByIDParams) WithTimeout(timeout time.Duration) *GetPromotionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get promotions by ID params
func (o *GetPromotionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get promotions by ID params
func (o *GetPromotionsByIDParams) WithContext(ctx context.Context) *GetPromotionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get promotions by ID params
func (o *GetPromotionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get promotions by ID params
func (o *GetPromotionsByIDParams) WithHTTPClient(client *http.Client) *GetPromotionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get promotions by ID params
func (o *GetPromotionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get promotions by ID params
func (o *GetPromotionsByIDParams) WithID(id string) *GetPromotionsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get promotions by ID params
func (o *GetPromotionsByIDParams) SetID(id string) {
	o.ID = id
}

// WithLocale adds the locale to the get promotions by ID params
func (o *GetPromotionsByIDParams) WithLocale(locale *string) *GetPromotionsByIDParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get promotions by ID params
func (o *GetPromotionsByIDParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetPromotionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
