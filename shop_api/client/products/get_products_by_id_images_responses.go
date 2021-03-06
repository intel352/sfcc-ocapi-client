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

// GetProductsByIDImagesReader is a Reader for the GetProductsByIDImages structure.
type GetProductsByIDImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetProductsByIDImagesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetProductsByIDImagesDefault creates a GetProductsByIDImagesDefault with default headers values
func NewGetProductsByIDImagesDefault(code int) *GetProductsByIDImagesDefault {
	return &GetProductsByIDImagesDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDImagesDefault handles this case with default header values.

GetProductsByIDImagesDefault get products by ID images default
*/
type GetProductsByIDImagesDefault struct {
	_statusCode int

	Payload *models.Product
}

// Code gets the status code for the get products by ID images default response
func (o *GetProductsByIDImagesDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDImagesDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}/images][%d] getProductsByIDImages default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
