// Code generated by go-swagger; DO NOT EDIT.

package ocapi_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostOcapiConfigsByIDReader is a Reader for the PostOcapiConfigsByID structure.
type PostOcapiConfigsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOcapiConfigsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostOcapiConfigsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostOcapiConfigsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostOcapiConfigsByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostOcapiConfigsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOcapiConfigsByIDBadRequest creates a PostOcapiConfigsByIDBadRequest with default headers values
func NewPostOcapiConfigsByIDBadRequest() *PostOcapiConfigsByIDBadRequest {
	return &PostOcapiConfigsByIDBadRequest{}
}

/*PostOcapiConfigsByIDBadRequest handles this case with default header values.

Indicates that the resulting config is not valid or Write operation on self is not allowed. (The clientId being used for calling this API should be different from target clientId)
*/
type PostOcapiConfigsByIDBadRequest struct {
}

func (o *PostOcapiConfigsByIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /ocapi_configs/{clientId}][%d] postOcapiConfigsByIdBadRequest ", 400)
}

func (o *PostOcapiConfigsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOcapiConfigsByIDNotFound creates a PostOcapiConfigsByIDNotFound with default headers values
func NewPostOcapiConfigsByIDNotFound() *PostOcapiConfigsByIDNotFound {
	return &PostOcapiConfigsByIDNotFound{}
}

/*PostOcapiConfigsByIDNotFound handles this case with default header values.

Indicates that at least one of the target sites is unknown
*/
type PostOcapiConfigsByIDNotFound struct {
}

func (o *PostOcapiConfigsByIDNotFound) Error() string {
	return fmt.Sprintf("[POST /ocapi_configs/{clientId}][%d] postOcapiConfigsByIdNotFound ", 404)
}

func (o *PostOcapiConfigsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOcapiConfigsByIDConflict creates a PostOcapiConfigsByIDConflict with default headers values
func NewPostOcapiConfigsByIDConflict() *PostOcapiConfigsByIDConflict {
	return &PostOcapiConfigsByIDConflict{}
}

/*PostOcapiConfigsByIDConflict handles this case with default header values.

Indicates that the client id already exists in original OCAPI configuration
*/
type PostOcapiConfigsByIDConflict struct {
}

func (o *PostOcapiConfigsByIDConflict) Error() string {
	return fmt.Sprintf("[POST /ocapi_configs/{clientId}][%d] postOcapiConfigsByIdConflict ", 409)
}

func (o *PostOcapiConfigsByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOcapiConfigsByIDDefault creates a PostOcapiConfigsByIDDefault with default headers values
func NewPostOcapiConfigsByIDDefault(code int) *PostOcapiConfigsByIDDefault {
	return &PostOcapiConfigsByIDDefault{
		_statusCode: code,
	}
}

/*PostOcapiConfigsByIDDefault handles this case with default header values.

PostOcapiConfigsByIDDefault post ocapi configs by ID default
*/
type PostOcapiConfigsByIDDefault struct {
	_statusCode int

	Payload *models.OcapiConfigsAPIResponse
}

// Code gets the status code for the post ocapi configs by ID default response
func (o *PostOcapiConfigsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PostOcapiConfigsByIDDefault) Error() string {
	return fmt.Sprintf("[POST /ocapi_configs/{clientId}][%d] postOcapiConfigsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PostOcapiConfigsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OcapiConfigsAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
