// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSitesByIDCampaignsByIDCouponsByIDReader is a Reader for the DeleteSitesByIDCampaignsByIDCouponsByID structure.
type DeleteSitesByIDCampaignsByIDCouponsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSitesByIDCampaignsByIDCouponsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSitesByIDCampaignsByIDCouponsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSitesByIDCampaignsByIDCouponsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSitesByIDCampaignsByIDCouponsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSitesByIDCampaignsByIDCouponsByIDNoContent creates a DeleteSitesByIDCampaignsByIDCouponsByIDNoContent with default headers values
func NewDeleteSitesByIDCampaignsByIDCouponsByIDNoContent() *DeleteSitesByIDCampaignsByIDCouponsByIDNoContent {
	return &DeleteSitesByIDCampaignsByIDCouponsByIDNoContent{}
}

/*DeleteSitesByIDCampaignsByIDCouponsByIDNoContent handles this case with default header values.

DeleteSitesByIDCampaignsByIDCouponsByIDNoContent delete sites by Id campaigns by Id coupons by Id no content
*/
type DeleteSitesByIDCampaignsByIDCouponsByIDNoContent struct {
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/campaigns/{campaign_id}/coupons/{coupon_id}][%d] deleteSitesByIdCampaignsByIdCouponsByIdNoContent ", 204)
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSitesByIDCampaignsByIDCouponsByIDBadRequest creates a DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest with default headers values
func NewDeleteSitesByIDCampaignsByIDCouponsByIDBadRequest() *DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest {
	return &DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest{}
}

/*DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest handles this case with default header values.

Indicates some parameter constraint violation occurs
*/
type DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest struct {
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/campaigns/{campaign_id}/coupons/{coupon_id}][%d] deleteSitesByIdCampaignsByIdCouponsByIdBadRequest ", 400)
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSitesByIDCampaignsByIDCouponsByIDNotFound creates a DeleteSitesByIDCampaignsByIDCouponsByIDNotFound with default headers values
func NewDeleteSitesByIDCampaignsByIDCouponsByIDNotFound() *DeleteSitesByIDCampaignsByIDCouponsByIDNotFound {
	return &DeleteSitesByIDCampaignsByIDCouponsByIDNotFound{}
}

/*DeleteSitesByIDCampaignsByIDCouponsByIDNotFound handles this case with default header values.

Indicates that the campaign with the given campaign id
 				is unknown.
*/
type DeleteSitesByIDCampaignsByIDCouponsByIDNotFound struct {
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/campaigns/{campaign_id}/coupons/{coupon_id}][%d] deleteSitesByIdCampaignsByIdCouponsByIdNotFound ", 404)
}

func (o *DeleteSitesByIDCampaignsByIDCouponsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}