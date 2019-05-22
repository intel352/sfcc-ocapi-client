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

// GetSitesByIDSlotConfigurationsReader is a Reader for the GetSitesByIDSlotConfigurations structure.
type GetSitesByIDSlotConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDSlotConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetSitesByIDSlotConfigurationsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetSitesByIDSlotConfigurationsDefault creates a GetSitesByIDSlotConfigurationsDefault with default headers values
func NewGetSitesByIDSlotConfigurationsDefault(code int) *GetSitesByIDSlotConfigurationsDefault {
	return &GetSitesByIDSlotConfigurationsDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDSlotConfigurationsDefault handles this case with default header values.

GetSitesByIDSlotConfigurationsDefault get sites by ID slot configurations default
*/
type GetSitesByIDSlotConfigurationsDefault struct {
	_statusCode int

	Payload *models.SlotConfigurations
}

// Code gets the status code for the get sites by ID slot configurations default response
func (o *GetSitesByIDSlotConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDSlotConfigurationsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/slot_configurations][%d] getSitesByIDSlotConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDSlotConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SlotConfigurations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}