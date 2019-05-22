// Code generated by go-swagger; DO NOT EDIT.

package custom_objects

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

// NewPatchCustomObjectsByIDByIDParams creates a new PatchCustomObjectsByIDByIDParams object
// with the default values initialized.
func NewPatchCustomObjectsByIDByIDParams() *PatchCustomObjectsByIDByIDParams {
	var ()
	return &PatchCustomObjectsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCustomObjectsByIDByIDParamsWithTimeout creates a new PatchCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCustomObjectsByIDByIDParamsWithTimeout(timeout time.Duration) *PatchCustomObjectsByIDByIDParams {
	var ()
	return &PatchCustomObjectsByIDByIDParams{

		timeout: timeout,
	}
}

// NewPatchCustomObjectsByIDByIDParamsWithContext creates a new PatchCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCustomObjectsByIDByIDParamsWithContext(ctx context.Context) *PatchCustomObjectsByIDByIDParams {
	var ()
	return &PatchCustomObjectsByIDByIDParams{

		Context: ctx,
	}
}

// NewPatchCustomObjectsByIDByIDParamsWithHTTPClient creates a new PatchCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCustomObjectsByIDByIDParamsWithHTTPClient(client *http.Client) *PatchCustomObjectsByIDByIDParams {
	var ()
	return &PatchCustomObjectsByIDByIDParams{
		HTTPClient: client,
	}
}

/*PatchCustomObjectsByIDByIDParams contains all the parameters to send to the API endpoint
for the patch custom objects by ID by ID operation typically these are written to a http.Request
*/
type PatchCustomObjectsByIDByIDParams struct {

	/*Body*/
	Body *models.CustomObject
	/*Key
	  the key attribute value of the Custom Object

	*/
	Key string
	/*ObjectType
	  the ID of the object type

	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithTimeout(timeout time.Duration) *PatchCustomObjectsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithContext(ctx context.Context) *PatchCustomObjectsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithHTTPClient(client *http.Client) *PatchCustomObjectsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithBody(body *models.CustomObject) *PatchCustomObjectsByIDByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetBody(body *models.CustomObject) {
	o.Body = body
}

// WithKey adds the key to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithKey(key string) *PatchCustomObjectsByIDByIDParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetKey(key string) {
	o.Key = key
}

// WithObjectType adds the objectType to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) WithObjectType(objectType string) *PatchCustomObjectsByIDByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the patch custom objects by ID by ID params
func (o *PatchCustomObjectsByIDByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCustomObjectsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
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