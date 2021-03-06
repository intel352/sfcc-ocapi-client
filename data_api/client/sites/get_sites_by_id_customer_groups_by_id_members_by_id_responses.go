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

// GetSitesByIDCustomerGroupsByIDMembersByIDReader is a Reader for the GetSitesByIDCustomerGroupsByIDMembersByID structure.
type GetSitesByIDCustomerGroupsByIDMembersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDCustomerGroupsByIDMembersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSitesByIDCustomerGroupsByIDMembersByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDCustomerGroupsByIDMembersByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDCustomerGroupsByIDMembersByIDNotFound creates a GetSitesByIDCustomerGroupsByIDMembersByIDNotFound with default headers values
func NewGetSitesByIDCustomerGroupsByIDMembersByIDNotFound() *GetSitesByIDCustomerGroupsByIDMembersByIDNotFound {
	return &GetSitesByIDCustomerGroupsByIDMembersByIDNotFound{}
}

/*GetSitesByIDCustomerGroupsByIDMembersByIDNotFound handles this case with default header values.

Thrown in case the customer group does not exist matching the given id or Thrown in case the customer group member did not exist, or the group is not a static group
*/
type GetSitesByIDCustomerGroupsByIDMembersByIDNotFound struct {
}

func (o *GetSitesByIDCustomerGroupsByIDMembersByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/customer_groups/{id}/members/{customer_no}][%d] getSitesByIdCustomerGroupsByIdMembersByIdNotFound ", 404)
}

func (o *GetSitesByIDCustomerGroupsByIDMembersByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDCustomerGroupsByIDMembersByIDDefault creates a GetSitesByIDCustomerGroupsByIDMembersByIDDefault with default headers values
func NewGetSitesByIDCustomerGroupsByIDMembersByIDDefault(code int) *GetSitesByIDCustomerGroupsByIDMembersByIDDefault {
	return &GetSitesByIDCustomerGroupsByIDMembersByIDDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDCustomerGroupsByIDMembersByIDDefault handles this case with default header values.

GetSitesByIDCustomerGroupsByIDMembersByIDDefault get sites by ID customer groups by ID members by ID default
*/
type GetSitesByIDCustomerGroupsByIDMembersByIDDefault struct {
	_statusCode int

	Payload *models.CustomerGroupMember
}

// Code gets the status code for the get sites by ID customer groups by ID members by ID default response
func (o *GetSitesByIDCustomerGroupsByIDMembersByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDCustomerGroupsByIDMembersByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/customer_groups/{id}/members/{customer_no}][%d] getSitesByIDCustomerGroupsByIDMembersByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDCustomerGroupsByIDMembersByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerGroupMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
