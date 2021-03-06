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

// PostCatalogsByIDCategorySearchReader is a Reader for the PostCatalogsByIDCategorySearch structure.
type PostCatalogsByIDCategorySearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCatalogsByIDCategorySearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCatalogsByIDCategorySearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCatalogsByIDCategorySearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCatalogsByIDCategorySearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCatalogsByIDCategorySearchBadRequest creates a PostCatalogsByIDCategorySearchBadRequest with default headers values
func NewPostCatalogsByIDCategorySearchBadRequest() *PostCatalogsByIDCategorySearchBadRequest {
	return &PostCatalogsByIDCategorySearchBadRequest{}
}

/*PostCatalogsByIDCategorySearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostCatalogsByIDCategorySearchBadRequest struct {
}

func (o *PostCatalogsByIDCategorySearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/category_search][%d] postCatalogsByIdCategorySearchBadRequest ", 400)
}

func (o *PostCatalogsByIDCategorySearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCatalogsByIDCategorySearchNotFound creates a PostCatalogsByIDCategorySearchNotFound with default headers values
func NewPostCatalogsByIDCategorySearchNotFound() *PostCatalogsByIDCategorySearchNotFound {
	return &PostCatalogsByIDCategorySearchNotFound{}
}

/*PostCatalogsByIDCategorySearchNotFound handles this case with default header values.

Thrown if the specified catalog does not exist.
*/
type PostCatalogsByIDCategorySearchNotFound struct {
}

func (o *PostCatalogsByIDCategorySearchNotFound) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/category_search][%d] postCatalogsByIdCategorySearchNotFound ", 404)
}

func (o *PostCatalogsByIDCategorySearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCatalogsByIDCategorySearchDefault creates a PostCatalogsByIDCategorySearchDefault with default headers values
func NewPostCatalogsByIDCategorySearchDefault(code int) *PostCatalogsByIDCategorySearchDefault {
	return &PostCatalogsByIDCategorySearchDefault{
		_statusCode: code,
	}
}

/*PostCatalogsByIDCategorySearchDefault handles this case with default header values.

PostCatalogsByIDCategorySearchDefault post catalogs by ID category search default
*/
type PostCatalogsByIDCategorySearchDefault struct {
	_statusCode int

	Payload *models.CategorySearchResult
}

// Code gets the status code for the post catalogs by ID category search default response
func (o *PostCatalogsByIDCategorySearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCatalogsByIDCategorySearchDefault) Error() string {
	return fmt.Sprintf("[POST /catalogs/{catalog_id}/category_search][%d] postCatalogsByIDCategorySearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCatalogsByIDCategorySearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategorySearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
