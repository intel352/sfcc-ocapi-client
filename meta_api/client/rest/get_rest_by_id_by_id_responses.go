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

// GetRestByIDByIDReader is a Reader for the GetRestByIDByID structure.
type GetRestByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetRestByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRestByIDByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRestByIDByIDNotFound creates a GetRestByIDByIDNotFound with default headers values
func NewGetRestByIDByIDNotFound() *GetRestByIDByIDNotFound {
	return &GetRestByIDByIDNotFound{}
}

/*GetRestByIDByIDNotFound handles this case with default header values.

in case of an invalid API identifier
*/
type GetRestByIDByIDNotFound struct {
}

func (o *GetRestByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /rest/{api_name}/{version}][%d] getRestByIdByIdNotFound ", 404)
}

func (o *GetRestByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRestByIDByIDDefault creates a GetRestByIDByIDDefault with default headers values
func NewGetRestByIDByIDDefault(code int) *GetRestByIDByIDDefault {
	return &GetRestByIDByIDDefault{
		_statusCode: code,
	}
}

/*GetRestByIDByIDDefault handles this case with default header values.

GetRestByIDByIDDefault get rest by ID by ID default
*/
type GetRestByIDByIDDefault struct {
	_statusCode int

	Payload *models.Swagger
}

// Code gets the status code for the get rest by ID by ID default response
func (o *GetRestByIDByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRestByIDByIDDefault) Error() string {
	return fmt.Sprintf("[GET /rest/{api_name}/{version}][%d] getRestByIDByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestByIDByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Swagger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
