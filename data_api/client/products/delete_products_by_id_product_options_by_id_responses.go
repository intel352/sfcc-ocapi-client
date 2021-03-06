// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProductsByIDProductOptionsByIDReader is a Reader for the DeleteProductsByIDProductOptionsByID structure.
type DeleteProductsByIDProductOptionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductsByIDProductOptionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProductsByIDProductOptionsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteProductsByIDProductOptionsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductsByIDProductOptionsByIDNoContent creates a DeleteProductsByIDProductOptionsByIDNoContent with default headers values
func NewDeleteProductsByIDProductOptionsByIDNoContent() *DeleteProductsByIDProductOptionsByIDNoContent {
	return &DeleteProductsByIDProductOptionsByIDNoContent{}
}

/*DeleteProductsByIDProductOptionsByIDNoContent handles this case with default header values.

DeleteProductsByIDProductOptionsByIDNoContent delete products by Id product options by Id no content
*/
type DeleteProductsByIDProductOptionsByIDNoContent struct {
}

func (o *DeleteProductsByIDProductOptionsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/product_options/{id}][%d] deleteProductsByIdProductOptionsByIdNoContent ", 204)
}

func (o *DeleteProductsByIDProductOptionsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDProductOptionsByIDNotFound creates a DeleteProductsByIDProductOptionsByIDNotFound with default headers values
func NewDeleteProductsByIDProductOptionsByIDNotFound() *DeleteProductsByIDProductOptionsByIDNotFound {
	return &DeleteProductsByIDProductOptionsByIDNotFound{}
}

/*DeleteProductsByIDProductOptionsByIDNotFound handles this case with default header values.

Indicates the local or shared product option is not found. or Indicates the product is not found.
*/
type DeleteProductsByIDProductOptionsByIDNotFound struct {
}

func (o *DeleteProductsByIDProductOptionsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/product_options/{id}][%d] deleteProductsByIdProductOptionsByIdNotFound ", 404)
}

func (o *DeleteProductsByIDProductOptionsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
