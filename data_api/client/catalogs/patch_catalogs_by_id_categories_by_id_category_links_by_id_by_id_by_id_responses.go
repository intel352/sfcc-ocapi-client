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

// PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader is a Reader for the PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByID structure.
type PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound creates a PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound with default headers values
func NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound() *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound {
	return &PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound{}
}

/*PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound handles this case with default header values.

Thrown in case the source catalog or the target catalog do not exist or Thrown in case the source category or the target category do not exist or Thrown in case the category link does not exist from the source catalog/category to the destination catalog/category with the given type.
*/
type PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound struct {
}

func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}/categories/{category_id}/category_links/{target_catalog_id}/{target_category_id}/{type}][%d] patchCatalogsByIdCategoriesByIdCategoryLinksByIdByIdByIdNotFound ", 404)
}

func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault creates a PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault with default headers values
func NewPatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault(code int) *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault {
	return &PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault{
		_statusCode: code,
	}
}

/*PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault handles this case with default header values.

PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault patch catalogs by ID categories by ID category links by ID by ID by ID default
*/
type PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault struct {
	_statusCode int

	Payload *models.CategoryLink
}

// Code gets the status code for the patch catalogs by ID categories by ID category links by ID by ID by ID default response
func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}/categories/{category_id}/category_links/{target_catalog_id}/{target_category_id}/{type}][%d] patchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}