// Code generated by go-swagger; DO NOT EDIT.

package sites

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

// NewDeleteSitesByIDCustomObjectsByIDByIDParams creates a new DeleteSitesByIDCustomObjectsByIDByIDParams object
// with the default values initialized.
func NewDeleteSitesByIDCustomObjectsByIDByIDParams() *DeleteSitesByIDCustomObjectsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCustomObjectsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithTimeout creates a new DeleteSitesByIDCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithTimeout(timeout time.Duration) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCustomObjectsByIDByIDParams{

		timeout: timeout,
	}
}

// NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithContext creates a new DeleteSitesByIDCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithContext(ctx context.Context) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCustomObjectsByIDByIDParams{

		Context: ctx,
	}
}

// NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithHTTPClient creates a new DeleteSitesByIDCustomObjectsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSitesByIDCustomObjectsByIDByIDParamsWithHTTPClient(client *http.Client) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	var ()
	return &DeleteSitesByIDCustomObjectsByIDByIDParams{
		HTTPClient: client,
	}
}

/*DeleteSitesByIDCustomObjectsByIDByIDParams contains all the parameters to send to the API endpoint
for the delete sites by ID custom objects by ID by ID operation typically these are written to a http.Request
*/
type DeleteSitesByIDCustomObjectsByIDByIDParams struct {

	/*Key
	  the key attribute value of the Custom Object

	*/
	Key string
	/*ObjectType
	  the ID of the object type

	*/
	ObjectType string
	/*SiteID
	  the ID of the site

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithTimeout(timeout time.Duration) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithContext(ctx context.Context) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithHTTPClient(client *http.Client) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKey adds the key to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithKey(key string) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetKey(key string) {
	o.Key = key
}

// WithObjectType adds the objectType to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithObjectType(objectType string) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WithSiteID adds the siteID to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WithSiteID(siteID string) *DeleteSitesByIDCustomObjectsByIDByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete sites by ID custom objects by ID by ID params
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSitesByIDCustomObjectsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	// path param object_type
	if err := r.SetPathParam("object_type", o.ObjectType); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
