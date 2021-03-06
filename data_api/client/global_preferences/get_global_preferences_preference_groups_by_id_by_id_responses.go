// Code generated by go-swagger; DO NOT EDIT.

package global_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetGlobalPreferencesPreferenceGroupsByIDByIDReader is a Reader for the GetGlobalPreferencesPreferenceGroupsByIDByID structure.
type GetGlobalPreferencesPreferenceGroupsByIDByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetGlobalPreferencesPreferenceGroupsByIDByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetGlobalPreferencesPreferenceGroupsByIDByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGlobalPreferencesPreferenceGroupsByIDByIDNotFound creates a GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound with default headers values
func NewGetGlobalPreferencesPreferenceGroupsByIDByIDNotFound() *GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound {
	return &GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound{}
}

/*GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound handles this case with default header values.

Indicates the preference group is not found.
*/
type GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound struct {
}

func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /global_preferences/preference_groups/{group_id}/{instance_type}][%d] getGlobalPreferencesPreferenceGroupsByIdByIdNotFound ", 404)
}

func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGlobalPreferencesPreferenceGroupsByIDByIDDefault creates a GetGlobalPreferencesPreferenceGroupsByIDByIDDefault with default headers values
func NewGetGlobalPreferencesPreferenceGroupsByIDByIDDefault(code int) *GetGlobalPreferencesPreferenceGroupsByIDByIDDefault {
	return &GetGlobalPreferencesPreferenceGroupsByIDByIDDefault{
		_statusCode: code,
	}
}

/*GetGlobalPreferencesPreferenceGroupsByIDByIDDefault handles this case with default header values.

GetGlobalPreferencesPreferenceGroupsByIDByIDDefault get global preferences preference groups by ID by ID default
*/
type GetGlobalPreferencesPreferenceGroupsByIDByIDDefault struct {
	_statusCode int

	Payload *models.OrganizationPreferences
}

// Code gets the status code for the get global preferences preference groups by ID by ID default response
func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDDefault) Error() string {
	return fmt.Sprintf("[GET /global_preferences/preference_groups/{group_id}/{instance_type}][%d] getGlobalPreferencesPreferenceGroupsByIDByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetGlobalPreferencesPreferenceGroupsByIDByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationPreferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
