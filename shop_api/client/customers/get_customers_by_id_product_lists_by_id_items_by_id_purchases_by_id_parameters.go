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

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams creates a new GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams object
// with the default values initialized.
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams() *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithTimeout creates a new GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithTimeout(timeout time.Duration) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams{

		timeout: timeout,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithContext creates a new GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithContext(ctx context.Context) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams{

		Context: ctx,
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithHTTPClient creates a new GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParamsWithHTTPClient(client *http.Client) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	var ()
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams{
		HTTPClient: client,
	}
}

/*GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams contains all the parameters to send to the API endpoint
for the get customers by ID product lists by ID items by ID purchases by ID operation typically these are written to a http.Request
*/
type GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams struct {

	/*CustomerID
	  The id of the customer to retrieve the product list items for.

	*/
	CustomerID string
	/*ItemID
	  The id of the product list item to retrieve.

	*/
	ItemID string
	/*ListID
	  The id of the product list.

	*/
	ListID string
	/*PurchaseID
	  The id of the product list item purchase to retrieve.

	*/
	PurchaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithTimeout(timeout time.Duration) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithContext(ctx context.Context) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithHTTPClient(client *http.Client) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithCustomerID(customerID string) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithItemID adds the itemID to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithItemID(itemID string) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithListID adds the listID to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithListID(listID string) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetListID(listID string) {
	o.ListID = listID
}

// WithPurchaseID adds the purchaseID to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WithPurchaseID(purchaseID string) *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams {
	o.SetPurchaseID(purchaseID)
	return o
}

// SetPurchaseID adds the purchaseId to the get customers by ID product lists by ID items by ID purchases by ID params
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) SetPurchaseID(purchaseID string) {
	o.PurchaseID = purchaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param purchase_id
	if err := r.SetPathParam("purchase_id", o.PurchaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}