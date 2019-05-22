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

// NewGetProductsByIdsParams creates a new GetProductsByIdsParams object
// with the default values initialized.
func NewGetProductsByIdsParams() *GetProductsByIdsParams {
	var ()
	return &GetProductsByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsByIdsParamsWithTimeout creates a new GetProductsByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsByIdsParamsWithTimeout(timeout time.Duration) *GetProductsByIdsParams {
	var ()
	return &GetProductsByIdsParams{

		timeout: timeout,
	}
}

// NewGetProductsByIdsParamsWithContext creates a new GetProductsByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsByIdsParamsWithContext(ctx context.Context) *GetProductsByIdsParams {
	var ()
	return &GetProductsByIdsParams{

		Context: ctx,
	}
}

// NewGetProductsByIdsParamsWithHTTPClient creates a new GetProductsByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsByIdsParamsWithHTTPClient(client *http.Client) *GetProductsByIdsParams {
	var ()
	return &GetProductsByIdsParams{
		HTTPClient: client,
	}
}

/*GetProductsByIdsParams contains all the parameters to send to the API endpoint
for the get products by ids operation typically these are written to a http.Request
*/
type GetProductsByIdsParams struct {

	/*AllImages*/
	AllImages *bool
	/*Currency*/
	Currency *string
	/*Expand*/
	Expand []string
	/*Ids*/
	Ids []string
	/*InventoryIds*/
	InventoryIds []string
	/*Locale*/
	Locale *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products by ids params
func (o *GetProductsByIdsParams) WithTimeout(timeout time.Duration) *GetProductsByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products by ids params
func (o *GetProductsByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products by ids params
func (o *GetProductsByIdsParams) WithContext(ctx context.Context) *GetProductsByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products by ids params
func (o *GetProductsByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products by ids params
func (o *GetProductsByIdsParams) WithHTTPClient(client *http.Client) *GetProductsByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products by ids params
func (o *GetProductsByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllImages adds the allImages to the get products by ids params
func (o *GetProductsByIdsParams) WithAllImages(allImages *bool) *GetProductsByIdsParams {
	o.SetAllImages(allImages)
	return o
}

// SetAllImages adds the allImages to the get products by ids params
func (o *GetProductsByIdsParams) SetAllImages(allImages *bool) {
	o.AllImages = allImages
}

// WithCurrency adds the currency to the get products by ids params
func (o *GetProductsByIdsParams) WithCurrency(currency *string) *GetProductsByIdsParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get products by ids params
func (o *GetProductsByIdsParams) SetCurrency(currency *string) {
	o.Currency = currency
}

// WithExpand adds the expand to the get products by ids params
func (o *GetProductsByIdsParams) WithExpand(expand []string) *GetProductsByIdsParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get products by ids params
func (o *GetProductsByIdsParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithIds adds the ids to the get products by ids params
func (o *GetProductsByIdsParams) WithIds(ids []string) *GetProductsByIdsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get products by ids params
func (o *GetProductsByIdsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithInventoryIds adds the inventoryIds to the get products by ids params
func (o *GetProductsByIdsParams) WithInventoryIds(inventoryIds []string) *GetProductsByIdsParams {
	o.SetInventoryIds(inventoryIds)
	return o
}

// SetInventoryIds adds the inventoryIds to the get products by ids params
func (o *GetProductsByIdsParams) SetInventoryIds(inventoryIds []string) {
	o.InventoryIds = inventoryIds
}

// WithLocale adds the locale to the get products by ids params
func (o *GetProductsByIdsParams) WithLocale(locale *string) *GetProductsByIdsParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get products by ids params
func (o *GetProductsByIdsParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllImages != nil {

		// query param all_images
		var qrAllImages bool
		if o.AllImages != nil {
			qrAllImages = *o.AllImages
		}
		qAllImages := swag.FormatBool(qrAllImages)
		if qAllImages != "" {
			if err := r.SetQueryParam("all_images", qAllImages); err != nil {
				return err
			}
		}

	}

	if o.Currency != nil {

		// query param currency
		var qrCurrency string
		if o.Currency != nil {
			qrCurrency = *o.Currency
		}
		qCurrency := qrCurrency
		if qCurrency != "" {
			if err := r.SetQueryParam("currency", qCurrency); err != nil {
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

	valuesInventoryIds := o.InventoryIds

	joinedInventoryIds := swag.JoinByFormat(valuesInventoryIds, "")
	// query array param inventory_ids
	if err := r.SetQueryParam("inventory_ids", joinedInventoryIds...); err != nil {
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