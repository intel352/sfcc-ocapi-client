// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSitesByIDCampaignsByIDPromotionsByIDReader is a Reader for the DeleteSitesByIDCampaignsByIDPromotionsByID structure.
type DeleteSitesByIDCampaignsByIDPromotionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSitesByIDCampaignsByIDPromotionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSitesByIDCampaignsByIDPromotionsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSitesByIDCampaignsByIDPromotionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSitesByIDCampaignsByIDPromotionsByIDNoContent creates a DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent with default headers values
func NewDeleteSitesByIDCampaignsByIDPromotionsByIDNoContent() *DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent {
	return &DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent{}
}

/*DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent handles this case with default header values.

DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent delete sites by Id campaigns by Id promotions by Id no content
*/
type DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent struct {
}

func (o *DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/campaigns/{campaign_id}/promotions/{promotion_id}][%d] deleteSitesByIdCampaignsByIdPromotionsByIdNoContent ", 204)
}

func (o *DeleteSitesByIDCampaignsByIDPromotionsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSitesByIDCampaignsByIDPromotionsByIDNotFound creates a DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound with default headers values
func NewDeleteSitesByIDCampaignsByIDPromotionsByIDNotFound() *DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound {
	return &DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound{}
}

/*DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound handles this case with default header values.

thrown when site with the given site id is not found. or Indicates that the campaign with the given campaign ID is not found. or Indicates that the campaign with the given promotion ID is not found.
*/
type DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound struct {
}

func (o *DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/campaigns/{campaign_id}/promotions/{promotion_id}][%d] deleteSitesByIdCampaignsByIdPromotionsByIdNotFound ", 404)
}

func (o *DeleteSitesByIDCampaignsByIDPromotionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
