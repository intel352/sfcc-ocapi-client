// Code generated by go-swagger; DO NOT EDIT.

package baskets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PatchBasketsByIDPaymentInstrumentsByIDReader is a Reader for the PatchBasketsByIDPaymentInstrumentsByID structure.
type PatchBasketsByIDPaymentInstrumentsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBasketsByIDPaymentInstrumentsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchBasketsByIDPaymentInstrumentsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchBasketsByIDPaymentInstrumentsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchBasketsByIDPaymentInstrumentsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBasketsByIDPaymentInstrumentsByIDBadRequest creates a PatchBasketsByIDPaymentInstrumentsByIDBadRequest with default headers values
func NewPatchBasketsByIDPaymentInstrumentsByIDBadRequest() *PatchBasketsByIDPaymentInstrumentsByIDBadRequest {
	return &PatchBasketsByIDPaymentInstrumentsByIDBadRequest{}
}

/*PatchBasketsByIDPaymentInstrumentsByIDBadRequest handles this case with default header values.

Indicates that the provided payment method is invalid or not applicable. or Indicates that the basket payment instrument with the given
             id already is permanently masked. Please see
             dw.order.PaymentInstrument.isPermanentlyMasked() for detailed information. or Indicates that the customer assigned to the basket does not match the verified
             customer represented by the JWT token, not relevant when using OAuth.
*/
type PatchBasketsByIDPaymentInstrumentsByIDBadRequest struct {
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /baskets/{basket_id}/payment_instruments/{payment_instrument_id}][%d] patchBasketsByIdPaymentInstrumentsByIdBadRequest ", 400)
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBasketsByIDPaymentInstrumentsByIDNotFound creates a PatchBasketsByIDPaymentInstrumentsByIDNotFound with default headers values
func NewPatchBasketsByIDPaymentInstrumentsByIDNotFound() *PatchBasketsByIDPaymentInstrumentsByIDNotFound {
	return &PatchBasketsByIDPaymentInstrumentsByIDNotFound{}
}

/*PatchBasketsByIDPaymentInstrumentsByIDNotFound handles this case with default header values.

Indicates that the basket with the given basket id is unknown. or Indicates that the payment instrument with the given payment
             instrument number is unknown.
*/
type PatchBasketsByIDPaymentInstrumentsByIDNotFound struct {
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /baskets/{basket_id}/payment_instruments/{payment_instrument_id}][%d] patchBasketsByIdPaymentInstrumentsByIdNotFound ", 404)
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBasketsByIDPaymentInstrumentsByIDDefault creates a PatchBasketsByIDPaymentInstrumentsByIDDefault with default headers values
func NewPatchBasketsByIDPaymentInstrumentsByIDDefault(code int) *PatchBasketsByIDPaymentInstrumentsByIDDefault {
	return &PatchBasketsByIDPaymentInstrumentsByIDDefault{
		_statusCode: code,
	}
}

/*PatchBasketsByIDPaymentInstrumentsByIDDefault handles this case with default header values.

PatchBasketsByIDPaymentInstrumentsByIDDefault patch baskets by ID payment instruments by ID default
*/
type PatchBasketsByIDPaymentInstrumentsByIDDefault struct {
	_statusCode int

	Payload *models.Basket
}

// Code gets the status code for the patch baskets by ID payment instruments by ID default response
func (o *PatchBasketsByIDPaymentInstrumentsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /baskets/{basket_id}/payment_instruments/{payment_instrument_id}][%d] patchBasketsByIDPaymentInstrumentsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBasketsByIDPaymentInstrumentsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Basket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
