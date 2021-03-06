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

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParams creates a new PostCustomersByIDProductListsByIDItemsByIDPurchasesParams object
// with the default values initialized.
func NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParams() *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	var ()
	return &PostCustomersByIDProductListsByIDItemsByIDPurchasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithTimeout creates a new PostCustomersByIDProductListsByIDItemsByIDPurchasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithTimeout(timeout time.Duration) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	var ()
	return &PostCustomersByIDProductListsByIDItemsByIDPurchasesParams{

		timeout: timeout,
	}
}

// NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithContext creates a new PostCustomersByIDProductListsByIDItemsByIDPurchasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithContext(ctx context.Context) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	var ()
	return &PostCustomersByIDProductListsByIDItemsByIDPurchasesParams{

		Context: ctx,
	}
}

// NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithHTTPClient creates a new PostCustomersByIDProductListsByIDItemsByIDPurchasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersByIDProductListsByIDItemsByIDPurchasesParamsWithHTTPClient(client *http.Client) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	var ()
	return &PostCustomersByIDProductListsByIDItemsByIDPurchasesParams{
		HTTPClient: client,
	}
}

/*PostCustomersByIDProductListsByIDItemsByIDPurchasesParams contains all the parameters to send to the API endpoint
for the post customers by ID product lists by ID items by ID purchases operation typically these are written to a http.Request
*/
type PostCustomersByIDProductListsByIDItemsByIDPurchasesParams struct {

	/*Body*/
	Body *models.CustomerProductListItemPurchase
	/*CustomerID
	  The id of the customer - owner of the product list.

	*/
	CustomerID string
	/*ItemID
	  The id of the product list item where to add the purchase.

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

// WithTimeout adds the timeout to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithTimeout(timeout time.Duration) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithContext(ctx context.Context) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithHTTPClient(client *http.Client) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithBody(body *models.CustomerProductListItemPurchase) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetBody(body *models.CustomerProductListItemPurchase) {
	o.Body = body
}

// WithCustomerID adds the customerID to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithCustomerID(customerID string) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithItemID adds the itemID to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithItemID(itemID string) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithListID adds the listID to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WithListID(listID string) *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the post customers by ID product lists by ID items by ID purchases params
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersByIDProductListsByIDItemsByIDPurchasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
