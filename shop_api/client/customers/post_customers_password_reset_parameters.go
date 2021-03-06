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

// NewPostCustomersPasswordResetParams creates a new PostCustomersPasswordResetParams object
// with the default values initialized.
func NewPostCustomersPasswordResetParams() *PostCustomersPasswordResetParams {
	var ()
	return &PostCustomersPasswordResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersPasswordResetParamsWithTimeout creates a new PostCustomersPasswordResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersPasswordResetParamsWithTimeout(timeout time.Duration) *PostCustomersPasswordResetParams {
	var ()
	return &PostCustomersPasswordResetParams{

		timeout: timeout,
	}
}

// NewPostCustomersPasswordResetParamsWithContext creates a new PostCustomersPasswordResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersPasswordResetParamsWithContext(ctx context.Context) *PostCustomersPasswordResetParams {
	var ()
	return &PostCustomersPasswordResetParams{

		Context: ctx,
	}
}

// NewPostCustomersPasswordResetParamsWithHTTPClient creates a new PostCustomersPasswordResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersPasswordResetParamsWithHTTPClient(client *http.Client) *PostCustomersPasswordResetParams {
	var ()
	return &PostCustomersPasswordResetParams{
		HTTPClient: client,
	}
}

/*PostCustomersPasswordResetParams contains all the parameters to send to the API endpoint
for the post customers password reset operation typically these are written to a http.Request
*/
type PostCustomersPasswordResetParams struct {

	/*Body*/
	Body *models.PasswordReset

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers password reset params
func (o *PostCustomersPasswordResetParams) WithTimeout(timeout time.Duration) *PostCustomersPasswordResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers password reset params
func (o *PostCustomersPasswordResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers password reset params
func (o *PostCustomersPasswordResetParams) WithContext(ctx context.Context) *PostCustomersPasswordResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers password reset params
func (o *PostCustomersPasswordResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers password reset params
func (o *PostCustomersPasswordResetParams) WithHTTPClient(client *http.Client) *PostCustomersPasswordResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers password reset params
func (o *PostCustomersPasswordResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customers password reset params
func (o *PostCustomersPasswordResetParams) WithBody(body *models.PasswordReset) *PostCustomersPasswordResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customers password reset params
func (o *PostCustomersPasswordResetParams) SetBody(body *models.PasswordReset) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersPasswordResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
