// Code generated by go-swagger; DO NOT EDIT.

package orders

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

// NewDeleteOrdersByIDPaymentInstrumentsByIDParams creates a new DeleteOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized.
func NewDeleteOrdersByIDPaymentInstrumentsByIDParams() *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteOrdersByIDPaymentInstrumentsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithTimeout creates a new DeleteOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithTimeout(timeout time.Duration) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteOrdersByIDPaymentInstrumentsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithContext creates a new DeleteOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithContext(ctx context.Context) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteOrdersByIDPaymentInstrumentsByIDParams{

		Context: ctx,
	}
}

// NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithHTTPClient creates a new DeleteOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOrdersByIDPaymentInstrumentsByIDParamsWithHTTPClient(client *http.Client) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteOrdersByIDPaymentInstrumentsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteOrdersByIDPaymentInstrumentsByIDParams contains all the parameters to send to the API endpoint
for the delete orders by ID payment instruments by ID operation typically these are written to a http.Request
*/
type DeleteOrdersByIDPaymentInstrumentsByIDParams struct {

	/*OrderNo
	  the order number

	*/
	OrderNo string
	/*PaymentInstrumentID
	  the id of the payment instrument to be updated

	*/
	PaymentInstrumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WithTimeout(timeout time.Duration) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WithContext(ctx context.Context) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WithHTTPClient(client *http.Client) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderNo adds the orderNo to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WithOrderNo(orderNo string) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) SetOrderNo(orderNo string) {
	o.OrderNo = orderNo
}

// WithPaymentInstrumentID adds the paymentInstrumentID to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WithPaymentInstrumentID(paymentInstrumentID string) *DeleteOrdersByIDPaymentInstrumentsByIDParams {
	o.SetPaymentInstrumentID(paymentInstrumentID)
	return o
}

// SetPaymentInstrumentID adds the paymentInstrumentId to the delete orders by ID payment instruments by ID params
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) SetPaymentInstrumentID(paymentInstrumentID string) {
	o.PaymentInstrumentID = paymentInstrumentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrdersByIDPaymentInstrumentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param order_no
	if err := r.SetPathParam("order_no", o.OrderNo); err != nil {
		return err
	}

	// path param payment_instrument_id
	if err := r.SetPathParam("payment_instrument_id", o.PaymentInstrumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
