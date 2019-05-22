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

// NewPutUsersByIDParams creates a new PutUsersByIDParams object
// with the default values initialized.
func NewPutUsersByIDParams() *PutUsersByIDParams {
	var ()
	return &PutUsersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUsersByIDParamsWithTimeout creates a new PutUsersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUsersByIDParamsWithTimeout(timeout time.Duration) *PutUsersByIDParams {
	var ()
	return &PutUsersByIDParams{

		timeout: timeout,
	}
}

// NewPutUsersByIDParamsWithContext creates a new PutUsersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUsersByIDParamsWithContext(ctx context.Context) *PutUsersByIDParams {
	var ()
	return &PutUsersByIDParams{

		Context: ctx,
	}
}

// NewPutUsersByIDParamsWithHTTPClient creates a new PutUsersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUsersByIDParamsWithHTTPClient(client *http.Client) *PutUsersByIDParams {
	var ()
	return &PutUsersByIDParams{
		HTTPClient: client,
	}
}

/*PutUsersByIDParams contains all the parameters to send to the API endpoint
for the put users by ID operation typically these are written to a http.Request
*/
type PutUsersByIDParams struct {

	/*Body*/
	Body *models.User
	/*Login
	  login of the user

	*/
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put users by ID params
func (o *PutUsersByIDParams) WithTimeout(timeout time.Duration) *PutUsersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put users by ID params
func (o *PutUsersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put users by ID params
func (o *PutUsersByIDParams) WithContext(ctx context.Context) *PutUsersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put users by ID params
func (o *PutUsersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put users by ID params
func (o *PutUsersByIDParams) WithHTTPClient(client *http.Client) *PutUsersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put users by ID params
func (o *PutUsersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put users by ID params
func (o *PutUsersByIDParams) WithBody(body *models.User) *PutUsersByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put users by ID params
func (o *PutUsersByIDParams) SetBody(body *models.User) {
	o.Body = body
}

// WithLogin adds the login to the put users by ID params
func (o *PutUsersByIDParams) WithLogin(login string) *PutUsersByIDParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the put users by ID params
func (o *PutUsersByIDParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *PutUsersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param login
	if err := r.SetPathParam("login", o.Login); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
