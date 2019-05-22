// Code generated by go-swagger; DO NOT EDIT.

package order_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new order search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for order search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostOrderSearch Searches for orders.

 The query attribute specifies a complex query that can be used to narrow down the search.

 Note that search fields are mandatory now and no default ones are supported.

 As the old order search version, the new one always uses Search Service too and the for that reason Order
 Incremental Indexing should be enabled. Otherwise HTTP 500 response will occur.

 The supported search fields are:

 affiliate_partner_i_d
 affiliate_partner_name
 business_type
 channel_type
 confirmation_status (String)
 created_by
 creation_date
 currency_code
 customer_email
 customer_name
 customer_no
 export_after
 export_status (String)
 external_order_no
 external_order_status
 global_party_id
 last_modified
 order_no
 original_order_no
 payment_status (String)
 replaced_order_no
 replacement_order_no
 shipping_status (String)
 status (String)
 total_gross_price
 total_net_price
 order.has_holds
 coupon_line_items.coupon_code
 coupon_line_items.coupon_id
 holds.type
 invoices.status
 order_items.status
 payment_instruments.credit_card_type
 payment_instruments.payment_method_id
 product_items.product_id
 return_cases.return_case_number
 shipments.shipping_method_id
 shipping_orders.shipping_order_number

 The sort order of the retrieved orders could be specified by the "sorts" parameter. It is a list of objects
 presenting field name and sort direction ("asc" or "desc").

 Custom attributes could be used as search fields and as sort fields too. A prefix "c_" has to be added to them.
*/
func (a *Client) PostOrderSearch(params *PostOrderSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrderSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postOrderSearch",
		Method:             "POST",
		PathPattern:        "/order_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOrderSearchReader{formats: a.formats},
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
