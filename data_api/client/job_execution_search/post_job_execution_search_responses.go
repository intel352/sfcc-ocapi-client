// Code generated by go-swagger; DO NOT EDIT.

package job_execution_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PostJobExecutionSearchReader is a Reader for the PostJobExecutionSearch structure.
type PostJobExecutionSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJobExecutionSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewPostJobExecutionSearchDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewPostJobExecutionSearchDefault creates a PostJobExecutionSearchDefault with default headers values
func NewPostJobExecutionSearchDefault(code int) *PostJobExecutionSearchDefault {
	return &PostJobExecutionSearchDefault{
		_statusCode: code,
	}
}

/*PostJobExecutionSearchDefault handles this case with default header values.

PostJobExecutionSearchDefault post job execution search default
*/
type PostJobExecutionSearchDefault struct {
	_statusCode int

	Payload *models.JobExecutionSearchResult
}

// Code gets the status code for the post job execution search default response
func (o *PostJobExecutionSearchDefault) Code() int {
	return o._statusCode
}

func (o *PostJobExecutionSearchDefault) Error() string {
	return fmt.Sprintf("[POST /job_execution_search][%d] postJobExecutionSearch default  %+v", o._statusCode, o.Payload)
}

func (o *PostJobExecutionSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobExecutionSearchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
