// Code generated by go-swagger; DO NOT EDIT.

package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PatchCatalogsByIDReader is a Reader for the PatchCatalogsByID structure.
type PatchCatalogsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCatalogsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchCatalogsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchCatalogsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCatalogsByIDNotFound creates a PatchCatalogsByIDNotFound with default headers values
func NewPatchCatalogsByIDNotFound() *PatchCatalogsByIDNotFound {
	return &PatchCatalogsByIDNotFound{}
}

/*PatchCatalogsByIDNotFound handles this case with default header values.

Thrown in case the catalog id is unknown.
*/
type PatchCatalogsByIDNotFound struct {
}

func (o *PatchCatalogsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}][%d] patchCatalogsByIdNotFound ", 404)
}

func (o *PatchCatalogsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCatalogsByIDDefault creates a PatchCatalogsByIDDefault with default headers values
func NewPatchCatalogsByIDDefault(code int) *PatchCatalogsByIDDefault {
	return &PatchCatalogsByIDDefault{
		_statusCode: code,
	}
}

/*PatchCatalogsByIDDefault handles this case with default header values.

PatchCatalogsByIDDefault patch catalogs by ID default
*/
type PatchCatalogsByIDDefault struct {
	_statusCode int

	Payload *models.Catalog
}

// Code gets the status code for the patch catalogs by ID default response
func (o *PatchCatalogsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCatalogsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /catalogs/{catalog_id}][%d] patchCatalogsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCatalogsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Catalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
