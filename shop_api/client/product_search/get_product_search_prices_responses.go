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

// GetProductSearchPricesReader is a Reader for the GetProductSearchPrices structure.
type GetProductSearchPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductSearchPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetProductSearchPricesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductSearchPricesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductSearchPricesBadRequest creates a GetProductSearchPricesBadRequest with default headers values
func NewGetProductSearchPricesBadRequest() *GetProductSearchPricesBadRequest {
	return &GetProductSearchPricesBadRequest{}
}

/*GetProductSearchPricesBadRequest handles this case with default header values.

Thrown if a price refinement parameter is malformed.
*/
type GetProductSearchPricesBadRequest struct {
}

func (o *GetProductSearchPricesBadRequest) Error() string {
	return fmt.Sprintf("[GET /product_search/prices][%d] getProductSearchPricesBadRequest ", 400)
}

func (o *GetProductSearchPricesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductSearchPricesDefault creates a GetProductSearchPricesDefault with default headers values
func NewGetProductSearchPricesDefault(code int) *GetProductSearchPricesDefault {
	return &GetProductSearchPricesDefault{
		_statusCode: code,
	}
}

/*GetProductSearchPricesDefault handles this case with default header values.

GetProductSearchPricesDefault get product search prices default
*/
type GetProductSearchPricesDefault struct {
	_statusCode int

	Payload *models.ProductSearchResult
}

// Code gets the status code for the get product search prices default response
func (o *GetProductSearchPricesDefault) Code() int {
	return o._statusCode
}

func (o *GetProductSearchPricesDefault) Error() string {
	return fmt.Sprintf("[GET /product_search/prices][%d] getProductSearchPrices default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductSearchPricesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}