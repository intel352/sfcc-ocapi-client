// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader is a Reader for the PutSitesByIDCampaignsByIDSlotConfigurationsByIDByID structure.
type PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent creates a PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent with default headers values
func NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent() *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent {
	return &PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent{}
}

/*PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent handles this case with default header values.

PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent put sites by Id campaigns by Id slot configurations by Id by Id no content
*/
type PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent struct {
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] putSitesByIdCampaignsByIdSlotConfigurationsByIdByIdNoContent ", 204)
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest creates a PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest with default headers values
func NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest() *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest {
	return &PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest{}
}

/*PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest handles this case with default header values.

Indicates the slot context type is not one of "global", "category", or "folder" or Indicates the ID in the URL does not match the ID in the request
*/
type PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest struct {
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] putSitesByIdCampaignsByIdSlotConfigurationsByIdByIdBadRequest ", 400)
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound creates a PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound with default headers values
func NewPutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound() *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound {
	return &PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound{}
}

/*PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound handles this case with default header values.

thrown when site with the given site id is not found. or Indicates that the campaign with the given campaign ID is not found. or Indicates that the slot with the given slot ID and context type is not found. or Indicates that the slot with the given slot ID, slot configuration ID and context type is not found.
*/
type PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound struct {
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] putSitesByIdCampaignsByIdSlotConfigurationsByIdByIdNotFound ", 404)
}

func (o *PutSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
