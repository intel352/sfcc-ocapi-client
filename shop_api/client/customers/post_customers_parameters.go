// Code generated by go-swagger; DO NOT EDIT.

package customers

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

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPostCustomersParams creates a new PostCustomersParams object
// with the default values initialized.
func NewPostCustomersParams() *PostCustomersParams {
	var ()
	return &PostCustomersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersParamsWithTimeout creates a new PostCustomersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersParamsWithTimeout(timeout time.Duration) *PostCustomersParams {
	var ()
	return &PostCustomersParams{

		timeout: timeout,
	}
}

// NewPostCustomersParamsWithContext creates a new PostCustomersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersParamsWithContext(ctx context.Context) *PostCustomersParams {
	var ()
	return &PostCustomersParams{

		Context: ctx,
	}
}

// NewPostCustomersParamsWithHTTPClient creates a new PostCustomersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersParamsWithHTTPClient(client *http.Client) *PostCustomersParams {
	var ()
	return &PostCustomersParams{
		HTTPClient: client,
	}
}

/*PostCustomersParams contains all the parameters to send to the API endpoint
for the post customers operation typically these are written to a http.Request
*/
type PostCustomersParams struct {

	/*Body*/
	Body *models.CustomerRegistration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers params
func (o *PostCustomersParams) WithTimeout(timeout time.Duration) *PostCustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers params
func (o *PostCustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers params
func (o *PostCustomersParams) WithContext(ctx context.Context) *PostCustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers params
func (o *PostCustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers params
func (o *PostCustomersParams) WithHTTPClient(client *http.Client) *PostCustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers params
func (o *PostCustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customers params
func (o *PostCustomersParams) WithBody(body *models.CustomerRegistration) *PostCustomersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customers params
func (o *PostCustomersParams) SetBody(body *models.CustomerRegistration) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}