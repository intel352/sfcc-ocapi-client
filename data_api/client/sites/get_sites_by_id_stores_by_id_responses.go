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

// GetSitesByIDStoresByIDReader is a Reader for the GetSitesByIDStoresByID structure.
type GetSitesByIDStoresByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDStoresByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetSitesByIDStoresByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDStoresByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDStoresByIDNotFound creates a GetSitesByIDStoresByIDNotFound with default headers values
func NewGetSitesByIDStoresByIDNotFound() *GetSitesByIDStoresByIDNotFound {
	return &GetSitesByIDStoresByIDNotFound{}
}

/*GetSitesByIDStoresByIDNotFound handles this case with default header values.

Thrown in case the store does not exist matching the given id
*/
type GetSitesByIDStoresByIDNotFound struct {
}

func (o *GetSitesByIDStoresByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/stores/{id}][%d] getSitesByIdStoresByIdNotFound ", 404)
}

func (o *GetSitesByIDStoresByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDStoresByIDDefault creates a GetSitesByIDStoresByIDDefault with default headers values
func NewGetSitesByIDStoresByIDDefault(code int) *GetSitesByIDStoresByIDDefault {
	return &GetSitesByIDStoresByIDDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDStoresByIDDefault handles this case with default header values.

GetSitesByIDStoresByIDDefault get sites by ID stores by ID default
*/
type GetSitesByIDStoresByIDDefault struct {
	_statusCode int

	Payload *models.Store
}

// Code gets the status code for the get sites by ID stores by ID default response
func (o *GetSitesByIDStoresByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDStoresByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/stores/{id}][%d] getSitesByIDStoresByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDStoresByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Store)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
