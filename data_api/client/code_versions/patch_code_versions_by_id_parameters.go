// Code generated by go-swagger; DO NOT EDIT.

package code_versions

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

// NewPatchCodeVersionsByIDParams creates a new PatchCodeVersionsByIDParams object
// with the default values initialized.
func NewPatchCodeVersionsByIDParams() *PatchCodeVersionsByIDParams {
	var ()
	return &PatchCodeVersionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCodeVersionsByIDParamsWithTimeout creates a new PatchCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCodeVersionsByIDParamsWithTimeout(timeout time.Duration) *PatchCodeVersionsByIDParams {
	var ()
	return &PatchCodeVersionsByIDParams{

		timeout: timeout,
	}
}

// NewPatchCodeVersionsByIDParamsWithContext creates a new PatchCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCodeVersionsByIDParamsWithContext(ctx context.Context) *PatchCodeVersionsByIDParams {
	var ()
	return &PatchCodeVersionsByIDParams{

		Context: ctx,
	}
}

// NewPatchCodeVersionsByIDParamsWithHTTPClient creates a new PatchCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCodeVersionsByIDParamsWithHTTPClient(client *http.Client) *PatchCodeVersionsByIDParams {
	var ()
	return &PatchCodeVersionsByIDParams{
		HTTPClient: client,
	}
}

/*PatchCodeVersionsByIDParams contains all the parameters to send to the API endpoint
for the patch code versions by ID operation typically these are written to a http.Request
*/
type PatchCodeVersionsByIDParams struct {

	/*Body*/
	Body *models.CodeVersion
	/*CodeVersionID
	  The id of the code version.

	*/
	CodeVersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) WithTimeout(timeout time.Duration) *PatchCodeVersionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) WithContext(ctx context.Context) *PatchCodeVersionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) WithHTTPClient(client *http.Client) *PatchCodeVersionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) WithBody(body *models.CodeVersion) *PatchCodeVersionsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) SetBody(body *models.CodeVersion) {
	o.Body = body
}

// WithCodeVersionID adds the codeVersionID to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) WithCodeVersionID(codeVersionID string) *PatchCodeVersionsByIDParams {
	o.SetCodeVersionID(codeVersionID)
	return o
}

// SetCodeVersionID adds the codeVersionId to the patch code versions by ID params
func (o *PatchCodeVersionsByIDParams) SetCodeVersionID(codeVersionID string) {
	o.CodeVersionID = codeVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCodeVersionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param code_version_id
	if err := r.SetPathParam("code_version_id", o.CodeVersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
