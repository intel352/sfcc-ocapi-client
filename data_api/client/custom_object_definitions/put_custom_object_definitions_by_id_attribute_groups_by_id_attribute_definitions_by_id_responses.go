// Code generated by go-swagger; DO NOT EDIT.

package custom_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDReader is a Reader for the PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID structure.
type PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent creates a PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent with default headers values
func NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent() *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent {
	return &PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent{}
}

/*PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent handles this case with default header values.

PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent put custom object definitions by Id attribute groups by Id attribute definitions by Id no content
*/
type PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent struct {
}

func (o *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /custom_object_definitions/{object_type}/attribute_groups/{group_id}/attribute_definitions/{def_id}][%d] putCustomObjectDefinitionsByIdAttributeGroupsByIdAttributeDefinitionsByIdNoContent ", 204)
}

func (o *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound creates a PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound with default headers values
func NewPutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound() *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound {
	return &PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound{}
}

/*PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound handles this case with default header values.

Indicates the specified custom object is not found. or Indicates the specified attribute group is not found. or Indicates the specified attribute definition is not found.
*/
type PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound struct {
}

func (o *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /custom_object_definitions/{object_type}/attribute_groups/{group_id}/attribute_definitions/{def_id}][%d] putCustomObjectDefinitionsByIdAttributeGroupsByIdAttributeDefinitionsByIdNotFound ", 404)
}

func (o *PutCustomObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}