// Code generated by go-swagger; DO NOT EDIT.

package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDReader is a Reader for the DeleteCatalogsByIDSharedVariationAttributesByIDValuesByID structure.
type DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent creates a DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent with default headers values
func NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent() *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent {
	return &DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent{}
}

/*DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent handles this case with default header values.

DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent delete catalogs by Id shared variation attributes by Id values by Id no content
*/
type DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent struct {
}

func (o *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/shared_variation_attributes/{attribute_id}/values/{id}][%d] deleteCatalogsByIdSharedVariationAttributesByIdValuesByIdNoContent ", 204)
}

func (o *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound creates a DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound with default headers values
func NewDeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound() *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound {
	return &DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound{}
}

/*DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound handles this case with default header values.

if the catalog id specified cannot be found or if the attribute id specified is not a valid variation attribute or if the value id specified does not a valid value id
*/
type DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound struct {
}

func (o *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/shared_variation_attributes/{attribute_id}/values/{id}][%d] deleteCatalogsByIdSharedVariationAttributesByIdValuesByIdNotFound ", 404)
}

func (o *DeleteCatalogsByIDSharedVariationAttributesByIDValuesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}