// Code generated by go-swagger; DO NOT EDIT.

package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader is a Reader for the DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByID structure.
type DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent creates a DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent with default headers values
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent() *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent {
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent{}
}

/*DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent handles this case with default header values.

DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent delete catalogs by Id categories by Id category links by Id by Id by Id no content
*/
type DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent struct {
}

func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/categories/{category_id}/category_links/{target_catalog_id}/{target_category_id}/{type}][%d] deleteCatalogsByIdCategoriesByIdCategoryLinksByIdByIdByIdNoContent ", 204)
}

func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound creates a DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound with default headers values
func NewDeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound() *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound {
	return &DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound{}
}

/*DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound handles this case with default header values.

Thrown in case the source catalog or the target catalog do not exist or Thrown in case the source category or the target category do not exist or Thrown in case the category link does not exist from the source catalog/category to the destination catalog/category with the given type.
*/
type DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound struct {
}

func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/categories/{category_id}/category_links/{target_catalog_id}/{target_category_id}/{type}][%d] deleteCatalogsByIdCategoriesByIdCategoryLinksByIdByIdByIdNotFound ", 404)
}

func (o *DeleteCatalogsByIDCategoriesByIDCategoryLinksByIDByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
