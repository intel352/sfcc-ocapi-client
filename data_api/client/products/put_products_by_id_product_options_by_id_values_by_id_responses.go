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

// PutProductsByIDProductOptionsByIDValuesByIDReader is a Reader for the PutProductsByIDProductOptionsByIDValuesByID structure.
type PutProductsByIDProductOptionsByIDValuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutProductsByIDProductOptionsByIDValuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPutProductsByIDProductOptionsByIDValuesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutProductsByIDProductOptionsByIDValuesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutProductsByIDProductOptionsByIDValuesByIDNotFound creates a PutProductsByIDProductOptionsByIDValuesByIDNotFound with default headers values
func NewPutProductsByIDProductOptionsByIDValuesByIDNotFound() *PutProductsByIDProductOptionsByIDValuesByIDNotFound {
	return &PutProductsByIDProductOptionsByIDValuesByIDNotFound{}
}

/*PutProductsByIDProductOptionsByIDValuesByIDNotFound handles this case with default header values.

Indicates the product is not found. or Indicates the local product option is not found.
*/
type PutProductsByIDProductOptionsByIDValuesByIDNotFound struct {
}

func (o *PutProductsByIDProductOptionsByIDValuesByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /products/{product_id}/product_options/{option_id}/values/{id}][%d] putProductsByIdProductOptionsByIdValuesByIdNotFound ", 404)
}

func (o *PutProductsByIDProductOptionsByIDValuesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutProductsByIDProductOptionsByIDValuesByIDDefault creates a PutProductsByIDProductOptionsByIDValuesByIDDefault with default headers values
func NewPutProductsByIDProductOptionsByIDValuesByIDDefault(code int) *PutProductsByIDProductOptionsByIDValuesByIDDefault {
	return &PutProductsByIDProductOptionsByIDValuesByIDDefault{
		_statusCode: code,
	}
}

/*PutProductsByIDProductOptionsByIDValuesByIDDefault handles this case with default header values.

PutProductsByIDProductOptionsByIDValuesByIDDefault put products by ID product options by ID values by ID default
*/
type PutProductsByIDProductOptionsByIDValuesByIDDefault struct {
	_statusCode int

	Payload *models.ProductOptionValue
}

// Code gets the status code for the put products by ID product options by ID values by ID default response
func (o *PutProductsByIDProductOptionsByIDValuesByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutProductsByIDProductOptionsByIDValuesByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /products/{product_id}/product_options/{option_id}/values/{id}][%d] putProductsByIDProductOptionsByIDValuesByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutProductsByIDProductOptionsByIDValuesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOptionValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
