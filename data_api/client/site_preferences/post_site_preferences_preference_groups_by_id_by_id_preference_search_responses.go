// Code generated by go-swagger; DO NOT EDIT.

package site_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchReader is a Reader for the PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearch structure.
type PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest creates a PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest with default headers values
func NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest() *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest {
	return &PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest{}
}

/*PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest handles this case with default header values.

Indicates the query is ill-formed.
*/
type PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest struct {
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /site_preferences/preference_groups/{group_id}/{instance_type}/preference_search][%d] postSitePreferencesPreferenceGroupsByIdByIdPreferenceSearchBadRequest ", 400)
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound creates a PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound with default headers values
func NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound() *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound {
	return &PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound{}
}

/*PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound handles this case with default header values.

Indicates the preference group is not found.
*/
type PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound struct {
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /site_preferences/preference_groups/{group_id}/{instance_type}/preference_search][%d] postSitePreferencesPreferenceGroupsByIdByIdPreferenceSearchNotFound ", 404)
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault creates a PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault with default headers values
func NewPostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault(code int) *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault {
	return &PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault{
		_statusCode: code,
	}
}

/*PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault handles this case with default header values.

PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault post site preferences preference groups by ID by ID preference search default
*/
type PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault struct {
	_statusCode int

	Payload *models.PreferenceValueSearchResult
}

// Code gets the status code for the post site preferences preference groups by ID by ID preference search default response
func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault) Error() string {
	return fmt.Sprintf("[POST /site_preferences/preference_groups/{group_id}/{instance_type}/preference_search][%d] postSitePreferencesPreferenceGroupsByIDByIDPreferenceSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostSitePreferencesPreferenceGroupsByIDByIDPreferenceSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PreferenceValueSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
