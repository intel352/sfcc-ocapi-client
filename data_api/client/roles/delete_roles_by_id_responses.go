// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRolesByIDReader is a Reader for the DeleteRolesByID structure.
type DeleteRolesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRolesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRolesByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteRolesByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteRolesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRolesByIDNoContent creates a DeleteRolesByIDNoContent with default headers values
func NewDeleteRolesByIDNoContent() *DeleteRolesByIDNoContent {
	return &DeleteRolesByIDNoContent{}
}

/*DeleteRolesByIDNoContent handles this case with default header values.

DeleteRolesByIDNoContent delete roles by Id no content
*/
type DeleteRolesByIDNoContent struct {
}

func (o *DeleteRolesByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /roles/{id}][%d] deleteRolesByIdNoContent ", 204)
}

func (o *DeleteRolesByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRolesByIDForbidden creates a DeleteRolesByIDForbidden with default headers values
func NewDeleteRolesByIDForbidden() *DeleteRolesByIDForbidden {
	return &DeleteRolesByIDForbidden{}
}

/*DeleteRolesByIDForbidden handles this case with default header values.

Thrown if deletion of the given role is not allowed
*/
type DeleteRolesByIDForbidden struct {
}

func (o *DeleteRolesByIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /roles/{id}][%d] deleteRolesByIdForbidden ", 403)
}

func (o *DeleteRolesByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRolesByIDNotFound creates a DeleteRolesByIDNotFound with default headers values
func NewDeleteRolesByIDNotFound() *DeleteRolesByIDNotFound {
	return &DeleteRolesByIDNotFound{}
}

/*DeleteRolesByIDNotFound handles this case with default header values.

Thrown if the given role does not exist
*/
type DeleteRolesByIDNotFound struct {
}

func (o *DeleteRolesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /roles/{id}][%d] deleteRolesByIdNotFound ", 404)
}

func (o *DeleteRolesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
