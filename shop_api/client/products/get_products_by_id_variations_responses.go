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

// GetProductsByIDVariationsReader is a Reader for the GetProductsByIDVariations structure.
type GetProductsByIDVariationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDVariationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetProductsByIDVariationsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetProductsByIDVariationsDefault creates a GetProductsByIDVariationsDefault with default headers values
func NewGetProductsByIDVariationsDefault(code int) *GetProductsByIDVariationsDefault {
	return &GetProductsByIDVariationsDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDVariationsDefault handles this case with default header values.

GetProductsByIDVariationsDefault get products by ID variations default
*/
type GetProductsByIDVariationsDefault struct {
	_statusCode int

	Payload *models.Product
}

// Code gets the status code for the get products by ID variations default response
func (o *GetProductsByIDVariationsDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDVariationsDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}/variations][%d] getProductsByIDVariations default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDVariationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}