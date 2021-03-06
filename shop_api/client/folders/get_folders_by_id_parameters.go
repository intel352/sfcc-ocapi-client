// Code generated by go-swagger; DO NOT EDIT.

package folders

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

// NewGetFoldersByIDParams creates a new GetFoldersByIDParams object
// with the default values initialized.
func NewGetFoldersByIDParams() *GetFoldersByIDParams {
	var ()
	return &GetFoldersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFoldersByIDParamsWithTimeout creates a new GetFoldersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFoldersByIDParamsWithTimeout(timeout time.Duration) *GetFoldersByIDParams {
	var ()
	return &GetFoldersByIDParams{

		timeout: timeout,
	}
}

// NewGetFoldersByIDParamsWithContext creates a new GetFoldersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFoldersByIDParamsWithContext(ctx context.Context) *GetFoldersByIDParams {
	var ()
	return &GetFoldersByIDParams{

		Context: ctx,
	}
}

// NewGetFoldersByIDParamsWithHTTPClient creates a new GetFoldersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFoldersByIDParamsWithHTTPClient(client *http.Client) *GetFoldersByIDParams {
	var ()
	return &GetFoldersByIDParams{
		HTTPClient: client,
	}
}

/*GetFoldersByIDParams contains all the parameters to send to the API endpoint
for the get folders by ID operation typically these are written to a http.Request
*/
type GetFoldersByIDParams struct {

	/*ID
	  The id of the requested content folder.

	*/
	ID string
	/*Levels*/
	Levels *int32
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get folders by ID params
func (o *GetFoldersByIDParams) WithTimeout(timeout time.Duration) *GetFoldersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get folders by ID params
func (o *GetFoldersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get folders by ID params
func (o *GetFoldersByIDParams) WithContext(ctx context.Context) *GetFoldersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get folders by ID params
func (o *GetFoldersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get folders by ID params
func (o *GetFoldersByIDParams) WithHTTPClient(client *http.Client) *GetFoldersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get folders by ID params
func (o *GetFoldersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get folders by ID params
func (o *GetFoldersByIDParams) WithID(id string) *GetFoldersByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get folders by ID params
func (o *GetFoldersByIDParams) SetID(id string) {
	o.ID = id
}

// WithLevels adds the levels to the get folders by ID params
func (o *GetFoldersByIDParams) WithLevels(levels *int32) *GetFoldersByIDParams {
	o.SetLevels(levels)
	return o
}

// SetLevels adds the levels to the get folders by ID params
func (o *GetFoldersByIDParams) SetLevels(levels *int32) {
	o.Levels = levels
}

// WithLocale adds the locale to the get folders by ID params
func (o *GetFoldersByIDParams) WithLocale(locale *string) *GetFoldersByIDParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get folders by ID params
func (o *GetFoldersByIDParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetFoldersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Levels != nil {

		// query param levels
		var qrLevels int32
		if o.Levels != nil {
			qrLevels = *o.Levels
		}
		qLevels := swag.FormatInt32(qrLevels)
		if qLevels != "" {
			if err := r.SetQueryParam("levels", qLevels); err != nil {
				return err
			}
		}

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
