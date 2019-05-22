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

// PostCustomerListsByIDCustomersReader is a Reader for the PostCustomerListsByIDCustomers structure.
type PostCustomerListsByIDCustomersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomerListsByIDCustomersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCustomerListsByIDCustomersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostCustomerListsByIDCustomersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCustomerListsByIDCustomersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomerListsByIDCustomersBadRequest creates a PostCustomerListsByIDCustomersBadRequest with default headers values
func NewPostCustomerListsByIDCustomersBadRequest() *PostCustomerListsByIDCustomersBadRequest {
	return &PostCustomerListsByIDCustomersBadRequest{}
}

/*PostCustomerListsByIDCustomersBadRequest handles this case with default header values.

Indicates the login does not match the login acceptance criteria. or Indicates the login is already in use. or Indicates that the mandatory credentials are missing in the input document. or Indicates that the mandatory login property is missing in the input document.
*/
type PostCustomerListsByIDCustomersBadRequest struct {
}

func (o *PostCustomerListsByIDCustomersBadRequest) Error() string {
	return fmt.Sprintf("[POST /customer_lists/{list_id}/customers][%d] postCustomerListsByIdCustomersBadRequest ", 400)
}

func (o *PostCustomerListsByIDCustomersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomerListsByIDCustomersNotFound creates a PostCustomerListsByIDCustomersNotFound with default headers values
func NewPostCustomerListsByIDCustomersNotFound() *PostCustomerListsByIDCustomersNotFound {
	return &PostCustomerListsByIDCustomersNotFound{}
}

/*PostCustomerListsByIDCustomersNotFound handles this case with default header values.

Indicates that the customer list with the given customer list id is unknown.
*/
type PostCustomerListsByIDCustomersNotFound struct {
}

func (o *PostCustomerListsByIDCustomersNotFound) Error() string {
	return fmt.Sprintf("[POST /customer_lists/{list_id}/customers][%d] postCustomerListsByIdCustomersNotFound ", 404)
}

func (o *PostCustomerListsByIDCustomersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomerListsByIDCustomersDefault creates a PostCustomerListsByIDCustomersDefault with default headers values
func NewPostCustomerListsByIDCustomersDefault(code int) *PostCustomerListsByIDCustomersDefault {
	return &PostCustomerListsByIDCustomersDefault{
		_statusCode: code,
	}
}

/*PostCustomerListsByIDCustomersDefault handles this case with default header values.

PostCustomerListsByIDCustomersDefault post customer lists by ID customers default
*/
type PostCustomerListsByIDCustomersDefault struct {
	_statusCode int

	Payload *models.Customer
}

// Code gets the status code for the post customer lists by ID customers default response
func (o *PostCustomerListsByIDCustomersDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomerListsByIDCustomersDefault) Error() string {
	return fmt.Sprintf("[POST /customer_lists/{list_id}/customers][%d] postCustomerListsByIDCustomers default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomerListsByIDCustomersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}