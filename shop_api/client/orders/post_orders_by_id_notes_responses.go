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

// PostOrdersByIDNotesReader is a Reader for the PostOrdersByIDNotes structure.
type PostOrdersByIDNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrdersByIDNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPostOrdersByIDNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostOrdersByIDNotesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOrdersByIDNotesNotFound creates a PostOrdersByIDNotesNotFound with default headers values
func NewPostOrdersByIDNotesNotFound() *PostOrdersByIDNotesNotFound {
	return &PostOrdersByIDNotesNotFound{}
}

/*PostOrdersByIDNotesNotFound handles this case with default header values.

Thrown if the order with the given order number is unknown.
*/
type PostOrdersByIDNotesNotFound struct {
}

func (o *PostOrdersByIDNotesNotFound) Error() string {
	return fmt.Sprintf("[POST /orders/{order_no}/notes][%d] postOrdersByIdNotesNotFound ", 404)
}

func (o *PostOrdersByIDNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOrdersByIDNotesDefault creates a PostOrdersByIDNotesDefault with default headers values
func NewPostOrdersByIDNotesDefault(code int) *PostOrdersByIDNotesDefault {
	return &PostOrdersByIDNotesDefault{
		_statusCode: code,
	}
}

/*PostOrdersByIDNotesDefault handles this case with default header values.

PostOrdersByIDNotesDefault post orders by ID notes default
*/
type PostOrdersByIDNotesDefault struct {
	_statusCode int

	Payload *models.Order
}

// Code gets the status code for the post orders by ID notes default response
func (o *PostOrdersByIDNotesDefault) Code() int {
	return o._statusCode
}

func (o *PostOrdersByIDNotesDefault) Error() string {
	return fmt.Sprintf("[POST /orders/{order_no}/notes][%d] postOrdersByIDNotes default  %+v", o._statusCode, o.Payload)
}

func (o *PostOrdersByIDNotesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}