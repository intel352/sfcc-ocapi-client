// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProductsByIDVariationGroupsByIDReader is a Reader for the DeleteProductsByIDVariationGroupsByID structure.
type DeleteProductsByIDVariationGroupsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductsByIDVariationGroupsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProductsByIDVariationGroupsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteProductsByIDVariationGroupsByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteProductsByIDVariationGroupsByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProductsByIDVariationGroupsByIDNoContent creates a DeleteProductsByIDVariationGroupsByIDNoContent with default headers values
func NewDeleteProductsByIDVariationGroupsByIDNoContent() *DeleteProductsByIDVariationGroupsByIDNoContent {
	return &DeleteProductsByIDVariationGroupsByIDNoContent{}
}

/*DeleteProductsByIDVariationGroupsByIDNoContent handles this case with default header values.

DeleteProductsByIDVariationGroupsByIDNoContent delete products by Id variation groups by Id no content
*/
type DeleteProductsByIDVariationGroupsByIDNoContent struct {
}

func (o *DeleteProductsByIDVariationGroupsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /products/{master_product_id}/variation_groups/{id}][%d] deleteProductsByIdVariationGroupsByIdNoContent ", 204)
}

func (o *DeleteProductsByIDVariationGroupsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDVariationGroupsByIDBadRequest creates a DeleteProductsByIDVariationGroupsByIDBadRequest with default headers values
func NewDeleteProductsByIDVariationGroupsByIDBadRequest() *DeleteProductsByIDVariationGroupsByIDBadRequest {
	return &DeleteProductsByIDVariationGroupsByIDBadRequest{}
}

/*DeleteProductsByIDVariationGroupsByIDBadRequest handles this case with default header values.

Indicates the master product id does not represent a master product. or Indicates product specified is not a variation group.
*/
type DeleteProductsByIDVariationGroupsByIDBadRequest struct {
}

func (o *DeleteProductsByIDVariationGroupsByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /products/{master_product_id}/variation_groups/{id}][%d] deleteProductsByIdVariationGroupsByIdBadRequest ", 400)
}

func (o *DeleteProductsByIDVariationGroupsByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductsByIDVariationGroupsByIDNotFound creates a DeleteProductsByIDVariationGroupsByIDNotFound with default headers values
func NewDeleteProductsByIDVariationGroupsByIDNotFound() *DeleteProductsByIDVariationGroupsByIDNotFound {
	return &DeleteProductsByIDVariationGroupsByIDNotFound{}
}

/*DeleteProductsByIDVariationGroupsByIDNotFound handles this case with default header values.

Indicates either the master product or the variation group product cannot be found. or Indicates the product does not belong to the master product.
*/
type DeleteProductsByIDVariationGroupsByIDNotFound struct {
}

func (o *DeleteProductsByIDVariationGroupsByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /products/{master_product_id}/variation_groups/{id}][%d] deleteProductsByIdVariationGroupsByIdNotFound ", 404)
}

func (o *DeleteProductsByIDVariationGroupsByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
