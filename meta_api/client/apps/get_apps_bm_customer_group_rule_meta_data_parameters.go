// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewGetAppsBmCustomerGroupRuleMetaDataParams creates a new GetAppsBmCustomerGroupRuleMetaDataParams object
// with the default values initialized.
func NewGetAppsBmCustomerGroupRuleMetaDataParams() *GetAppsBmCustomerGroupRuleMetaDataParams {

	return &GetAppsBmCustomerGroupRuleMetaDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsBmCustomerGroupRuleMetaDataParamsWithTimeout creates a new GetAppsBmCustomerGroupRuleMetaDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsBmCustomerGroupRuleMetaDataParamsWithTimeout(timeout time.Duration) *GetAppsBmCustomerGroupRuleMetaDataParams {

	return &GetAppsBmCustomerGroupRuleMetaDataParams{

		timeout: timeout,
	}
}

// NewGetAppsBmCustomerGroupRuleMetaDataParamsWithContext creates a new GetAppsBmCustomerGroupRuleMetaDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsBmCustomerGroupRuleMetaDataParamsWithContext(ctx context.Context) *GetAppsBmCustomerGroupRuleMetaDataParams {

	return &GetAppsBmCustomerGroupRuleMetaDataParams{

		Context: ctx,
	}
}

// NewGetAppsBmCustomerGroupRuleMetaDataParamsWithHTTPClient creates a new GetAppsBmCustomerGroupRuleMetaDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppsBmCustomerGroupRuleMetaDataParamsWithHTTPClient(client *http.Client) *GetAppsBmCustomerGroupRuleMetaDataParams {

	return &GetAppsBmCustomerGroupRuleMetaDataParams{
		HTTPClient: client,
	}
}

/*GetAppsBmCustomerGroupRuleMetaDataParams contains all the parameters to send to the API endpoint
for the get apps bm customer group rule meta data operation typically these are written to a http.Request
*/
type GetAppsBmCustomerGroupRuleMetaDataParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) WithTimeout(timeout time.Duration) *GetAppsBmCustomerGroupRuleMetaDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) WithContext(ctx context.Context) *GetAppsBmCustomerGroupRuleMetaDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) WithHTTPClient(client *http.Client) *GetAppsBmCustomerGroupRuleMetaDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apps bm customer group rule meta data params
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsBmCustomerGroupRuleMetaDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
