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

// PatchSitesByIDAbTestsByIDReader is a Reader for the PatchSitesByIDAbTestsByID structure.
type PatchSitesByIDAbTestsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDAbTestsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchSitesByIDAbTestsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSitesByIDAbTestsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchSitesByIDAbTestsByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDAbTestsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDAbTestsByIDBadRequest creates a PatchSitesByIDAbTestsByIDBadRequest with default headers values
func NewPatchSitesByIDAbTestsByIDBadRequest() *PatchSitesByIDAbTestsByIDBadRequest {
	return &PatchSitesByIDAbTestsByIDBadRequest{}
}

/*PatchSitesByIDAbTestsByIDBadRequest handles this case with default header values.

Thrown if an invalid email is specified. or Thrown if an invalid pipeline is specified. or Thrown if there is missing pipline_call or categories in the specified trigger of type pipeline_call or category_view_page respectively.
*/
type PatchSitesByIDAbTestsByIDBadRequest struct {
}

func (o *PatchSitesByIDAbTestsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/ab_tests/{id}][%d] patchSitesByIdAbTestsByIdBadRequest ", 400)
}

func (o *PatchSitesByIDAbTestsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDAbTestsByIDNotFound creates a PatchSitesByIDAbTestsByIDNotFound with default headers values
func NewPatchSitesByIDAbTestsByIDNotFound() *PatchSitesByIDAbTestsByIDNotFound {
	return &PatchSitesByIDAbTestsByIDNotFound{}
}

/*PatchSitesByIDAbTestsByIDNotFound handles this case with default header values.

Indicates that site specified with the given id is unknown. or Thrown in case the A/B Test does not exist matching the given id
*/
type PatchSitesByIDAbTestsByIDNotFound struct {
}

func (o *PatchSitesByIDAbTestsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/ab_tests/{id}][%d] patchSitesByIdAbTestsByIdNotFound ", 404)
}

func (o *PatchSitesByIDAbTestsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDAbTestsByIDConflict creates a PatchSitesByIDAbTestsByIDConflict with default headers values
func NewPatchSitesByIDAbTestsByIDConflict() *PatchSitesByIDAbTestsByIDConflict {
	return &PatchSitesByIDAbTestsByIDConflict{}
}

/*PatchSitesByIDAbTestsByIDConflict handles this case with default header values.

Thrown if an A/B Test exists already in the site with the given identifier specified in the body that is different from the identifier in the path.
*/
type PatchSitesByIDAbTestsByIDConflict struct {
}

func (o *PatchSitesByIDAbTestsByIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/ab_tests/{id}][%d] patchSitesByIdAbTestsByIdConflict ", 409)
}

func (o *PatchSitesByIDAbTestsByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDAbTestsByIDDefault creates a PatchSitesByIDAbTestsByIDDefault with default headers values
func NewPatchSitesByIDAbTestsByIDDefault(code int) *PatchSitesByIDAbTestsByIDDefault {
	return &PatchSitesByIDAbTestsByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDAbTestsByIDDefault handles this case with default header values.

PatchSitesByIDAbTestsByIDDefault patch sites by ID ab tests by ID default
*/
type PatchSitesByIDAbTestsByIDDefault struct {
	_statusCode int

	Payload *models.AbTest
}

// Code gets the status code for the patch sites by ID ab tests by ID default response
func (o *PatchSitesByIDAbTestsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDAbTestsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/ab_tests/{id}][%d] patchSitesByIDAbTestsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDAbTestsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AbTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}