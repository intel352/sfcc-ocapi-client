// Code generated by go-swagger; DO NOT EDIT.

package inventory_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PatchInventoryListsByIDProductInventoryRecordsByIDReader is a Reader for the PatchInventoryListsByIDProductInventoryRecordsByID structure.
type PatchInventoryListsByIDProductInventoryRecordsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchInventoryListsByIDProductInventoryRecordsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewPatchInventoryListsByIDProductInventoryRecordsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchInventoryListsByIDProductInventoryRecordsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchInventoryListsByIDProductInventoryRecordsByIDNotFound creates a PatchInventoryListsByIDProductInventoryRecordsByIDNotFound with default headers values
func NewPatchInventoryListsByIDProductInventoryRecordsByIDNotFound() *PatchInventoryListsByIDProductInventoryRecordsByIDNotFound {
	return &PatchInventoryListsByIDProductInventoryRecordsByIDNotFound{}
}

/*PatchInventoryListsByIDProductInventoryRecordsByIDNotFound handles this case with default header values.

Indicates the product inventory record does not exist. or Indicates the inventory list does not exist.
*/
type PatchInventoryListsByIDProductInventoryRecordsByIDNotFound struct {
}

func (o *PatchInventoryListsByIDProductInventoryRecordsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /inventory_lists/{inventory_list_id}/product_inventory_records/{product_id}][%d] patchInventoryListsByIdProductInventoryRecordsByIdNotFound ", 404)
}

func (o *PatchInventoryListsByIDProductInventoryRecordsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchInventoryListsByIDProductInventoryRecordsByIDDefault creates a PatchInventoryListsByIDProductInventoryRecordsByIDDefault with default headers values
func NewPatchInventoryListsByIDProductInventoryRecordsByIDDefault(code int) *PatchInventoryListsByIDProductInventoryRecordsByIDDefault {
	return &PatchInventoryListsByIDProductInventoryRecordsByIDDefault{
		_statusCode: code,
	}
}

/*PatchInventoryListsByIDProductInventoryRecordsByIDDefault handles this case with default header values.

PatchInventoryListsByIDProductInventoryRecordsByIDDefault patch inventory lists by ID product inventory records by ID default
*/
type PatchInventoryListsByIDProductInventoryRecordsByIDDefault struct {
	_statusCode int

	Payload *models.ProductInventoryRecord
}

// Code gets the status code for the patch inventory lists by ID product inventory records by ID default response
func (o *PatchInventoryListsByIDProductInventoryRecordsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchInventoryListsByIDProductInventoryRecordsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /inventory_lists/{inventory_list_id}/product_inventory_records/{product_id}][%d] patchInventoryListsByIDProductInventoryRecordsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchInventoryListsByIDProductInventoryRecordsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductInventoryRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
