// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSitesByIDGiftCertificatesByIDReader is a Reader for the DeleteSitesByIDGiftCertificatesByID structure.
type DeleteSitesByIDGiftCertificatesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSitesByIDGiftCertificatesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSitesByIDGiftCertificatesByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSitesByIDGiftCertificatesByIDNoContent creates a DeleteSitesByIDGiftCertificatesByIDNoContent with default headers values
func NewDeleteSitesByIDGiftCertificatesByIDNoContent() *DeleteSitesByIDGiftCertificatesByIDNoContent {
	return &DeleteSitesByIDGiftCertificatesByIDNoContent{}
}

/*DeleteSitesByIDGiftCertificatesByIDNoContent handles this case with default header values.

DeleteSitesByIDGiftCertificatesByIDNoContent delete sites by Id gift certificates by Id no content
*/
type DeleteSitesByIDGiftCertificatesByIDNoContent struct {
}

func (o *DeleteSitesByIDGiftCertificatesByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/gift_certificates/{merchant_id}][%d] deleteSitesByIdGiftCertificatesByIdNoContent ", 204)
}

func (o *DeleteSitesByIDGiftCertificatesByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
