// Code generated by go-swagger; DO NOT EDIT.

package system_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDReader is a Reader for the DeleteSystemObjectDefinitionsByIDAttributeGroupsByID structure.
type DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent creates a DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent with default headers values
func NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent() *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent {
	return &DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent{}
}

/*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent handles this case with default header values.

DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent delete system object definitions by Id attribute groups by Id no content
*/
type DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent struct {
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /system_object_definitions/{object_type}/attribute_groups/{id}][%d] deleteSystemObjectDefinitionsByIdAttributeGroupsByIdNoContent ", 204)
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest creates a DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest with default headers values
func NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest() *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest {
	return &DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest{}
}

/*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest handles this case with default header values.

Thrown when trying to delete an internal attribute group
*/
type DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest struct {
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /system_object_definitions/{object_type}/attribute_groups/{id}][%d] deleteSystemObjectDefinitionsByIdAttributeGroupsByIdBadRequest ", 400)
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound creates a DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound with default headers values
func NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound() *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound {
	return &DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound{}
}

/*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound handles this case with default header values.

Thrown in case the object type cannot be found or Thrown in case the attribute group cannot be found
*/
type DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound struct {
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /system_object_definitions/{object_type}/attribute_groups/{id}][%d] deleteSystemObjectDefinitionsByIdAttributeGroupsByIdNotFound ", 404)
}

func (o *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
