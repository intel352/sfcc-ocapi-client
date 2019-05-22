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

// GetSitesByIDSourceCodeGroupsReader is a Reader for the GetSitesByIDSourceCodeGroups structure.
type GetSitesByIDSourceCodeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDSourceCodeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetSitesByIDSourceCodeGroupsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetSitesByIDSourceCodeGroupsDefault creates a GetSitesByIDSourceCodeGroupsDefault with default headers values
func NewGetSitesByIDSourceCodeGroupsDefault(code int) *GetSitesByIDSourceCodeGroupsDefault {
	return &GetSitesByIDSourceCodeGroupsDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDSourceCodeGroupsDefault handles this case with default header values.

GetSitesByIDSourceCodeGroupsDefault get sites by ID source code groups default
*/
type GetSitesByIDSourceCodeGroupsDefault struct {
	_statusCode int

	Payload *models.SourceCodeGroups
}

// Code gets the status code for the get sites by ID source code groups default response
func (o *GetSitesByIDSourceCodeGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDSourceCodeGroupsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/source_code_groups][%d] getSitesByIDSourceCodeGroups default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDSourceCodeGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceCodeGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}