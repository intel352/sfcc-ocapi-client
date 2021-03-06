// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostSitesByIDSourceCodeGroupSearchReader is a Reader for the PostSitesByIDSourceCodeGroupSearch structure.
type PostSitesByIDSourceCodeGroupSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesByIDSourceCodeGroupSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostSitesByIDSourceCodeGroupSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSitesByIDSourceCodeGroupSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSitesByIDSourceCodeGroupSearchBadRequest creates a PostSitesByIDSourceCodeGroupSearchBadRequest with default headers values
func NewPostSitesByIDSourceCodeGroupSearchBadRequest() *PostSitesByIDSourceCodeGroupSearchBadRequest {
	return &PostSitesByIDSourceCodeGroupSearchBadRequest{}
}

/*PostSitesByIDSourceCodeGroupSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostSitesByIDSourceCodeGroupSearchBadRequest struct {
}

func (o *PostSitesByIDSourceCodeGroupSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/source_code_group_search][%d] postSitesByIdSourceCodeGroupSearchBadRequest ", 400)
}

func (o *PostSitesByIDSourceCodeGroupSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitesByIDSourceCodeGroupSearchDefault creates a PostSitesByIDSourceCodeGroupSearchDefault with default headers values
func NewPostSitesByIDSourceCodeGroupSearchDefault(code int) *PostSitesByIDSourceCodeGroupSearchDefault {
	return &PostSitesByIDSourceCodeGroupSearchDefault{
		_statusCode: code,
	}
}

/*PostSitesByIDSourceCodeGroupSearchDefault handles this case with default header values.

PostSitesByIDSourceCodeGroupSearchDefault post sites by ID source code group search default
*/
type PostSitesByIDSourceCodeGroupSearchDefault struct {
	_statusCode int

	Payload *models.SourceCodeGroupSearchResult
}

// Code gets the status code for the post sites by ID source code group search default response
func (o *PostSitesByIDSourceCodeGroupSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostSitesByIDSourceCodeGroupSearchDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/source_code_group_search][%d] postSitesByIDSourceCodeGroupSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitesByIDSourceCodeGroupSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceCodeGroupSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
