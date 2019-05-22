// Code generated by go-swagger; DO NOT EDIT.

package category_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostCategorySearchReader is a Reader for the PostCategorySearch structure.
type PostCategorySearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCategorySearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCategorySearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCategorySearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCategorySearchBadRequest creates a PostCategorySearchBadRequest with default headers values
func NewPostCategorySearchBadRequest() *PostCategorySearchBadRequest {
	return &PostCategorySearchBadRequest{}
}

/*PostCategorySearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostCategorySearchBadRequest struct {
}

func (o *PostCategorySearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /category_search][%d] postCategorySearchBadRequest ", 400)
}

func (o *PostCategorySearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCategorySearchDefault creates a PostCategorySearchDefault with default headers values
func NewPostCategorySearchDefault(code int) *PostCategorySearchDefault {
	return &PostCategorySearchDefault{
		_statusCode: code,
	}
}

/*PostCategorySearchDefault handles this case with default header values.

PostCategorySearchDefault post category search default
*/
type PostCategorySearchDefault struct {
	_statusCode int

	Payload *models.CategorySearchResult
}

// Code gets the status code for the post category search default response
func (o *PostCategorySearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCategorySearchDefault) Error() string {
	return fmt.Sprintf("[POST /category_search][%d] postCategorySearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCategorySearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategorySearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}