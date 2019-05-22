// Code generated by go-swagger; DO NOT EDIT.

package system_object_definitions

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

// NewGetSystemObjectDefinitionsByIDParams creates a new GetSystemObjectDefinitionsByIDParams object
// with the default values initialized.
func NewGetSystemObjectDefinitionsByIDParams() *GetSystemObjectDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemObjectDefinitionsByIDParamsWithTimeout creates a new GetSystemObjectDefinitionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemObjectDefinitionsByIDParamsWithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDParams{

		timeout: timeout,
	}
}

// NewGetSystemObjectDefinitionsByIDParamsWithContext creates a new GetSystemObjectDefinitionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemObjectDefinitionsByIDParamsWithContext(ctx context.Context) *GetSystemObjectDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDParams{

		Context: ctx,
	}
}

// NewGetSystemObjectDefinitionsByIDParamsWithHTTPClient creates a new GetSystemObjectDefinitionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemObjectDefinitionsByIDParamsWithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDParams{
		HTTPClient: client,
	}
}

/*GetSystemObjectDefinitionsByIDParams contains all the parameters to send to the API endpoint
for the get system object definitions by ID operation typically these are written to a http.Request
*/
type GetSystemObjectDefinitionsByIDParams struct {

	/*ObjectType
	  The id of the object type for the requested system object.

	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) WithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) WithContext(ctx context.Context) *GetSystemObjectDefinitionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) WithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithObjectType adds the objectType to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) WithObjectType(objectType string) *GetSystemObjectDefinitionsByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get system object definitions by ID params
func (o *GetSystemObjectDefinitionsByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemObjectDefinitionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param object_type
	if err := r.SetPathParam("object_type", o.ObjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}