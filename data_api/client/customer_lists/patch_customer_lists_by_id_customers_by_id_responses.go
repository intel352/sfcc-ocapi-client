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

// PatchCustomerListsByIDCustomersByIDReader is a Reader for the PatchCustomerListsByIDCustomersByID structure.
type PatchCustomerListsByIDCustomersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCustomerListsByIDCustomersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchCustomerListsByIDCustomersByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchCustomerListsByIDCustomersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchCustomerListsByIDCustomersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCustomerListsByIDCustomersByIDBadRequest creates a PatchCustomerListsByIDCustomersByIDBadRequest with default headers values
func NewPatchCustomerListsByIDCustomersByIDBadRequest() *PatchCustomerListsByIDCustomersByIDBadRequest {
	return &PatchCustomerListsByIDCustomersByIDBadRequest{}
}

/*PatchCustomerListsByIDCustomersByIDBadRequest handles this case with default header values.

Indicates the login does not match the login acceptance criteria. or Indicates the login is already in use.
*/
type PatchCustomerListsByIDCustomersByIDBadRequest struct {
}

func (o *PatchCustomerListsByIDCustomersByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /customer_lists/{list_id}/customers/{customer_no}][%d] patchCustomerListsByIdCustomersByIdBadRequest ", 400)
}

func (o *PatchCustomerListsByIDCustomersByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomerListsByIDCustomersByIDNotFound creates a PatchCustomerListsByIDCustomersByIDNotFound with default headers values
func NewPatchCustomerListsByIDCustomersByIDNotFound() *PatchCustomerListsByIDCustomersByIDNotFound {
	return &PatchCustomerListsByIDCustomersByIDNotFound{}
}

/*PatchCustomerListsByIDCustomersByIDNotFound handles this case with default header values.

Indicates that the customer with the given customer number is unknown. or Indicates that the customer list with the given customer list id is unknown.
*/
type PatchCustomerListsByIDCustomersByIDNotFound struct {
}

func (o *PatchCustomerListsByIDCustomersByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /customer_lists/{list_id}/customers/{customer_no}][%d] patchCustomerListsByIdCustomersByIdNotFound ", 404)
}

func (o *PatchCustomerListsByIDCustomersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCustomerListsByIDCustomersByIDDefault creates a PatchCustomerListsByIDCustomersByIDDefault with default headers values
func NewPatchCustomerListsByIDCustomersByIDDefault(code int) *PatchCustomerListsByIDCustomersByIDDefault {
	return &PatchCustomerListsByIDCustomersByIDDefault{
		_statusCode: code,
	}
}

/*PatchCustomerListsByIDCustomersByIDDefault handles this case with default header values.

PatchCustomerListsByIDCustomersByIDDefault patch customer lists by ID customers by ID default
*/
type PatchCustomerListsByIDCustomersByIDDefault struct {
	_statusCode int

	Payload *models.Customer
}

// Code gets the status code for the patch customer lists by ID customers by ID default response
func (o *PatchCustomerListsByIDCustomersByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCustomerListsByIDCustomersByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /customer_lists/{list_id}/customers/{customer_no}][%d] patchCustomerListsByIDCustomersByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCustomerListsByIDCustomersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
