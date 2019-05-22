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

// PatchSitesByIDStoresByIDReader is a Reader for the PatchSitesByIDStoresByID structure.
type PatchSitesByIDStoresByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDStoresByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchSitesByIDStoresByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDStoresByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDStoresByIDNotFound creates a PatchSitesByIDStoresByIDNotFound with default headers values
func NewPatchSitesByIDStoresByIDNotFound() *PatchSitesByIDStoresByIDNotFound {
	return &PatchSitesByIDStoresByIDNotFound{}
}

/*PatchSitesByIDStoresByIDNotFound handles this case with default header values.

Thrown in case the store does not exist matching the given id
*/
type PatchSitesByIDStoresByIDNotFound struct {
}

func (o *PatchSitesByIDStoresByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/stores/{id}][%d] patchSitesByIdStoresByIdNotFound ", 404)
}

func (o *PatchSitesByIDStoresByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDStoresByIDDefault creates a PatchSitesByIDStoresByIDDefault with default headers values
func NewPatchSitesByIDStoresByIDDefault(code int) *PatchSitesByIDStoresByIDDefault {
	return &PatchSitesByIDStoresByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDStoresByIDDefault handles this case with default header values.

PatchSitesByIDStoresByIDDefault patch sites by ID stores by ID default
*/
type PatchSitesByIDStoresByIDDefault struct {
	_statusCode int

	Payload *models.Store
}

// Code gets the status code for the patch sites by ID stores by ID default response
func (o *PatchSitesByIDStoresByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDStoresByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/stores/{id}][%d] patchSitesByIDStoresByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDStoresByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Store)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}