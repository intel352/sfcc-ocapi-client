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

// GetProductsByIDVariationGroupsByIDReader is a Reader for the GetProductsByIDVariationGroupsByID structure.
type GetProductsByIDVariationGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDVariationGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetProductsByIDVariationGroupsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProductsByIDVariationGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductsByIDVariationGroupsByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsByIDVariationGroupsByIDBadRequest creates a GetProductsByIDVariationGroupsByIDBadRequest with default headers values
func NewGetProductsByIDVariationGroupsByIDBadRequest() *GetProductsByIDVariationGroupsByIDBadRequest {
	return &GetProductsByIDVariationGroupsByIDBadRequest{}
}

/*GetProductsByIDVariationGroupsByIDBadRequest handles this case with default header values.

Indicates the master product id does not represent a master product. or Indicates the product specified is not a variation group.
*/
type GetProductsByIDVariationGroupsByIDBadRequest struct {
}

func (o *GetProductsByIDVariationGroupsByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variation_groups/{id}][%d] getProductsByIdVariationGroupsByIdBadRequest ", 400)
}

func (o *GetProductsByIDVariationGroupsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDVariationGroupsByIDNotFound creates a GetProductsByIDVariationGroupsByIDNotFound with default headers values
func NewGetProductsByIDVariationGroupsByIDNotFound() *GetProductsByIDVariationGroupsByIDNotFound {
	return &GetProductsByIDVariationGroupsByIDNotFound{}
}

/*GetProductsByIDVariationGroupsByIDNotFound handles this case with default header values.

Indicates either the master product or the variation group product cannot be found. or Indicates the product does not belong to the master product.
*/
type GetProductsByIDVariationGroupsByIDNotFound struct {
}

func (o *GetProductsByIDVariationGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variation_groups/{id}][%d] getProductsByIdVariationGroupsByIdNotFound ", 404)
}

func (o *GetProductsByIDVariationGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDVariationGroupsByIDDefault creates a GetProductsByIDVariationGroupsByIDDefault with default headers values
func NewGetProductsByIDVariationGroupsByIDDefault(code int) *GetProductsByIDVariationGroupsByIDDefault {
	return &GetProductsByIDVariationGroupsByIDDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDVariationGroupsByIDDefault handles this case with default header values.

GetProductsByIDVariationGroupsByIDDefault get products by ID variation groups by ID default
*/
type GetProductsByIDVariationGroupsByIDDefault struct {
	_statusCode int

	Payload *models.VariationGroup
}

// Code gets the status code for the get products by ID variation groups by ID default response
func (o *GetProductsByIDVariationGroupsByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDVariationGroupsByIDDefault) Error() string {
	return fmt.Sprintf("[GET /products/{master_product_id}/variation_groups/{id}][%d] getProductsByIDVariationGroupsByID default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDVariationGroupsByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariationGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
