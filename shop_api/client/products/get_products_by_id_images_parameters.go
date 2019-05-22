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

// NewGetProductsByIDImagesParams creates a new GetProductsByIDImagesParams object
// with the default values initialized.
func NewGetProductsByIDImagesParams() *GetProductsByIDImagesParams {
	var ()
	return &GetProductsByIDImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsByIDImagesParamsWithTimeout creates a new GetProductsByIDImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsByIDImagesParamsWithTimeout(timeout time.Duration) *GetProductsByIDImagesParams {
	var ()
	return &GetProductsByIDImagesParams{

		timeout: timeout,
	}
}

// NewGetProductsByIDImagesParamsWithContext creates a new GetProductsByIDImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductsByIDImagesParamsWithContext(ctx context.Context) *GetProductsByIDImagesParams {
	var ()
	return &GetProductsByIDImagesParams{

		Context: ctx,
	}
}

// NewGetProductsByIDImagesParamsWithHTTPClient creates a new GetProductsByIDImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductsByIDImagesParamsWithHTTPClient(client *http.Client) *GetProductsByIDImagesParams {
	var ()
	return &GetProductsByIDImagesParams{
		HTTPClient: client,
	}
}

/*GetProductsByIDImagesParams contains all the parameters to send to the API endpoint
for the get products by ID images operation typically these are written to a http.Request
*/
type GetProductsByIDImagesParams struct {

	/*AllImages*/
	AllImages *bool
	/*ID
	  The requested product id.

	*/
	ID string
	/*Locale*/
	Locale *string
	/*VariationAttribute*/
	VariationAttribute *string
	/*ViewType*/
	ViewType []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithTimeout(timeout time.Duration) *GetProductsByIDImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithContext(ctx context.Context) *GetProductsByIDImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithHTTPClient(client *http.Client) *GetProductsByIDImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllImages adds the allImages to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithAllImages(allImages *bool) *GetProductsByIDImagesParams {
	o.SetAllImages(allImages)
	return o
}

// SetAllImages adds the allImages to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetAllImages(allImages *bool) {
	o.AllImages = allImages
}

// WithID adds the id to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithID(id string) *GetProductsByIDImagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetID(id string) {
	o.ID = id
}

// WithLocale adds the locale to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithLocale(locale *string) *GetProductsByIDImagesParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WithVariationAttribute adds the variationAttribute to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithVariationAttribute(variationAttribute *string) *GetProductsByIDImagesParams {
	o.SetVariationAttribute(variationAttribute)
	return o
}

// SetVariationAttribute adds the variationAttribute to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetVariationAttribute(variationAttribute *string) {
	o.VariationAttribute = variationAttribute
}

// WithViewType adds the viewType to the get products by ID images params
func (o *GetProductsByIDImagesParams) WithViewType(viewType []string) *GetProductsByIDImagesParams {
	o.SetViewType(viewType)
	return o
}

// SetViewType adds the viewType to the get products by ID images params
func (o *GetProductsByIDImagesParams) SetViewType(viewType []string) {
	o.ViewType = viewType
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsByIDImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.VariationAttribute != nil {

		// query param variation_attribute
		var qrVariationAttribute string
		if o.VariationAttribute != nil {
			qrVariationAttribute = *o.VariationAttribute
		}
		qVariationAttribute := qrVariationAttribute
		if qVariationAttribute != "" {
			if err := r.SetQueryParam("variation_attribute", qVariationAttribute); err != nil {
				return err
			}
		}

	}

	valuesViewType := o.ViewType

	joinedViewType := swag.JoinByFormat(valuesViewType, "")
	// query array param view_type
	if err := r.SetQueryParam("view_type", joinedViewType...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}