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

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPatchOrdersByIDParams creates a new PatchOrdersByIDParams object
// with the default values initialized.
func NewPatchOrdersByIDParams() *PatchOrdersByIDParams {
	var ()
	return &PatchOrdersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOrdersByIDParamsWithTimeout creates a new PatchOrdersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchOrdersByIDParamsWithTimeout(timeout time.Duration) *PatchOrdersByIDParams {
	var ()
	return &PatchOrdersByIDParams{

		timeout: timeout,
	}
}

// NewPatchOrdersByIDParamsWithContext creates a new PatchOrdersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchOrdersByIDParamsWithContext(ctx context.Context) *PatchOrdersByIDParams {
	var ()
	return &PatchOrdersByIDParams{

		Context: ctx,
	}
}

// NewPatchOrdersByIDParamsWithHTTPClient creates a new PatchOrdersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchOrdersByIDParamsWithHTTPClient(client *http.Client) *PatchOrdersByIDParams {
	var ()
	return &PatchOrdersByIDParams{
		HTTPClient: client,
	}
}

/*PatchOrdersByIDParams contains all the parameters to send to the API endpoint
for the patch orders by ID operation typically these are written to a http.Request
*/
type PatchOrdersByIDParams struct {

	/*Body*/
	Body *models.Order
	/*OrderNo
	  the order number

	*/
	OrderNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch orders by ID params
func (o *PatchOrdersByIDParams) WithTimeout(timeout time.Duration) *PatchOrdersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch orders by ID params
func (o *PatchOrdersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch orders by ID params
func (o *PatchOrdersByIDParams) WithContext(ctx context.Context) *PatchOrdersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch orders by ID params
func (o *PatchOrdersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch orders by ID params
func (o *PatchOrdersByIDParams) WithHTTPClient(client *http.Client) *PatchOrdersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch orders by ID params
func (o *PatchOrdersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch orders by ID params
func (o *PatchOrdersByIDParams) WithBody(body *models.Order) *PatchOrdersByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch orders by ID params
func (o *PatchOrdersByIDParams) SetBody(body *models.Order) {
	o.Body = body
}

// WithOrderNo adds the orderNo to the patch orders by ID params
func (o *PatchOrdersByIDParams) WithOrderNo(orderNo string) *PatchOrdersByIDParams {
	o.SetOrderNo(orderNo)
	return o
}

// SetOrderNo adds the orderNo to the patch orders by ID params
func (o *PatchOrdersByIDParams) SetOrderNo(orderNo string) {
	o.OrderNo = orderNo
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOrdersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}