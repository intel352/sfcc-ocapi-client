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

// PostSitesByIDSlotSearchReader is a Reader for the PostSitesByIDSlotSearch structure.
type PostSitesByIDSlotSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitesByIDSlotSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostSitesByIDSlotSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSitesByIDSlotSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSitesByIDSlotSearchBadRequest creates a PostSitesByIDSlotSearchBadRequest with default headers values
func NewPostSitesByIDSlotSearchBadRequest() *PostSitesByIDSlotSearchBadRequest {
	return &PostSitesByIDSlotSearchBadRequest{}
}

/*PostSitesByIDSlotSearchBadRequest handles this case with default header values.

Thrown if the query is ill-formed.
*/
type PostSitesByIDSlotSearchBadRequest struct {
}

func (o *PostSitesByIDSlotSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/slot_search][%d] postSitesByIdSlotSearchBadRequest ", 400)
}

func (o *PostSitesByIDSlotSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitesByIDSlotSearchDefault creates a PostSitesByIDSlotSearchDefault with default headers values
func NewPostSitesByIDSlotSearchDefault(code int) *PostSitesByIDSlotSearchDefault {
	return &PostSitesByIDSlotSearchDefault{
		_statusCode: code,
	}
}

/*PostSitesByIDSlotSearchDefault handles this case with default header values.

PostSitesByIDSlotSearchDefault post sites by ID slot search default
*/
type PostSitesByIDSlotSearchDefault struct {
	_statusCode int

	Payload *models.SlotSearchResult
}

// Code gets the status code for the post sites by ID slot search default response
func (o *PostSitesByIDSlotSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostSitesByIDSlotSearchDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/slot_search][%d] postSitesByIDSlotSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitesByIDSlotSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SlotSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}