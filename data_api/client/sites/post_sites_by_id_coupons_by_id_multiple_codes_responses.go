// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostSitesByIDCouponsByIDMultipleCodesReader is a Reader for the PostSitesByIDCouponsByIDMultipleCodes structure.
type PostSitesByIDCouponsByIDMultipleCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesByIDCouponsByIDMultipleCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostSitesByIDCouponsByIDMultipleCodesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSitesByIDCouponsByIDMultipleCodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostSitesByIDCouponsByIDMultipleCodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSitesByIDCouponsByIDMultipleCodesNoContent creates a PostSitesByIDCouponsByIDMultipleCodesNoContent with default headers values
func NewPostSitesByIDCouponsByIDMultipleCodesNoContent() *PostSitesByIDCouponsByIDMultipleCodesNoContent {
	return &PostSitesByIDCouponsByIDMultipleCodesNoContent{}
}

/*PostSitesByIDCouponsByIDMultipleCodesNoContent handles this case with default header values.

PostSitesByIDCouponsByIDMultipleCodesNoContent post sites by Id coupons by Id multiple codes no content
*/
type PostSitesByIDCouponsByIDMultipleCodesNoContent struct {
}

func (o *PostSitesByIDCouponsByIDMultipleCodesNoContent) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/coupons/{coupon_id}/multiple_codes][%d] postSitesByIdCouponsByIdMultipleCodesNoContent ", 204)
}

func (o *PostSitesByIDCouponsByIDMultipleCodesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitesByIDCouponsByIDMultipleCodesBadRequest creates a PostSitesByIDCouponsByIDMultipleCodesBadRequest with default headers values
func NewPostSitesByIDCouponsByIDMultipleCodesBadRequest() *PostSitesByIDCouponsByIDMultipleCodesBadRequest {
	return &PostSitesByIDCouponsByIDMultipleCodesBadRequest{}
}

/*PostSitesByIDCouponsByIDMultipleCodesBadRequest handles this case with default header values.

Thrown if the coupon code could not be created.
*/
type PostSitesByIDCouponsByIDMultipleCodesBadRequest struct {
}

func (o *PostSitesByIDCouponsByIDMultipleCodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/coupons/{coupon_id}/multiple_codes][%d] postSitesByIdCouponsByIdMultipleCodesBadRequest ", 400)
}

func (o *PostSitesByIDCouponsByIDMultipleCodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitesByIDCouponsByIDMultipleCodesNotFound creates a PostSitesByIDCouponsByIDMultipleCodesNotFound with default headers values
func NewPostSitesByIDCouponsByIDMultipleCodesNotFound() *PostSitesByIDCouponsByIDMultipleCodesNotFound {
	return &PostSitesByIDCouponsByIDMultipleCodesNotFound{}
}

/*PostSitesByIDCouponsByIDMultipleCodesNotFound handles this case with default header values.

Thrown if the coupon does not exist matching the given id.
*/
type PostSitesByIDCouponsByIDMultipleCodesNotFound struct {
}

func (o *PostSitesByIDCouponsByIDMultipleCodesNotFound) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/coupons/{coupon_id}/multiple_codes][%d] postSitesByIdCouponsByIdMultipleCodesNotFound ", 404)
}

func (o *PostSitesByIDCouponsByIDMultipleCodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
