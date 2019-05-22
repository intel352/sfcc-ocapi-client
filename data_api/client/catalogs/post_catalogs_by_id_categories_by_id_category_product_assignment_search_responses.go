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

// PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchReader is a Reader for the PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearch structure.
type PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest creates a PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest with default headers values
func NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest() *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest {
	return &PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest{}
}

/*PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest struct {
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/categories/{category_id}/category_product_assignment_search][%d] postCatalogsByIdCategoriesByIdCategoryProductAssignmentSearchBadRequest ", 400)
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound creates a PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound with default headers values
func NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound() *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound {
	return &PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound{}
}

/*PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound handles this case with default header values.

Thrown if the category does not exist matching the given catalog_id. or Thrown if the catalog does not exist matching the given category_id.
*/
type PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound struct {
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/categories/{category_id}/category_product_assignment_search][%d] postCatalogsByIdCategoriesByIdCategoryProductAssignmentSearchNotFound ", 404)
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault creates a PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault with default headers values
func NewPostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault(code int) *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault {
	return &PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault{
		_statusCode: code,
	}
}

/*PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault handles this case with default header values.

PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault post catalogs by ID categories by ID category product assignment search default
*/
type PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault struct {
	_statusCode int

	Payload *models.CategoryProductAssignmentSearchResult
}

// Code gets the status code for the post catalogs by ID categories by ID category product assignment search default response
func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/categories/{category_id}/category_product_assignment_search][%d] postCatalogsByIDCategoriesByIDCategoryProductAssignmentSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCatalogsByIDCategoriesByIDCategoryProductAssignmentSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryProductAssignmentSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
