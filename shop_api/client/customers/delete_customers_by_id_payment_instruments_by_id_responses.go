// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCustomersByIDPaymentInstrumentsByIDReader is a Reader for the DeleteCustomersByIDPaymentInstrumentsByID structure.
type DeleteCustomersByIDPaymentInstrumentsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomersByIDPaymentInstrumentsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCustomersByIDPaymentInstrumentsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteCustomersByIDPaymentInstrumentsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCustomersByIDPaymentInstrumentsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDNoContent creates a DeleteCustomersByIDPaymentInstrumentsByIDNoContent with default headers values
func NewDeleteCustomersByIDPaymentInstrumentsByIDNoContent() *DeleteCustomersByIDPaymentInstrumentsByIDNoContent {
	return &DeleteCustomersByIDPaymentInstrumentsByIDNoContent{}
}

/*DeleteCustomersByIDPaymentInstrumentsByIDNoContent handles this case with default header values.

DeleteCustomersByIDPaymentInstrumentsByIDNoContent delete customers by Id payment instruments by Id no content
*/
type DeleteCustomersByIDPaymentInstrumentsByIDNoContent struct {
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customer_id}/payment_instruments/{payment_instrument_id}][%d] deleteCustomersByIdPaymentInstrumentsByIdNoContent ", 204)
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDBadRequest creates a DeleteCustomersByIDPaymentInstrumentsByIDBadRequest with default headers values
func NewDeleteCustomersByIDPaymentInstrumentsByIDBadRequest() *DeleteCustomersByIDPaymentInstrumentsByIDBadRequest {
	return &DeleteCustomersByIDPaymentInstrumentsByIDBadRequest{}
}

/*DeleteCustomersByIDPaymentInstrumentsByIDBadRequest handles this case with default header values.

Indicates that the customerId URL parameter does not match the
             verified customer represented by the JWT token, not relevant
             when using OAuth.
*/
type DeleteCustomersByIDPaymentInstrumentsByIDBadRequest struct {
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customer_id}/payment_instruments/{payment_instrument_id}][%d] deleteCustomersByIdPaymentInstrumentsByIdBadRequest ", 400)
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomersByIDPaymentInstrumentsByIDNotFound creates a DeleteCustomersByIDPaymentInstrumentsByIDNotFound with default headers values
func NewDeleteCustomersByIDPaymentInstrumentsByIDNotFound() *DeleteCustomersByIDPaymentInstrumentsByIDNotFound {
	return &DeleteCustomersByIDPaymentInstrumentsByIDNotFound{}
}

/*DeleteCustomersByIDPaymentInstrumentsByIDNotFound handles this case with default header values.

Indicates that the customer with the given customer id is
             unknown for the site.
*/
type DeleteCustomersByIDPaymentInstrumentsByIDNotFound struct {
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /customers/{customer_id}/payment_instruments/{payment_instrument_id}][%d] deleteCustomersByIdPaymentInstrumentsByIdNotFound ", 404)
}

func (o *DeleteCustomersByIDPaymentInstrumentsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
