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

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams creates a new DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized.
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams() *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithTimeout creates a new DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithTimeout(timeout time.Duration) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		timeout: timeout,
	}
}

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithContext creates a new DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithContext(ctx context.Context) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{

		Context: ctx,
	}
}

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithHTTPClient creates a new DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParamsWithHTTPClient(client *http.Client) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	var ()
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams{
		HTTPClient: client,
	}
}

/*DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams contains all the parameters to send to the API endpoint
for the delete catalogs by ID categories by ID category links by ID by ID by ID operation typically these are written to a http.Request
*/
type DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams struct {

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

// WithTimeout adds the timeout to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTimeout(timeout time.Duration) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithContext(ctx context.Context) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithHTTPClient(client *http.Client) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithCatalogID(catalogID string) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithCategoryID adds the categoryID to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithCategoryID(categoryID string) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithTargetCatalogID adds the targetCatalogID to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTargetCatalogID(targetCatalogID string) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTargetCatalogID(targetCatalogID)
	return o
}

// SetTargetCatalogID adds the targetCatalogId to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTargetCatalogID(targetCatalogID string) {
	o.TargetCatalogID = targetCatalogID
}

// WithTargetCategoryID adds the targetCategoryID to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithTargetCategoryID(targetCategoryID string) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetTargetCategoryID(targetCategoryID)
	return o
}

// SetTargetCategoryID adds the targetCategoryId to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetTargetCategoryID(targetCategoryID string) {
	o.TargetCategoryID = targetCategoryID
}

// WithType adds the typeVar to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WithType(typeVar string) *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the delete catalogs by ID categories by ID category links by ID by ID by ID params
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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