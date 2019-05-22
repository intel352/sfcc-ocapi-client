// Code generated by go-swagger; DO NOT EDIT.

package customers

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

// NewGetCustomersByIDProductListsByIDItemsByIDParams creates a new GetCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized.
func NewGetCustomersByIDProductListsByIDItemsByIDParams() *GetCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDParamsWithTimeout creates a new GetCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersByIDProductListsByIDItemsByIDParamsWithTimeout(timeout time.Duration) *GetCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDParams{

		timeout: timeout,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDParamsWithContext creates a new GetCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersByIDProductListsByIDItemsByIDParamsWithContext(ctx context.Context) *GetCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDParams{

		Context: ctx,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDParamsWithHTTPClient creates a new GetCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersByIDProductListsByIDItemsByIDParamsWithHTTPClient(client *http.Client) *GetCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDParams{
		HTTPClient: client,
	}
}

/*GetCustomersByIDProductListsByIDItemsByIDParams contains all the parameters to send to the API endpoint
for the get customers by ID product lists by ID items by ID operation typically these are written to a http.Request
*/
type GetCustomersByIDProductListsByIDItemsByIDParams struct {

	/*CustomerID
	  The id of the customer to retrieve the product list items for.

	*/
	CustomerID string
	/*Expand*/
	Expand []string
	/*ItemID
	  The id of the product list item to retrieve.

	*/
	ItemID string
	/*ListID
	  The id of the product list.

	*/
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithTimeout(timeout time.Duration) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithContext(ctx context.Context) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithHTTPClient(client *http.Client) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithCustomerID(customerID string) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithExpand adds the expand to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithExpand(expand []string) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithItemID adds the itemID to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithItemID(itemID string) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithListID adds the listID to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WithListID(listID string) *GetCustomersByIDProductListsByIDItemsByIDParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the get customers by ID product lists by ID items by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersByIDProductListsByIDItemsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
		return err
	}

	valuesExpand := o.Expand

	joinedExpand := swag.JoinByFormat(valuesExpand, "")
	// query array param expand
	if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
		return err
	}

	// path param item_id
	if err := r.SetPathParam("item_id", o.ItemID); err != nil {
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
