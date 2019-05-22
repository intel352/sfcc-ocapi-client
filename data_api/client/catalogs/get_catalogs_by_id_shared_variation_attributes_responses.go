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

// GetCatalogsByIDSharedVariationAttributesReader is a Reader for the GetCatalogsByIDSharedVariationAttributes structure.
type GetCatalogsByIDSharedVariationAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogsByIDSharedVariationAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCatalogsByIDSharedVariationAttributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCatalogsByIDSharedVariationAttributesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCatalogsByIDSharedVariationAttributesNotFound creates a GetCatalogsByIDSharedVariationAttributesNotFound with default headers values
func NewGetCatalogsByIDSharedVariationAttributesNotFound() *GetCatalogsByIDSharedVariationAttributesNotFound {
	return &GetCatalogsByIDSharedVariationAttributesNotFound{}
}

/*GetCatalogsByIDSharedVariationAttributesNotFound handles this case with default header values.

Indicates the catalog is not found.
*/
type GetCatalogsByIDSharedVariationAttributesNotFound struct {
}

func (o *GetCatalogsByIDSharedVariationAttributesNotFound) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/shared_variation_attributes][%d] getCatalogsByIdSharedVariationAttributesNotFound ", 404)
}

func (o *GetCatalogsByIDSharedVariationAttributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogsByIDSharedVariationAttributesDefault creates a GetCatalogsByIDSharedVariationAttributesDefault with default headers values
func NewGetCatalogsByIDSharedVariationAttributesDefault(code int) *GetCatalogsByIDSharedVariationAttributesDefault {
	return &GetCatalogsByIDSharedVariationAttributesDefault{
		_statusCode: code,
	}
}

/*GetCatalogsByIDSharedVariationAttributesDefault handles this case with default header values.

GetCatalogsByIDSharedVariationAttributesDefault get catalogs by ID shared variation attributes default
*/
type GetCatalogsByIDSharedVariationAttributesDefault struct {
	_statusCode int

	Payload *models.VariationAttributes
}

// Code gets the status code for the get catalogs by ID shared variation attributes default response
func (o *GetCatalogsByIDSharedVariationAttributesDefault) Code() int {
	return o._statusCode
}

func (o *GetCatalogsByIDSharedVariationAttributesDefault) Error() string {
	return fmt.Sprintf("[GET /catalogs/{catalog_id}/shared_variation_attributes][%d] getCatalogsByIDSharedVariationAttributes default  %+v", o._statusCode, o.Payload)
}

func (o *GetCatalogsByIDSharedVariationAttributesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariationAttributes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
