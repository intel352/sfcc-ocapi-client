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

// NewPutBasketsByIDAgentParams creates a new PutBasketsByIDAgentParams object
// with the default values initialized.
func NewPutBasketsByIDAgentParams() *PutBasketsByIDAgentParams {
	var ()
	return &PutBasketsByIDAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBasketsByIDAgentParamsWithTimeout creates a new PutBasketsByIDAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBasketsByIDAgentParamsWithTimeout(timeout time.Duration) *PutBasketsByIDAgentParams {
	var ()
	return &PutBasketsByIDAgentParams{

		timeout: timeout,
	}
}

// NewPutBasketsByIDAgentParamsWithContext creates a new PutBasketsByIDAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBasketsByIDAgentParamsWithContext(ctx context.Context) *PutBasketsByIDAgentParams {
	var ()
	return &PutBasketsByIDAgentParams{

		Context: ctx,
	}
}

// NewPutBasketsByIDAgentParamsWithHTTPClient creates a new PutBasketsByIDAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBasketsByIDAgentParamsWithHTTPClient(client *http.Client) *PutBasketsByIDAgentParams {
	var ()
	return &PutBasketsByIDAgentParams{
		HTTPClient: client,
	}
}

/*PutBasketsByIDAgentParams contains all the parameters to send to the API endpoint
for the put baskets by ID agent operation typically these are written to a http.Request
*/
type PutBasketsByIDAgentParams struct {

	/*BasketID
	  the basket id

	*/
	BasketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) WithTimeout(timeout time.Duration) *PutBasketsByIDAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) WithContext(ctx context.Context) *PutBasketsByIDAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) WithHTTPClient(client *http.Client) *PutBasketsByIDAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) WithBasketID(basketID string) *PutBasketsByIDAgentParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the put baskets by ID agent params
func (o *PutBasketsByIDAgentParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WriteToRequest writes these params to a swagger request
func (o *PutBasketsByIDAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param basket_id
	if err := r.SetPathParam("basket_id", o.BasketID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}