// Code generated by go-swagger; DO NOT EDIT.

package libraries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PutLibrariesByIDFoldersByIDReader is a Reader for the PutLibrariesByIDFoldersByID structure.
type PutLibrariesByIDFoldersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLibrariesByIDFoldersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutLibrariesByIDFoldersByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutLibrariesByIDFoldersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutLibrariesByIDFoldersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLibrariesByIDFoldersByIDBadRequest creates a PutLibrariesByIDFoldersByIDBadRequest with default headers values
func NewPutLibrariesByIDFoldersByIDBadRequest() *PutLibrariesByIDFoldersByIDBadRequest {
	return &PutLibrariesByIDFoldersByIDBadRequest{}
}

/*PutLibrariesByIDFoldersByIDBadRequest handles this case with default header values.

Indicates that the ID from the request body doesn't match the URL-Id.
*/
type PutLibrariesByIDFoldersByIDBadRequest struct {
}

func (o *PutLibrariesByIDFoldersByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /libraries/{library_id}/folders/{folder_id}][%d] putLibrariesByIdFoldersByIdBadRequest ", 400)
}

func (o *PutLibrariesByIDFoldersByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLibrariesByIDFoldersByIDNotFound creates a PutLibrariesByIDFoldersByIDNotFound with default headers values
func NewPutLibrariesByIDFoldersByIDNotFound() *PutLibrariesByIDFoldersByIDNotFound {
	return &PutLibrariesByIDFoldersByIDNotFound{}
}

/*PutLibrariesByIDFoldersByIDNotFound handles this case with default header values.

Indicates that library with the given id is unknown. or Indicates that the parent folder with the given id is unknown.
*/
type PutLibrariesByIDFoldersByIDNotFound struct {
}

func (o *PutLibrariesByIDFoldersByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /libraries/{library_id}/folders/{folder_id}][%d] putLibrariesByIdFoldersByIdNotFound ", 404)
}

func (o *PutLibrariesByIDFoldersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLibrariesByIDFoldersByIDDefault creates a PutLibrariesByIDFoldersByIDDefault with default headers values
func NewPutLibrariesByIDFoldersByIDDefault(code int) *PutLibrariesByIDFoldersByIDDefault {
	return &PutLibrariesByIDFoldersByIDDefault{
		_statusCode: code,
	}
}

/*PutLibrariesByIDFoldersByIDDefault handles this case with default header values.

PutLibrariesByIDFoldersByIDDefault put libraries by ID folders by ID default
*/
type PutLibrariesByIDFoldersByIDDefault struct {
	_statusCode int

	Payload *models.ContentFolder
}

// Code gets the status code for the put libraries by ID folders by ID default response
func (o *PutLibrariesByIDFoldersByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutLibrariesByIDFoldersByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /libraries/{library_id}/folders/{folder_id}][%d] putLibrariesByIDFoldersByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutLibrariesByIDFoldersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentFolder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
