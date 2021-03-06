// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSitesByIDCouponsByIDReader is a Reader for the DeleteSitesByIDCouponsByID structure.
type DeleteSitesByIDCouponsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSitesByIDCouponsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSitesByIDCouponsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSitesByIDCouponsByIDNoContent creates a DeleteSitesByIDCouponsByIDNoContent with default headers values
func NewDeleteSitesByIDCouponsByIDNoContent() *DeleteSitesByIDCouponsByIDNoContent {
	return &DeleteSitesByIDCouponsByIDNoContent{}
}

/*DeleteSitesByIDCouponsByIDNoContent handles this case with default header values.

DeleteSitesByIDCouponsByIDNoContent delete sites by Id coupons by Id no content
*/
type DeleteSitesByIDCouponsByIDNoContent struct {
}

func (o *DeleteSitesByIDCouponsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/coupons/{coupon_id}][%d] deleteSitesByIdCouponsByIdNoContent ", 204)
}

func (o *DeleteSitesByIDCouponsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
