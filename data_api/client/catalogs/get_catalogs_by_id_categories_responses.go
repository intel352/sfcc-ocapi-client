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

// GetCatalogsByIDCategoriesReader is a Reader for the GetCatalogsByIDCategories structure.
type GetCatalogsByIDCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogsByIDCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCatalogsByIDCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCatalogsByIDCategoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCatalogsByIDCategoriesNotFound creates a GetCatalogsByIDCategoriesNotFound with default headers values
func NewGetCatalogsByIDCategoriesNotFound() *GetCatalogsByIDCategoriesNotFound {
	return &GetCatalogsByIDCategoriesNotFound{}
}

/*GetCatalogsByIDCategoriesNotFound handles this case with default header values.

Indicates that the catalog is not provided in the request.
*/
type GetCatalogsByIDCategoriesNotFound struct {
}

func (o *GetCatalogsByIDCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/categories][%d] getCatalogsByIdCategoriesNotFound ", 404)
}

func (o *GetCatalogsByIDCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogsByIDCategoriesDefault creates a GetCatalogsByIDCategoriesDefault with default headers values
func NewGetCatalogsByIDCategoriesDefault(code int) *GetCatalogsByIDCategoriesDefault {
	return &GetCatalogsByIDCategoriesDefault{
		_statusCode: code,
	}
}

/*GetCatalogsByIDCategoriesDefault handles this case with default header values.

GetCatalogsByIDCategoriesDefault get catalogs by ID categories default
*/
type GetCatalogsByIDCategoriesDefault struct {
	_statusCode int

	Payload *models.Categories
}

// Code gets the status code for the get catalogs by ID categories default response
func (o *GetCatalogsByIDCategoriesDefault) Code() int {
	return o._statusCode
}

func (o *GetCatalogsByIDCategoriesDefault) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/categories][%d] getCatalogsByIDCategories default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogsByIDCategoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Categories)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}