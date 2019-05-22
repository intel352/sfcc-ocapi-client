// Code generated by go-swagger; DO NOT EDIT.

package ocapi_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteOcapiConfigsByIDReader is a Reader for the DeleteOcapiConfigsByID structure.
type DeleteOcapiConfigsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOcapiConfigsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOcapiConfigsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOcapiConfigsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOcapiConfigsByIDNoContent creates a DeleteOcapiConfigsByIDNoContent with default headers values
func NewDeleteOcapiConfigsByIDNoContent() *DeleteOcapiConfigsByIDNoContent {
	return &DeleteOcapiConfigsByIDNoContent{}
}

/*DeleteOcapiConfigsByIDNoContent handles this case with default header values.

DeleteOcapiConfigsByIDNoContent delete ocapi configs by Id no content
*/
type DeleteOcapiConfigsByIDNoContent struct {
}

func (o *DeleteOcapiConfigsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ocapi_configs/{clientId}][%d] deleteOcapiConfigsByIdNoContent ", 204)
}

func (o *DeleteOcapiConfigsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOcapiConfigsByIDBadRequest creates a DeleteOcapiConfigsByIDBadRequest with default headers values
func NewDeleteOcapiConfigsByIDBadRequest() *DeleteOcapiConfigsByIDBadRequest {
	return &DeleteOcapiConfigsByIDBadRequest{}
}

/*DeleteOcapiConfigsByIDBadRequest handles this case with default header values.

Indicates that the resulting config is not valid or Write operation on self is not allowed. (The clientId being used for calling this API should be different from target clientId)
*/
type DeleteOcapiConfigsByIDBadRequest struct {
}

func (o *DeleteOcapiConfigsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ocapi_configs/{clientId}][%d] deleteOcapiConfigsByIdBadRequest ", 400)
}

func (o *DeleteOcapiConfigsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
