// Code generated by go-swagger; DO NOT EDIT.

package customer_lists

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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPatchCustomerListsByIDCustomersByIDAddressesByIDParams creates a new PatchCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized.
func NewPatchCustomerListsByIDCustomersByIDAddressesByIDParams() *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &PatchCustomerListsByIDCustomersByIDAddressesByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithTimeout creates a new PatchCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithTimeout(timeout time.Duration) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &PatchCustomerListsByIDCustomersByIDAddressesByIDParams{

		timeout: timeout,
	}
}

// NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithContext creates a new PatchCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithContext(ctx context.Context) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &PatchCustomerListsByIDCustomersByIDAddressesByIDParams{

		Context: ctx,
	}
}

// NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithHTTPClient creates a new PatchCustomerListsByIDCustomersByIDAddressesByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCustomerListsByIDCustomersByIDAddressesByIDParamsWithHTTPClient(client *http.Client) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	var ()
	return &PatchCustomerListsByIDCustomersByIDAddressesByIDParams{
		HTTPClient: client,
	}
}

/*PatchCustomerListsByIDCustomersByIDAddressesByIDParams contains all the parameters to send to the API endpoint
for the patch customer lists by ID customers by ID addresses by ID operation typically these are written to a http.Request
*/
type PatchCustomerListsByIDCustomersByIDAddressesByIDParams struct {

	/*AddressID
	  The address id

	*/
	AddressID string
	/*Body*/
	Body *models.CustomerAddress
	/*CustomerNo
	  The customer number

	*/
	CustomerNo string
	/*ListID
	  The customer list id

	*/
	ListID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithTimeout(timeout time.Duration) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithContext(ctx context.Context) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithHTTPClient(client *http.Client) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressID adds the addressID to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithAddressID(addressID string) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetAddressID(addressID)
	return o
}

// SetAddressID adds the addressId to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetAddressID(addressID string) {
	o.AddressID = addressID
}

// WithBody adds the body to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithBody(body *models.CustomerAddress) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetBody(body *models.CustomerAddress) {
	o.Body = body
}

// WithCustomerNo adds the customerNo to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithCustomerNo(customerNo string) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetCustomerNo(customerNo)
	return o
}

// SetCustomerNo adds the customerNo to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetCustomerNo(customerNo string) {
	o.CustomerNo = customerNo
}

// WithListID adds the listID to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WithListID(listID string) *PatchCustomerListsByIDCustomersByIDAddressesByIDParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the patch customer lists by ID customers by ID addresses by ID params
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) SetListID(listID string) {
	o.ListID = listID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCustomerListsByIDCustomersByIDAddressesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address_id
	if err := r.SetPathParam("address_id", o.AddressID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param customer_no
	if err := r.SetPathParam("customer_no", o.CustomerNo); err != nil {
		return err
	}

	// path param list_id
	if err := r.SetPathParam("list_id", o.ListID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}