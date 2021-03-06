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

// PostSitesByIDGiftCertificatesReader is a Reader for the PostSitesByIDGiftCertificates structure.
type PostSitesByIDGiftCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesByIDGiftCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostSitesByIDGiftCertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSitesByIDGiftCertificatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSitesByIDGiftCertificatesBadRequest creates a PostSitesByIDGiftCertificatesBadRequest with default headers values
func NewPostSitesByIDGiftCertificatesBadRequest() *PostSitesByIDGiftCertificatesBadRequest {
	return &PostSitesByIDGiftCertificatesBadRequest{}
}

/*PostSitesByIDGiftCertificatesBadRequest handles this case with default header values.

if the gift certificate passed in is not valid (the
              argument indicates the field that was invalid). or If the amount specified is out of range or If merchant id is not unique or If recipient email address is invalid or If gift certificate status specified is invalid
*/
type PostSitesByIDGiftCertificatesBadRequest struct {
}

func (o *PostSitesByIDGiftCertificatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/gift_certificates][%d] postSitesByIdGiftCertificatesBadRequest ", 400)
}

func (o *PostSitesByIDGiftCertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitesByIDGiftCertificatesDefault creates a PostSitesByIDGiftCertificatesDefault with default headers values
func NewPostSitesByIDGiftCertificatesDefault(code int) *PostSitesByIDGiftCertificatesDefault {
	return &PostSitesByIDGiftCertificatesDefault{
		_statusCode: code,
	}
}

/*PostSitesByIDGiftCertificatesDefault handles this case with default header values.

PostSitesByIDGiftCertificatesDefault post sites by ID gift certificates default
*/
type PostSitesByIDGiftCertificatesDefault struct {
	_statusCode int

	Payload *models.GiftCertificate
}

// Code gets the status code for the post sites by ID gift certificates default response
func (o *PostSitesByIDGiftCertificatesDefault) Code() int {
	return o._statusCode
}

func (o *PostSitesByIDGiftCertificatesDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/gift_certificates][%d] postSitesByIDGiftCertificates default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitesByIDGiftCertificatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GiftCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
