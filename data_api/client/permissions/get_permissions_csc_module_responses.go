// Code generated by go-swagger; DO NOT EDIT.

package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetPermissionsCscModuleReader is a Reader for the GetPermissionsCscModule structure.
type GetPermissionsCscModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsCscModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetPermissionsCscModuleDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetPermissionsCscModuleDefault creates a GetPermissionsCscModuleDefault with default headers values
func NewGetPermissionsCscModuleDefault(code int) *GetPermissionsCscModuleDefault {
	return &GetPermissionsCscModuleDefault{
		_statusCode: code,
	}
}

/*GetPermissionsCscModuleDefault handles this case with default header values.

GetPermissionsCscModuleDefault get permissions csc module default
*/
type GetPermissionsCscModuleDefault struct {
	_statusCode int

	Payload *models.ModulePermissions
}

// Code gets the status code for the get permissions csc module default response
func (o *GetPermissionsCscModuleDefault) Code() int {
	return o._statusCode
}

func (o *GetPermissionsCscModuleDefault) Error() string {
	return fmt.Sprintf("[GET /permissions/csc/module][%d] getPermissionsCscModule default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionsCscModuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModulePermissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
