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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams creates a new GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams object
// with the default values initialized.
func NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams() *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	var ()
	return &GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithTimeout creates a new GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithTimeout(timeout time.Duration) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	var ()
	return &GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams{

		timeout: timeout,
	}
}

// NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithContext creates a new GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithContext(ctx context.Context) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	var ()
	return &GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams{

		Context: ctx,
	}
}

// NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithHTTPClient creates a new GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogsByIDSharedVariationAttributesByIDValuesByIDParamsWithHTTPClient(client *http.Client) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	var ()
	return &GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams{
		HTTPClient: client,
	}
}

/*GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams contains all the parameters to send to the API endpoint
for the get catalogs by ID shared variation attributes by ID values by ID operation typically these are written to a http.Request
*/
type GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams struct {

	/*AttributeID
	  The variation attribute custom ID

	*/
	AttributeID string
	/*CatalogID
	  The owning catalog ID.

	*/
	CatalogID string
	/*ID
	  The id of the variation attribute value

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithTimeout(timeout time.Duration) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithContext(ctx context.Context) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithHTTPClient(client *http.Client) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributeID adds the attributeID to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithAttributeID(attributeID string) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetAttributeID(attributeID)
	return o
}

// SetAttributeID adds the attributeId to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetAttributeID(attributeID string) {
	o.AttributeID = attributeID
}

// WithCatalogID adds the catalogID to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithCatalogID(catalogID string) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithID adds the id to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WithID(id string) *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get catalogs by ID shared variation attributes by ID values by ID params
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogsByIDSharedVariationAttributesByIDValuesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param attribute_id
	if err := r.SetPathParam("attribute_id", o.AttributeID); err != nil {
		return err
	}

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
