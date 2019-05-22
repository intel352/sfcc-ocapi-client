// Code generated by go-swagger; DO NOT EDIT.

package code_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetCodeVersionsByIDReader is a Reader for the GetCodeVersionsByID structure.
type GetCodeVersionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCodeVersionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCodeVersionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCodeVersionsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCodeVersionsByIDNotFound creates a GetCodeVersionsByIDNotFound with default headers values
func NewGetCodeVersionsByIDNotFound() *GetCodeVersionsByIDNotFound {
	return &GetCodeVersionsByIDNotFound{}
}

/*GetCodeVersionsByIDNotFound handles this case with default header values.

Indicates that a code version with the given id was not found.
*/
type GetCodeVersionsByIDNotFound struct {
}

func (o *GetCodeVersionsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /code_versions/{code_version_id}][%d] getCodeVersionsByIdNotFound ", 404)
}

func (o *GetCodeVersionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCodeVersionsByIDDefault creates a GetCodeVersionsByIDDefault with default headers values
func NewGetCodeVersionsByIDDefault(code int) *GetCodeVersionsByIDDefault {
	return &GetCodeVersionsByIDDefault{
		_statusCode: code,
	}
}

/*GetCodeVersionsByIDDefault handles this case with default header values.

GetCodeVersionsByIDDefault get code versions by ID default
*/
type GetCodeVersionsByIDDefault struct {
	_statusCode int

	Payload *models.CodeVersion
}

// Code gets the status code for the get code versions by ID default response
func (o *GetCodeVersionsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCodeVersionsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /code_versions/{code_version_id}][%d] getCodeVersionsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCodeVersionsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CodeVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
