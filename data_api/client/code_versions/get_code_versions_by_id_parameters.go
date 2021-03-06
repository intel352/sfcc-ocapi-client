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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCodeVersionsByIDParams creates a new GetCodeVersionsByIDParams object
// with the default values initialized.
func NewGetCodeVersionsByIDParams() *GetCodeVersionsByIDParams {
	var ()
	return &GetCodeVersionsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCodeVersionsByIDParamsWithTimeout creates a new GetCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCodeVersionsByIDParamsWithTimeout(timeout time.Duration) *GetCodeVersionsByIDParams {
	var ()
	return &GetCodeVersionsByIDParams{

		timeout: timeout,
	}
}

// NewGetCodeVersionsByIDParamsWithContext creates a new GetCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCodeVersionsByIDParamsWithContext(ctx context.Context) *GetCodeVersionsByIDParams {
	var ()
	return &GetCodeVersionsByIDParams{

		Context: ctx,
	}
}

// NewGetCodeVersionsByIDParamsWithHTTPClient creates a new GetCodeVersionsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCodeVersionsByIDParamsWithHTTPClient(client *http.Client) *GetCodeVersionsByIDParams {
	var ()
	return &GetCodeVersionsByIDParams{
		HTTPClient: client,
	}
}

/*GetCodeVersionsByIDParams contains all the parameters to send to the API endpoint
for the get code versions by ID operation typically these are written to a http.Request
*/
type GetCodeVersionsByIDParams struct {

	/*CodeVersionID
	  The id of the code version.

	*/
	CodeVersionID string
	/*Expand*/
	Expand []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) WithTimeout(timeout time.Duration) *GetCodeVersionsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) WithContext(ctx context.Context) *GetCodeVersionsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) WithHTTPClient(client *http.Client) *GetCodeVersionsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCodeVersionID adds the codeVersionID to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) WithCodeVersionID(codeVersionID string) *GetCodeVersionsByIDParams {
	o.SetCodeVersionID(codeVersionID)
	return o
}

// SetCodeVersionID adds the codeVersionId to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) SetCodeVersionID(codeVersionID string) {
	o.CodeVersionID = codeVersionID
}

// WithExpand adds the expand to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) WithExpand(expand []string) *GetCodeVersionsByIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get code versions by ID params
func (o *GetCodeVersionsByIDParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WriteToRequest writes these params to a swagger request
func (o *GetCodeVersionsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param code_version_id
	if err := r.SetPathParam("code_version_id", o.CodeVersionID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
