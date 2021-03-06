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

// PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader is a Reader for the PutCustomObjectDefinitionsByIDAttributeDefinitionsByID structure.
type PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest creates a PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest with default headers values
func NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest() *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest {
	return &PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest{}
}

/*PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest handles this case with default header values.

Thrown when trying to create a key attribute of a custom object or if the attribute definition could not be created. or if trying to create a localizable attribute with a non-localizable type or if the Id in request is not the same as the ID in document.
*/
type PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest struct {
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] putCustomObjectDefinitionsByIdAttributeDefinitionsByIdBadRequest ", 400)
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound creates a PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound with default headers values
func NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound() *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound {
	return &PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound{}
}

/*PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound handles this case with default header values.

Thrown in case the object type cannot be found
*/
type PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound struct {
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] putCustomObjectDefinitionsByIdAttributeDefinitionsByIdNotFound ", 404)
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault creates a PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault with default headers values
func NewPutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault(code int) *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault {
	return &PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault{
		_statusCode: code,
	}
}

/*PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault handles this case with default header values.

PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault put custom object definitions by ID attribute definitions by ID default
*/
type PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeDefinition
}

// Code gets the status code for the put custom object definitions by ID attribute definitions by ID default response
func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] putCustomObjectDefinitionsByIDAttributeDefinitionsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutCustomObjectDefinitionsByIDAttributeDefinitionsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
