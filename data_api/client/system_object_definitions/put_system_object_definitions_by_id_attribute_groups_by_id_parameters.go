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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParams creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized.
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParams() *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithTimeout creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithTimeout(timeout time.Duration) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithContext creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithContext(ctx context.Context) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams{

		Context: ctx,
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithHTTPClient creates a new PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParamsWithHTTPClient(client *http.Client) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	var ()
	return &PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams contains all the parameters to send to the API endpoint
for the put system object definitions by ID attribute groups by ID operation typically these are written to a http.Request
*/
type PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams struct {

	/*Body*/
	Body *models.ObjectAttributeGroup
	/*ID
	  The id of the attribute group to create.

	*/
	ID string
	/*ObjectType
	  the object type id that contains this attribute gorup

	*/
	ObjectType string
	/*XDwValidateExisting*/
	XDwValidateExisting *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithTimeout(timeout time.Duration) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithContext(ctx context.Context) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithHTTPClient(client *http.Client) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithBody(body *models.ObjectAttributeGroup) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetBody(body *models.ObjectAttributeGroup) {
	o.Body = body
}

// WithID adds the id to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithID(id string) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetID(id string) {
	o.ID = id
}

// WithObjectType adds the objectType to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithObjectType(objectType string) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WithXDwValidateExisting adds the xDwValidateExisting to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WithXDwValidateExisting(xDwValidateExisting *bool) *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams {
	o.SetXDwValidateExisting(xDwValidateExisting)
	return o
}

// SetXDwValidateExisting adds the xDwValidateExisting to the put system object definitions by ID attribute groups by ID params
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) SetXDwValidateExisting(xDwValidateExisting *bool) {
	o.XDwValidateExisting = xDwValidateExisting
}

// WriteToRequest writes these params to a swagger request
func (o *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.XDwValidateExisting != nil {

		// header param x-dw-validate-existing
		if err := r.SetHeaderParam("x-dw-validate-existing", swag.FormatBool(*o.XDwValidateExisting)); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
