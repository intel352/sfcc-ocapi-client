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

// GetSitesByIDSlotsByIDSlotConfigurationsByIDReader is a Reader for the GetSitesByIDSlotsByIDSlotConfigurationsByID structure.
type GetSitesByIDSlotsByIDSlotConfigurationsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDSlotsByIDSlotConfigurationsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest creates a GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest with default headers values
func NewGetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest() *GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest {
	return &GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest{}
}

/*GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest handles this case with default header values.

Thrown if the specified context type is invalid. or Thrown if the specified context id for given context 'category' or 'folder' is missing.
*/
type GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest struct {
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] getSitesByIdSlotsByIdSlotConfigurationsByIdBadRequest ", 400)
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound creates a GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound with default headers values
func NewGetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound() *GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound {
	return &GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound{}
}

/*GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound handles this case with default header values.

Throw if there was no slot with the specified id found for the
              requested site. or Thrown if there was no slot configuration found for the specified
              configuration id.
*/
type GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound struct {
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] getSitesByIdSlotsByIdSlotConfigurationsByIdNotFound ", 404)
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDSlotsByIDSlotConfigurationsByIDDefault creates a GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault with default headers values
func NewGetSitesByIDSlotsByIDSlotConfigurationsByIDDefault(code int) *GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault {
	return &GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault handles this case with default header values.

GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault get sites by ID slots by ID slot configurations by ID default
*/
type GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault struct {
	_statusCode int

	Payload *models.SlotConfiguration
}

// Code gets the status code for the get sites by ID slots by ID slot configurations by ID default response
func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/slots/{slot_id}/slot_configurations/{configuration_id}][%d] getSitesByIDSlotsByIDSlotConfigurationsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDSlotsByIDSlotConfigurationsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SlotConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
