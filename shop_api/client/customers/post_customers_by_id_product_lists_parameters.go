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

// NewPostCustomersByIDProductListsParams creates a new PostCustomersByIDProductListsParams object
// with the default values initialized.
func NewPostCustomersByIDProductListsParams() *PostCustomersByIDProductListsParams {
	var ()
	return &PostCustomersByIDProductListsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomersByIDProductListsParamsWithTimeout creates a new PostCustomersByIDProductListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomersByIDProductListsParamsWithTimeout(timeout time.Duration) *PostCustomersByIDProductListsParams {
	var ()
	return &PostCustomersByIDProductListsParams{

		timeout: timeout,
	}
}

// NewPostCustomersByIDProductListsParamsWithContext creates a new PostCustomersByIDProductListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomersByIDProductListsParamsWithContext(ctx context.Context) *PostCustomersByIDProductListsParams {
	var ()
	return &PostCustomersByIDProductListsParams{

		Context: ctx,
	}
}

// NewPostCustomersByIDProductListsParamsWithHTTPClient creates a new PostCustomersByIDProductListsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomersByIDProductListsParamsWithHTTPClient(client *http.Client) *PostCustomersByIDProductListsParams {
	var ()
	return &PostCustomersByIDProductListsParams{
		HTTPClient: client,
	}
}

/*PostCustomersByIDProductListsParams contains all the parameters to send to the API endpoint
for the post customers by ID product lists operation typically these are written to a http.Request
*/
type PostCustomersByIDProductListsParams struct {

	/*Body*/
	Body *models.CustomerProductList
	/*CustomerID
	  The customer id.

	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) WithTimeout(timeout time.Duration) *PostCustomersByIDProductListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) WithContext(ctx context.Context) *PostCustomersByIDProductListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) WithHTTPClient(client *http.Client) *PostCustomersByIDProductListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) WithBody(body *models.CustomerProductList) *PostCustomersByIDProductListsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) SetBody(body *models.CustomerProductList) {
	o.Body = body
}

// WithCustomerID adds the customerID to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) WithCustomerID(customerID string) *PostCustomersByIDProductListsParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the post customers by ID product lists params
func (o *PostCustomersByIDProductListsParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomersByIDProductListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
