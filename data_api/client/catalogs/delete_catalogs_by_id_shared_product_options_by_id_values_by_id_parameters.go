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

// NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams creates a new DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams object
// with the default values initialized.
func NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams() *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	var ()
	return &DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithTimeout creates a new DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithTimeout(timeout time.Duration) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	var ()
	return &DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams{

		timeout: timeout,
	}
}

// NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithContext creates a new DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithContext(ctx context.Context) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	var ()
	return &DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams{

		Context: ctx,
	}
}

// NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithHTTPClient creates a new DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParamsWithHTTPClient(client *http.Client) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	var ()
	return &DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams{
		HTTPClient: client,
	}
}

/*DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams contains all the parameters to send to the API endpoint
for the delete catalogs by ID shared product options by ID values by ID operation typically these are written to a http.Request
*/
type DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams struct {

	/*CatalogID
	  The id of the catalog.

	*/
	CatalogID string
	/*ID
	  The id of the shared product option value.

	*/
	ID string
	/*OptionID
	  The id of the shared product option.

	*/
	OptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithTimeout(timeout time.Duration) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithContext(ctx context.Context) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithHTTPClient(client *http.Client) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithCatalogID(catalogID string) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithID adds the id to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithID(id string) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetID(id string) {
	o.ID = id
}

// WithOptionID adds the optionID to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WithOptionID(optionID string) *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams {
	o.SetOptionID(optionID)
	return o
}

// SetOptionID adds the optionId to the delete catalogs by ID shared product options by ID values by ID params
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) SetOptionID(optionID string) {
	o.OptionID = optionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogsByIDSharedProductOptionsByIDValuesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param option_id
	if err := r.SetPathParam("option_id", o.OptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
