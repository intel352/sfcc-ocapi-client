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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCustomersByIDProductListsByIDItemsByIDParams creates a new DeleteCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized.
func NewDeleteCustomersByIDProductListsByIDItemsByIDParams() *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &DeleteCustomersByIDProductListsByIDItemsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithTimeout creates a new DeleteCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithTimeout(timeout time.Duration) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &DeleteCustomersByIDProductListsByIDItemsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithContext creates a new DeleteCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithContext(ctx context.Context) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &DeleteCustomersByIDProductListsByIDItemsByIDParams{

		Context: ctx,
	}
}

// NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithHTTPClient creates a new DeleteCustomersByIDProductListsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersByIDProductListsByIDItemsByIDParamsWithHTTPClient(client *http.Client) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	var ()
	return &DeleteCustomersByIDProductListsByIDItemsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersByIDProductListsByIDItemsByIDParams contains all the parameters to send to the API endpoint
for the delete customers by ID product lists by ID items by ID operation typically these are written to a http.Request
*/
type DeleteCustomersByIDProductListsByIDItemsByIDParams struct {

	/*CustomerID
	  The id of the owner of the product list

	*/
	CustomerID string
	/*ItemID
	  The id of the product list item to delete.

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

// WithTimeout adds the timeout to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithTimeout(timeout time.Duration) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithContext(ctx context.Context) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithHTTPClient(client *http.Client) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithCustomerID(customerID string) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithItemID adds the itemID to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithItemID(itemID string) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithListID adds the listID to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WithListID(listID string) *DeleteCustomersByIDProductListsByIDItemsByIDParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the delete customers by ID product lists by ID items by ID params
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersByIDProductListsByIDItemsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
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