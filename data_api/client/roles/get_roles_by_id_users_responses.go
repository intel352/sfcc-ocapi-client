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

// GetRolesByIDUsersReader is a Reader for the GetRolesByIDUsers structure.
type GetRolesByIDUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesByIDUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetRolesByIDUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRolesByIDUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRolesByIDUsersNotFound creates a GetRolesByIDUsersNotFound with default headers values
func NewGetRolesByIDUsersNotFound() *GetRolesByIDUsersNotFound {
	return &GetRolesByIDUsersNotFound{}
}

/*GetRolesByIDUsersNotFound handles this case with default header values.

Thrown if the access role with the given id does not exist.
*/
type GetRolesByIDUsersNotFound struct {
}

func (o *GetRolesByIDUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /roles/{id}/users][%d] getRolesByIdUsersNotFound ", 404)
}

func (o *GetRolesByIDUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRolesByIDUsersDefault creates a GetRolesByIDUsersDefault with default headers values
func NewGetRolesByIDUsersDefault(code int) *GetRolesByIDUsersDefault {
	return &GetRolesByIDUsersDefault{
		_statusCode: code,
	}
}

/*GetRolesByIDUsersDefault handles this case with default header values.

GetRolesByIDUsersDefault get roles by ID users default
*/
type GetRolesByIDUsersDefault struct {
	_statusCode int

	Payload *models.Users
}

// Code gets the status code for the get roles by ID users default response
func (o *GetRolesByIDUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetRolesByIDUsersDefault) Error() string {
	return fmt.Sprintf("[GET /roles/{id}/users][%d] getRolesByIDUsers default  %+v", o._statusCode, o.Payload)
}

func (o *GetRolesByIDUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Users)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
