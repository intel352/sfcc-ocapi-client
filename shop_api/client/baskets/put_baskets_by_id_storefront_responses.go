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

// PutBasketsByIDStorefrontReader is a Reader for the PutBasketsByIDStorefront structure.
type PutBasketsByIDStorefrontReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutBasketsByIDStorefrontReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutBasketsByIDStorefrontBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutBasketsByIDStorefrontNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutBasketsByIDStorefrontDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutBasketsByIDStorefrontBadRequest creates a PutBasketsByIDStorefrontBadRequest with default headers values
func NewPutBasketsByIDStorefrontBadRequest() *PutBasketsByIDStorefrontBadRequest {
	return &PutBasketsByIDStorefrontBadRequest{}
}

/*PutBasketsByIDStorefrontBadRequest handles this case with default header values.

Thrown if a storefront basket already exists and exchange is
             false.
*/
type PutBasketsByIDStorefrontBadRequest struct {
}

func (o *PutBasketsByIDStorefrontBadRequest) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/storefront][%d] putBasketsByIdStorefrontBadRequest ", 400)
}

func (o *PutBasketsByIDStorefrontBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBasketsByIDStorefrontNotFound creates a PutBasketsByIDStorefrontNotFound with default headers values
func NewPutBasketsByIDStorefrontNotFound() *PutBasketsByIDStorefrontNotFound {
	return &PutBasketsByIDStorefrontNotFound{}
}

/*PutBasketsByIDStorefrontNotFound handles this case with default header values.

Indicates that the basket with the given basket id is unknown.
*/
type PutBasketsByIDStorefrontNotFound struct {
}

func (o *PutBasketsByIDStorefrontNotFound) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/storefront][%d] putBasketsByIdStorefrontNotFound ", 404)
}

func (o *PutBasketsByIDStorefrontNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutBasketsByIDStorefrontDefault creates a PutBasketsByIDStorefrontDefault with default headers values
func NewPutBasketsByIDStorefrontDefault(code int) *PutBasketsByIDStorefrontDefault {
	return &PutBasketsByIDStorefrontDefault{
		_statusCode: code,
	}
}

/*PutBasketsByIDStorefrontDefault handles this case with default header values.

PutBasketsByIDStorefrontDefault put baskets by ID storefront default
*/
type PutBasketsByIDStorefrontDefault struct {
	_statusCode int

	Payload *models.Basket
}

// Code gets the status code for the put baskets by ID storefront default response
func (o *PutBasketsByIDStorefrontDefault) Code() int {
	return o._statusCode
}

func (o *PutBasketsByIDStorefrontDefault) Error() string {
	return fmt.Sprintf("[PUT /baskets/{basket_id}/storefront][%d] putBasketsByIDStorefront default  %+v", o._statusCode, o.Payload)
}

func (o *PutBasketsByIDStorefrontDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Basket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
