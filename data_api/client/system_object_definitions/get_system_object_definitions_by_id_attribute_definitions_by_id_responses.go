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

// GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader is a Reader for the GetSystemObjectDefinitionsByIDAttributeDefinitionsByID structure.
type GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound creates a GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound with default headers values
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound() *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound {
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound{}
}

/*GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound handles this case with default header values.

in case the object type does not match an existing system type or Thrown in case the attribute definition does not exist matching the given id
*/
type GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound struct {
}

func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /system_object_definitions/{object_type}/attribute_definitions/{id}][%d] getSystemObjectDefinitionsByIdAttributeDefinitionsByIdNotFound ", 404)
}

func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault creates a GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault with default headers values
func NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault(code int) *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault {
	return &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault{
		_statusCode: code,
	}
}

/*GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault handles this case with default header values.

GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault get system object definitions by ID attribute definitions by ID default
*/
type GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeDefinition
}

// Code gets the status code for the get system object definitions by ID attribute definitions by ID default response
func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /system_object_definitions/{object_type}/attribute_definitions/{id}][%d] getSystemObjectDefinitionsByIDAttributeDefinitionsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
