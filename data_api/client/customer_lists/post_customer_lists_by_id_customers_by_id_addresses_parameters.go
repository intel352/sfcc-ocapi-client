// Code generated by go-swagger; DO NOT EDIT.

package customer_lists

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

// NewPostCustomerListsByIDCustomersByIDAddressesParams creates a new PostCustomerListsByIDCustomersByIDAddressesParams object
// with the default values initialized.
func NewPostCustomerListsByIDCustomersByIDAddressesParams() *PostCustomerListsByIDCustomersByIDAddressesParams {
	var ()
	return &PostCustomerListsByIDCustomersByIDAddressesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomerListsByIDCustomersByIDAddressesParamsWithTimeout creates a new PostCustomerListsByIDCustomersByIDAddressesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomerListsByIDCustomersByIDAddressesParamsWithTimeout(timeout time.Duration) *PostCustomerListsByIDCustomersByIDAddressesParams {
	var ()
	return &PostCustomerListsByIDCustomersByIDAddressesParams{

		timeout: timeout,
	}
}

// NewPostCustomerListsByIDCustomersByIDAddressesParamsWithContext creates a new PostCustomerListsByIDCustomersByIDAddressesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomerListsByIDCustomersByIDAddressesParamsWithContext(ctx context.Context) *PostCustomerListsByIDCustomersByIDAddressesParams {
	var ()
	return &PostCustomerListsByIDCustomersByIDAddressesParams{

		Context: ctx,
	}
}

// NewPostCustomerListsByIDCustomersByIDAddressesParamsWithHTTPClient creates a new PostCustomerListsByIDCustomersByIDAddressesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomerListsByIDCustomersByIDAddressesParamsWithHTTPClient(client *http.Client) *PostCustomerListsByIDCustomersByIDAddressesParams {
	var ()
	return &PostCustomerListsByIDCustomersByIDAddressesParams{
		HTTPClient: client,
	}
}

/*PostCustomerListsByIDCustomersByIDAddressesParams contains all the parameters to send to the API endpoint
for the post customer lists by ID customers by ID addresses operation typically these are written to a http.Request
*/
type PostCustomerListsByIDCustomersByIDAddressesParams struct {

	/*Body*/
	Body *models.CustomerAddress
	/*CustomerNo
	  The customer number

	*/
	CustomerNo string
	/*ListID
	  The customer list id

	*/
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithTimeout(timeout time.Duration) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithContext(ctx context.Context) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithHTTPClient(client *http.Client) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithBody(body *models.CustomerAddress) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetBody(body *models.CustomerAddress) {
	o.Body = body
}

// WithCustomerNo adds the customerNo to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithCustomerNo(customerNo string) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetCustomerNo(customerNo)
	return o
}

// SetCustomerNo adds the customerNo to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetCustomerNo(customerNo string) {
	o.CustomerNo = customerNo
}

// WithListID adds the listID to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WithListID(listID string) *PostCustomerListsByIDCustomersByIDAddressesParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the post customer lists by ID customers by ID addresses params
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomerListsByIDCustomersByIDAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param customer_no
	if err := r.SetPathParam("customer_no", o.CustomerNo); err != nil {
		return err
	}

	// path param list_id
	if err := r.SetPathParam("list_id", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
