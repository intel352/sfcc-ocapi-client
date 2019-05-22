// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new products API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for products API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetProductsByID To access single products resource, you construct a URL using the template shown below. This template requires
 you to specify an Id (typically a SKU) for a product. In response, the server returns a corresponding Product
 document, provided the product is online and assigned to site catalog. The document contains variation attributes
 (including values) and the variant matrix; this data is provided for both the master and for the variant.
*/
func (a *Client) GetProductsByID(params *GetProductsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByID",
		Method:             "GET",
		PathPattern:        "/products/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDReader{formats: a.formats},
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
GetProductsByIDAvailability Access product availability information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDAvailability(params *GetProductsByIDAvailabilityParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDAvailabilityParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDAvailability",
		Method:             "GET",
		PathPattern:        "/products/{id}/availability",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDAvailabilityReader{formats: a.formats},
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
GetProductsByIDBundledProducts Access bundled product information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDBundledProducts(params *GetProductsByIDBundledProductsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDBundledProductsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDBundledProducts",
		Method:             "GET",
		PathPattern:        "/products/{id}/bundled_products",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDBundledProductsReader{formats: a.formats},
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
GetProductsByIDImages Access product image information of products that are online and assigned to site catalog. Filter the result by
 view type and variation values.
*/
func (a *Client) GetProductsByIDImages(params *GetProductsByIDImagesParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDImagesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDImages",
		Method:             "GET",
		PathPattern:        "/products/{id}/images",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDImagesReader{formats: a.formats},
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
GetProductsByIDLinks Access product link information of products that are online and assigned to site catalog. Filter the result by
 link type and link direction.
*/
func (a *Client) GetProductsByIDLinks(params *GetProductsByIDLinksParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDLinksParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDLinks",
		Method:             "GET",
		PathPattern:        "/products/{id}/links",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDLinksReader{formats: a.formats},
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
GetProductsByIDOptions Access product option information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDOptions(params *GetProductsByIDOptionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDOptionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDOptions",
		Method:             "GET",
		PathPattern:        "/products/{id}/options",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDOptionsReader{formats: a.formats},
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
GetProductsByIDPrices Access product price information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDPrices(params *GetProductsByIDPricesParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDPricesParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDPrices",
		Method:             "GET",
		PathPattern:        "/products/{id}/prices",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDPricesReader{formats: a.formats},
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
GetProductsByIDPromotions Access product promotion information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDPromotions(params *GetProductsByIDPromotionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDPromotionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDPromotions",
		Method:             "GET",
		PathPattern:        "/products/{id}/promotions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDPromotionsReader{formats: a.formats},
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
GetProductsByIDRecommendations Access product recommendation information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDRecommendations(params *GetProductsByIDRecommendationsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDRecommendationsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDRecommendations",
		Method:             "GET",
		PathPattern:        "/products/{id}/recommendations",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDRecommendationsReader{formats: a.formats},
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
GetProductsByIDSetProducts Access product set information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDSetProducts(params *GetProductsByIDSetProductsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDSetProductsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDSetProducts",
		Method:             "GET",
		PathPattern:        "/products/{id}/set_products",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDSetProductsReader{formats: a.formats},
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
GetProductsByIDShippingMethods Retrieves the applicable shipping methods for a certain product.
*/
func (a *Client) GetProductsByIDShippingMethods(params *GetProductsByIDShippingMethodsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDShippingMethodsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDShippingMethods",
		Method:             "GET",
		PathPattern:        "/products/{id}/shipping_methods",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDShippingMethodsReader{formats: a.formats},
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
GetProductsByIDVariations Access product variation information of products that are online and assigned to site catalog.
*/
func (a *Client) GetProductsByIDVariations(params *GetProductsByIDVariationsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIDVariationsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDVariations",
		Method:             "GET",
		PathPattern:        "/products/{id}/variations",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIDVariationsReader{formats: a.formats},
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
GetProductsByIds get products by ids API
*/
func (a *Client) GetProductsByIds(params *GetProductsByIdsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductsByIdsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductsByIDs",
		Method:             "GET",
		PathPattern:        "/products/({ids})",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductsByIdsReader{formats: a.formats},
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