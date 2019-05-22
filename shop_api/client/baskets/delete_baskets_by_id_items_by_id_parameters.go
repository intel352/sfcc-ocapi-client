// Code generated by go-swagger; DO NOT EDIT.

package baskets

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

// NewDeleteBasketsByIDItemsByIDParams creates a new DeleteBasketsByIDItemsByIDParams object
// with the default values initialized.
func NewDeleteBasketsByIDItemsByIDParams() *DeleteBasketsByIDItemsByIDParams {
	var ()
	return &DeleteBasketsByIDItemsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBasketsByIDItemsByIDParamsWithTimeout creates a new DeleteBasketsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBasketsByIDItemsByIDParamsWithTimeout(timeout time.Duration) *DeleteBasketsByIDItemsByIDParams {
	var ()
	return &DeleteBasketsByIDItemsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteBasketsByIDItemsByIDParamsWithContext creates a new DeleteBasketsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBasketsByIDItemsByIDParamsWithContext(ctx context.Context) *DeleteBasketsByIDItemsByIDParams {
	var ()
	return &DeleteBasketsByIDItemsByIDParams{

		Context: ctx,
	}
}

// NewDeleteBasketsByIDItemsByIDParamsWithHTTPClient creates a new DeleteBasketsByIDItemsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBasketsByIDItemsByIDParamsWithHTTPClient(client *http.Client) *DeleteBasketsByIDItemsByIDParams {
	var ()
	return &DeleteBasketsByIDItemsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteBasketsByIDItemsByIDParams contains all the parameters to send to the API endpoint
for the delete baskets by ID items by ID operation typically these are written to a http.Request
*/
type DeleteBasketsByIDItemsByIDParams struct {

	/*BasketID
	  the id of the basket to be modified

	*/
	BasketID string
	/*ItemID
	  the id of the product item to be removed

	*/
	ItemID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) WithTimeout(timeout time.Duration) *DeleteBasketsByIDItemsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) WithContext(ctx context.Context) *DeleteBasketsByIDItemsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) WithHTTPClient(client *http.Client) *DeleteBasketsByIDItemsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) WithBasketID(basketID string) *DeleteBasketsByIDItemsByIDParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WithItemID adds the itemID to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) WithItemID(itemID string) *DeleteBasketsByIDItemsByIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the delete baskets by ID items by ID params
func (o *DeleteBasketsByIDItemsByIDParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBasketsByIDItemsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param basket_id
	if err := r.SetPathParam("basket_id", o.BasketID); err != nil {
		return err
	}

	// path param item_id
	if err := r.SetPathParam("item_id", o.ItemID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
