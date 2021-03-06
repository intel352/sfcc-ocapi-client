// Code generated by go-swagger; DO NOT EDIT.

package custom_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader is a Reader for the DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByID structure.
type DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent creates a DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent with default headers values
func NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent() *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent {
	return &DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent{}
}

/*DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent handles this case with default header values.

DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent delete custom object definitions by Id attribute definitions by Id no content
*/
type DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent struct {
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] deleteCustomObjectDefinitionsByIdAttributeDefinitionsByIdNoContent ", 204)
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest creates a DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest with default headers values
func NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest() *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest {
	return &DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest{}
}

/*DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest handles this case with default header values.

Thrown when trying to delete a internal attribute of custom object or Thrown when trying to delete a key attribute of custom object
*/
type DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest struct {
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] deleteCustomObjectDefinitionsByIdAttributeDefinitionsByIdBadRequest ", 400)
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound creates a DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound with default headers values
func NewDeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound() *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound {
	return &DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound{}
}

/*DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound handles this case with default header values.

Thrown in case the object type cannot be found or Thrown in case the attribute definition does not exist matching the given id
*/
type DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound struct {
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /custom_object_definitions/{object_type}/attribute_definitions/{id}][%d] deleteCustomObjectDefinitionsByIdAttributeDefinitionsByIdNotFound ", 404)
}

func (o *DeleteCustomObjectDefinitionsByIDAttributeDefinitionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
