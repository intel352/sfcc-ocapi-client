// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetProductsByIDReader is a Reader for the GetProductsByID structure.
type GetProductsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetProductsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsByIDNotFound creates a GetProductsByIDNotFound with default headers values
func NewGetProductsByIDNotFound() *GetProductsByIDNotFound {
	return &GetProductsByIDNotFound{}
}

/*GetProductsByIDNotFound handles this case with default header values.

Indicates the product is not found.
*/
type GetProductsByIDNotFound struct {
}

func (o *GetProductsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsByIdNotFound ", 404)
}

func (o *GetProductsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDDefault creates a GetProductsByIDDefault with default headers values
func NewGetProductsByIDDefault(code int) *GetProductsByIDDefault {
	return &GetProductsByIDDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDDefault handles this case with default header values.

GetProductsByIDDefault get products by ID default
*/
type GetProductsByIDDefault struct {
	_statusCode int

	Payload *models.Product
}

// Code gets the status code for the get products by ID default response
func (o *GetProductsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}