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

// PostSitesByIDCartridgesReader is a Reader for the PostSitesByIDCartridges structure.
type PostSitesByIDCartridgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesByIDCartridgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostSitesByIDCartridgesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostSitesByIDCartridgesDefault creates a PostSitesByIDCartridgesDefault with default headers values
func NewPostSitesByIDCartridgesDefault(code int) *PostSitesByIDCartridgesDefault {
	return &PostSitesByIDCartridgesDefault{
		_statusCode: code,
	}
}

/*PostSitesByIDCartridgesDefault handles this case with default header values.

PostSitesByIDCartridgesDefault post sites by ID cartridges default
*/
type PostSitesByIDCartridgesDefault struct {
	_statusCode int

	Payload *models.CartridgePathAPIResponse
}

// Code gets the status code for the post sites by ID cartridges default response
func (o *PostSitesByIDCartridgesDefault) Code() int {
	return o._statusCode
}

func (o *PostSitesByIDCartridgesDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/cartridges][%d] postSitesByIDCartridges default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitesByIDCartridgesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CartridgePathAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
