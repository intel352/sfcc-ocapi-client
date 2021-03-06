// Code generated by go-swagger; DO NOT EDIT.

package catalog_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostCatalogSearchReader is a Reader for the PostCatalogSearch structure.
type PostCatalogSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCatalogSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCatalogSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCatalogSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCatalogSearchBadRequest creates a PostCatalogSearchBadRequest with default headers values
func NewPostCatalogSearchBadRequest() *PostCatalogSearchBadRequest {
	return &PostCatalogSearchBadRequest{}
}

/*PostCatalogSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostCatalogSearchBadRequest struct {
}

func (o *PostCatalogSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /catalog_search][%d] postCatalogSearchBadRequest ", 400)
}

func (o *PostCatalogSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCatalogSearchDefault creates a PostCatalogSearchDefault with default headers values
func NewPostCatalogSearchDefault(code int) *PostCatalogSearchDefault {
	return &PostCatalogSearchDefault{
		_statusCode: code,
	}
}

/*PostCatalogSearchDefault handles this case with default header values.

PostCatalogSearchDefault post catalog search default
*/
type PostCatalogSearchDefault struct {
	_statusCode int

	Payload *models.CatalogSearchResult
}

// Code gets the status code for the post catalog search default response
func (o *PostCatalogSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCatalogSearchDefault) Error() string {
	return fmt.Sprintf("[POST /catalog_search][%d] postCatalogSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCatalogSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
