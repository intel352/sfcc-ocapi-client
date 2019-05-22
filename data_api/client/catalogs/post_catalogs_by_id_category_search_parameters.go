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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPostCatalogsByIDCategorySearchParams creates a new PostCatalogsByIDCategorySearchParams object
// with the default values initialized.
func NewPostCatalogsByIDCategorySearchParams() *PostCatalogsByIDCategorySearchParams {
	var ()
	return &PostCatalogsByIDCategorySearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCatalogsByIDCategorySearchParamsWithTimeout creates a new PostCatalogsByIDCategorySearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCatalogsByIDCategorySearchParamsWithTimeout(timeout time.Duration) *PostCatalogsByIDCategorySearchParams {
	var ()
	return &PostCatalogsByIDCategorySearchParams{

		timeout: timeout,
	}
}

// NewPostCatalogsByIDCategorySearchParamsWithContext creates a new PostCatalogsByIDCategorySearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCatalogsByIDCategorySearchParamsWithContext(ctx context.Context) *PostCatalogsByIDCategorySearchParams {
	var ()
	return &PostCatalogsByIDCategorySearchParams{

		Context: ctx,
	}
}

// NewPostCatalogsByIDCategorySearchParamsWithHTTPClient creates a new PostCatalogsByIDCategorySearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCatalogsByIDCategorySearchParamsWithHTTPClient(client *http.Client) *PostCatalogsByIDCategorySearchParams {
	var ()
	return &PostCatalogsByIDCategorySearchParams{
		HTTPClient: client,
	}
}

/*PostCatalogsByIDCategorySearchParams contains all the parameters to send to the API endpoint
for the post catalogs by ID category search operation typically these are written to a http.Request
*/
type PostCatalogsByIDCategorySearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*CatalogID
	  The id of the catalog.

	*/
	CatalogID string
	/*Levels*/
	Levels *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithTimeout(timeout time.Duration) *PostCatalogsByIDCategorySearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithContext(ctx context.Context) *PostCatalogsByIDCategorySearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithHTTPClient(client *http.Client) *PostCatalogsByIDCategorySearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithBody(body *models.SearchRequest) *PostCatalogsByIDCategorySearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithCatalogID adds the catalogID to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithCatalogID(catalogID string) *PostCatalogsByIDCategorySearchParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetCatalogID(catalogID string) {
	o.CatalogID = catalogID
}

// WithLevels adds the levels to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) WithLevels(levels *int32) *PostCatalogsByIDCategorySearchParams {
	o.SetLevels(levels)
	return o
}

// SetLevels adds the levels to the post catalogs by ID category search params
func (o *PostCatalogsByIDCategorySearchParams) SetLevels(levels *int32) {
	o.Levels = levels
}

// WriteToRequest writes these params to a swagger request
func (o *PostCatalogsByIDCategorySearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
