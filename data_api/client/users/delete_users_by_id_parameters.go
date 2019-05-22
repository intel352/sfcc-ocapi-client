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

// NewDeleteUsersByIDParams creates a new DeleteUsersByIDParams object
// with the default values initialized.
func NewDeleteUsersByIDParams() *DeleteUsersByIDParams {
	var ()
	return &DeleteUsersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersByIDParamsWithTimeout creates a new DeleteUsersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUsersByIDParamsWithTimeout(timeout time.Duration) *DeleteUsersByIDParams {
	var ()
	return &DeleteUsersByIDParams{

		timeout: timeout,
	}
}

// NewDeleteUsersByIDParamsWithContext creates a new DeleteUsersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUsersByIDParamsWithContext(ctx context.Context) *DeleteUsersByIDParams {
	var ()
	return &DeleteUsersByIDParams{

		Context: ctx,
	}
}

// NewDeleteUsersByIDParamsWithHTTPClient creates a new DeleteUsersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUsersByIDParamsWithHTTPClient(client *http.Client) *DeleteUsersByIDParams {
	var ()
	return &DeleteUsersByIDParams{
		HTTPClient: client,
	}
}

/*DeleteUsersByIDParams contains all the parameters to send to the API endpoint
for the delete users by ID operation typically these are written to a http.Request
*/
type DeleteUsersByIDParams struct {

	/*Login
	  login of the user

	*/
	Login string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete users by ID params
func (o *DeleteUsersByIDParams) WithTimeout(timeout time.Duration) *DeleteUsersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users by ID params
func (o *DeleteUsersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users by ID params
func (o *DeleteUsersByIDParams) WithContext(ctx context.Context) *DeleteUsersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users by ID params
func (o *DeleteUsersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users by ID params
func (o *DeleteUsersByIDParams) WithHTTPClient(client *http.Client) *DeleteUsersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users by ID params
func (o *DeleteUsersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the delete users by ID params
func (o *DeleteUsersByIDParams) WithLogin(login string) *DeleteUsersByIDParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the delete users by ID params
func (o *DeleteUsersByIDParams) SetLogin(login string) {
	o.Login = login
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
