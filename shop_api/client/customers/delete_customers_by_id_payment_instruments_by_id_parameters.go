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

// NewDeleteCustomersByIDPaymentInstrumentsByIDParams creates a new DeleteCustomersByIDPaymentInstrumentsByIDParams object
// with the default values initialized.
func NewDeleteCustomersByIDPaymentInstrumentsByIDParams() *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteCustomersByIDPaymentInstrumentsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithTimeout creates a new DeleteCustomersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithTimeout(timeout time.Duration) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteCustomersByIDPaymentInstrumentsByIDParams{

		timeout: timeout,
	}
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithContext creates a new DeleteCustomersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithContext(ctx context.Context) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteCustomersByIDPaymentInstrumentsByIDParams{

		Context: ctx,
	}
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithHTTPClient creates a new DeleteCustomersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomersByIDPaymentInstrumentsByIDParamsWithHTTPClient(client *http.Client) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	var ()
	return &DeleteCustomersByIDPaymentInstrumentsByIDParams{
		HTTPClient: client,
	}
}

/*DeleteCustomersByIDPaymentInstrumentsByIDParams contains all the parameters to send to the API endpoint
for the delete customers by ID payment instruments by ID operation typically these are written to a http.Request
*/
type DeleteCustomersByIDPaymentInstrumentsByIDParams struct {

	/*CustomerID
	  the id of the customer to delete the payment instrument for

	*/
	CustomerID string
	/*PaymentInstrumentID
	  the id of the payment instrument to be deleted

	*/
	PaymentInstrumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WithTimeout(timeout time.Duration) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WithContext(ctx context.Context) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WithHTTPClient(client *http.Client) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WithCustomerID(customerID string) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithPaymentInstrumentID adds the paymentInstrumentID to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WithPaymentInstrumentID(paymentInstrumentID string) *DeleteCustomersByIDPaymentInstrumentsByIDParams {
	o.SetPaymentInstrumentID(paymentInstrumentID)
	return o
}

// SetPaymentInstrumentID adds the paymentInstrumentId to the delete customers by ID payment instruments by ID params
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) SetPaymentInstrumentID(paymentInstrumentID string) {
	o.PaymentInstrumentID = paymentInstrumentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomersByIDPaymentInstrumentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
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
