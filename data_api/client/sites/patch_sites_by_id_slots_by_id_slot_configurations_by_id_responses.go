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

// PatchSitesByIDSlotsByIDSlotConfigurationsByIDReader is a Reader for the PatchSitesByIDSlotsByIDSlotConfigurationsByID structure.
type PatchSitesByIDSlotsByIDSlotConfigurationsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest creates a PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest with default headers values
func NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest() *PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest {
	return &PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest{}
}

/*PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest handles this case with default header values.

Thrown if the specified context id for given context 'category' or 'folder' is missing. or Thrown if the specified context type is invalid.
*/
type PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest struct {
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] patchSitesByIdSlotsByIdSlotConfigurationsByIdBadRequest ", 400)
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound creates a PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound with default headers values
func NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound() *PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound {
	return &PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound{}
}

/*PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound handles this case with default header values.

Thrown if there was no slot with the given id found for the
              requested site. or Thrown if there was no slot configuration found for the specified configuration
              id. or Thrown if content asset does not exist in the library of the current domain
              but is assigned during update.
*/
type PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound struct {
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] patchSitesByIdSlotsByIdSlotConfigurationsByIdNotFound ", 404)
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict creates a PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict with default headers values
func NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict() *PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict {
	return &PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict{}
}

/*PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict handles this case with default header values.

Thrown if the configuration ID should be updated into one that is already
              assigned.
*/
type PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict struct {
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] patchSitesByIdSlotsByIdSlotConfigurationsByIdConflict ", 409)
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault creates a PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault with default headers values
func NewPatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault(code int) *PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault {
	return &PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault handles this case with default header values.

PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault patch sites by ID slots by ID slot configurations by ID default
*/
type PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault struct {
	_statusCode int

	Payload *models.SlotConfiguration
}

// Code gets the status code for the patch sites by ID slots by ID slot configurations by ID default response
func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] patchSitesByIDSlotsByIDSlotConfigurationsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDSlotsByIDSlotConfigurationsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SlotConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
