// Code generated by go-swagger; DO NOT EDIT.

package categories

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

// NewGetCategoriesByIdsParams creates a new GetCategoriesByIdsParams object
// with the default values initialized.
func NewGetCategoriesByIdsParams() *GetCategoriesByIdsParams {
	var ()
	return &GetCategoriesByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoriesByIdsParamsWithTimeout creates a new GetCategoriesByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCategoriesByIdsParamsWithTimeout(timeout time.Duration) *GetCategoriesByIdsParams {
	var ()
	return &GetCategoriesByIdsParams{

		timeout: timeout,
	}
}

// NewGetCategoriesByIdsParamsWithContext creates a new GetCategoriesByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCategoriesByIdsParamsWithContext(ctx context.Context) *GetCategoriesByIdsParams {
	var ()
	return &GetCategoriesByIdsParams{

		Context: ctx,
	}
}

// NewGetCategoriesByIdsParamsWithHTTPClient creates a new GetCategoriesByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCategoriesByIdsParamsWithHTTPClient(client *http.Client) *GetCategoriesByIdsParams {
	var ()
	return &GetCategoriesByIdsParams{
		HTTPClient: client,
	}
}

/*GetCategoriesByIdsParams contains all the parameters to send to the API endpoint
for the get categories by ids operation typically these are written to a http.Request
*/
type GetCategoriesByIdsParams struct {

	/*Ids*/
	Ids []string
	/*Levels*/
	Levels *int32
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithTimeout(timeout time.Duration) *GetCategoriesByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithContext(ctx context.Context) *GetCategoriesByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithHTTPClient(client *http.Client) *GetCategoriesByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithIds(ids []string) *GetCategoriesByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithLevels adds the levels to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithLevels(levels *int32) *GetCategoriesByIdsParams {
	o.SetLevels(levels)
	return o
}

// SetLevels adds the levels to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetLevels(levels *int32) {
	o.Levels = levels
}

// WithLocale adds the locale to the get categories by ids params
func (o *GetCategoriesByIdsParams) WithLocale(locale *string) *GetCategoriesByIdsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get categories by ids params
func (o *GetCategoriesByIdsParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoriesByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "")
	// path array param ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedIds) > 0 {
		if err := r.SetPathParam("ids", joinedIds[0]); err != nil {
			return err
		}
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
