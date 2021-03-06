// Code generated by go-swagger; DO NOT EDIT.

package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCatalogsByIDCategoriesByIDProductsByIDReader is a Reader for the DeleteCatalogsByIDCategoriesByIDProductsByID structure.
type DeleteCatalogsByIDCategoriesByIDProductsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCatalogsByIDCategoriesByIDProductsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCatalogsByIDCategoriesByIDProductsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteCatalogsByIDCategoriesByIDProductsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCatalogsByIDCategoriesByIDProductsByIDNoContent creates a DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent with default headers values
func NewDeleteCatalogsByIDCategoriesByIDProductsByIDNoContent() *DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent {
	return &DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent{}
}

/*DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent handles this case with default header values.

DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent delete catalogs by Id categories by Id products by Id no content
*/
type DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent struct {
}

func (o *DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/categories/{category_id}/products/{product_id}][%d] deleteCatalogsByIdCategoriesByIdProductsByIdNoContent ", 204)
}

func (o *DeleteCatalogsByIDCategoriesByIDProductsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCatalogsByIDCategoriesByIDProductsByIDNotFound creates a DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound with default headers values
func NewDeleteCatalogsByIDCategoriesByIDProductsByIDNotFound() *DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound {
	return &DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound{}
}

/*DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound handles this case with default header values.

Thrown if the category does not exist matching the given id. or Thrown if the catalog does not exist matching the given id.
*/
type DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound struct {
}

func (o *DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /catalogs/{catalog_id}/categories/{category_id}/products/{product_id}][%d] deleteCatalogsByIdCategoriesByIdProductsByIdNotFound ", 404)
}

func (o *DeleteCatalogsByIDCategoriesByIDProductsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
