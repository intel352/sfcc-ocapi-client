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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPutBasketsByIDShipmentsByIDShippingAddressParams creates a new PutBasketsByIDShipmentsByIDShippingAddressParams object
// with the default values initialized.
func NewPutBasketsByIDShipmentsByIDShippingAddressParams() *PutBasketsByIDShipmentsByIDShippingAddressParams {
	var ()
	return &PutBasketsByIDShipmentsByIDShippingAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithTimeout creates a new PutBasketsByIDShipmentsByIDShippingAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithTimeout(timeout time.Duration) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	var ()
	return &PutBasketsByIDShipmentsByIDShippingAddressParams{

		timeout: timeout,
	}
}

// NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithContext creates a new PutBasketsByIDShipmentsByIDShippingAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithContext(ctx context.Context) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	var ()
	return &PutBasketsByIDShipmentsByIDShippingAddressParams{

		Context: ctx,
	}
}

// NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithHTTPClient creates a new PutBasketsByIDShipmentsByIDShippingAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutBasketsByIDShipmentsByIDShippingAddressParamsWithHTTPClient(client *http.Client) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	var ()
	return &PutBasketsByIDShipmentsByIDShippingAddressParams{
		HTTPClient: client,
	}
}

/*PutBasketsByIDShipmentsByIDShippingAddressParams contains all the parameters to send to the API endpoint
for the put baskets by ID shipments by ID shipping address operation typically these are written to a http.Request
*/
type PutBasketsByIDShipmentsByIDShippingAddressParams struct {

	/*BasketID
	  The id of the basket to be modified.

	*/
	BasketID string
	/*Body*/
	Body *models.OrderAddress
	/*CustomerAddressID*/
	CustomerAddressID *string
	/*ShipmentID
	  The id of the shipment to be modified.

	*/
	ShipmentID string
	/*UseAsBilling*/
	UseAsBilling *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithTimeout(timeout time.Duration) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithContext(ctx context.Context) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithHTTPClient(client *http.Client) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBasketID adds the basketID to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithBasketID(basketID string) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetBasketID(basketID)
	return o
}

// SetBasketID adds the basketId to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetBasketID(basketID string) {
	o.BasketID = basketID
}

// WithBody adds the body to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithBody(body *models.OrderAddress) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetBody(body *models.OrderAddress) {
	o.Body = body
}

// WithCustomerAddressID adds the customerAddressID to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithCustomerAddressID(customerAddressID *string) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetCustomerAddressID(customerAddressID)
	return o
}

// SetCustomerAddressID adds the customerAddressId to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetCustomerAddressID(customerAddressID *string) {
	o.CustomerAddressID = customerAddressID
}

// WithShipmentID adds the shipmentID to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithShipmentID(shipmentID string) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WithUseAsBilling adds the useAsBilling to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WithUseAsBilling(useAsBilling *bool) *PutBasketsByIDShipmentsByIDShippingAddressParams {
	o.SetUseAsBilling(useAsBilling)
	return o
}

// SetUseAsBilling adds the useAsBilling to the put baskets by ID shipments by ID shipping address params
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) SetUseAsBilling(useAsBilling *bool) {
	o.UseAsBilling = useAsBilling
}

// WriteToRequest writes these params to a swagger request
func (o *PutBasketsByIDShipmentsByIDShippingAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomerAddressID != nil {

		// query param customer_address_id
		var qrCustomerAddressID string
		if o.CustomerAddressID != nil {
			qrCustomerAddressID = *o.CustomerAddressID
		}
		qCustomerAddressID := qrCustomerAddressID
		if qCustomerAddressID != "" {
			if err := r.SetQueryParam("customer_address_id", qCustomerAddressID); err != nil {
				return err
			}
		}

	}

	// path param shipment_id
	if err := r.SetPathParam("shipment_id", o.ShipmentID); err != nil {
		return err
	}

	if o.UseAsBilling != nil {

		// query param use_as_billing
		var qrUseAsBilling bool
		if o.UseAsBilling != nil {
			qrUseAsBilling = *o.UseAsBilling
		}
		qUseAsBilling := swag.FormatBool(qrUseAsBilling)
		if qUseAsBilling != "" {
			if err := r.SetQueryParam("use_as_billing", qUseAsBilling); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}