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

// PutSitesByIDAbTestsByIDReader is a Reader for the PutSitesByIDAbTestsByID structure.
type PutSitesByIDAbTestsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSitesByIDAbTestsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPutSitesByIDAbTestsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutSitesByIDAbTestsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPutSitesByIDAbTestsByIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutSitesByIDAbTestsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutSitesByIDAbTestsByIDBadRequest creates a PutSitesByIDAbTestsByIDBadRequest with default headers values
func NewPutSitesByIDAbTestsByIDBadRequest() *PutSitesByIDAbTestsByIDBadRequest {
	return &PutSitesByIDAbTestsByIDBadRequest{}
}

/*PutSitesByIDAbTestsByIDBadRequest handles this case with default header values.

Thrown if an invalid email is specified. or Thrown if an invalid pipeline is specified. or Thrown if there is missing pipline_call or categories in the specified trigger of type pipeline_call or category_view_page respectively.
*/
type PutSitesByIDAbTestsByIDBadRequest struct {
}

func (o *PutSitesByIDAbTestsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/ab_tests/{id}][%d] putSitesByIdAbTestsByIdBadRequest ", 400)
}

func (o *PutSitesByIDAbTestsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDAbTestsByIDNotFound creates a PutSitesByIDAbTestsByIDNotFound with default headers values
func NewPutSitesByIDAbTestsByIDNotFound() *PutSitesByIDAbTestsByIDNotFound {
	return &PutSitesByIDAbTestsByIDNotFound{}
}

/*PutSitesByIDAbTestsByIDNotFound handles this case with default header values.

Indicates that site specified with the given id is unknown. or Thrown in case the A/B Test does not exist matching the given id
*/
type PutSitesByIDAbTestsByIDNotFound struct {
}

func (o *PutSitesByIDAbTestsByIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/ab_tests/{id}][%d] putSitesByIdAbTestsByIdNotFound ", 404)
}

func (o *PutSitesByIDAbTestsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDAbTestsByIDConflict creates a PutSitesByIDAbTestsByIDConflict with default headers values
func NewPutSitesByIDAbTestsByIDConflict() *PutSitesByIDAbTestsByIDConflict {
	return &PutSitesByIDAbTestsByIDConflict{}
}

/*PutSitesByIDAbTestsByIDConflict handles this case with default header values.

If an A/B Test exists already in the site with the given identifier and the header
x-dw-validate-existing=true is passed in with the request.
*/
type PutSitesByIDAbTestsByIDConflict struct {
}

func (o *PutSitesByIDAbTestsByIDConflict) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/ab_tests/{id}][%d] putSitesByIdAbTestsByIdConflict ", 409)
}

func (o *PutSitesByIDAbTestsByIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutSitesByIDAbTestsByIDDefault creates a PutSitesByIDAbTestsByIDDefault with default headers values
func NewPutSitesByIDAbTestsByIDDefault(code int) *PutSitesByIDAbTestsByIDDefault {
	return &PutSitesByIDAbTestsByIDDefault{
		_statusCode: code,
	}
}

/*PutSitesByIDAbTestsByIDDefault handles this case with default header values.

PutSitesByIDAbTestsByIDDefault put sites by ID ab tests by ID default
*/
type PutSitesByIDAbTestsByIDDefault struct {
	_statusCode int

	Payload *models.AbTest
}

// Code gets the status code for the put sites by ID ab tests by ID default response
func (o *PutSitesByIDAbTestsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PutSitesByIDAbTestsByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/ab_tests/{id}][%d] putSitesByIDAbTestsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PutSitesByIDAbTestsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AbTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}