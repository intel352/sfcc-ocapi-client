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

// NewGetBasketsByIDApproachingDiscountsParams creates a new GetBasketsByIDApproachingDiscountsParams object
// with the default values initialized.
func NewGetBasketsByIDApproachingDiscountsParams() *GetBasketsByIDApproachingDiscountsParams {
	var ()
	return &GetBasketsByIDApproachingDiscountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBasketsByIDApproachingDiscountsParamsWithTimeout creates a new GetBasketsByIDApproachingDiscountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBasketsByIDApproachingDiscountsParamsWithTimeout(timeout time.Duration) *GetBasketsByIDApproachingDiscountsParams {
	var ()
	return &GetBasketsByIDApproachingDiscountsParams{

		timeout: timeout,
	}
}

// NewGetBasketsByIDApproachingDiscountsParamsWithContext creates a new GetBasketsByIDApproachingDiscountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBasketsByIDApproachingDiscountsParamsWithContext(ctx context.Context) *GetBasketsByIDApproachingDiscountsParams {
	var ()
	return &GetBasketsByIDApproachingDiscountsParams{

		Context: ctx,
	}
}

// NewGetBasketsByIDApproachingDiscountsParamsWithHTTPClient creates a new GetBasketsByIDApproachingDiscountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBasketsByIDApproachingDiscountsParamsWithHTTPClient(client *http.Client) *GetBasketsByIDApproachingDiscountsParams {
	var ()
	return &GetBasketsByIDApproachingDiscountsParams{
		HTTPClient: client,
	}
}

/*GetBasketsByIDApproachingDiscountsParams contains all the parameters to send to the API endpoint
for the get baskets by ID approaching discounts operation typically these are written to a http.Request
*/
type GetBasketsByIDApproachingDiscountsParams struct {

	/*BasketID
	  The id of the basket to be checked.

	*/
	BasketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) WithTimeout(timeout time.Duration) *GetBasketsByIDApproachingDiscountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) WithContext(ctx context.Context) *GetBasketsByIDApproachingDiscountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) WithHTTPClient(client *http.Client) *GetBasketsByIDApproachingDiscountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) WithBasketID(basketID string) *GetBasketsByIDApproachingDiscountsParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the get baskets by ID approaching discounts params
func (o *GetBasketsByIDApproachingDiscountsParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBasketsByIDApproachingDiscountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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