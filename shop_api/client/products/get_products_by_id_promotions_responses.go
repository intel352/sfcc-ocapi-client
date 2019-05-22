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

// GetProductsByIDPromotionsReader is a Reader for the GetProductsByIDPromotions structure.
type GetProductsByIDPromotionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDPromotionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetProductsByIDPromotionsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetProductsByIDPromotionsDefault creates a GetProductsByIDPromotionsDefault with default headers values
func NewGetProductsByIDPromotionsDefault(code int) *GetProductsByIDPromotionsDefault {
	return &GetProductsByIDPromotionsDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDPromotionsDefault handles this case with default header values.

GetProductsByIDPromotionsDefault get products by ID promotions default
*/
type GetProductsByIDPromotionsDefault struct {
	_statusCode int

	Payload *models.Product
}

// Code gets the status code for the get products by ID promotions default response
func (o *GetProductsByIDPromotionsDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDPromotionsDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}/promotions][%d] getProductsByIDPromotions default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDPromotionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
