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

// GetInventoryListsByIDProductInventoryRecordsByIDReader is a Reader for the GetInventoryListsByIDProductInventoryRecordsByID structure.
type GetInventoryListsByIDProductInventoryRecordsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryListsByIDProductInventoryRecordsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 404:
		result := NewGetInventoryListsByIDProductInventoryRecordsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInventoryListsByIDProductInventoryRecordsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInventoryListsByIDProductInventoryRecordsByIDNotFound creates a GetInventoryListsByIDProductInventoryRecordsByIDNotFound with default headers values
func NewGetInventoryListsByIDProductInventoryRecordsByIDNotFound() *GetInventoryListsByIDProductInventoryRecordsByIDNotFound {
	return &GetInventoryListsByIDProductInventoryRecordsByIDNotFound{}
}

/*GetInventoryListsByIDProductInventoryRecordsByIDNotFound handles this case with default header values.

Indicates the product inventory record does not exist. or Indicates the inventory list does not exist.
*/
type GetInventoryListsByIDProductInventoryRecordsByIDNotFound struct {
}

func (o *GetInventoryListsByIDProductInventoryRecordsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /inventory_lists/{inventory_list_id}/product_inventory_records/{product_id}][%d] getInventoryListsByIdProductInventoryRecordsByIdNotFound ", 404)
}

func (o *GetInventoryListsByIDProductInventoryRecordsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInventoryListsByIDProductInventoryRecordsByIDDefault creates a GetInventoryListsByIDProductInventoryRecordsByIDDefault with default headers values
func NewGetInventoryListsByIDProductInventoryRecordsByIDDefault(code int) *GetInventoryListsByIDProductInventoryRecordsByIDDefault {
	return &GetInventoryListsByIDProductInventoryRecordsByIDDefault{
		_statusCode: code,
	}
}

/*GetInventoryListsByIDProductInventoryRecordsByIDDefault handles this case with default header values.

GetInventoryListsByIDProductInventoryRecordsByIDDefault get inventory lists by ID product inventory records by ID default
*/
type GetInventoryListsByIDProductInventoryRecordsByIDDefault struct {
	_statusCode int

	Payload *models.ProductInventoryRecord
}

// Code gets the status code for the get inventory lists by ID product inventory records by ID default response
func (o *GetInventoryListsByIDProductInventoryRecordsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetInventoryListsByIDProductInventoryRecordsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /inventory_lists/{inventory_list_id}/product_inventory_records/{product_id}][%d] getInventoryListsByIDProductInventoryRecordsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetInventoryListsByIDProductInventoryRecordsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductInventoryRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
