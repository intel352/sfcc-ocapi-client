// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader is a Reader for the PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByID structure.
type PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent creates a PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent with default headers values
func NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent() *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent {
	return &PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent{}
}

/*PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent handles this case with default header values.

PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent patch sites by Id campaigns by Id slot configurations by Id by Id no content
*/
type PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent struct {
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] patchSitesByIdCampaignsByIdSlotConfigurationsByIdByIdNoContent ", 204)
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest creates a PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest with default headers values
func NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest() *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest {
	return &PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest{}
}

/*PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest handles this case with default header values.

Indicates the slot context type is not one of "global", "category", or "folder"
*/
type PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest struct {
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] patchSitesByIdCampaignsByIdSlotConfigurationsByIdByIdBadRequest ", 400)
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound creates a PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound with default headers values
func NewPatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound() *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound {
	return &PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound{}
}

/*PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound handles this case with default header values.

thrown when site with the given site id is not found. or Indicates that the campaign with the given campaign ID is not found. or Indicates that the slot with the given slot ID and context type is not found. or Indicates that the slot with the given slot ID, slot configuration ID and context type is not found.
*/
type PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound struct {
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/campaigns/{campaign_id}/slot_configurations/{slot_id}/{slot_config_id}][%d] patchSitesByIdCampaignsByIdSlotConfigurationsByIdByIdNotFound ", 404)
}

func (o *PatchSitesByIDCampaignsByIDSlotConfigurationsByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}