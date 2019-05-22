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

// PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader is a Reader for the PutSystemObjectDefinitionsByIDAttributeDefinitionsByID structure.
type PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest creates a PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest with default headers values
func NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest() *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest {
	return &PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest{}
}

/*PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest handles this case with default header values.

if the attribute definition could not be created. or Thrown when trying to create an internal attribute definition or if a attribute definition exists already with the given attribute definition id. or if the Id in request is not the same as the ID in document.
*/
type PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest struct {
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /system_object_definitions/{object_type}/attribute_definitions/{id}][%d] putSystemObjectDefinitionsByIdAttributeDefinitionsByIdBadRequest ", 400)
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound creates a PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound with default headers values
func NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound() *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound {
	return &PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound{}
}

/*PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound handles this case with default header values.

Thrown in case the object type cannot be found
*/
type PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound struct {
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /system_object_definitions/{object_type}/attribute_definitions/{id}][%d] putSystemObjectDefinitionsByIdAttributeDefinitionsByIdNotFound ", 404)
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault creates a PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault with default headers values
func NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault(code int) *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault {
	return &PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault{
		_statusCode: code,
	}
}

/*PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault handles this case with default header values.

PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault put system object definitions by ID attribute definitions by ID default
*/
type PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeDefinition
}

// Code gets the status code for the put system object definitions by ID attribute definitions by ID default response
func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /system_object_definitions/{object_type}/attribute_definitions/{id}][%d] putSystemObjectDefinitionsByIDAttributeDefinitionsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}