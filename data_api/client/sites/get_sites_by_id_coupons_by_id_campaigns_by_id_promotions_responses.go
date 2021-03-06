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

// GetSitesByIDCouponsByIDCampaignsByIDPromotionsReader is a Reader for the GetSitesByIDCouponsByIDCampaignsByIDPromotions structure.
type GetSitesByIDCouponsByIDCampaignsByIDPromotionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound creates a GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound with default headers values
func NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound() *GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound {
	return &GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound{}
}

/*GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound handles this case with default header values.

Thrown if the coupon does not exist matching the given id. or Thrown if the campaign does not exist matching the given id.
*/
type GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound struct {
}

func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/coupons/{coupon_id}/campaigns/{campaign_id}/promotions][%d] getSitesByIdCouponsByIdCampaignsByIdPromotionsNotFound ", 404)
}

func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault creates a GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault with default headers values
func NewGetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault(code int) *GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault {
	return &GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault handles this case with default header values.

GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault get sites by ID coupons by ID campaigns by ID promotions default
*/
type GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault struct {
	_statusCode int

	Payload *models.Promotions
}

// Code gets the status code for the get sites by ID coupons by ID campaigns by ID promotions default response
func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/coupons/{coupon_id}/campaigns/{campaign_id}/promotions][%d] getSitesByIDCouponsByIDCampaignsByIDPromotions default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDCouponsByIDCampaignsByIDPromotionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Promotions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
