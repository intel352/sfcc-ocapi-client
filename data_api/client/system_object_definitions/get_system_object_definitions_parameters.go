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

// NewGetSystemObjectDefinitionsParams creates a new GetSystemObjectDefinitionsParams object
// with the default values initialized.
func NewGetSystemObjectDefinitionsParams() *GetSystemObjectDefinitionsParams {
	var ()
	return &GetSystemObjectDefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSystemObjectDefinitionsParamsWithTimeout creates a new GetSystemObjectDefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSystemObjectDefinitionsParamsWithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsParams {
	var ()
	return &GetSystemObjectDefinitionsParams{

		timeout: timeout,
	}
}

// NewGetSystemObjectDefinitionsParamsWithContext creates a new GetSystemObjectDefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSystemObjectDefinitionsParamsWithContext(ctx context.Context) *GetSystemObjectDefinitionsParams {
	var ()
	return &GetSystemObjectDefinitionsParams{

		Context: ctx,
	}
}

// NewGetSystemObjectDefinitionsParamsWithHTTPClient creates a new GetSystemObjectDefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSystemObjectDefinitionsParamsWithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsParams {
	var ()
	return &GetSystemObjectDefinitionsParams{
		HTTPClient: client,
	}
}

/*GetSystemObjectDefinitionsParams contains all the parameters to send to the API endpoint
for the get system object definitions operation typically these are written to a http.Request
*/
type GetSystemObjectDefinitionsParams struct {

	/*Count
	  Optional count for retrieving only a subset of the items (default is 25).

	*/
	Count *int32
	/*Select
	  The property selector.

	*/
	Select *string
	/*Start
	  Optional start index for retrieving the items from a given index (default 0).

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithTimeout(timeout time.Duration) *GetSystemObjectDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithContext(ctx context.Context) *GetSystemObjectDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithHTTPClient(client *http.Client) *GetSystemObjectDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithCount(count *int32) *GetSystemObjectDefinitionsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetCount(count *int32) {
	o.Count = count
}

// WithSelect adds the selectVar to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithSelect(selectVar *string) *GetSystemObjectDefinitionsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetSelect(selectVar *string) {
	o.Select = selectVar
}

// WithStart adds the start to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) WithStart(start *int32) *GetSystemObjectDefinitionsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get system object definitions params
func (o *GetSystemObjectDefinitionsParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetSystemObjectDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

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

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}