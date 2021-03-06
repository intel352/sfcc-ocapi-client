// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// DeleteOrdersByIDNotesByIDReader is a Reader for the DeleteOrdersByIDNotesByID structure.
type DeleteOrdersByIDNotesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrdersByIDNotesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewDeleteOrdersByIDNotesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteOrdersByIDNotesByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrdersByIDNotesByIDNotFound creates a DeleteOrdersByIDNotesByIDNotFound with default headers values
func NewDeleteOrdersByIDNotesByIDNotFound() *DeleteOrdersByIDNotesByIDNotFound {
	return &DeleteOrdersByIDNotesByIDNotFound{}
}

/*DeleteOrdersByIDNotesByIDNotFound handles this case with default header values.

Indicates that the order with the given order number is unknown. or Indicates that the order with the given order number is unknown.
*/
type DeleteOrdersByIDNotesByIDNotFound struct {
}

func (o *DeleteOrdersByIDNotesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orders/{order_no}/notes/{note_id}][%d] deleteOrdersByIdNotesByIdNotFound ", 404)
}

func (o *DeleteOrdersByIDNotesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrdersByIDNotesByIDDefault creates a DeleteOrdersByIDNotesByIDDefault with default headers values
func NewDeleteOrdersByIDNotesByIDDefault(code int) *DeleteOrdersByIDNotesByIDDefault {
	return &DeleteOrdersByIDNotesByIDDefault{
		_statusCode: code,
	}
}

/*DeleteOrdersByIDNotesByIDDefault handles this case with default header values.

DeleteOrdersByIDNotesByIDDefault delete orders by ID notes by ID default
*/
type DeleteOrdersByIDNotesByIDDefault struct {
	_statusCode int

	Payload *models.Order
}

// Code gets the status code for the delete orders by ID notes by ID default response
func (o *DeleteOrdersByIDNotesByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrdersByIDNotesByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /orders/{order_no}/notes/{note_id}][%d] deleteOrdersByIDNotesByID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrdersByIDNotesByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
