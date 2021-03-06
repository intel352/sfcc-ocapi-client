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

// PutSitesByIDCampaignsByIDReader is a Reader for the PutSitesByIDCampaignsByID structure.
type PutSitesByIDCampaignsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSitesByIDCampaignsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutSitesByIDCampaignsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutSitesByIDCampaignsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutSitesByIDCampaignsByIDBadRequest creates a PutSitesByIDCampaignsByIDBadRequest with default headers values
func NewPutSitesByIDCampaignsByIDBadRequest() *PutSitesByIDCampaignsByIDBadRequest {
	return &PutSitesByIDCampaignsByIDBadRequest{}
}

/*PutSitesByIDCampaignsByIDBadRequest handles this case with default header values.

if the Id in request is not the same as the ID in document.
*/
type PutSitesByIDCampaignsByIDBadRequest struct {
}

func (o *PutSitesByIDCampaignsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/campaigns/{campaign_id}][%d] putSitesByIdCampaignsByIdBadRequest ", 400)
}

func (o *PutSitesByIDCampaignsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDCampaignsByIDDefault creates a PutSitesByIDCampaignsByIDDefault with default headers values
func NewPutSitesByIDCampaignsByIDDefault(code int) *PutSitesByIDCampaignsByIDDefault {
	return &PutSitesByIDCampaignsByIDDefault{
		_statusCode: code,
	}
}

/*PutSitesByIDCampaignsByIDDefault handles this case with default header values.

PutSitesByIDCampaignsByIDDefault put sites by ID campaigns by ID default
*/
type PutSitesByIDCampaignsByIDDefault struct {
	_statusCode int

	Payload *models.Campaign
}

// Code gets the status code for the put sites by ID campaigns by ID default response
func (o *PutSitesByIDCampaignsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutSitesByIDCampaignsByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/campaigns/{campaign_id}][%d] putSitesByIDCampaignsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutSitesByIDCampaignsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
