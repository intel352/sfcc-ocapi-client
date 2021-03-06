// Code generated by go-swagger; DO NOT EDIT.

package inventory_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetInventoryListsByIDReader is a Reader for the GetInventoryListsByID structure.
type GetInventoryListsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryListsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetInventoryListsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInventoryListsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInventoryListsByIDNotFound creates a GetInventoryListsByIDNotFound with default headers values
func NewGetInventoryListsByIDNotFound() *GetInventoryListsByIDNotFound {
	return &GetInventoryListsByIDNotFound{}
}

/*GetInventoryListsByIDNotFound handles this case with default header values.

Thrown in case the inventory list does not exist matching the given id
*/
type GetInventoryListsByIDNotFound struct {
}

func (o *GetInventoryListsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /inventory_lists/{id}][%d] getInventoryListsByIdNotFound ", 404)
}

func (o *GetInventoryListsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryListsByIDDefault creates a GetInventoryListsByIDDefault with default headers values
func NewGetInventoryListsByIDDefault(code int) *GetInventoryListsByIDDefault {
	return &GetInventoryListsByIDDefault{
		_statusCode: code,
	}
}

/*GetInventoryListsByIDDefault handles this case with default header values.

GetInventoryListsByIDDefault get inventory lists by ID default
*/
type GetInventoryListsByIDDefault struct {
	_statusCode int

	Payload *models.InventoryList
}

// Code gets the status code for the get inventory lists by ID default response
func (o *GetInventoryListsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetInventoryListsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /inventory_lists/{id}][%d] getInventoryListsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetInventoryListsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
