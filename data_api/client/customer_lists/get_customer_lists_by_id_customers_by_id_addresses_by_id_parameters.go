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
)

// NewGetCustomerListsByIDCustomersByIDAddressesByIDParams creates a new GetCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized.
func NewGetCustomerListsByIDCustomersByIDAddressesByIDParams() *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &GetCustomerListsByIDCustomersByIDAddressesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithTimeout creates a new GetCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithTimeout(timeout time.Duration) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &GetCustomerListsByIDCustomersByIDAddressesByIDParams{

		timeout: timeout,
	}
}

// NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithContext creates a new GetCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithContext(ctx context.Context) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &GetCustomerListsByIDCustomersByIDAddressesByIDParams{

		Context: ctx,
	}
}

// NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithHTTPClient creates a new GetCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomerListsByIDCustomersByIDAddressesByIDParamsWithHTTPClient(client *http.Client) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &GetCustomerListsByIDCustomersByIDAddressesByIDParams{
		HTTPClient: client,
	}
}

/*GetCustomerListsByIDCustomersByIDAddressesByIDParams contains all the parameters to send to the API endpoint
for the get customer lists by ID customers by ID addresses by ID operation typically these are written to a http.Request
*/
type GetCustomerListsByIDCustomersByIDAddressesByIDParams struct {

	/*AddressID
	  The address id

	*/
	AddressID string
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

// WithTimeout adds the timeout to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithTimeout(timeout time.Duration) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithContext(ctx context.Context) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithHTTPClient(client *http.Client) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithAddressID(addressID string) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WithCustomerNo adds the customerNo to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithCustomerNo(customerNo string) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetCustomerNo(customerNo)
	return o
}

// SetCustomerNo adds the customerNo to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetCustomerNo(customerNo string) {
	o.CustomerNo = customerNo
}

// WithListID adds the listID to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WithListID(listID string) *GetCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the get customer lists by ID customers by ID addresses by ID params
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerListsByIDCustomersByIDAddressesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address_id
	if err := r.SetPathParam("address_id", o.AddressID); err != nil {
		return err
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