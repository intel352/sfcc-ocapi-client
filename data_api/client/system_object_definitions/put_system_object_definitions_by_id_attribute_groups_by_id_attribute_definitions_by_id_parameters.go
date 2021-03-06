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

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams object
// with the default values initialized.
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams() *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithTimeout creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithTimeout(timeout time.Duration) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams{

		timeout: timeout,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithContext creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithContext(ctx context.Context) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams{

		Context: ctx,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithHTTPClient creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParamsWithHTTPClient(client *http.Client) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams{
		HTTPClient: client,
	}
}

/*PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams contains all the parameters to send to the API endpoint
for the put system object definitions by ID attribute groups by ID attribute definitions by ID operation typically these are written to a http.Request
*/
type PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams struct {

	/*DefID
	  The ID of the attribute definition.

	*/
	DefID string
	/*GroupID
	  The ID of the attribute group.

	*/
	GroupID string
	/*ObjectType
	  The ID of the system object that contains the attribute definition and attribute group.

	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithTimeout(timeout time.Duration) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithContext(ctx context.Context) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithHTTPClient(client *http.Client) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefID adds the defID to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithDefID(defID string) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetDefID(defID)
	return o
}

// SetDefID adds the defId to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetDefID(defID string) {
	o.DefID = defID
}

// WithGroupID adds the groupID to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithGroupID(groupID string) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithObjectType adds the objectType to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WithObjectType(objectType string) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the put system object definitions by ID attribute groups by ID attribute definitions by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param def_id
	if err := r.SetPathParam("def_id", o.DefID); err != nil {
		return err
	}

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	// path param object_type
	if err := r.SetPathParam("object_type", o.ObjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
