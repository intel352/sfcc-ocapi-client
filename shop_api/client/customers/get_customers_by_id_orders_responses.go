// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetCustomersByIDOrdersReader is a Reader for the GetCustomersByIDOrders structure.
type GetCustomersByIDOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersByIDOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetCustomersByIDOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCustomersByIDOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomersByIDOrdersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersByIDOrdersBadRequest creates a GetCustomersByIDOrdersBadRequest with default headers values
func NewGetCustomersByIDOrdersBadRequest() *GetCustomersByIDOrdersBadRequest {
	return &GetCustomersByIDOrdersBadRequest{}
}

/*GetCustomersByIDOrdersBadRequest handles this case with default header values.

Indicates that the customerId URL parameter does not match the verified customer
             represented by the JWT token, not relevant when using OAuth. or Thrown if specified status is unknown.
*/
type GetCustomersByIDOrdersBadRequest struct {
}

func (o *GetCustomersByIDOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/orders][%d] getCustomersByIdOrdersBadRequest ", 400)
}

func (o *GetCustomersByIDOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDOrdersNotFound creates a GetCustomersByIDOrdersNotFound with default headers values
func NewGetCustomersByIDOrdersNotFound() *GetCustomersByIDOrdersNotFound {
	return &GetCustomersByIDOrdersNotFound{}
}

/*GetCustomersByIDOrdersNotFound handles this case with default header values.

Indicates that the customer with the given customer id is unknown for the site.
*/
type GetCustomersByIDOrdersNotFound struct {
}

func (o *GetCustomersByIDOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/orders][%d] getCustomersByIdOrdersNotFound ", 404)
}

func (o *GetCustomersByIDOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDOrdersDefault creates a GetCustomersByIDOrdersDefault with default headers values
func NewGetCustomersByIDOrdersDefault(code int) *GetCustomersByIDOrdersDefault {
	return &GetCustomersByIDOrdersDefault{
		_statusCode: code,
	}
}

/*GetCustomersByIDOrdersDefault handles this case with default header values.

GetCustomersByIDOrdersDefault get customers by ID orders default
*/
type GetCustomersByIDOrdersDefault struct {
	_statusCode int

	Payload *models.CustomerOrderResult
}

// Code gets the status code for the get customers by ID orders default response
func (o *GetCustomersByIDOrdersDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersByIDOrdersDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/orders][%d] getCustomersByIDOrders default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersByIDOrdersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerOrderResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
