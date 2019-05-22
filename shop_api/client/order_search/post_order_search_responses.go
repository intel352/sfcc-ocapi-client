// Code generated by go-swagger; DO NOT EDIT.

package order_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PostOrderSearchReader is a Reader for the PostOrderSearch structure.
type PostOrderSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrderSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostOrderSearchDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostOrderSearchDefault creates a PostOrderSearchDefault with default headers values
func NewPostOrderSearchDefault(code int) *PostOrderSearchDefault {
	return &PostOrderSearchDefault{
		_statusCode: code,
	}
}

/*PostOrderSearchDefault handles this case with default header values.

PostOrderSearchDefault post order search default
*/
type PostOrderSearchDefault struct {
	_statusCode int

	Payload *models.OrderSearchResult
}

// Code gets the status code for the post order search default response
func (o *PostOrderSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostOrderSearchDefault) Error() string {
	return fmt.Sprintf("[POST /order_search][%d] postOrderSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostOrderSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
