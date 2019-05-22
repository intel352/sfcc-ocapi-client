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

// GetSitesByIDCouponsByIDPromotionsReader is a Reader for the GetSitesByIDCouponsByIDPromotions structure.
type GetSitesByIDCouponsByIDPromotionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDCouponsByIDPromotionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSitesByIDCouponsByIDPromotionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDCouponsByIDPromotionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDCouponsByIDPromotionsNotFound creates a GetSitesByIDCouponsByIDPromotionsNotFound with default headers values
func NewGetSitesByIDCouponsByIDPromotionsNotFound() *GetSitesByIDCouponsByIDPromotionsNotFound {
	return &GetSitesByIDCouponsByIDPromotionsNotFound{}
}

/*GetSitesByIDCouponsByIDPromotionsNotFound handles this case with default header values.

Thrown if the coupon does not exist matching the given id.
*/
type GetSitesByIDCouponsByIDPromotionsNotFound struct {
}

func (o *GetSitesByIDCouponsByIDPromotionsNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/coupons/{coupon_id}/promotions][%d] getSitesByIdCouponsByIdPromotionsNotFound ", 404)
}

func (o *GetSitesByIDCouponsByIDPromotionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDCouponsByIDPromotionsDefault creates a GetSitesByIDCouponsByIDPromotionsDefault with default headers values
func NewGetSitesByIDCouponsByIDPromotionsDefault(code int) *GetSitesByIDCouponsByIDPromotionsDefault {
	return &GetSitesByIDCouponsByIDPromotionsDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDCouponsByIDPromotionsDefault handles this case with default header values.

GetSitesByIDCouponsByIDPromotionsDefault get sites by ID coupons by ID promotions default
*/
type GetSitesByIDCouponsByIDPromotionsDefault struct {
	_statusCode int

	Payload *models.Promotions
}

// Code gets the status code for the get sites by ID coupons by ID promotions default response
func (o *GetSitesByIDCouponsByIDPromotionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDCouponsByIDPromotionsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/coupons/{coupon_id}/promotions][%d] getSitesByIDCouponsByIDPromotions default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDCouponsByIDPromotionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Promotions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
