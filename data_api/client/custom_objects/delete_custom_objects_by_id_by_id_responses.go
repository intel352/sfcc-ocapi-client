// Code generated by go-swagger; DO NOT EDIT.

package custom_objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCustomObjectsByIDByIDReader is a Reader for the DeleteCustomObjectsByIDByID structure.
type DeleteCustomObjectsByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomObjectsByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCustomObjectsByIDByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteCustomObjectsByIDByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCustomObjectsByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomObjectsByIDByIDNoContent creates a DeleteCustomObjectsByIDByIDNoContent with default headers values
func NewDeleteCustomObjectsByIDByIDNoContent() *DeleteCustomObjectsByIDByIDNoContent {
	return &DeleteCustomObjectsByIDByIDNoContent{}
}

/*DeleteCustomObjectsByIDByIDNoContent handles this case with default header values.

DeleteCustomObjectsByIDByIDNoContent delete custom objects by Id by Id no content
*/
type DeleteCustomObjectsByIDByIDNoContent struct {
}

func (o *DeleteCustomObjectsByIDByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /custom_objects/{object_type}/{key}][%d] deleteCustomObjectsByIdByIdNoContent ", 204)
}

func (o *DeleteCustomObjectsByIDByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomObjectsByIDByIDBadRequest creates a DeleteCustomObjectsByIDByIDBadRequest with default headers values
func NewDeleteCustomObjectsByIDByIDBadRequest() *DeleteCustomObjectsByIDByIDBadRequest {
	return &DeleteCustomObjectsByIDByIDBadRequest{}
}

/*DeleteCustomObjectsByIDByIDBadRequest handles this case with default header values.

If the object key must be an integer, but the path parameter has an invalid format
*/
type DeleteCustomObjectsByIDByIDBadRequest struct {
}

func (o *DeleteCustomObjectsByIDByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /custom_objects/{object_type}/{key}][%d] deleteCustomObjectsByIdByIdBadRequest ", 400)
}

func (o *DeleteCustomObjectsByIDByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomObjectsByIDByIDNotFound creates a DeleteCustomObjectsByIDByIDNotFound with default headers values
func NewDeleteCustomObjectsByIDByIDNotFound() *DeleteCustomObjectsByIDByIDNotFound {
	return &DeleteCustomObjectsByIDByIDNotFound{}
}

/*DeleteCustomObjectsByIDByIDNotFound handles this case with default header values.

For an unknown object type ID
*/
type DeleteCustomObjectsByIDByIDNotFound struct {
}

func (o *DeleteCustomObjectsByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /custom_objects/{object_type}/{key}][%d] deleteCustomObjectsByIdByIdNotFound ", 404)
}

func (o *DeleteCustomObjectsByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
