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

// PatchProductsByIDProductOptionsByIDValuesByIDReader is a Reader for the PatchProductsByIDProductOptionsByIDValuesByID structure.
type PatchProductsByIDProductOptionsByIDValuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProductsByIDProductOptionsByIDValuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchProductsByIDProductOptionsByIDValuesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchProductsByIDProductOptionsByIDValuesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchProductsByIDProductOptionsByIDValuesByIDNotFound creates a PatchProductsByIDProductOptionsByIDValuesByIDNotFound with default headers values
func NewPatchProductsByIDProductOptionsByIDValuesByIDNotFound() *PatchProductsByIDProductOptionsByIDValuesByIDNotFound {
	return &PatchProductsByIDProductOptionsByIDValuesByIDNotFound{}
}

/*PatchProductsByIDProductOptionsByIDValuesByIDNotFound handles this case with default header values.

Indicates the product is not found. or Indicates the local product option is not found. or Indicates the local product option value is not found.
*/
type PatchProductsByIDProductOptionsByIDValuesByIDNotFound struct {
}

func (o *PatchProductsByIDProductOptionsByIDValuesByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /products/{product_id}/product_options/{option_id}/values/{id}][%d] patchProductsByIdProductOptionsByIdValuesByIdNotFound ", 404)
}

func (o *PatchProductsByIDProductOptionsByIDValuesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProductsByIDProductOptionsByIDValuesByIDDefault creates a PatchProductsByIDProductOptionsByIDValuesByIDDefault with default headers values
func NewPatchProductsByIDProductOptionsByIDValuesByIDDefault(code int) *PatchProductsByIDProductOptionsByIDValuesByIDDefault {
	return &PatchProductsByIDProductOptionsByIDValuesByIDDefault{
		_statusCode: code,
	}
}

/*PatchProductsByIDProductOptionsByIDValuesByIDDefault handles this case with default header values.

PatchProductsByIDProductOptionsByIDValuesByIDDefault patch products by ID product options by ID values by ID default
*/
type PatchProductsByIDProductOptionsByIDValuesByIDDefault struct {
	_statusCode int

	Payload *models.ProductOptionValue
}

// Code gets the status code for the patch products by ID product options by ID values by ID default response
func (o *PatchProductsByIDProductOptionsByIDValuesByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchProductsByIDProductOptionsByIDValuesByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /products/{product_id}/product_options/{option_id}/values/{id}][%d] patchProductsByIDProductOptionsByIDValuesByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProductsByIDProductOptionsByIDValuesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOptionValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
