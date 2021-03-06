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

// GetRestReader is a Reader for the GetRest structure.
type GetRestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetRestDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetRestDefault creates a GetRestDefault with default headers values
func NewGetRestDefault(code int) *GetRestDefault {
	return &GetRestDefault{
		_statusCode: code,
	}
}

/*GetRestDefault handles this case with default header values.

GetRestDefault get rest default
*/
type GetRestDefault struct {
	_statusCode int

	Payload *models.APIResult
}

// Code gets the status code for the get rest default response
func (o *GetRestDefault) Code() int {
	return o._statusCode
}

func (o *GetRestDefault) Error() string {
	return fmt.Sprintf("[GET /rest][%d] getRest default  %+v", o._statusCode, o.Payload)
}

func (o *GetRestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
