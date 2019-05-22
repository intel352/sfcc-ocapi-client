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

// GetCatalogsByIDCategoriesByIDCategoryLinksReader is a Reader for the GetCatalogsByIDCategoriesByIDCategoryLinks structure.
type GetCatalogsByIDCategoriesByIDCategoryLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogsByIDCategoriesByIDCategoryLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCatalogsByIDCategoriesByIDCategoryLinksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCatalogsByIDCategoriesByIDCategoryLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCatalogsByIDCategoriesByIDCategoryLinksNotFound creates a GetCatalogsByIDCategoriesByIDCategoryLinksNotFound with default headers values
func NewGetCatalogsByIDCategoriesByIDCategoryLinksNotFound() *GetCatalogsByIDCategoriesByIDCategoryLinksNotFound {
	return &GetCatalogsByIDCategoriesByIDCategoryLinksNotFound{}
}

/*GetCatalogsByIDCategoriesByIDCategoryLinksNotFound handles this case with default header values.

Thrown in case the source catalog does not exist or Thrown in case the source category does not exist
*/
type GetCatalogsByIDCategoriesByIDCategoryLinksNotFound struct {
}

func (o *GetCatalogsByIDCategoriesByIDCategoryLinksNotFound) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/categories/{category_id}/category_links][%d] getCatalogsByIdCategoriesByIdCategoryLinksNotFound ", 404)
}

func (o *GetCatalogsByIDCategoriesByIDCategoryLinksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogsByIDCategoriesByIDCategoryLinksDefault creates a GetCatalogsByIDCategoriesByIDCategoryLinksDefault with default headers values
func NewGetCatalogsByIDCategoriesByIDCategoryLinksDefault(code int) *GetCatalogsByIDCategoriesByIDCategoryLinksDefault {
	return &GetCatalogsByIDCategoriesByIDCategoryLinksDefault{
		_statusCode: code,
	}
}

/*GetCatalogsByIDCategoriesByIDCategoryLinksDefault handles this case with default header values.

GetCatalogsByIDCategoriesByIDCategoryLinksDefault get catalogs by ID categories by ID category links default
*/
type GetCatalogsByIDCategoriesByIDCategoryLinksDefault struct {
	_statusCode int

	Payload *models.CategoryLinks
}

// Code gets the status code for the get catalogs by ID categories by ID category links default response
func (o *GetCatalogsByIDCategoriesByIDCategoryLinksDefault) Code() int {
	return o._statusCode
}

func (o *GetCatalogsByIDCategoriesByIDCategoryLinksDefault) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/categories/{category_id}/category_links][%d] getCatalogsByIDCategoriesByIDCategoryLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogsByIDCategoriesByIDCategoryLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryLinks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}