// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// PatchProductsByIDVariationGroupsByIDReader is a Reader for the PatchProductsByIDVariationGroupsByID structure.
type PatchProductsByIDVariationGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchProductsByIDVariationGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPatchProductsByIDVariationGroupsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchProductsByIDVariationGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchProductsByIDVariationGroupsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchProductsByIDVariationGroupsByIDBadRequest creates a PatchProductsByIDVariationGroupsByIDBadRequest with default headers values
func NewPatchProductsByIDVariationGroupsByIDBadRequest() *PatchProductsByIDVariationGroupsByIDBadRequest {
	return &PatchProductsByIDVariationGroupsByIDBadRequest{}
}

/*PatchProductsByIDVariationGroupsByIDBadRequest handles this case with default header values.

Indicates the master product id does not represent a master product. or Indicates product specified is not a variation group. or Indicates the values passed are not valid for the variation attributes.
*/
type PatchProductsByIDVariationGroupsByIDBadRequest struct {
}

func (o *PatchProductsByIDVariationGroupsByIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /products/{master_product_id}/variation_groups/{id}][%d] patchProductsByIdVariationGroupsByIdBadRequest ", 400)
}

func (o *PatchProductsByIDVariationGroupsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProductsByIDVariationGroupsByIDNotFound creates a PatchProductsByIDVariationGroupsByIDNotFound with default headers values
func NewPatchProductsByIDVariationGroupsByIDNotFound() *PatchProductsByIDVariationGroupsByIDNotFound {
	return &PatchProductsByIDVariationGroupsByIDNotFound{}
}

/*PatchProductsByIDVariationGroupsByIDNotFound handles this case with default header values.

Indicates either the master product or the variation group product cannot be found. or Indicates the product does not belong to the master product.
*/
type PatchProductsByIDVariationGroupsByIDNotFound struct {
}

func (o *PatchProductsByIDVariationGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /products/{master_product_id}/variation_groups/{id}][%d] patchProductsByIdVariationGroupsByIdNotFound ", 404)
}

func (o *PatchProductsByIDVariationGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchProductsByIDVariationGroupsByIDDefault creates a PatchProductsByIDVariationGroupsByIDDefault with default headers values
func NewPatchProductsByIDVariationGroupsByIDDefault(code int) *PatchProductsByIDVariationGroupsByIDDefault {
	return &PatchProductsByIDVariationGroupsByIDDefault{
		_statusCode: code,
	}
}

/*PatchProductsByIDVariationGroupsByIDDefault handles this case with default header values.

PatchProductsByIDVariationGroupsByIDDefault patch products by ID variation groups by ID default
*/
type PatchProductsByIDVariationGroupsByIDDefault struct {
	_statusCode int

	Payload *models.VariationGroup
}

// Code gets the status code for the patch products by ID variation groups by ID default response
func (o *PatchProductsByIDVariationGroupsByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchProductsByIDVariationGroupsByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /products/{master_product_id}/variation_groups/{id}][%d] patchProductsByIDVariationGroupsByID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchProductsByIDVariationGroupsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariationGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
