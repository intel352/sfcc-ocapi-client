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

// PatchSitesByIDCustomObjectsByIDByIDReader is a Reader for the PatchSitesByIDCustomObjectsByIDByID structure.
type PatchSitesByIDCustomObjectsByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDCustomObjectsByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchSitesByIDCustomObjectsByIDByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSitesByIDCustomObjectsByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDCustomObjectsByIDByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDCustomObjectsByIDByIDBadRequest creates a PatchSitesByIDCustomObjectsByIDByIDBadRequest with default headers values
func NewPatchSitesByIDCustomObjectsByIDByIDBadRequest() *PatchSitesByIDCustomObjectsByIDByIDBadRequest {
	return &PatchSitesByIDCustomObjectsByIDByIDBadRequest{}
}

/*PatchSitesByIDCustomObjectsByIDByIDBadRequest handles this case with default header values.

If the object key must be an integer, but the path parameter has an invalid format
*/
type PatchSitesByIDCustomObjectsByIDByIDBadRequest struct {
}

func (o *PatchSitesByIDCustomObjectsByIDByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/custom_objects/{object_type}/{key}][%d] patchSitesByIdCustomObjectsByIdByIdBadRequest ", 400)
}

func (o *PatchSitesByIDCustomObjectsByIDByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCustomObjectsByIDByIDNotFound creates a PatchSitesByIDCustomObjectsByIDByIDNotFound with default headers values
func NewPatchSitesByIDCustomObjectsByIDByIDNotFound() *PatchSitesByIDCustomObjectsByIDByIDNotFound {
	return &PatchSitesByIDCustomObjectsByIDByIDNotFound{}
}

/*PatchSitesByIDCustomObjectsByIDByIDNotFound handles this case with default header values.

For an unknown object key or For an unknown object type ID
*/
type PatchSitesByIDCustomObjectsByIDByIDNotFound struct {
}

func (o *PatchSitesByIDCustomObjectsByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/custom_objects/{object_type}/{key}][%d] patchSitesByIdCustomObjectsByIdByIdNotFound ", 404)
}

func (o *PatchSitesByIDCustomObjectsByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCustomObjectsByIDByIDDefault creates a PatchSitesByIDCustomObjectsByIDByIDDefault with default headers values
func NewPatchSitesByIDCustomObjectsByIDByIDDefault(code int) *PatchSitesByIDCustomObjectsByIDByIDDefault {
	return &PatchSitesByIDCustomObjectsByIDByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDCustomObjectsByIDByIDDefault handles this case with default header values.

PatchSitesByIDCustomObjectsByIDByIDDefault patch sites by ID custom objects by ID by ID default
*/
type PatchSitesByIDCustomObjectsByIDByIDDefault struct {
	_statusCode int

	Payload *models.CustomObject
}

// Code gets the status code for the patch sites by ID custom objects by ID by ID default response
func (o *PatchSitesByIDCustomObjectsByIDByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDCustomObjectsByIDByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/custom_objects/{object_type}/{key}][%d] patchSitesByIDCustomObjectsByIDByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDCustomObjectsByIDByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
