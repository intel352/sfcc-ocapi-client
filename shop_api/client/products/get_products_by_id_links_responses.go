// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetProductsByIDLinksReader is a Reader for the GetProductsByIDLinks structure.
type GetProductsByIDLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductsByIDLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetProductsByIDLinksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetProductsByIDLinksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProductsByIDLinksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProductsByIDLinksBadRequest creates a GetProductsByIDLinksBadRequest with default headers values
func NewGetProductsByIDLinksBadRequest() *GetProductsByIDLinksBadRequest {
	return &GetProductsByIDLinksBadRequest{}
}

/*GetProductsByIDLinksBadRequest handles this case with default header values.

400 indicates unknown product link type code or 400 indicates unknown link direction
*/
type GetProductsByIDLinksBadRequest struct {
}

func (o *GetProductsByIDLinksBadRequest) Error() string {
	return fmt.Sprintf("[GET /products/{id}/links][%d] getProductsByIdLinksBadRequest ", 400)
}

func (o *GetProductsByIDLinksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDLinksNotFound creates a GetProductsByIDLinksNotFound with default headers values
func NewGetProductsByIDLinksNotFound() *GetProductsByIDLinksNotFound {
	return &GetProductsByIDLinksNotFound{}
}

/*GetProductsByIDLinksNotFound handles this case with default header values.

404 No product with given id found
*/
type GetProductsByIDLinksNotFound struct {
}

func (o *GetProductsByIDLinksNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}/links][%d] getProductsByIdLinksNotFound ", 404)
}

func (o *GetProductsByIDLinksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductsByIDLinksDefault creates a GetProductsByIDLinksDefault with default headers values
func NewGetProductsByIDLinksDefault(code int) *GetProductsByIDLinksDefault {
	return &GetProductsByIDLinksDefault{
		_statusCode: code,
	}
}

/*GetProductsByIDLinksDefault handles this case with default header values.

GetProductsByIDLinksDefault get products by ID links default
*/
type GetProductsByIDLinksDefault struct {
	_statusCode int

	Payload *models.Product
}

// Code gets the status code for the get products by ID links default response
func (o *GetProductsByIDLinksDefault) Code() int {
	return o._statusCode
}

func (o *GetProductsByIDLinksDefault) Error() string {
	return fmt.Sprintf("[GET /products/{id}/links][%d] getProductsByIDLinks default  %+v", o._statusCode, o.Payload)
}

func (o *GetProductsByIDLinksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
