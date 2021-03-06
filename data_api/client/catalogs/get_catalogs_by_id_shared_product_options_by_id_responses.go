// Code generated by go-swagger; DO NOT EDIT.

package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetCatalogsByIDSharedProductOptionsByIDReader is a Reader for the GetCatalogsByIDSharedProductOptionsByID structure.
type GetCatalogsByIDSharedProductOptionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogsByIDSharedProductOptionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCatalogsByIDSharedProductOptionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCatalogsByIDSharedProductOptionsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCatalogsByIDSharedProductOptionsByIDNotFound creates a GetCatalogsByIDSharedProductOptionsByIDNotFound with default headers values
func NewGetCatalogsByIDSharedProductOptionsByIDNotFound() *GetCatalogsByIDSharedProductOptionsByIDNotFound {
	return &GetCatalogsByIDSharedProductOptionsByIDNotFound{}
}

/*GetCatalogsByIDSharedProductOptionsByIDNotFound handles this case with default header values.

Indicates the shared product option is not found. or Indicates the catalog is not found.
*/
type GetCatalogsByIDSharedProductOptionsByIDNotFound struct {
}

func (o *GetCatalogsByIDSharedProductOptionsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/shared_product_options/{id}][%d] getCatalogsByIdSharedProductOptionsByIdNotFound ", 404)
}

func (o *GetCatalogsByIDSharedProductOptionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogsByIDSharedProductOptionsByIDDefault creates a GetCatalogsByIDSharedProductOptionsByIDDefault with default headers values
func NewGetCatalogsByIDSharedProductOptionsByIDDefault(code int) *GetCatalogsByIDSharedProductOptionsByIDDefault {
	return &GetCatalogsByIDSharedProductOptionsByIDDefault{
		_statusCode: code,
	}
}

/*GetCatalogsByIDSharedProductOptionsByIDDefault handles this case with default header values.

GetCatalogsByIDSharedProductOptionsByIDDefault get catalogs by ID shared product options by ID default
*/
type GetCatalogsByIDSharedProductOptionsByIDDefault struct {
	_statusCode int

	Payload *models.ProductOption
}

// Code gets the status code for the get catalogs by ID shared product options by ID default response
func (o *GetCatalogsByIDSharedProductOptionsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCatalogsByIDSharedProductOptionsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/shared_product_options/{id}][%d] getCatalogsByIDSharedProductOptionsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogsByIDSharedProductOptionsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductOption)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
