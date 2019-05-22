// Code generated by go-swagger; DO NOT EDIT.

package custom_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostCustomObjectDefinitionsByIDAttributeGroupSearchReader is a Reader for the PostCustomObjectDefinitionsByIDAttributeGroupSearch structure.
type PostCustomObjectDefinitionsByIDAttributeGroupSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCustomObjectDefinitionsByIDAttributeGroupSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest creates a PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest with default headers values
func NewPostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest() *PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest {
	return &PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest{}
}

/*PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest struct {
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /custom_object_definitions/{object_type}/attribute_group_search][%d] postCustomObjectDefinitionsByIdAttributeGroupSearchBadRequest ", 400)
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound creates a PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound with default headers values
func NewPostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound() *PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound {
	return &PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound{}
}

/*PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound handles this case with default header values.

Thrown in case the object type cannot be found
*/
type PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound struct {
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /custom_object_definitions/{object_type}/attribute_group_search][%d] postCustomObjectDefinitionsByIdAttributeGroupSearchNotFound ", 404)
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomObjectDefinitionsByIDAttributeGroupSearchDefault creates a PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault with default headers values
func NewPostCustomObjectDefinitionsByIDAttributeGroupSearchDefault(code int) *PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault {
	return &PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault{
		_statusCode: code,
	}
}

/*PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault handles this case with default header values.

PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault post custom object definitions by ID attribute group search default
*/
type PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeGroupSearchResult
}

// Code gets the status code for the post custom object definitions by ID attribute group search default response
func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault) Error() string {
	return fmt.Sprintf("[POST /custom_object_definitions/{object_type}/attribute_group_search][%d] postCustomObjectDefinitionsByIDAttributeGroupSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomObjectDefinitionsByIDAttributeGroupSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeGroupSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}