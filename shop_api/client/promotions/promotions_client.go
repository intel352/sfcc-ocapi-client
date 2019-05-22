// Code generated by go-swagger; DO NOT EDIT.

package promotions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new promotions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for promotions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPromotions Handles get promotion by filter criteria Returns an array of enabled promotions matching specified filter
 criteria. In the request URL, you must provide a campaign_id parameter, and you can optionally specify a date
 range by providing start_date and end_date parameters. Both parameters are required to specify a date range:
 omitting one causes the server to return a MissingParameterException fault. Each request returns only enabled
 promotions; the server does not consider promotion qualifiers or schedules.
*/
func (a *Client) GetPromotions(params *GetPromotionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPromotionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPromotions",
		Method:             "GET",
		PathPattern:        "/promotions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPromotionsReader{formats: a.formats},
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
GetPromotionsByID Returns an enabled promotion using a specified id. Each request returns a response only for an enabled promotion;
 the server does not consider promotion qualifiers or schedules.
*/
func (a *Client) GetPromotionsByID(params *GetPromotionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPromotionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPromotionsByID",
		Method:             "GET",
		PathPattern:        "/promotions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPromotionsByIDReader{formats: a.formats},
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
GetPromotionsByIds get promotions by ids API
*/
func (a *Client) GetPromotionsByIds(params *GetPromotionsByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPromotionsByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPromotionsByIDs",
		Method:             "GET",
		PathPattern:        "/promotions/({ids})",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPromotionsByIdsReader{formats: a.formats},
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
