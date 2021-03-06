// Code generated by go-swagger; DO NOT EDIT.

package product_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetProductListsByIDItemsByIDReader is a Reader for the GetProductListsByIDItemsByID structure.
type GetProductListsByIDItemsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductListsByIDItemsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetProductListsByIDItemsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductListsByIDItemsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductListsByIDItemsByIDNotFound creates a GetProductListsByIDItemsByIDNotFound with default headers values
func NewGetProductListsByIDItemsByIDNotFound() *GetProductListsByIDItemsByIDNotFound {
	return &GetProductListsByIDItemsByIDNotFound{}
}

/*GetProductListsByIDItemsByIDNotFound handles this case with default header values.

Indicates that this list doesn't exist. or Indicates that this product list item doesn't exist.
*/
type GetProductListsByIDItemsByIDNotFound struct {
}

func (o *GetProductListsByIDItemsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /product_lists/{list_id}/items/{item_id}][%d] getProductListsByIdItemsByIdNotFound ", 404)
}

func (o *GetProductListsByIDItemsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductListsByIDItemsByIDDefault creates a GetProductListsByIDItemsByIDDefault with default headers values
func NewGetProductListsByIDItemsByIDDefault(code int) *GetProductListsByIDItemsByIDDefault {
	return &GetProductListsByIDItemsByIDDefault{
		_statusCode: code,
	}
}

/*GetProductListsByIDItemsByIDDefault handles this case with default header values.

GetProductListsByIDItemsByIDDefault get product lists by ID items by ID default
*/
type GetProductListsByIDItemsByIDDefault struct {
	_statusCode int

	Payload *models.PublicProductListItem
}

// Code gets the status code for the get product lists by ID items by ID default response
func (o *GetProductListsByIDItemsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProductListsByIDItemsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /product_lists/{list_id}/items/{item_id}][%d] getProductListsByIDItemsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductListsByIDItemsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicProductListItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
