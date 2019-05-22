// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PatchOrdersByIDPaymentInstrumentsByIDReader is a Reader for the PatchOrdersByIDPaymentInstrumentsByID structure.
type PatchOrdersByIDPaymentInstrumentsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrdersByIDPaymentInstrumentsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchOrdersByIDPaymentInstrumentsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchOrdersByIDPaymentInstrumentsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchOrdersByIDPaymentInstrumentsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchOrdersByIDPaymentInstrumentsByIDBadRequest creates a PatchOrdersByIDPaymentInstrumentsByIDBadRequest with default headers values
func NewPatchOrdersByIDPaymentInstrumentsByIDBadRequest() *PatchOrdersByIDPaymentInstrumentsByIDBadRequest {
	return &PatchOrdersByIDPaymentInstrumentsByIDBadRequest{}
}

/*PatchOrdersByIDPaymentInstrumentsByIDBadRequest handles this case with default header values.

Indicates that the basket payment instrument with the given
             id already is permanently masked. or Indicates that the provided payment method is invalid or not applicable.
*/
type PatchOrdersByIDPaymentInstrumentsByIDBadRequest struct {
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /orders/{order_no}/payment_instruments/{payment_instrument_id}][%d] patchOrdersByIdPaymentInstrumentsByIdBadRequest ", 400)
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOrdersByIDPaymentInstrumentsByIDNotFound creates a PatchOrdersByIDPaymentInstrumentsByIDNotFound with default headers values
func NewPatchOrdersByIDPaymentInstrumentsByIDNotFound() *PatchOrdersByIDPaymentInstrumentsByIDNotFound {
	return &PatchOrdersByIDPaymentInstrumentsByIDNotFound{}
}

/*PatchOrdersByIDPaymentInstrumentsByIDNotFound handles this case with default header values.

Indicates that the order with the given order number is unknown. or Indicates that the payment instrument with the given payment
             instrument number is unknown.
*/
type PatchOrdersByIDPaymentInstrumentsByIDNotFound struct {
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /orders/{order_no}/payment_instruments/{payment_instrument_id}][%d] patchOrdersByIdPaymentInstrumentsByIdNotFound ", 404)
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOrdersByIDPaymentInstrumentsByIDDefault creates a PatchOrdersByIDPaymentInstrumentsByIDDefault with default headers values
func NewPatchOrdersByIDPaymentInstrumentsByIDDefault(code int) *PatchOrdersByIDPaymentInstrumentsByIDDefault {
	return &PatchOrdersByIDPaymentInstrumentsByIDDefault{
		_statusCode: code,
	}
}

/*PatchOrdersByIDPaymentInstrumentsByIDDefault handles this case with default header values.

PatchOrdersByIDPaymentInstrumentsByIDDefault patch orders by ID payment instruments by ID default
*/
type PatchOrdersByIDPaymentInstrumentsByIDDefault struct {
	_statusCode int

	Payload *models.Order
}

// Code gets the status code for the patch orders by ID payment instruments by ID default response
func (o *PatchOrdersByIDPaymentInstrumentsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /orders/{order_no}/payment_instruments/{payment_instrument_id}][%d] patchOrdersByIDPaymentInstrumentsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchOrdersByIDPaymentInstrumentsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}