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

// PatchSitesByIDCouponsByIDReader is a Reader for the PatchSitesByIDCouponsByID structure.
type PatchSitesByIDCouponsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSitesByIDCouponsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchSitesByIDCouponsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchSitesByIDCouponsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchSitesByIDCouponsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSitesByIDCouponsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchSitesByIDCouponsByIDBadRequest creates a PatchSitesByIDCouponsByIDBadRequest with default headers values
func NewPatchSitesByIDCouponsByIDBadRequest() *PatchSitesByIDCouponsByIDBadRequest {
	return &PatchSitesByIDCouponsByIDBadRequest{}
}

/*PatchSitesByIDCouponsByIDBadRequest handles this case with default header values.

Thrown if you try to update certain fields in a coupon after
              redeeming or exporting it. or Thrown if the code provided is in use by another coupon. or Thrown whenever a quota regarding system-generated coupons is violated
*/
type PatchSitesByIDCouponsByIDBadRequest struct {
}

func (o *PatchSitesByIDCouponsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/coupons/{coupon_id}][%d] patchSitesByIdCouponsByIdBadRequest ", 400)
}

func (o *PatchSitesByIDCouponsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCouponsByIDForbidden creates a PatchSitesByIDCouponsByIDForbidden with default headers values
func NewPatchSitesByIDCouponsByIDForbidden() *PatchSitesByIDCouponsByIDForbidden {
	return &PatchSitesByIDCouponsByIDForbidden{}
}

/*PatchSitesByIDCouponsByIDForbidden handles this case with default header values.

Thrown when trying to update a broken coupon, that should be deleted and re-created. or Thrown when trying to update a legacy coupon.
*/
type PatchSitesByIDCouponsByIDForbidden struct {
}

func (o *PatchSitesByIDCouponsByIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/coupons/{coupon_id}][%d] patchSitesByIdCouponsByIdForbidden ", 403)
}

func (o *PatchSitesByIDCouponsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCouponsByIDNotFound creates a PatchSitesByIDCouponsByIDNotFound with default headers values
func NewPatchSitesByIDCouponsByIDNotFound() *PatchSitesByIDCouponsByIDNotFound {
	return &PatchSitesByIDCouponsByIDNotFound{}
}

/*PatchSitesByIDCouponsByIDNotFound handles this case with default header values.

Thrown if the coupon does not exist matching the given id.
*/
type PatchSitesByIDCouponsByIDNotFound struct {
}

func (o *PatchSitesByIDCouponsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/coupons/{coupon_id}][%d] patchSitesByIdCouponsByIdNotFound ", 404)
}

func (o *PatchSitesByIDCouponsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchSitesByIDCouponsByIDDefault creates a PatchSitesByIDCouponsByIDDefault with default headers values
func NewPatchSitesByIDCouponsByIDDefault(code int) *PatchSitesByIDCouponsByIDDefault {
	return &PatchSitesByIDCouponsByIDDefault{
		_statusCode: code,
	}
}

/*PatchSitesByIDCouponsByIDDefault handles this case with default header values.

PatchSitesByIDCouponsByIDDefault patch sites by ID coupons by ID default
*/
type PatchSitesByIDCouponsByIDDefault struct {
	_statusCode int

	Payload *models.Coupon
}

// Code gets the status code for the patch sites by ID coupons by ID default response
func (o *PatchSitesByIDCouponsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSitesByIDCouponsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /sites/{site_id}/coupons/{coupon_id}][%d] patchSitesByIDCouponsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSitesByIDCouponsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Coupon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
