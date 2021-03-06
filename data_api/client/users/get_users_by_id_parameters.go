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
)

// NewGetUsersByIDParams creates a new GetUsersByIDParams object
// with the default values initialized.
func NewGetUsersByIDParams() *GetUsersByIDParams {
	var ()
	return &GetUsersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersByIDParamsWithTimeout creates a new GetUsersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersByIDParamsWithTimeout(timeout time.Duration) *GetUsersByIDParams {
	var ()
	return &GetUsersByIDParams{

		timeout: timeout,
	}
}

// NewGetUsersByIDParamsWithContext creates a new GetUsersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersByIDParamsWithContext(ctx context.Context) *GetUsersByIDParams {
	var ()
	return &GetUsersByIDParams{

		Context: ctx,
	}
}

// NewGetUsersByIDParamsWithHTTPClient creates a new GetUsersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersByIDParamsWithHTTPClient(client *http.Client) *GetUsersByIDParams {
	var ()
	return &GetUsersByIDParams{
		HTTPClient: client,
	}
}

/*GetUsersByIDParams contains all the parameters to send to the API endpoint
for the get users by ID operation typically these are written to a http.Request
*/
type GetUsersByIDParams struct {

	/*Login
	  login of the user

	*/
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users by ID params
func (o *GetUsersByIDParams) WithTimeout(timeout time.Duration) *GetUsersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users by ID params
func (o *GetUsersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users by ID params
func (o *GetUsersByIDParams) WithContext(ctx context.Context) *GetUsersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users by ID params
func (o *GetUsersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users by ID params
func (o *GetUsersByIDParams) WithHTTPClient(client *http.Client) *GetUsersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users by ID params
func (o *GetUsersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the get users by ID params
func (o *GetUsersByIDParams) WithLogin(login string) *GetUsersByIDParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get users by ID params
func (o *GetUsersByIDParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param login
	if err := r.SetPathParam("login", o.Login); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
