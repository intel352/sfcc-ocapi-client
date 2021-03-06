// Code generated by go-swagger; DO NOT EDIT.

package customer_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostCustomerListsByIDCustomerSearchReader is a Reader for the PostCustomerListsByIDCustomerSearch structure.
type PostCustomerListsByIDCustomerSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCustomerListsByIDCustomerSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostCustomerListsByIDCustomerSearchDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostCustomerListsByIDCustomerSearchDefault creates a PostCustomerListsByIDCustomerSearchDefault with default headers values
func NewPostCustomerListsByIDCustomerSearchDefault(code int) *PostCustomerListsByIDCustomerSearchDefault {
	return &PostCustomerListsByIDCustomerSearchDefault{
		_statusCode: code,
	}
}

/*PostCustomerListsByIDCustomerSearchDefault handles this case with default header values.

PostCustomerListsByIDCustomerSearchDefault post customer lists by ID customer search default
*/
type PostCustomerListsByIDCustomerSearchDefault struct {
	_statusCode int

	Payload *models.CustomerSearchResult
}

// Code gets the status code for the post customer lists by ID customer search default response
func (o *PostCustomerListsByIDCustomerSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostCustomerListsByIDCustomerSearchDefault) Error() string {
	return fmt.Sprintf("[POST /customer_lists/{customer_list_id}/customer_search][%d] postCustomerListsByIDCustomerSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostCustomerListsByIDCustomerSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
