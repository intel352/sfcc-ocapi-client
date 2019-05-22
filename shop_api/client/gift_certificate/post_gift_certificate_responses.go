// Code generated by go-swagger; DO NOT EDIT.

package gift_certificate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PostGiftCertificateReader is a Reader for the PostGiftCertificate structure.
type PostGiftCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGiftCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPostGiftCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostGiftCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostGiftCertificateNotFound creates a PostGiftCertificateNotFound with default headers values
func NewPostGiftCertificateNotFound() *PostGiftCertificateNotFound {
	return &PostGiftCertificateNotFound{}
}

/*PostGiftCertificateNotFound handles this case with default header values.

Thrown if the given gift certificate code is not valid.
*/
type PostGiftCertificateNotFound struct {
}

func (o *PostGiftCertificateNotFound) Error() string {
	return fmt.Sprintf("[POST /gift_certificate][%d] postGiftCertificateNotFound ", 404)
}

func (o *PostGiftCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGiftCertificateDefault creates a PostGiftCertificateDefault with default headers values
func NewPostGiftCertificateDefault(code int) *PostGiftCertificateDefault {
	return &PostGiftCertificateDefault{
		_statusCode: code,
	}
}

/*PostGiftCertificateDefault handles this case with default header values.

PostGiftCertificateDefault post gift certificate default
*/
type PostGiftCertificateDefault struct {
	_statusCode int

	Payload *models.GiftCertificate
}

// Code gets the status code for the post gift certificate default response
func (o *PostGiftCertificateDefault) Code() int {
	return o._statusCode
}

func (o *PostGiftCertificateDefault) Error() string {
	return fmt.Sprintf("[POST /gift_certificate][%d] postGiftCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *PostGiftCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GiftCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
