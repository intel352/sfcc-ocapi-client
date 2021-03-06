// Code generated by go-swagger; DO NOT EDIT.

package system_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostSystemObjectDefinitionsByIDAttributeGroupSearchReader is a Reader for the PostSystemObjectDefinitionsByIDAttributeGroupSearch structure.
type PostSystemObjectDefinitionsByIDAttributeGroupSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSystemObjectDefinitionsByIDAttributeGroupSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest creates a PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest with default headers values
func NewPostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest() *PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest {
	return &PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest{}
}

/*PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest struct {
}

func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /system_object_definitions/{object_type}/attribute_group_search][%d] postSystemObjectDefinitionsByIdAttributeGroupSearchBadRequest ", 400)
}

func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSystemObjectDefinitionsByIDAttributeGroupSearchDefault creates a PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault with default headers values
func NewPostSystemObjectDefinitionsByIDAttributeGroupSearchDefault(code int) *PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault {
	return &PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault{
		_statusCode: code,
	}
}

/*PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault handles this case with default header values.

PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault post system object definitions by ID attribute group search default
*/
type PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeGroupSearchResult
}

// Code gets the status code for the post system object definitions by ID attribute group search default response
func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault) Error() string {
	return fmt.Sprintf("[POST /system_object_definitions/{object_type}/attribute_group_search][%d] postSystemObjectDefinitionsByIDAttributeGroupSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostSystemObjectDefinitionsByIDAttributeGroupSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeGroupSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
