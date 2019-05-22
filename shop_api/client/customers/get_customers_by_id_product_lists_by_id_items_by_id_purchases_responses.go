// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// GetCustomersByIDProductListsByIDItemsByIDPurchasesReader is a Reader for the GetCustomersByIDProductListsByIDItemsByIDPurchases structure.
type GetCustomersByIDProductListsByIDItemsByIDPurchasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCustomersByIDProductListsByIDItemsByIDPurchasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest creates a GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest with default headers values
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest() *GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest {
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest{}
}

/*GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest handles this case with default header values.

Indicates that the customerId URL parameter does not match the verified customer
             represented by the JWT token, not relevant when using OAuth.
*/
type GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest struct {
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/product_lists/{list_id}/items/{item_id}/purchases][%d] getCustomersByIdProductListsByIdItemsByIdPurchasesBadRequest ", 400)
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound creates a GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound with default headers values
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound() *GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound {
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound{}
}

/*GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound handles this case with default header values.

Indicates that the customer with the given customer id is unknown for the site. or Indicates that the product list with the given list id is unknown for the
             site and the customer. or Indicates that the product list item with the given item id is unknown
             for the site, the customer and the product list.
*/
type GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound struct {
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/product_lists/{list_id}/items/{item_id}/purchases][%d] getCustomersByIdProductListsByIdItemsByIdPurchasesNotFound ", 404)
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomersByIDProductListsByIDItemsByIDPurchasesDefault creates a GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault with default headers values
func NewGetCustomersByIDProductListsByIDItemsByIDPurchasesDefault(code int) *GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault {
	return &GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault{
		_statusCode: code,
	}
}

/*GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault handles this case with default header values.

GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault get customers by ID product lists by ID items by ID purchases default
*/
type GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault struct {
	_statusCode int

	Payload *models.CustomerProductListItemPurchaseResult
}

// Code gets the status code for the get customers by ID product lists by ID items by ID purchases default response
func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault) Code() int {
	return o._statusCode
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault) Error() string {
	return fmt.Sprintf("[GET /customers/{customer_id}/product_lists/{list_id}/items/{item_id}/purchases][%d] getCustomersByIDProductListsByIDItemsByIDPurchases default  %+v", o._statusCode, o.Payload)
}

func (o *GetCustomersByIDProductListsByIDItemsByIDPurchasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerProductListItemPurchaseResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}