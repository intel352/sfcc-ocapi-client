// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetProductsByIDShippingMethodsReader is a Reader for the GetProductsByIDShippingMethods structure.
type GetProductsByIDShippingMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDShippingMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetProductsByIDShippingMethodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductsByIDShippingMethodsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsByIDShippingMethodsNotFound creates a GetProductsByIDShippingMethodsNotFound with default headers values
func NewGetProductsByIDShippingMethodsNotFound() *GetProductsByIDShippingMethodsNotFound {
	return &GetProductsByIDShippingMethodsNotFound{}
}

/*GetProductsByIDShippingMethodsNotFound handles this case with default header values.

Indicates that the product with the given id is unknown.
*/
type GetProductsByIDShippingMethodsNotFound struct {
}

func (o *GetProductsByIDShippingMethodsNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}/shipping_methods][%d] getProductsByIdShippingMethodsNotFound ", 404)
}

func (o *GetProductsByIDShippingMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDShippingMethodsDefault creates a GetProductsByIDShippingMethodsDefault with default headers values
func NewGetProductsByIDShippingMethodsDefault(code int) *GetProductsByIDShippingMethodsDefault {
	return &GetProductsByIDShippingMethodsDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDShippingMethodsDefault handles this case with default header values.

GetProductsByIDShippingMethodsDefault get products by ID shipping methods default
*/
type GetProductsByIDShippingMethodsDefault struct {
	_statusCode int

	Payload *models.ShippingMethodResult
}

// Code gets the status code for the get products by ID shipping methods default response
func (o *GetProductsByIDShippingMethodsDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDShippingMethodsDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}/shipping_methods][%d] getProductsByIDShippingMethods default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDShippingMethodsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShippingMethodResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
