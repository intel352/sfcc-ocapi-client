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

// PatchCustomObjectDefinitionsByIDAttributeGroupsByIDReader is a Reader for the PatchCustomObjectDefinitionsByIDAttributeGroupsByID structure.
type PatchCustomObjectDefinitionsByIDAttributeGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound creates a PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound with default headers values
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound() *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound {
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound{}
}

/*PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound handles this case with default header values.

Thrown in case the object type cannot be found or Thrown in case the attribute group does not exist matching the given id
*/
type PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound struct {
}

func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /custom_object_definitions/{object_type}/attribute_groups/{id}][%d] patchCustomObjectDefinitionsByIdAttributeGroupsByIdNotFound ", 404)
}

func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault creates a PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault with default headers values
func NewPatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault(code int) *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault {
	return &PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault{
		_statusCode: code,
	}
}

/*PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault handles this case with default header values.

PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault patch custom object definitions by ID attribute groups by ID default
*/
type PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault struct {
	_statusCode int

	Payload *models.ObjectAttributeGroup
}

// Code gets the status code for the patch custom object definitions by ID attribute groups by ID default response
func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /custom_object_definitions/{object_type}/attribute_groups/{id}][%d] patchCustomObjectDefinitionsByIDAttributeGroupsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCustomObjectDefinitionsByIDAttributeGroupsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectAttributeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}