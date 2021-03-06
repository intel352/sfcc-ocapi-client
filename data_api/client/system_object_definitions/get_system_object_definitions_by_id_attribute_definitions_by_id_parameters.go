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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams creates a new GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams object
// with the default values initialized.
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams() *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithTimeout creates a new GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams{

		timeout: timeout,
	}
}

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithContext creates a new GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithContext(ctx context.Context) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams{

		Context: ctx,
	}
}

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithHTTPClient creates a new GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParamsWithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	var ()
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams{
		HTTPClient: client,
	}
}

/*GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams contains all the parameters to send to the API endpoint
for the get system object definitions by ID attribute definitions by ID operation typically these are written to a http.Request
*/
type GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams struct {

	/*Expand*/
	Expand []string
	/*ID
	  The id of the requested attribute definition.

	*/
	ID string
	/*ObjectType
	  The object type id that contains these definitions

	*/
	ObjectType string
	/*Select*/
	Select *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithContext(ctx context.Context) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithExpand(expand []string) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithID adds the id to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithID(id string) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetID(id string) {
	o.ID = id
}

// WithObjectType adds the objectType to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithObjectType(objectType string) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WithSelect adds the selectVar to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WithSelect(selectVar *string) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get system object definitions by ID attribute definitions by ID params
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param object_type
	if err := r.SetPathParam("object_type", o.ObjectType); err != nil {
		return err
	}

	if o.Select != nil {

		// query param select
		var qrSelect string
		if o.Select != nil {
			qrSelect = *o.Select
		}
		qSelect := qrSelect
		if qSelect != "" {
			if err := r.SetQueryParam("select", qSelect); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
