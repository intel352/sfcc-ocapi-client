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

// GetLibrariesByIDContentByIDFoldersReader is a Reader for the GetLibrariesByIDContentByIDFolders structure.
type GetLibrariesByIDContentByIDFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibrariesByIDContentByIDFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetLibrariesByIDContentByIDFoldersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetLibrariesByIDContentByIDFoldersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLibrariesByIDContentByIDFoldersNotFound creates a GetLibrariesByIDContentByIDFoldersNotFound with default headers values
func NewGetLibrariesByIDContentByIDFoldersNotFound() *GetLibrariesByIDContentByIDFoldersNotFound {
	return &GetLibrariesByIDContentByIDFoldersNotFound{}
}

/*GetLibrariesByIDContentByIDFoldersNotFound handles this case with default header values.

Indicates that library with the given id is unknown. or Indicates that the content asset with the given id is unknown.
*/
type GetLibrariesByIDContentByIDFoldersNotFound struct {
}

func (o *GetLibrariesByIDContentByIDFoldersNotFound) Error() string {
	return fmt.Sprintf("[GET /libraries/{library_id}/content/{content_id}/folders][%d] getLibrariesByIdContentByIdFoldersNotFound ", 404)
}

func (o *GetLibrariesByIDContentByIDFoldersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLibrariesByIDContentByIDFoldersDefault creates a GetLibrariesByIDContentByIDFoldersDefault with default headers values
func NewGetLibrariesByIDContentByIDFoldersDefault(code int) *GetLibrariesByIDContentByIDFoldersDefault {
	return &GetLibrariesByIDContentByIDFoldersDefault{
		_statusCode: code,
	}
}

/*GetLibrariesByIDContentByIDFoldersDefault handles this case with default header values.

GetLibrariesByIDContentByIDFoldersDefault get libraries by ID content by ID folders default
*/
type GetLibrariesByIDContentByIDFoldersDefault struct {
	_statusCode int

	Payload *models.ContentFolderResult
}

// Code gets the status code for the get libraries by ID content by ID folders default response
func (o *GetLibrariesByIDContentByIDFoldersDefault) Code() int {
	return o._statusCode
}

func (o *GetLibrariesByIDContentByIDFoldersDefault) Error() string {
	return fmt.Sprintf("[GET /libraries/{library_id}/content/{content_id}/folders][%d] getLibrariesByIDContentByIDFolders default  %+v", o._statusCode, o.Payload)
}

func (o *GetLibrariesByIDContentByIDFoldersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentFolderResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
