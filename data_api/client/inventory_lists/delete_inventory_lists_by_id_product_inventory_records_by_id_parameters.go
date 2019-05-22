// Code generated by go-swagger; DO NOT EDIT.

package inventory_lists

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

// NewDeleteInventoryListsByIDProductInventoryRecordsByIDParams creates a new DeleteInventoryListsByIDProductInventoryRecordsByIDParams object
// with the default values initialized.
func NewDeleteInventoryListsByIDProductInventoryRecordsByIDParams() *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	var ()
	return &DeleteInventoryListsByIDProductInventoryRecordsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithTimeout creates a new DeleteInventoryListsByIDProductInventoryRecordsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithTimeout(timeout time.Duration) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	var ()
	return &DeleteInventoryListsByIDProductInventoryRecordsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithContext creates a new DeleteInventoryListsByIDProductInventoryRecordsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithContext(ctx context.Context) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	var ()
	return &DeleteInventoryListsByIDProductInventoryRecordsByIDParams{

		Context: ctx,
	}
}

// NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithHTTPClient creates a new DeleteInventoryListsByIDProductInventoryRecordsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInventoryListsByIDProductInventoryRecordsByIDParamsWithHTTPClient(client *http.Client) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	var ()
	return &DeleteInventoryListsByIDProductInventoryRecordsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteInventoryListsByIDProductInventoryRecordsByIDParams contains all the parameters to send to the API endpoint
for the delete inventory lists by ID product inventory records by ID operation typically these are written to a http.Request
*/
type DeleteInventoryListsByIDProductInventoryRecordsByIDParams struct {

	/*InventoryListID
	  The inventory list ID

	*/
	InventoryListID string
	/*ProductID
	  The product ID

	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WithTimeout(timeout time.Duration) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WithContext(ctx context.Context) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WithHTTPClient(client *http.Client) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryListID adds the inventoryListID to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WithInventoryListID(inventoryListID string) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	o.SetInventoryListID(inventoryListID)
	return o
}

// SetInventoryListID adds the inventoryListId to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) SetInventoryListID(inventoryListID string) {
	o.InventoryListID = inventoryListID
}

// WithProductID adds the productID to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WithProductID(productID string) *DeleteInventoryListsByIDProductInventoryRecordsByIDParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the delete inventory lists by ID product inventory records by ID params
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInventoryListsByIDProductInventoryRecordsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inventory_list_id
	if err := r.SetPathParam("inventory_list_id", o.InventoryListID); err != nil {
		return err
	}

	// path param product_id
	if err := r.SetPathParam("product_id", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
