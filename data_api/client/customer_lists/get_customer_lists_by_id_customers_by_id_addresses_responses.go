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

// GetCustomerListsByIDCustomersByIDAddressesReader is a Reader for the GetCustomerListsByIDCustomersByIDAddresses structure.
type GetCustomerListsByIDCustomersByIDAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerListsByIDCustomersByIDAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetCustomerListsByIDCustomersByIDAddressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomerListsByIDCustomersByIDAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomerListsByIDCustomersByIDAddressesNotFound creates a GetCustomerListsByIDCustomersByIDAddressesNotFound with default headers values
func NewGetCustomerListsByIDCustomersByIDAddressesNotFound() *GetCustomerListsByIDCustomersByIDAddressesNotFound {
	return &GetCustomerListsByIDCustomersByIDAddressesNotFound{}
}

/*GetCustomerListsByIDCustomersByIDAddressesNotFound handles this case with default header values.

Indicates that the customer with the given customer number is unknown. or Indicates that the customer list with the given customer list id is unknown.
*/
type GetCustomerListsByIDCustomersByIDAddressesNotFound struct {
}

func (o *GetCustomerListsByIDCustomersByIDAddressesNotFound) Error() string {
	return fmt.Sprintf("[GET /customer_lists/{list_id}/customers/{customer_no}/addresses][%d] getCustomerListsByIdCustomersByIdAddressesNotFound ", 404)
}

func (o *GetCustomerListsByIDCustomersByIDAddressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomerListsByIDCustomersByIDAddressesDefault creates a GetCustomerListsByIDCustomersByIDAddressesDefault with default headers values
func NewGetCustomerListsByIDCustomersByIDAddressesDefault(code int) *GetCustomerListsByIDCustomersByIDAddressesDefault {
	return &GetCustomerListsByIDCustomersByIDAddressesDefault{
		_statusCode: code,
	}
}

/*GetCustomerListsByIDCustomersByIDAddressesDefault handles this case with default header values.

GetCustomerListsByIDCustomersByIDAddressesDefault get customer lists by ID customers by ID addresses default
*/
type GetCustomerListsByIDCustomersByIDAddressesDefault struct {
	_statusCode int

	Payload *models.CustomerAddressResult
}

// Code gets the status code for the get customer lists by ID customers by ID addresses default response
func (o *GetCustomerListsByIDCustomersByIDAddressesDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomerListsByIDCustomersByIDAddressesDefault) Error() string {
	return fmt.Sprintf("[GET /customer_lists/{list_id}/customers/{customer_no}/addresses][%d] getCustomerListsByIDCustomersByIDAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomerListsByIDCustomersByIDAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerAddressResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
