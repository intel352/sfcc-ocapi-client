// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteOrdersByIDNotesByID Removes an order note.
*/
func (a *Client) DeleteOrdersByIDNotesByID(params *DeleteOrdersByIDNotesByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrdersByIDNotesByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrdersByIDNotesByID",
		Method:             "DELETE",
		PathPattern:        "/orders/{order_no}/notes/{note_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOrdersByIDNotesByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteOrdersByIDPaymentInstrumentsByID Removes a payment instrument of an order.
*/
func (a *Client) DeleteOrdersByIDPaymentInstrumentsByID(params *DeleteOrdersByIDPaymentInstrumentsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrdersByIDPaymentInstrumentsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrdersByIDPaymentInstrumentsByID",
		Method:             "DELETE",
		PathPattern:        "/orders/{order_no}/payment_instruments/{payment_instrument_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOrdersByIDPaymentInstrumentsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetOrdersByID Gets information for an order.
*/
func (a *Client) GetOrdersByID(params *GetOrdersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrdersByID",
		Method:             "GET",
		PathPattern:        "/orders/{order_no}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrdersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetOrdersByIDNotes Retrieves notes for an order.
*/
func (a *Client) GetOrdersByIDNotes(params *GetOrdersByIDNotesParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersByIDNotesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrdersByIDNotes",
		Method:             "GET",
		PathPattern:        "/orders/{order_no}/notes",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrdersByIDNotesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetOrdersByIDPaymentMethods Gets the applicable payment methods for an existing order considering the open payment amount only.
*/
func (a *Client) GetOrdersByIDPaymentMethods(params *GetOrdersByIDPaymentMethodsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersByIDPaymentMethodsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrdersByIDPaymentMethods",
		Method:             "GET",
		PathPattern:        "/orders/{order_no}/payment_methods",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOrdersByIDPaymentMethodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PatchOrdersByID Update an order.

 Considered fields for update are status (same status transitions are possible as for dw.order.Order.setStatus(int
 status) plus CREATED to FAILED) and custom properties. During the call the correct channel type will be assured to be set for the order
 in a successful call. Without agent context the channel type will be storefront otherwise callcenter.
*/
func (a *Client) PatchOrdersByID(params *PatchOrdersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrdersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchOrdersByID",
		Method:             "PATCH",
		PathPattern:        "/orders/{order_no}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchOrdersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PatchOrdersByIDPaymentInstrumentsByID Updates a payment instrument of an order and passes the order and updated payment instrument to the correct
 payment authorizeCreditcard or authorize hook.

 Details:

 The payment instrument is updated with the provided details. The payment method must be applicable for the
 order see GET /baskets/{basket_id}/payment_methods, if the payment method is 'CREDIT_CARD' a
 payment_card must be specified in the request.


 Order authorization:


 To authorize the order one of two possible customization hooks is called and an
 dw.order.OrderPaymentInstrument is passed as an input argument.


 Which hook is called?


 If the request includes a payment_card or the dw.order.OrderPaymentInstrument
 contains a creditCardType the customization hook dw.order.payment.authorizeCreditCard is called.
 See dw.order.hooks.PaymentHooks.authorizeCreditCard(order : Order, paymentDetails : OrderPaymentInstrument, cvn : String) : Status.
 Otherwise dw.order.payment.authorize is called.
 See dw.order.hooks.PaymentHooks.authorize(order : Order, paymentDetails : OrderPaymentInstrument) : Status.


 What is the dw.order.OrderPaymentInstrument input argument passed to the hook?


 If the request contains a customer_payment_instrument_id the
 dw.order.OrderPaymentInstrument is copied from the customer payment instrument (An exception is thrown
 if none was found).
 Otherwise the data from the request document is passed (payment_card or
 payment_bank_account etc. information).


 Note: the amount and the security_code (cvn) contained in the
 payment_card data will be propagated if available to
 dw.order.payment.authorizeCreditCard even if the dw.order.OrderPaymentInstrument is
 resolved from a customer payment instrument.


 Customization hook dw.ocapi.shop.order.afterPatchPaymentInstrument is called. The default
 implementation places the order if the order status is CREATED and the authorization amount equals or exceeds the
 order total. Placing the order (equivalent to calling dw.order.OrderMgr.placeOrder(order : Order)
 in the scripting API) results in the order being changed to status NEW and prepared for export.

*/
func (a *Client) PatchOrdersByIDPaymentInstrumentsByID(params *PatchOrdersByIDPaymentInstrumentsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrdersByIDPaymentInstrumentsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchOrdersByIDPaymentInstrumentsByID",
		Method:             "PATCH",
		PathPattern:        "/orders/{order_no}/payment_instruments/{payment_instrument_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchOrdersByIDPaymentInstrumentsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostOrders Submits an order based on a prepared basket.  Note: If the basket has been submitted
 using Order Center (considered by it's client id) the channel type will
 be set to "Call Center". In case another channel type was set by a script
 before submitting the basket, the channel type will be reset to
 "Call Center" and a warning will be logged.
 The only considered value from the request body is basket_id.
*/
func (a *Client) PostOrders(params *PostOrdersParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOrders",
		Method:             "POST",
		PathPattern:        "/orders",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostOrdersByIDNotes Adds a note to an existing order.
*/
func (a *Client) PostOrdersByIDNotes(params *PostOrdersByIDNotesParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersByIDNotesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOrdersByIDNotes",
		Method:             "POST",
		PathPattern:        "/orders/{order_no}/notes",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOrdersByIDNotesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostOrdersByIDPaymentInstruments Adds a payment instrument to an order. It is possible either to supply the full payment information or only a
 customer payment instrument id and amount. In case the customer payment instrument id was set all the other
 properties (except amount) are ignored and the payment data is resolved from the stored customer payment
 information. An attempt is made to authorize the order by passing it to the authorize or authorizeCreditCard
 hook.

 Details:

 The payment instrument is added with the provided details or the details from the customer payment
 instrument. The payment method must be applicable for the order see GET
 /baskets/{basket_id}/payment_methods, if the payment method is 'CREDIT_CARD' a
 payment_card must be specified in the request.


 Order authorization:


 To authorize the order one of two possible customization hooks is called and an
 dw.order.OrderPaymentInstrument is passed as an input argument.


 Which hook is called?


 If the request includes a payment_card or the dw.order.OrderPaymentInstrument
 contains a creditCardType the customization hook dw.order.payment.authorizeCreditCard is called.
  See dw.order.hooks.PaymentHooks.authorizeCreditCard(order : Order, paymentDetails : OrderPaymentInstrument, cvn : String) : Status.
 Otherwise dw.order.payment.authorize is called. See dw.order.hooks.PaymentHooks.authorize(order : Order, paymentDetails : OrderPaymentInstrument) : Status.


 What is the dw.order.OrderPaymentInstrument input argument passed to the hook?


 If the request contains a customer_payment_instrument_id the
 dw.order.OrderPaymentInstrument is copied from the customer payment instrument (An exception is thrown
 if none was found).
 Otherwise the data from the request document is passed (payment_card or
 payment_bank_account etc. information).


 Note: the amount and the security_code (cvn) contained in the
 payment_card data will be propagated if available to
 dw.order.payment.authorizeCreditCard even if the dw.order.OrderPaymentInstrument is
 resolved from a customer payment instrument.


 Customization hook dw.ocapi.shop.order.afterPostPaymentInstrument is called. The default
 implementation places the order if the order status is CREATED and the authorization amount equals or exceeds the
 order total. Placing the order (equivalent to calling dw.order.OrderMgr.placeOrder(order : Order) in the
 scripting API) results in the order being changed to status NEW and prepared for export.

*/
func (a *Client) PostOrdersByIDPaymentInstruments(params *PostOrdersByIDPaymentInstrumentsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrdersByIDPaymentInstrumentsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOrdersByIDPaymentInstruments",
		Method:             "POST",
		PathPattern:        "/orders/{order_no}/payment_instruments",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOrdersByIDPaymentInstrumentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutOrdersByID Submits an order with a given order number, based on a prepared basket.  The
 only considered value from the request body is basket_id. This resource is available for OAuth
 authentication and requires no user i.e. it supports server-server communication with client grant authentication
 and no user is specified.
*/
func (a *Client) PutOrdersByID(params *PutOrdersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOrdersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putOrdersByID",
		Method:             "PUT",
		PathPattern:        "/orders/{order_no}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutOrdersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
