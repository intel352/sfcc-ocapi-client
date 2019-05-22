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

// GetPermissionsBmLocaleReader is a Reader for the GetPermissionsBmLocale structure.
type GetPermissionsBmLocaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsBmLocaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetPermissionsBmLocaleDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetPermissionsBmLocaleDefault creates a GetPermissionsBmLocaleDefault with default headers values
func NewGetPermissionsBmLocaleDefault(code int) *GetPermissionsBmLocaleDefault {
	return &GetPermissionsBmLocaleDefault{
		_statusCode: code,
	}
}

/*GetPermissionsBmLocaleDefault handles this case with default header values.

GetPermissionsBmLocaleDefault get permissions bm locale default
*/
type GetPermissionsBmLocaleDefault struct {
	_statusCode int

	Payload *models.LocalePermissions
}

// Code gets the status code for the get permissions bm locale default response
func (o *GetPermissionsBmLocaleDefault) Code() int {
	return o._statusCode
}

func (o *GetPermissionsBmLocaleDefault) Error() string {
	return fmt.Sprintf("[GET /permissions/bm/locale][%d] getPermissionsBmLocale default  %+v", o._statusCode, o.Payload)
}

func (o *GetPermissionsBmLocaleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalePermissions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}