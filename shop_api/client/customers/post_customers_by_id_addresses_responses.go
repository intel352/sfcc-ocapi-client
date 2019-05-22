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

// PostCustomersByIDAddressesReader is a Reader for the PostCustomersByIDAddresses structure.
type PostCustomersByIDAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersByIDAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCustomersByIDAddressesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCustomersByIDAddressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCustomersByIDAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersByIDAddressesBadRequest creates a PostCustomersByIDAddressesBadRequest with default headers values
func NewPostCustomersByIDAddressesBadRequest() *PostCustomersByIDAddressesBadRequest {
	return &PostCustomersByIDAddressesBadRequest{}
}

/*PostCustomersByIDAddressesBadRequest handles this case with default header values.

Indicates that the customerId URL parameter does not match the verified customer
             represented by the JWT token, not relevant when using OAuth. or Indicates that address name is not provided or it's blank. or Indicates that the provided address name is already used for the customer.
*/
type PostCustomersByIDAddressesBadRequest struct {
}

func (o *PostCustomersByIDAddressesBadRequest) Error() string {
	return fmt.Sprintf("[POST /customers/{customer_id}/addresses][%d] postCustomersByIdAddressesBadRequest ", 400)
}

func (o *PostCustomersByIDAddressesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersByIDAddressesNotFound creates a PostCustomersByIDAddressesNotFound with default headers values
func NewPostCustomersByIDAddressesNotFound() *PostCustomersByIDAddressesNotFound {
	return &PostCustomersByIDAddressesNotFound{}
}

/*PostCustomersByIDAddressesNotFound handles this case with default header values.

Indicates that the customer with the given customer id is unknown.
*/
type PostCustomersByIDAddressesNotFound struct {
}

func (o *PostCustomersByIDAddressesNotFound) Error() string {
	return fmt.Sprintf("[POST /customers/{customer_id}/addresses][%d] postCustomersByIdAddressesNotFound ", 404)
}

func (o *PostCustomersByIDAddressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersByIDAddressesDefault creates a PostCustomersByIDAddressesDefault with default headers values
func NewPostCustomersByIDAddressesDefault(code int) *PostCustomersByIDAddressesDefault {
	return &PostCustomersByIDAddressesDefault{
		_statusCode: code,
	}
}

/*PostCustomersByIDAddressesDefault handles this case with default header values.

PostCustomersByIDAddressesDefault post customers by ID addresses default
*/
type PostCustomersByIDAddressesDefault struct {
	_statusCode int

	Payload *models.CustomerAddress
}

// Code gets the status code for the post customers by ID addresses default response
func (o *PostCustomersByIDAddressesDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersByIDAddressesDefault) Error() string {
	return fmt.Sprintf("[POST /customers/{customer_id}/addresses][%d] postCustomersByIDAddresses default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersByIDAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}