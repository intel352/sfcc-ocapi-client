// Code generated by go-swagger; DO NOT EDIT.

package baskets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PutBasketsByIDShipmentsByIDShippingAddressReader is a Reader for the PutBasketsByIDShipmentsByIDShippingAddress structure.
type PutBasketsByIDShipmentsByIDShippingAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBasketsByIDShipmentsByIDShippingAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutBasketsByIDShipmentsByIDShippingAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutBasketsByIDShipmentsByIDShippingAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutBasketsByIDShipmentsByIDShippingAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutBasketsByIDShipmentsByIDShippingAddressBadRequest creates a PutBasketsByIDShipmentsByIDShippingAddressBadRequest with default headers values
func NewPutBasketsByIDShipmentsByIDShippingAddressBadRequest() *PutBasketsByIDShipmentsByIDShippingAddressBadRequest {
	return &PutBasketsByIDShipmentsByIDShippingAddressBadRequest{}
}

/*PutBasketsByIDShipmentsByIDShippingAddressBadRequest handles this case with default header values.

Indicates that both customer_address_id and address body was
             provided. or Thrown if the shipment with the given shipment id is unknown. or Indicates that the customer assigned to the basket does not
             match the verified customer represented by the JWT token (not
             relevant when using OAuth). or Indicates that an customerAddressId was provided but either
             an anonymous or no customer was set.
*/
type PutBasketsByIDShipmentsByIDShippingAddressBadRequest struct {
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressBadRequest) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/shipments/{shipment_id}/shipping_address][%d] putBasketsByIdShipmentsByIdShippingAddressBadRequest ", 400)
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBasketsByIDShipmentsByIDShippingAddressNotFound creates a PutBasketsByIDShipmentsByIDShippingAddressNotFound with default headers values
func NewPutBasketsByIDShipmentsByIDShippingAddressNotFound() *PutBasketsByIDShipmentsByIDShippingAddressNotFound {
	return &PutBasketsByIDShipmentsByIDShippingAddressNotFound{}
}

/*PutBasketsByIDShipmentsByIDShippingAddressNotFound handles this case with default header values.

Indicates that the basket with the given basket id is unknown. or Indicates that the address specified by customer_address_id is
             unknown.
*/
type PutBasketsByIDShipmentsByIDShippingAddressNotFound struct {
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressNotFound) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/shipments/{shipment_id}/shipping_address][%d] putBasketsByIdShipmentsByIdShippingAddressNotFound ", 404)
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBasketsByIDShipmentsByIDShippingAddressDefault creates a PutBasketsByIDShipmentsByIDShippingAddressDefault with default headers values
func NewPutBasketsByIDShipmentsByIDShippingAddressDefault(code int) *PutBasketsByIDShipmentsByIDShippingAddressDefault {
	return &PutBasketsByIDShipmentsByIDShippingAddressDefault{
		_statusCode: code,
	}
}

/*PutBasketsByIDShipmentsByIDShippingAddressDefault handles this case with default header values.

PutBasketsByIDShipmentsByIDShippingAddressDefault put baskets by ID shipments by ID shipping address default
*/
type PutBasketsByIDShipmentsByIDShippingAddressDefault struct {
	_statusCode int

	Payload *models.Basket
}

// Code gets the status code for the put baskets by ID shipments by ID shipping address default response
func (o *PutBasketsByIDShipmentsByIDShippingAddressDefault) Code() int {
	return o._statusCode
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressDefault) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/shipments/{shipment_id}/shipping_address][%d] putBasketsByIDShipmentsByIDShippingAddress default  %+v", o._statusCode, o.Payload)
}

func (o *PutBasketsByIDShipmentsByIDShippingAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Basket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}