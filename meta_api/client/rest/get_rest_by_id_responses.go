// Code generated by go-swagger; DO NOT EDIT.

package rest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/meta_api/models"
)

// GetRestByIDReader is a Reader for the GetRestByID structure.
type GetRestByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetRestByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRestByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRestByIDNotFound creates a GetRestByIDNotFound with default headers values
func NewGetRestByIDNotFound() *GetRestByIDNotFound {
	return &GetRestByIDNotFound{}
}

/*GetRestByIDNotFound handles this case with default header values.

in case of an invalid API name
*/
type GetRestByIDNotFound struct {
}

func (o *GetRestByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /rest/{api_name}][%d] getRestByIdNotFound ", 404)
}

func (o *GetRestByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRestByIDDefault creates a GetRestByIDDefault with default headers values
func NewGetRestByIDDefault(code int) *GetRestByIDDefault {
	return &GetRestByIDDefault{
		_statusCode: code,
	}
}

/*GetRestByIDDefault handles this case with default header values.

GetRestByIDDefault get rest by ID default
*/
type GetRestByIDDefault struct {
	_statusCode int

	Payload *models.VersionResult
}

// Code gets the status code for the get rest by ID default response
func (o *GetRestByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRestByIDDefault) Error() string {
	return fmt.Sprintf("[GET /rest/{api_name}][%d] getRestByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
