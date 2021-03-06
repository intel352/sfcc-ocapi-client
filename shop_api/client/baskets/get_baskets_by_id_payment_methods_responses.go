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

// GetBasketsByIDPaymentMethodsReader is a Reader for the GetBasketsByIDPaymentMethods structure.
type GetBasketsByIDPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBasketsByIDPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetBasketsByIDPaymentMethodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBasketsByIDPaymentMethodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetBasketsByIDPaymentMethodsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBasketsByIDPaymentMethodsBadRequest creates a GetBasketsByIDPaymentMethodsBadRequest with default headers values
func NewGetBasketsByIDPaymentMethodsBadRequest() *GetBasketsByIDPaymentMethodsBadRequest {
	return &GetBasketsByIDPaymentMethodsBadRequest{}
}

/*GetBasketsByIDPaymentMethodsBadRequest handles this case with default header values.

Indicates that the customer assigned to the basket does not match the
                verified customer represented by the JWT token, not relevant
                when using OAuth.
*/
type GetBasketsByIDPaymentMethodsBadRequest struct {
}

func (o *GetBasketsByIDPaymentMethodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /baskets/{basket_id}/payment_methods][%d] getBasketsByIdPaymentMethodsBadRequest ", 400)
}

func (o *GetBasketsByIDPaymentMethodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBasketsByIDPaymentMethodsNotFound creates a GetBasketsByIDPaymentMethodsNotFound with default headers values
func NewGetBasketsByIDPaymentMethodsNotFound() *GetBasketsByIDPaymentMethodsNotFound {
	return &GetBasketsByIDPaymentMethodsNotFound{}
}

/*GetBasketsByIDPaymentMethodsNotFound handles this case with default header values.

Indicates that the basket with the given basket id is
                unknown.
*/
type GetBasketsByIDPaymentMethodsNotFound struct {
}

func (o *GetBasketsByIDPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[GET /baskets/{basket_id}/payment_methods][%d] getBasketsByIdPaymentMethodsNotFound ", 404)
}

func (o *GetBasketsByIDPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBasketsByIDPaymentMethodsDefault creates a GetBasketsByIDPaymentMethodsDefault with default headers values
func NewGetBasketsByIDPaymentMethodsDefault(code int) *GetBasketsByIDPaymentMethodsDefault {
	return &GetBasketsByIDPaymentMethodsDefault{
		_statusCode: code,
	}
}

/*GetBasketsByIDPaymentMethodsDefault handles this case with default header values.

GetBasketsByIDPaymentMethodsDefault get baskets by ID payment methods default
*/
type GetBasketsByIDPaymentMethodsDefault struct {
	_statusCode int

	Payload *models.PaymentMethodResult
}

// Code gets the status code for the get baskets by ID payment methods default response
func (o *GetBasketsByIDPaymentMethodsDefault) Code() int {
	return o._statusCode
}

func (o *GetBasketsByIDPaymentMethodsDefault) Error() string {
	return fmt.Sprintf("[GET /baskets/{basket_id}/payment_methods][%d] getBasketsByIDPaymentMethods default  %+v", o._statusCode, o.Payload)
}

func (o *GetBasketsByIDPaymentMethodsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
