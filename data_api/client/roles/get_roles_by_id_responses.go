// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetRolesByIDReader is a Reader for the GetRolesByID structure.
type GetRolesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetRolesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRolesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRolesByIDNotFound creates a GetRolesByIDNotFound with default headers values
func NewGetRolesByIDNotFound() *GetRolesByIDNotFound {
	return &GetRolesByIDNotFound{}
}

/*GetRolesByIDNotFound handles this case with default header values.

Thrown if the access role with the given id does not exist.
*/
type GetRolesByIDNotFound struct {
}

func (o *GetRolesByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /roles/{id}][%d] getRolesByIdNotFound ", 404)
}

func (o *GetRolesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesByIDDefault creates a GetRolesByIDDefault with default headers values
func NewGetRolesByIDDefault(code int) *GetRolesByIDDefault {
	return &GetRolesByIDDefault{
		_statusCode: code,
	}
}

/*GetRolesByIDDefault handles this case with default header values.

GetRolesByIDDefault get roles by ID default
*/
type GetRolesByIDDefault struct {
	_statusCode int

	Payload *models.Role
}

// Code gets the status code for the get roles by ID default response
func (o *GetRolesByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRolesByIDDefault) Error() string {
	return fmt.Sprintf("[GET /roles/{id}][%d] getRolesByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetRolesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
