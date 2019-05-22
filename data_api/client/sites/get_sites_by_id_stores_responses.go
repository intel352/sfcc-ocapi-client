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

// GetSitesByIDStoresReader is a Reader for the GetSitesByIDStores structure.
type GetSitesByIDStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetSitesByIDStoresDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetSitesByIDStoresDefault creates a GetSitesByIDStoresDefault with default headers values
func NewGetSitesByIDStoresDefault(code int) *GetSitesByIDStoresDefault {
	return &GetSitesByIDStoresDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDStoresDefault handles this case with default header values.

GetSitesByIDStoresDefault get sites by ID stores default
*/
type GetSitesByIDStoresDefault struct {
	_statusCode int

	Payload *models.Stores
}

// Code gets the status code for the get sites by ID stores default response
func (o *GetSitesByIDStoresDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDStoresDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/stores][%d] getSitesByIDStores default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDStoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stores)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}