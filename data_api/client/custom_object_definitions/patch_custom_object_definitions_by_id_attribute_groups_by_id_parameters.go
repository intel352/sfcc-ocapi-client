// Code generated by go-swagger; DO NOT EDIT.

package custom_object_definitions

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

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams creates a new PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized.
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams() *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithTimeout creates a new PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithTimeout(timeout time.Duration) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithContext creates a new PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithContext(ctx context.Context) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams{

		Context: ctx,
	}
}

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithHTTPClient creates a new PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDParamsWithHTTPClient(client *http.Client) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams contains all the parameters to send to the API endpoint
for the patch custom object definitions by ID attribute groups by ID operation typically these are written to a http.Request
*/
type PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams struct {

	/*IfMatch*/
	IfMatch *string
	/*Body*/
	Body *models.ObjectAttributeGroup
	/*ID
	  The id of the requested attribute group.

	*/
	ID string
	/*ObjectType
	  the object type id that contains this attribute gorup

	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithTimeout(timeout time.Duration) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithContext(ctx context.Context) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithHTTPClient(client *http.Client) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithIfMatch(ifMatch *string) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithBody(body *models.ObjectAttributeGroup) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetBody(body *models.ObjectAttributeGroup) {
	o.Body = body
}

// WithID adds the id to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithID(id string) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetID(id string) {
	o.ID = id
}

// WithObjectType adds the objectType to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WithObjectType(objectType string) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the patch custom object definitions by ID attribute groups by ID params
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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