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

// PatchSitesByIDSourceCodeGroupsByIDReader is a Reader for the PatchSitesByIDSourceCodeGroupsByID structure.
type PatchSitesByIDSourceCodeGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDSourceCodeGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchSitesByIDSourceCodeGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDSourceCodeGroupsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDSourceCodeGroupsByIDNotFound creates a PatchSitesByIDSourceCodeGroupsByIDNotFound with default headers values
func NewPatchSitesByIDSourceCodeGroupsByIDNotFound() *PatchSitesByIDSourceCodeGroupsByIDNotFound {
	return &PatchSitesByIDSourceCodeGroupsByIDNotFound{}
}

/*PatchSitesByIDSourceCodeGroupsByIDNotFound handles this case with default header values.

Thrown in case the source code group does not exist matching the given id
*/
type PatchSitesByIDSourceCodeGroupsByIDNotFound struct {
}

func (o *PatchSitesByIDSourceCodeGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/source_code_groups/{id}][%d] patchSitesByIdSourceCodeGroupsByIdNotFound ", 404)
}

func (o *PatchSitesByIDSourceCodeGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDSourceCodeGroupsByIDDefault creates a PatchSitesByIDSourceCodeGroupsByIDDefault with default headers values
func NewPatchSitesByIDSourceCodeGroupsByIDDefault(code int) *PatchSitesByIDSourceCodeGroupsByIDDefault {
	return &PatchSitesByIDSourceCodeGroupsByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDSourceCodeGroupsByIDDefault handles this case with default header values.

PatchSitesByIDSourceCodeGroupsByIDDefault patch sites by ID source code groups by ID default
*/
type PatchSitesByIDSourceCodeGroupsByIDDefault struct {
	_statusCode int

	Payload *models.SourceCodeGroup
}

// Code gets the status code for the patch sites by ID source code groups by ID default response
func (o *PatchSitesByIDSourceCodeGroupsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDSourceCodeGroupsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/source_code_groups/{id}][%d] patchSitesByIDSourceCodeGroupsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDSourceCodeGroupsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceCodeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
