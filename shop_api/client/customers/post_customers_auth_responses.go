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

// PostCustomersAuthReader is a Reader for the PostCustomersAuth structure.
type PostCustomersAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomersAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostCustomersAuthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostCustomersAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostCustomersAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostCustomersAuthBadRequest creates a PostCustomersAuthBadRequest with default headers values
func NewPostCustomersAuthBadRequest() *PostCustomersAuthBadRequest {
	return &PostCustomersAuthBadRequest{}
}

/*PostCustomersAuthBadRequest handles this case with default header values.

Indicates that no HTTP Authorization:Basic header was
             provided.
*/
type PostCustomersAuthBadRequest struct {
}

func (o *PostCustomersAuthBadRequest) Error() string {
	return fmt.Sprintf("[POST /customers/auth][%d] postCustomersAuthBadRequest ", 400)
}

func (o *PostCustomersAuthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersAuthUnauthorized creates a PostCustomersAuthUnauthorized with default headers values
func NewPostCustomersAuthUnauthorized() *PostCustomersAuthUnauthorized {
	return &PostCustomersAuthUnauthorized{}
}

/*PostCustomersAuthUnauthorized handles this case with default header values.

Indicates in case of type credentials the username is unknown or the password does
             not match. In case of type session the session is not active anymore or
             the dwsecuretoken value is invalid. In both cases the customer is disabled or locked.
*/
type PostCustomersAuthUnauthorized struct {
}

func (o *PostCustomersAuthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /customers/auth][%d] postCustomersAuthUnauthorized ", 401)
}

func (o *PostCustomersAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCustomersAuthDefault creates a PostCustomersAuthDefault with default headers values
func NewPostCustomersAuthDefault(code int) *PostCustomersAuthDefault {
	return &PostCustomersAuthDefault{
		_statusCode: code,
	}
}

/*PostCustomersAuthDefault handles this case with default header values.

PostCustomersAuthDefault post customers auth default
*/
type PostCustomersAuthDefault struct {
	_statusCode int

	Payload *models.Customer
}

// Code gets the status code for the post customers auth default response
func (o *PostCustomersAuthDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomersAuthDefault) Error() string {
	return fmt.Sprintf("[POST /customers/auth][%d] postCustomersAuth default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomersAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Customer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
