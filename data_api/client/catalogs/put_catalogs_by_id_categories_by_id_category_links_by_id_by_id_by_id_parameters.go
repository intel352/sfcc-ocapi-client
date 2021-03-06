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

// NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams creates a new PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized.
func NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams() *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithTimeout creates a new PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithTimeout(timeout time.Duration) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		timeout: timeout,
	}
}

// NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithContext creates a new PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithContext(ctx context.Context) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		Context: ctx,
	}
}

// NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithHTTPClient creates a new PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithHTTPClient(client *http.Client) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{
		HTTPClient: client,
	}
}

/*PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams contains all the parameters to send to the API endpoint
for the put catalogs by ID categories by ID category links by ID by ID by ID operation typically these are written to a http.Request
*/
type PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams struct {

	/*Body*/
	Body *models.CategoryLink
	/*CatalogID
	  The id of the source catalog.

	*/
	CatalogID string
	/*CategoryID
	  The id of the source category.

	*/
	CategoryID string
	/*TargetCatalogID
	  The id of the target catalog.

	*/
	TargetCatalogID string
	/*TargetCategoryID
	  The id of the target category.

	*/
	TargetCategoryID string
	/*Type
	  the link type

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTimeout(timeout time.Duration) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithContext(ctx context.Context) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithHTTPClient(client *http.Client) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithBody(body *models.CategoryLink) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetBody(body *models.CategoryLink) {
	o.Body = body
}

// WithCatalogID adds the catalogID to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithCatalogID(catalogID string) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithCategoryID adds the categoryID to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithCategoryID(categoryID string) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithTargetCatalogID adds the targetCatalogID to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTargetCatalogID(targetCatalogID string) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTargetCatalogID(targetCatalogID)
	return o
}

// SetTargetCatalogID adds the targetCatalogId to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTargetCatalogID(targetCatalogID string) {
	o.TargetCatalogID = targetCatalogID
}

// WithTargetCategoryID adds the targetCategoryID to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTargetCategoryID(targetCategoryID string) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTargetCategoryID(targetCategoryID)
	return o
}

// SetTargetCategoryID adds the targetCategoryId to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTargetCategoryID(targetCategoryID string) {
	o.TargetCategoryID = targetCategoryID
}

// WithType adds the typeVar to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithType(typeVar string) *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put catalogs by ID categories by ID category links by ID by ID by ID params
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PutCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param category_id
	if err := r.SetPathParam("category_id", o.CategoryID); err != nil {
		return err
	}

	// path param target_catalog_id
	if err := r.SetPathParam("target_catalog_id", o.TargetCatalogID); err != nil {
		return err
	}

	// path param target_category_id
	if err := r.SetPathParam("target_category_id", o.TargetCategoryID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
