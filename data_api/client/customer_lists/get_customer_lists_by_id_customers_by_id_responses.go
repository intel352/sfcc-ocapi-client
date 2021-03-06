// Code generated by go-swagger; DO NOT EDIT.

package customer_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetCustomerListsByIDCustomersByIDReader is a Reader for the GetCustomerListsByIDCustomersByID structure.
type GetCustomerListsByIDCustomersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerListsByIDCustomersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCustomerListsByIDCustomersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomerListsByIDCustomersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomerListsByIDCustomersByIDNotFound creates a GetCustomerListsByIDCustomersByIDNotFound with default headers values
func NewGetCustomerListsByIDCustomersByIDNotFound() *GetCustomerListsByIDCustomersByIDNotFound {
	return &GetCustomerListsByIDCustomersByIDNotFound{}
}

/*GetCustomerListsByIDCustomersByIDNotFound handles this case with default header values.

Indicates that the customer with the given customer number is unknown. or Indicates that the customer list with the given customer list id is unknown.
*/
type GetCustomerListsByIDCustomersByIDNotFound struct {
}

func (o *GetCustomerListsByIDCustomersByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /customer_lists/{list_id}/customers/{customer_no}][%d] getCustomerListsByIdCustomersByIdNotFound ", 404)
}

func (o *GetCustomerListsByIDCustomersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomerListsByIDCustomersByIDDefault creates a GetCustomerListsByIDCustomersByIDDefault with default headers values
func NewGetCustomerListsByIDCustomersByIDDefault(code int) *GetCustomerListsByIDCustomersByIDDefault {
	return &GetCustomerListsByIDCustomersByIDDefault{
		_statusCode: code,
	}
}

/*GetCustomerListsByIDCustomersByIDDefault handles this case with default header values.

GetCustomerListsByIDCustomersByIDDefault get customer lists by ID customers by ID default
*/
type GetCustomerListsByIDCustomersByIDDefault struct {
	_statusCode int

	Payload *models.Customer
}

// Code gets the status code for the get customer lists by ID customers by ID default response
func (o *GetCustomerListsByIDCustomersByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomerListsByIDCustomersByIDDefault) Error() string {
	return fmt.Sprintf("[GET /customer_lists/{list_id}/customers/{customer_no}][%d] getCustomerListsByIDCustomersByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomerListsByIDCustomersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
