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

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPostBasketsByIDShipmentsParams creates a new PostBasketsByIDShipmentsParams object
// with the default values initialized.
func NewPostBasketsByIDShipmentsParams() *PostBasketsByIDShipmentsParams {
	var ()
	return &PostBasketsByIDShipmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBasketsByIDShipmentsParamsWithTimeout creates a new PostBasketsByIDShipmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBasketsByIDShipmentsParamsWithTimeout(timeout time.Duration) *PostBasketsByIDShipmentsParams {
	var ()
	return &PostBasketsByIDShipmentsParams{

		timeout: timeout,
	}
}

// NewPostBasketsByIDShipmentsParamsWithContext creates a new PostBasketsByIDShipmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBasketsByIDShipmentsParamsWithContext(ctx context.Context) *PostBasketsByIDShipmentsParams {
	var ()
	return &PostBasketsByIDShipmentsParams{

		Context: ctx,
	}
}

// NewPostBasketsByIDShipmentsParamsWithHTTPClient creates a new PostBasketsByIDShipmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBasketsByIDShipmentsParamsWithHTTPClient(client *http.Client) *PostBasketsByIDShipmentsParams {
	var ()
	return &PostBasketsByIDShipmentsParams{
		HTTPClient: client,
	}
}

/*PostBasketsByIDShipmentsParams contains all the parameters to send to the API endpoint
for the post baskets by ID shipments operation typically these are written to a http.Request
*/
type PostBasketsByIDShipmentsParams struct {

	/*BasketID
	  the id of the basket to be modified

	*/
	BasketID string
	/*Body*/
	Body *models.Shipment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) WithTimeout(timeout time.Duration) *PostBasketsByIDShipmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) WithContext(ctx context.Context) *PostBasketsByIDShipmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) WithHTTPClient(client *http.Client) *PostBasketsByIDShipmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) WithBasketID(basketID string) *PostBasketsByIDShipmentsParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WithBody adds the body to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) WithBody(body *models.Shipment) *PostBasketsByIDShipmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post baskets by ID shipments params
func (o *PostBasketsByIDShipmentsParams) SetBody(body *models.Shipment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBasketsByIDShipmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param basket_id
	if err := r.SetPathParam("basket_id", o.BasketID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}