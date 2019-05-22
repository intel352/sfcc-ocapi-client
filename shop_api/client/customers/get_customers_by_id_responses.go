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

// GetCustomersByIDReader is a Reader for the GetCustomersByID structure.
type GetCustomersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetCustomersByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCustomersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersByIDBadRequest creates a GetCustomersByIDBadRequest with default headers values
func NewGetCustomersByIDBadRequest() *GetCustomersByIDBadRequest {
	return &GetCustomersByIDBadRequest{}
}

/*GetCustomersByIDBadRequest handles this case with default header values.

If customerId URL parameter does not match the verified
             customer represented by the JWT token (not relevant when
             using OAuth).
*/
type GetCustomersByIDBadRequest struct {
}

func (o *GetCustomersByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}][%d] getCustomersByIdBadRequest ", 400)
}

func (o *GetCustomersByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDNotFound creates a GetCustomersByIDNotFound with default headers values
func NewGetCustomersByIDNotFound() *GetCustomersByIDNotFound {
	return &GetCustomersByIDNotFound{}
}

/*GetCustomersByIDNotFound handles this case with default header values.

Indicates that the customer with the given customer id is
             unknown.
*/
type GetCustomersByIDNotFound struct {
}

func (o *GetCustomersByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}][%d] getCustomersByIdNotFound ", 404)
}

func (o *GetCustomersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDDefault creates a GetCustomersByIDDefault with default headers values
func NewGetCustomersByIDDefault(code int) *GetCustomersByIDDefault {
	return &GetCustomersByIDDefault{
		_statusCode: code,
	}
}

/*GetCustomersByIDDefault handles this case with default header values.

GetCustomersByIDDefault get customers by ID default
*/
type GetCustomersByIDDefault struct {
	_statusCode int

	Payload *models.Customer
}

// Code gets the status code for the get customers by ID default response
func (o *GetCustomersByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersByIDDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}][%d] getCustomersByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}