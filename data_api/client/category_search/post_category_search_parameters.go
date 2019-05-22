// Code generated by go-swagger; DO NOT EDIT.

package category_search

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

// NewPostCategorySearchParams creates a new PostCategorySearchParams object
// with the default values initialized.
func NewPostCategorySearchParams() *PostCategorySearchParams {
	var ()
	return &PostCategorySearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCategorySearchParamsWithTimeout creates a new PostCategorySearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCategorySearchParamsWithTimeout(timeout time.Duration) *PostCategorySearchParams {
	var ()
	return &PostCategorySearchParams{

		timeout: timeout,
	}
}

// NewPostCategorySearchParamsWithContext creates a new PostCategorySearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCategorySearchParamsWithContext(ctx context.Context) *PostCategorySearchParams {
	var ()
	return &PostCategorySearchParams{

		Context: ctx,
	}
}

// NewPostCategorySearchParamsWithHTTPClient creates a new PostCategorySearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCategorySearchParamsWithHTTPClient(client *http.Client) *PostCategorySearchParams {
	var ()
	return &PostCategorySearchParams{
		HTTPClient: client,
	}
}

/*PostCategorySearchParams contains all the parameters to send to the API endpoint
for the post category search operation typically these are written to a http.Request
*/
type PostCategorySearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*Levels*/
	Levels *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post category search params
func (o *PostCategorySearchParams) WithTimeout(timeout time.Duration) *PostCategorySearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post category search params
func (o *PostCategorySearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post category search params
func (o *PostCategorySearchParams) WithContext(ctx context.Context) *PostCategorySearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post category search params
func (o *PostCategorySearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post category search params
func (o *PostCategorySearchParams) WithHTTPClient(client *http.Client) *PostCategorySearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post category search params
func (o *PostCategorySearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post category search params
func (o *PostCategorySearchParams) WithBody(body *models.SearchRequest) *PostCategorySearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post category search params
func (o *PostCategorySearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithLevels adds the levels to the post category search params
func (o *PostCategorySearchParams) WithLevels(levels *int32) *PostCategorySearchParams {
	o.SetLevels(levels)
	return o
}

// SetLevels adds the levels to the post category search params
func (o *PostCategorySearchParams) SetLevels(levels *int32) {
	o.Levels = levels
}

// WriteToRequest writes these params to a swagger request
func (o *PostCategorySearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
