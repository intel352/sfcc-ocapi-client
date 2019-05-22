// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PutSitesByIDCustomerGroupsByIDMembersByIDReader is a Reader for the PutSitesByIDCustomerGroupsByIDMembersByID structure.
type PutSitesByIDCustomerGroupsByIDMembersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSitesByIDCustomerGroupsByIDMembersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 403:
		result := NewPutSitesByIDCustomerGroupsByIDMembersByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutSitesByIDCustomerGroupsByIDMembersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutSitesByIDCustomerGroupsByIDMembersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutSitesByIDCustomerGroupsByIDMembersByIDForbidden creates a PutSitesByIDCustomerGroupsByIDMembersByIDForbidden with default headers values
func NewPutSitesByIDCustomerGroupsByIDMembersByIDForbidden() *PutSitesByIDCustomerGroupsByIDMembersByIDForbidden {
	return &PutSitesByIDCustomerGroupsByIDMembersByIDForbidden{}
}

/*PutSitesByIDCustomerGroupsByIDMembersByIDForbidden handles this case with default header values.

Thrown when the customer group is not a static group.
*/
type PutSitesByIDCustomerGroupsByIDMembersByIDForbidden struct {
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/customer_groups/{id}/members/{customer_no}][%d] putSitesByIdCustomerGroupsByIdMembersByIdForbidden ", 403)
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDCustomerGroupsByIDMembersByIDNotFound creates a PutSitesByIDCustomerGroupsByIDMembersByIDNotFound with default headers values
func NewPutSitesByIDCustomerGroupsByIDMembersByIDNotFound() *PutSitesByIDCustomerGroupsByIDMembersByIDNotFound {
	return &PutSitesByIDCustomerGroupsByIDMembersByIDNotFound{}
}

/*PutSitesByIDCustomerGroupsByIDMembersByIDNotFound handles this case with default header values.

Thrown in case the customer group does not exist matching the given id. or Thrown in case the customer list did not exist, or the customer does not exist in in the list.
*/
type PutSitesByIDCustomerGroupsByIDMembersByIDNotFound struct {
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/customer_groups/{id}/members/{customer_no}][%d] putSitesByIdCustomerGroupsByIdMembersByIdNotFound ", 404)
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDCustomerGroupsByIDMembersByIDDefault creates a PutSitesByIDCustomerGroupsByIDMembersByIDDefault with default headers values
func NewPutSitesByIDCustomerGroupsByIDMembersByIDDefault(code int) *PutSitesByIDCustomerGroupsByIDMembersByIDDefault {
	return &PutSitesByIDCustomerGroupsByIDMembersByIDDefault{
		_statusCode: code,
	}
}

/*PutSitesByIDCustomerGroupsByIDMembersByIDDefault handles this case with default header values.

PutSitesByIDCustomerGroupsByIDMembersByIDDefault put sites by ID customer groups by ID members by ID default
*/
type PutSitesByIDCustomerGroupsByIDMembersByIDDefault struct {
	_statusCode int

	Payload *models.CustomerGroupMember
}

// Code gets the status code for the put sites by ID customer groups by ID members by ID default response
func (o *PutSitesByIDCustomerGroupsByIDMembersByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/customer_groups/{id}/members/{customer_no}][%d] putSitesByIDCustomerGroupsByIDMembersByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutSitesByIDCustomerGroupsByIDMembersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerGroupMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}