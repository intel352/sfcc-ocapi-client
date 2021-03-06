// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/meta_api/models"
)

// GetAppsByIDUIByIDReader is a Reader for the GetAppsByIDUIByID structure.
type GetAppsByIDUIByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsByIDUIByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewGetAppsByIDUIByIDDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewGetAppsByIDUIByIDDefault creates a GetAppsByIDUIByIDDefault with default headers values
func NewGetAppsByIDUIByIDDefault(code int) *GetAppsByIDUIByIDDefault {
	return &GetAppsByIDUIByIDDefault{
		_statusCode: code,
	}
}

/*GetAppsByIDUIByIDDefault handles this case with default header values.

GetAppsByIDUIByIDDefault get apps by ID Ui by ID default
*/
type GetAppsByIDUIByIDDefault struct {
	_statusCode int

	Payload *models.Uicustomization
}

// Code gets the status code for the get apps by ID Ui by ID default response
func (o *GetAppsByIDUIByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAppsByIDUIByIDDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{application}/ui/{template}][%d] getAppsByIDUiByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppsByIDUIByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Uicustomization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
