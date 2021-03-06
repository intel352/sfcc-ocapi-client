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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCustomersByIDOrdersParams creates a new GetCustomersByIDOrdersParams object
// with the default values initialized.
func NewGetCustomersByIDOrdersParams() *GetCustomersByIDOrdersParams {
	var ()
	return &GetCustomersByIDOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomersByIDOrdersParamsWithTimeout creates a new GetCustomersByIDOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomersByIDOrdersParamsWithTimeout(timeout time.Duration) *GetCustomersByIDOrdersParams {
	var ()
	return &GetCustomersByIDOrdersParams{

		timeout: timeout,
	}
}

// NewGetCustomersByIDOrdersParamsWithContext creates a new GetCustomersByIDOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomersByIDOrdersParamsWithContext(ctx context.Context) *GetCustomersByIDOrdersParams {
	var ()
	return &GetCustomersByIDOrdersParams{

		Context: ctx,
	}
}

// NewGetCustomersByIDOrdersParamsWithHTTPClient creates a new GetCustomersByIDOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomersByIDOrdersParamsWithHTTPClient(client *http.Client) *GetCustomersByIDOrdersParams {
	var ()
	return &GetCustomersByIDOrdersParams{
		HTTPClient: client,
	}
}

/*GetCustomersByIDOrdersParams contains all the parameters to send to the API endpoint
for the get customers by ID orders operation typically these are written to a http.Request
*/
type GetCustomersByIDOrdersParams struct {

	/*Count*/
	Count *int32
	/*CrossSites*/
	CrossSites *bool
	/*CustomerID
	  the customer uuid

	*/
	CustomerID string
	/*From*/
	From *string
	/*Start*/
	Start *int32
	/*Status*/
	Status *string
	/*Until*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithTimeout(timeout time.Duration) *GetCustomersByIDOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithContext(ctx context.Context) *GetCustomersByIDOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithHTTPClient(client *http.Client) *GetCustomersByIDOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithCount(count *int32) *GetCustomersByIDOrdersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetCount(count *int32) {
	o.Count = count
}

// WithCrossSites adds the crossSites to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithCrossSites(crossSites *bool) *GetCustomersByIDOrdersParams {
	o.SetCrossSites(crossSites)
	return o
}

// SetCrossSites adds the crossSites to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetCrossSites(crossSites *bool) {
	o.CrossSites = crossSites
}

// WithCustomerID adds the customerID to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithCustomerID(customerID string) *GetCustomersByIDOrdersParams {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithFrom adds the from to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithFrom(from *string) *GetCustomersByIDOrdersParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetFrom(from *string) {
	o.From = from
}

// WithStart adds the start to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithStart(start *int32) *GetCustomersByIDOrdersParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetStart(start *int32) {
	o.Start = start
}

// WithStatus adds the status to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithStatus(status *string) *GetCustomersByIDOrdersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetStatus(status *string) {
	o.Status = status
}

// WithUntil adds the until to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) WithUntil(until *string) *GetCustomersByIDOrdersParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the get customers by ID orders params
func (o *GetCustomersByIDOrdersParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomersByIDOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.CrossSites != nil {

		// query param cross-sites
		var qrCrossSites bool
		if o.CrossSites != nil {
			qrCrossSites = *o.CrossSites
		}
		qCrossSites := swag.FormatBool(qrCrossSites)
		if qCrossSites != "" {
			if err := r.SetQueryParam("cross-sites", qCrossSites); err != nil {
				return err
			}
		}

	}

	// path param customer_id
	if err := r.SetPathParam("customer_id", o.CustomerID); err != nil {
		return err
	}

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Until != nil {

		// query param until
		var qrUntil string
		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {
			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
