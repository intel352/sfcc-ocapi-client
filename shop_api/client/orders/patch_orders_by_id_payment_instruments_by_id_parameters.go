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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPatchOrdersByIDPaymentInstrumentsByIDParams creates a new PatchOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized.
func NewPatchOrdersByIDPaymentInstrumentsByIDParams() *PatchOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &PatchOrdersByIDPaymentInstrumentsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithTimeout creates a new PatchOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithTimeout(timeout time.Duration) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &PatchOrdersByIDPaymentInstrumentsByIDParams{

		timeout: timeout,
	}
}

// NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithContext creates a new PatchOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithContext(ctx context.Context) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &PatchOrdersByIDPaymentInstrumentsByIDParams{

		Context: ctx,
	}
}

// NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithHTTPClient creates a new PatchOrdersByIDPaymentInstrumentsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchOrdersByIDPaymentInstrumentsByIDParamsWithHTTPClient(client *http.Client) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	var ()
	return &PatchOrdersByIDPaymentInstrumentsByIDParams{
		HTTPClient: client,
	}
}

/*PatchOrdersByIDPaymentInstrumentsByIDParams contains all the parameters to send to the API endpoint
for the patch orders by ID payment instruments by ID operation typically these are written to a http.Request
*/
type PatchOrdersByIDPaymentInstrumentsByIDParams struct {

	/*Body*/
	Body *models.OrderPaymentInstrumentRequest
	/*OrderNo
	  the order number

	*/
	OrderNo string
	/*PaymentInstrumentID
	  the id of the payment instrument to be updated

	*/
	PaymentInstrumentID string
	/*SkipAuthorization*/
	SkipAuthorization *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithTimeout(timeout time.Duration) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithContext(ctx context.Context) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithHTTPClient(client *http.Client) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithBody(body *models.OrderPaymentInstrumentRequest) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetBody(body *models.OrderPaymentInstrumentRequest) {
	o.Body = body
}

// WithOrderNo adds the orderNo to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithOrderNo(orderNo string) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetOrderNo(orderNo string) {
	o.OrderNo = orderNo
}

// WithPaymentInstrumentID adds the paymentInstrumentID to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithPaymentInstrumentID(paymentInstrumentID string) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetPaymentInstrumentID(paymentInstrumentID)
	return o
}

// SetPaymentInstrumentID adds the paymentInstrumentId to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetPaymentInstrumentID(paymentInstrumentID string) {
	o.PaymentInstrumentID = paymentInstrumentID
}

// WithSkipAuthorization adds the skipAuthorization to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WithSkipAuthorization(skipAuthorization *bool) *PatchOrdersByIDPaymentInstrumentsByIDParams {
	o.SetSkipAuthorization(skipAuthorization)
	return o
}

// SetSkipAuthorization adds the skipAuthorization to the patch orders by ID payment instruments by ID params
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) SetSkipAuthorization(skipAuthorization *bool) {
	o.SkipAuthorization = skipAuthorization
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOrdersByIDPaymentInstrumentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param order_no
	if err := r.SetPathParam("order_no", o.OrderNo); err != nil {
		return err
	}

	// path param payment_instrument_id
	if err := r.SetPathParam("payment_instrument_id", o.PaymentInstrumentID); err != nil {
		return err
	}

	if o.SkipAuthorization != nil {

		// query param skip_authorization
		var qrSkipAuthorization bool
		if o.SkipAuthorization != nil {
			qrSkipAuthorization = *o.SkipAuthorization
		}
		qSkipAuthorization := swag.FormatBool(qrSkipAuthorization)
		if qSkipAuthorization != "" {
			if err := r.SetQueryParam("skip_authorization", qSkipAuthorization); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}