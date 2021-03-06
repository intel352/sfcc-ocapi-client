// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewPostRolesByIDUserSearchParams creates a new PostRolesByIDUserSearchParams object
// with the default values initialized.
func NewPostRolesByIDUserSearchParams() *PostRolesByIDUserSearchParams {
	var ()
	return &PostRolesByIDUserSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRolesByIDUserSearchParamsWithTimeout creates a new PostRolesByIDUserSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRolesByIDUserSearchParamsWithTimeout(timeout time.Duration) *PostRolesByIDUserSearchParams {
	var ()
	return &PostRolesByIDUserSearchParams{

		timeout: timeout,
	}
}

// NewPostRolesByIDUserSearchParamsWithContext creates a new PostRolesByIDUserSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRolesByIDUserSearchParamsWithContext(ctx context.Context) *PostRolesByIDUserSearchParams {
	var ()
	return &PostRolesByIDUserSearchParams{

		Context: ctx,
	}
}

// NewPostRolesByIDUserSearchParamsWithHTTPClient creates a new PostRolesByIDUserSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRolesByIDUserSearchParamsWithHTTPClient(client *http.Client) *PostRolesByIDUserSearchParams {
	var ()
	return &PostRolesByIDUserSearchParams{
		HTTPClient: client,
	}
}

/*PostRolesByIDUserSearchParams contains all the parameters to send to the API endpoint
for the post roles by ID user search operation typically these are written to a http.Request
*/
type PostRolesByIDUserSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) WithTimeout(timeout time.Duration) *PostRolesByIDUserSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) WithContext(ctx context.Context) *PostRolesByIDUserSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) WithHTTPClient(client *http.Client) *PostRolesByIDUserSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) WithBody(body *models.SearchRequest) *PostRolesByIDUserSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithID adds the id to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) WithID(id string) *PostRolesByIDUserSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post roles by ID user search params
func (o *PostRolesByIDUserSearchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRolesByIDUserSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
