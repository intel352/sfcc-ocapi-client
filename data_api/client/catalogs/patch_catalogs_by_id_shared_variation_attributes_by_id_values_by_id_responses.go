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

// PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDReader is a Reader for the PatchCatalogsByIDSharedVariationAttributesByIDValuesByID structure.
type PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound creates a PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound with default headers values
func NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound() *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound {
	return &PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound{}
}

/*PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound handles this case with default header values.

if the catalog id specified cannot be found or if the attribute id specified is not a valid variation attribute or if the value id specified does not a valid value id
*/
type PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound struct {
}

func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}/shared_variation_attributes/{attribute_id}/values/{id}][%d] patchCatalogsByIdSharedVariationAttributesByIdValuesByIdNotFound ", 404)
}

func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault creates a PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault with default headers values
func NewPatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault(code int) *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault {
	return &PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault{
		_statusCode: code,
	}
}

/*PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault handles this case with default header values.

PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault patch catalogs by ID shared variation attributes by ID values by ID default
*/
type PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault struct {
	_statusCode int

	Payload *models.VariationAttributeValue
}

// Code gets the status code for the patch catalogs by ID shared variation attributes by ID values by ID default response
func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}/shared_variation_attributes/{attribute_id}/values/{id}][%d] patchCatalogsByIDSharedVariationAttributesByIDValuesByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCatalogsByIDSharedVariationAttributesByIDValuesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariationAttributeValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
