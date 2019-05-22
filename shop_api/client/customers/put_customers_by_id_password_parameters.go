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

// NewPutCustomersByIDPasswordParams creates a new PutCustomersByIDPasswordParams object
// with the default values initialized.
func NewPutCustomersByIDPasswordParams() *PutCustomersByIDPasswordParams {
	var ()
	return &PutCustomersByIDPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCustomersByIDPasswordParamsWithTimeout creates a new PutCustomersByIDPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCustomersByIDPasswordParamsWithTimeout(timeout time.Duration) *PutCustomersByIDPasswordParams {
	var ()
	return &PutCustomersByIDPasswordParams{

		timeout: timeout,
	}
}

// NewPutCustomersByIDPasswordParamsWithContext creates a new PutCustomersByIDPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCustomersByIDPasswordParamsWithContext(ctx context.Context) *PutCustomersByIDPasswordParams {
	var ()
	return &PutCustomersByIDPasswordParams{

		Context: ctx,
	}
}

// NewPutCustomersByIDPasswordParamsWithHTTPClient creates a new PutCustomersByIDPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCustomersByIDPasswordParamsWithHTTPClient(client *http.Client) *PutCustomersByIDPasswordParams {
	var ()
	return &PutCustomersByIDPasswordParams{
		HTTPClient: client,
	}
}

/*PutCustomersByIDPasswordParams contains all the parameters to send to the API endpoint
for the put customers by ID password operation typically these are written to a http.Request
*/
type PutCustomersByIDPasswordParams struct {

	/*Body*/
	Body *models.PasswordChangeRequest
	/*CustomerID
	  the customer id

	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) WithTimeout(timeout time.Duration) *PutCustomersByIDPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) WithContext(ctx context.Context) *PutCustomersByIDPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) WithHTTPClient(client *http.Client) *PutCustomersByIDPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) WithBody(body *models.PasswordChangeRequest) *PutCustomersByIDPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) SetBody(body *models.PasswordChangeRequest) {
	o.Body = body
}

// WithCustomerID adds the customerID to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) WithCustomerID(customerID string) *PutCustomersByIDPasswordParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the put customers by ID password params
func (o *PutCustomersByIDPasswordParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCustomersByIDPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
