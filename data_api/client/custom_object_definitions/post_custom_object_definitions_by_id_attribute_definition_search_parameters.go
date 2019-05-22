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

// NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams creates a new PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams object
// with the default values initialized.
func NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams() *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	var ()
	return &PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithTimeout creates a new PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithTimeout(timeout time.Duration) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	var ()
	return &PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams{

		timeout: timeout,
	}
}

// NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithContext creates a new PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithContext(ctx context.Context) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	var ()
	return &PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams{

		Context: ctx,
	}
}

// NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithHTTPClient creates a new PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomObjectDefinitionsByIDAttributeDefinitionSearchParamsWithHTTPClient(client *http.Client) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	var ()
	return &PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams{
		HTTPClient: client,
	}
}

/*PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams contains all the parameters to send to the API endpoint
for the post custom object definitions by ID attribute definition search operation typically these are written to a http.Request
*/
type PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams struct {

	/*Body*/
	Body *models.SearchRequest
	/*ObjectType
	  The type of object

	*/
	ObjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WithTimeout(timeout time.Duration) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WithContext(ctx context.Context) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WithHTTPClient(client *http.Client) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WithBody(body *models.SearchRequest) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithObjectType adds the objectType to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WithObjectType(objectType string) *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the post custom object definitions by ID attribute definition search params
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomObjectDefinitionsByIDAttributeDefinitionSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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