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

// NewGetBasketsByIDNotesParams creates a new GetBasketsByIDNotesParams object
// with the default values initialized.
func NewGetBasketsByIDNotesParams() *GetBasketsByIDNotesParams {
	var ()
	return &GetBasketsByIDNotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBasketsByIDNotesParamsWithTimeout creates a new GetBasketsByIDNotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBasketsByIDNotesParamsWithTimeout(timeout time.Duration) *GetBasketsByIDNotesParams {
	var ()
	return &GetBasketsByIDNotesParams{

		timeout: timeout,
	}
}

// NewGetBasketsByIDNotesParamsWithContext creates a new GetBasketsByIDNotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBasketsByIDNotesParamsWithContext(ctx context.Context) *GetBasketsByIDNotesParams {
	var ()
	return &GetBasketsByIDNotesParams{

		Context: ctx,
	}
}

// NewGetBasketsByIDNotesParamsWithHTTPClient creates a new GetBasketsByIDNotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBasketsByIDNotesParamsWithHTTPClient(client *http.Client) *GetBasketsByIDNotesParams {
	var ()
	return &GetBasketsByIDNotesParams{
		HTTPClient: client,
	}
}

/*GetBasketsByIDNotesParams contains all the parameters to send to the API endpoint
for the get baskets by ID notes operation typically these are written to a http.Request
*/
type GetBasketsByIDNotesParams struct {

	/*BasketID
	  The id of the basket for which you want to retrieve the notes.

	*/
	BasketID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) WithTimeout(timeout time.Duration) *GetBasketsByIDNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) WithContext(ctx context.Context) *GetBasketsByIDNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) WithHTTPClient(client *http.Client) *GetBasketsByIDNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) WithBasketID(basketID string) *GetBasketsByIDNotesParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the get baskets by ID notes params
func (o *GetBasketsByIDNotesParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBasketsByIDNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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