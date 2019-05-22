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

// GetProductsByIDVariationsByIDReader is a Reader for the GetProductsByIDVariationsByID structure.
type GetProductsByIDVariationsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDVariationsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetProductsByIDVariationsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProductsByIDVariationsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductsByIDVariationsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsByIDVariationsByIDBadRequest creates a GetProductsByIDVariationsByIDBadRequest with default headers values
func NewGetProductsByIDVariationsByIDBadRequest() *GetProductsByIDVariationsByIDBadRequest {
	return &GetProductsByIDVariationsByIDBadRequest{}
}

/*GetProductsByIDVariationsByIDBadRequest handles this case with default header values.

Indicates the master product id does not represent a master product. or Indicates the product specified is not a variation.
*/
type GetProductsByIDVariationsByIDBadRequest struct {
}

func (o *GetProductsByIDVariationsByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variations/{id}][%d] getProductsByIdVariationsByIdBadRequest ", 400)
}

func (o *GetProductsByIDVariationsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDVariationsByIDNotFound creates a GetProductsByIDVariationsByIDNotFound with default headers values
func NewGetProductsByIDVariationsByIDNotFound() *GetProductsByIDVariationsByIDNotFound {
	return &GetProductsByIDVariationsByIDNotFound{}
}

/*GetProductsByIDVariationsByIDNotFound handles this case with default header values.

Indicates either the master product or the variation product cannot be found. or Indicates the product does not belong to the master product.
*/
type GetProductsByIDVariationsByIDNotFound struct {
}

func (o *GetProductsByIDVariationsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variations/{id}][%d] getProductsByIdVariationsByIdNotFound ", 404)
}

func (o *GetProductsByIDVariationsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDVariationsByIDDefault creates a GetProductsByIDVariationsByIDDefault with default headers values
func NewGetProductsByIDVariationsByIDDefault(code int) *GetProductsByIDVariationsByIDDefault {
	return &GetProductsByIDVariationsByIDDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDVariationsByIDDefault handles this case with default header values.

GetProductsByIDVariationsByIDDefault get products by ID variations by ID default
*/
type GetProductsByIDVariationsByIDDefault struct {
	_statusCode int

	Payload *models.Variant
}

// Code gets the status code for the get products by ID variations by ID default response
func (o *GetProductsByIDVariationsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDVariationsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variations/{id}][%d] getProductsByIDVariationsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDVariationsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}