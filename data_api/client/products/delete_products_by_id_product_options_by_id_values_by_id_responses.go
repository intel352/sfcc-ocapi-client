// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProductsByIDProductOptionsByIDValuesByIDReader is a Reader for the DeleteProductsByIDProductOptionsByIDValuesByID structure.
type DeleteProductsByIDProductOptionsByIDValuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductsByIDProductOptionsByIDValuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProductsByIDProductOptionsByIDValuesByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteProductsByIDProductOptionsByIDValuesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductsByIDProductOptionsByIDValuesByIDNoContent creates a DeleteProductsByIDProductOptionsByIDValuesByIDNoContent with default headers values
func NewDeleteProductsByIDProductOptionsByIDValuesByIDNoContent() *DeleteProductsByIDProductOptionsByIDValuesByIDNoContent {
	return &DeleteProductsByIDProductOptionsByIDValuesByIDNoContent{}
}

/*DeleteProductsByIDProductOptionsByIDValuesByIDNoContent handles this case with default header values.

DeleteProductsByIDProductOptionsByIDValuesByIDNoContent delete products by Id product options by Id values by Id no content
*/
type DeleteProductsByIDProductOptionsByIDValuesByIDNoContent struct {
}

func (o *DeleteProductsByIDProductOptionsByIDValuesByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/product_options/{option_id}/values/{id}][%d] deleteProductsByIdProductOptionsByIdValuesByIdNoContent ", 204)
}

func (o *DeleteProductsByIDProductOptionsByIDValuesByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDProductOptionsByIDValuesByIDNotFound creates a DeleteProductsByIDProductOptionsByIDValuesByIDNotFound with default headers values
func NewDeleteProductsByIDProductOptionsByIDValuesByIDNotFound() *DeleteProductsByIDProductOptionsByIDValuesByIDNotFound {
	return &DeleteProductsByIDProductOptionsByIDValuesByIDNotFound{}
}

/*DeleteProductsByIDProductOptionsByIDValuesByIDNotFound handles this case with default header values.

Indicates the product is not found. or Indicates the local product option is not found. or Indicates the local product option value is not found.
*/
type DeleteProductsByIDProductOptionsByIDValuesByIDNotFound struct {
}

func (o *DeleteProductsByIDProductOptionsByIDValuesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/product_options/{option_id}/values/{id}][%d] deleteProductsByIdProductOptionsByIdValuesByIdNotFound ", 404)
}

func (o *DeleteProductsByIDProductOptionsByIDValuesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
