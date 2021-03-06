// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPatchUsersThisPasswordParams creates a new PatchUsersThisPasswordParams object
// with the default values initialized.
func NewPatchUsersThisPasswordParams() *PatchUsersThisPasswordParams {
	var ()
	return &PatchUsersThisPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUsersThisPasswordParamsWithTimeout creates a new PatchUsersThisPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUsersThisPasswordParamsWithTimeout(timeout time.Duration) *PatchUsersThisPasswordParams {
	var ()
	return &PatchUsersThisPasswordParams{

		timeout: timeout,
	}
}

// NewPatchUsersThisPasswordParamsWithContext creates a new PatchUsersThisPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUsersThisPasswordParamsWithContext(ctx context.Context) *PatchUsersThisPasswordParams {
	var ()
	return &PatchUsersThisPasswordParams{

		Context: ctx,
	}
}

// NewPatchUsersThisPasswordParamsWithHTTPClient creates a new PatchUsersThisPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUsersThisPasswordParamsWithHTTPClient(client *http.Client) *PatchUsersThisPasswordParams {
	var ()
	return &PatchUsersThisPasswordParams{
		HTTPClient: client,
	}
}

/*PatchUsersThisPasswordParams contains all the parameters to send to the API endpoint
for the patch users this password operation typically these are written to a http.Request
*/
type PatchUsersThisPasswordParams struct {

	/*Body*/
	Body *models.PasswordChangeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch users this password params
func (o *PatchUsersThisPasswordParams) WithTimeout(timeout time.Duration) *PatchUsersThisPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch users this password params
func (o *PatchUsersThisPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch users this password params
func (o *PatchUsersThisPasswordParams) WithContext(ctx context.Context) *PatchUsersThisPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch users this password params
func (o *PatchUsersThisPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch users this password params
func (o *PatchUsersThisPasswordParams) WithHTTPClient(client *http.Client) *PatchUsersThisPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch users this password params
func (o *PatchUsersThisPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch users this password params
func (o *PatchUsersThisPasswordParams) WithBody(body *models.PasswordChangeRequest) *PatchUsersThisPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch users this password params
func (o *PatchUsersThisPasswordParams) SetBody(body *models.PasswordChangeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUsersThisPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
