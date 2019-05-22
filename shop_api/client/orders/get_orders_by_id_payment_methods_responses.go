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

// GetOrdersByIDPaymentMethodsReader is a Reader for the GetOrdersByIDPaymentMethods structure.
type GetOrdersByIDPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersByIDPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetOrdersByIDPaymentMethodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrdersByIDPaymentMethodsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetOrdersByIDPaymentMethodsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrdersByIDPaymentMethodsBadRequest creates a GetOrdersByIDPaymentMethodsBadRequest with default headers values
func NewGetOrdersByIDPaymentMethodsBadRequest() *GetOrdersByIDPaymentMethodsBadRequest {
	return &GetOrdersByIDPaymentMethodsBadRequest{}
}

/*GetOrdersByIDPaymentMethodsBadRequest handles this case with default header values.

Indicates that the customer assigned to the order does not
                match the verified customer represented by the JWT token, not
                relevant when using OAuth.
*/
type GetOrdersByIDPaymentMethodsBadRequest struct {
}

func (o *GetOrdersByIDPaymentMethodsBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/{order_no}/payment_methods][%d] getOrdersByIdPaymentMethodsBadRequest ", 400)
}

func (o *GetOrdersByIDPaymentMethodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrdersByIDPaymentMethodsNotFound creates a GetOrdersByIDPaymentMethodsNotFound with default headers values
func NewGetOrdersByIDPaymentMethodsNotFound() *GetOrdersByIDPaymentMethodsNotFound {
	return &GetOrdersByIDPaymentMethodsNotFound{}
}

/*GetOrdersByIDPaymentMethodsNotFound handles this case with default header values.

Indicates that the order with the given order number is unknown.
*/
type GetOrdersByIDPaymentMethodsNotFound struct {
}

func (o *GetOrdersByIDPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/{order_no}/payment_methods][%d] getOrdersByIdPaymentMethodsNotFound ", 404)
}

func (o *GetOrdersByIDPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrdersByIDPaymentMethodsDefault creates a GetOrdersByIDPaymentMethodsDefault with default headers values
func NewGetOrdersByIDPaymentMethodsDefault(code int) *GetOrdersByIDPaymentMethodsDefault {
	return &GetOrdersByIDPaymentMethodsDefault{
		_statusCode: code,
	}
}

/*GetOrdersByIDPaymentMethodsDefault handles this case with default header values.

GetOrdersByIDPaymentMethodsDefault get orders by ID payment methods default
*/
type GetOrdersByIDPaymentMethodsDefault struct {
	_statusCode int

	Payload *models.PaymentMethodResult
}

// Code gets the status code for the get orders by ID payment methods default response
func (o *GetOrdersByIDPaymentMethodsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrdersByIDPaymentMethodsDefault) Error() string {
	return fmt.Sprintf("[GET /orders/{order_no}/payment_methods][%d] getOrdersByIDPaymentMethods default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrdersByIDPaymentMethodsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}