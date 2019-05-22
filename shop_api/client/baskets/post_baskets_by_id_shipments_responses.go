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

// PostBasketsByIDShipmentsReader is a Reader for the PostBasketsByIDShipments structure.
type PostBasketsByIDShipmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBasketsByIDShipmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostBasketsByIDShipmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostBasketsByIDShipmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostBasketsByIDShipmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostBasketsByIDShipmentsBadRequest creates a PostBasketsByIDShipmentsBadRequest with default headers values
func NewPostBasketsByIDShipmentsBadRequest() *PostBasketsByIDShipmentsBadRequest {
	return &PostBasketsByIDShipmentsBadRequest{}
}

/*PostBasketsByIDShipmentsBadRequest handles this case with default header values.

Indicates that a shipment id is not provided. or Indicates that a shipment with the provided id already
             exists for the basket. or Indicates that a shipment with the provided shipment number
             already exists for the basket. or Indicates that a shipping method with an id was specified
             which is not a valid shipping method id. or Indicates that the customer assigned to the basket does not match the
             verified customer represented by the JWT token, not relevant
             when using OAuth.
*/
type PostBasketsByIDShipmentsBadRequest struct {
}

func (o *PostBasketsByIDShipmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/shipments][%d] postBasketsByIdShipmentsBadRequest ", 400)
}

func (o *PostBasketsByIDShipmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBasketsByIDShipmentsNotFound creates a PostBasketsByIDShipmentsNotFound with default headers values
func NewPostBasketsByIDShipmentsNotFound() *PostBasketsByIDShipmentsNotFound {
	return &PostBasketsByIDShipmentsNotFound{}
}

/*PostBasketsByIDShipmentsNotFound handles this case with default header values.

Indicates that the basket with the given basket id is
             unknown.
*/
type PostBasketsByIDShipmentsNotFound struct {
}

func (o *PostBasketsByIDShipmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/shipments][%d] postBasketsByIdShipmentsNotFound ", 404)
}

func (o *PostBasketsByIDShipmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBasketsByIDShipmentsDefault creates a PostBasketsByIDShipmentsDefault with default headers values
func NewPostBasketsByIDShipmentsDefault(code int) *PostBasketsByIDShipmentsDefault {
	return &PostBasketsByIDShipmentsDefault{
		_statusCode: code,
	}
}

/*PostBasketsByIDShipmentsDefault handles this case with default header values.

PostBasketsByIDShipmentsDefault post baskets by ID shipments default
*/
type PostBasketsByIDShipmentsDefault struct {
	_statusCode int

	Payload *models.Basket
}

// Code gets the status code for the post baskets by ID shipments default response
func (o *PostBasketsByIDShipmentsDefault) Code() int {
	return o._statusCode
}

func (o *PostBasketsByIDShipmentsDefault) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/shipments][%d] postBasketsByIDShipments default  %+v", o._statusCode, o.Payload)
}

func (o *PostBasketsByIDShipmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Basket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
