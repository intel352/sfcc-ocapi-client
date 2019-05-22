// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProductsByIDVariationAttributesByIDReader is a Reader for the DeleteProductsByIDVariationAttributesByID structure.
type DeleteProductsByIDVariationAttributesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductsByIDVariationAttributesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProductsByIDVariationAttributesByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteProductsByIDVariationAttributesByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteProductsByIDVariationAttributesByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductsByIDVariationAttributesByIDNoContent creates a DeleteProductsByIDVariationAttributesByIDNoContent with default headers values
func NewDeleteProductsByIDVariationAttributesByIDNoContent() *DeleteProductsByIDVariationAttributesByIDNoContent {
	return &DeleteProductsByIDVariationAttributesByIDNoContent{}
}

/*DeleteProductsByIDVariationAttributesByIDNoContent handles this case with default header values.

DeleteProductsByIDVariationAttributesByIDNoContent delete products by Id variation attributes by Id no content
*/
type DeleteProductsByIDVariationAttributesByIDNoContent struct {
}

func (o *DeleteProductsByIDVariationAttributesByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/variation_attributes/{id}][%d] deleteProductsByIdVariationAttributesByIdNoContent ", 204)
}

func (o *DeleteProductsByIDVariationAttributesByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDVariationAttributesByIDBadRequest creates a DeleteProductsByIDVariationAttributesByIDBadRequest with default headers values
func NewDeleteProductsByIDVariationAttributesByIDBadRequest() *DeleteProductsByIDVariationAttributesByIDBadRequest {
	return &DeleteProductsByIDVariationAttributesByIDBadRequest{}
}

/*DeleteProductsByIDVariationAttributesByIDBadRequest handles this case with default header values.

Indicates the master product is not found.
*/
type DeleteProductsByIDVariationAttributesByIDBadRequest struct {
}

func (o *DeleteProductsByIDVariationAttributesByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/variation_attributes/{id}][%d] deleteProductsByIdVariationAttributesByIdBadRequest ", 400)
}

func (o *DeleteProductsByIDVariationAttributesByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDVariationAttributesByIDNotFound creates a DeleteProductsByIDVariationAttributesByIDNotFound with default headers values
func NewDeleteProductsByIDVariationAttributesByIDNotFound() *DeleteProductsByIDVariationAttributesByIDNotFound {
	return &DeleteProductsByIDVariationAttributesByIDNotFound{}
}

/*DeleteProductsByIDVariationAttributesByIDNotFound handles this case with default header values.

Indicates that the variation attribute is not found. or Indicates the product is not found.
*/
type DeleteProductsByIDVariationAttributesByIDNotFound struct {
}

func (o *DeleteProductsByIDVariationAttributesByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{product_id}/variation_attributes/{id}][%d] deleteProductsByIdVariationAttributesByIdNotFound ", 404)
}

func (o *DeleteProductsByIDVariationAttributesByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}