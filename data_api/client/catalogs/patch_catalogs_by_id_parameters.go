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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPatchCatalogsByIDParams creates a new PatchCatalogsByIDParams object
// with the default values initialized.
func NewPatchCatalogsByIDParams() *PatchCatalogsByIDParams {
	var ()
	return &PatchCatalogsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCatalogsByIDParamsWithTimeout creates a new PatchCatalogsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCatalogsByIDParamsWithTimeout(timeout time.Duration) *PatchCatalogsByIDParams {
	var ()
	return &PatchCatalogsByIDParams{

		timeout: timeout,
	}
}

// NewPatchCatalogsByIDParamsWithContext creates a new PatchCatalogsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCatalogsByIDParamsWithContext(ctx context.Context) *PatchCatalogsByIDParams {
	var ()
	return &PatchCatalogsByIDParams{

		Context: ctx,
	}
}

// NewPatchCatalogsByIDParamsWithHTTPClient creates a new PatchCatalogsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCatalogsByIDParamsWithHTTPClient(client *http.Client) *PatchCatalogsByIDParams {
	var ()
	return &PatchCatalogsByIDParams{
		HTTPClient: client,
	}
}

/*PatchCatalogsByIDParams contains all the parameters to send to the API endpoint
for the patch catalogs by ID operation typically these are written to a http.Request
*/
type PatchCatalogsByIDParams struct {

	/*Body*/
	Body *models.Catalog
	/*CatalogID
	  The id of the requested catalog.

	*/
	CatalogID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) WithTimeout(timeout time.Duration) *PatchCatalogsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) WithContext(ctx context.Context) *PatchCatalogsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) WithHTTPClient(client *http.Client) *PatchCatalogsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) WithBody(body *models.Catalog) *PatchCatalogsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) SetBody(body *models.Catalog) {
	o.Body = body
}

// WithCatalogID adds the catalogID to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) WithCatalogID(catalogID string) *PatchCatalogsByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the patch catalogs by ID params
func (o *PatchCatalogsByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCatalogsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param catalog_id
	if err := r.SetPathParam("catalog_id", o.CatalogID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
