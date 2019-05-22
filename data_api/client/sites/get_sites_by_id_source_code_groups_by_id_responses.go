// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetSitesByIDSourceCodeGroupsByIDReader is a Reader for the GetSitesByIDSourceCodeGroupsByID structure.
type GetSitesByIDSourceCodeGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDSourceCodeGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSitesByIDSourceCodeGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDSourceCodeGroupsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDSourceCodeGroupsByIDNotFound creates a GetSitesByIDSourceCodeGroupsByIDNotFound with default headers values
func NewGetSitesByIDSourceCodeGroupsByIDNotFound() *GetSitesByIDSourceCodeGroupsByIDNotFound {
	return &GetSitesByIDSourceCodeGroupsByIDNotFound{}
}

/*GetSitesByIDSourceCodeGroupsByIDNotFound handles this case with default header values.

Thrown in case the source code group does not exist matching the given id
*/
type GetSitesByIDSourceCodeGroupsByIDNotFound struct {
}

func (o *GetSitesByIDSourceCodeGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/source_code_groups/{id}][%d] getSitesByIdSourceCodeGroupsByIdNotFound ", 404)
}

func (o *GetSitesByIDSourceCodeGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDSourceCodeGroupsByIDDefault creates a GetSitesByIDSourceCodeGroupsByIDDefault with default headers values
func NewGetSitesByIDSourceCodeGroupsByIDDefault(code int) *GetSitesByIDSourceCodeGroupsByIDDefault {
	return &GetSitesByIDSourceCodeGroupsByIDDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDSourceCodeGroupsByIDDefault handles this case with default header values.

GetSitesByIDSourceCodeGroupsByIDDefault get sites by ID source code groups by ID default
*/
type GetSitesByIDSourceCodeGroupsByIDDefault struct {
	_statusCode int

	Payload *models.SourceCodeGroup
}

// Code gets the status code for the get sites by ID source code groups by ID default response
func (o *GetSitesByIDSourceCodeGroupsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDSourceCodeGroupsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/source_code_groups/{id}][%d] getSitesByIDSourceCodeGroupsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDSourceCodeGroupsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceCodeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
