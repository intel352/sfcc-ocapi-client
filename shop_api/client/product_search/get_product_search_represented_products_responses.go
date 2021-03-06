// Code generated by go-swagger; DO NOT EDIT.

package product_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetProductSearchRepresentedProductsReader is a Reader for the GetProductSearchRepresentedProducts structure.
type GetProductSearchRepresentedProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductSearchRepresentedProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetProductSearchRepresentedProductsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductSearchRepresentedProductsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductSearchRepresentedProductsBadRequest creates a GetProductSearchRepresentedProductsBadRequest with default headers values
func NewGetProductSearchRepresentedProductsBadRequest() *GetProductSearchRepresentedProductsBadRequest {
	return &GetProductSearchRepresentedProductsBadRequest{}
}

/*GetProductSearchRepresentedProductsBadRequest handles this case with default header values.

Thrown if a price refinement parameter is malformed.
*/
type GetProductSearchRepresentedProductsBadRequest struct {
}

func (o *GetProductSearchRepresentedProductsBadRequest) Error() string {
	return fmt.Sprintf("[GET /product_search/represented_products][%d] getProductSearchRepresentedProductsBadRequest ", 400)
}

func (o *GetProductSearchRepresentedProductsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductSearchRepresentedProductsDefault creates a GetProductSearchRepresentedProductsDefault with default headers values
func NewGetProductSearchRepresentedProductsDefault(code int) *GetProductSearchRepresentedProductsDefault {
	return &GetProductSearchRepresentedProductsDefault{
		_statusCode: code,
	}
}

/*GetProductSearchRepresentedProductsDefault handles this case with default header values.

GetProductSearchRepresentedProductsDefault get product search represented products default
*/
type GetProductSearchRepresentedProductsDefault struct {
	_statusCode int

	Payload *models.ProductSearchResult
}

// Code gets the status code for the get product search represented products default response
func (o *GetProductSearchRepresentedProductsDefault) Code() int {
	return o._statusCode
}

func (o *GetProductSearchRepresentedProductsDefault) Error() string {
	return fmt.Sprintf("[GET /product_search/represented_products][%d] getProductSearchRepresentedProducts default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductSearchRepresentedProductsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
