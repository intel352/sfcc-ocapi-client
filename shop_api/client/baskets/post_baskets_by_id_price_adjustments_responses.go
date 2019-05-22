// Code generated by go-swagger; DO NOT EDIT.

package baskets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// PostBasketsByIDPriceAdjustmentsReader is a Reader for the PostBasketsByIDPriceAdjustments structure.
type PostBasketsByIDPriceAdjustmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostBasketsByIDPriceAdjustmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewPostBasketsByIDPriceAdjustmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostBasketsByIDPriceAdjustmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostBasketsByIDPriceAdjustmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostBasketsByIDPriceAdjustmentsBadRequest creates a PostBasketsByIDPriceAdjustmentsBadRequest with default headers values
func NewPostBasketsByIDPriceAdjustmentsBadRequest() *PostBasketsByIDPriceAdjustmentsBadRequest {
	return &PostBasketsByIDPriceAdjustmentsBadRequest{}
}

/*PostBasketsByIDPriceAdjustmentsBadRequest handles this case with default header values.

Indicates that a fixed price adjustment was already created for the given level. or Indicates that a fixed price adjustment was added at order level. or Indicates that a promotion id was used twice. or Indicates that a system promotion id was used as a manual promotion id. or Indicates that more than one hundred price adjustments would have been created. or Indicates that the price adjustment limit is exceeded.
*/
type PostBasketsByIDPriceAdjustmentsBadRequest struct {
}

func (o *PostBasketsByIDPriceAdjustmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/price_adjustments][%d] postBasketsByIdPriceAdjustmentsBadRequest ", 400)
}

func (o *PostBasketsByIDPriceAdjustmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBasketsByIDPriceAdjustmentsNotFound creates a PostBasketsByIDPriceAdjustmentsNotFound with default headers values
func NewPostBasketsByIDPriceAdjustmentsNotFound() *PostBasketsByIDPriceAdjustmentsNotFound {
	return &PostBasketsByIDPriceAdjustmentsNotFound{}
}

/*PostBasketsByIDPriceAdjustmentsNotFound handles this case with default header values.

Indicates that the basket with the given basket id is
             unknown.
*/
type PostBasketsByIDPriceAdjustmentsNotFound struct {
}

func (o *PostBasketsByIDPriceAdjustmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/price_adjustments][%d] postBasketsByIdPriceAdjustmentsNotFound ", 404)
}

func (o *PostBasketsByIDPriceAdjustmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostBasketsByIDPriceAdjustmentsDefault creates a PostBasketsByIDPriceAdjustmentsDefault with default headers values
func NewPostBasketsByIDPriceAdjustmentsDefault(code int) *PostBasketsByIDPriceAdjustmentsDefault {
	return &PostBasketsByIDPriceAdjustmentsDefault{
		_statusCode: code,
	}
}

/*PostBasketsByIDPriceAdjustmentsDefault handles this case with default header values.

PostBasketsByIDPriceAdjustmentsDefault post baskets by ID price adjustments default
*/
type PostBasketsByIDPriceAdjustmentsDefault struct {
	_statusCode int

	Payload *models.Basket
}

// Code gets the status code for the post baskets by ID price adjustments default response
func (o *PostBasketsByIDPriceAdjustmentsDefault) Code() int {
	return o._statusCode
}

func (o *PostBasketsByIDPriceAdjustmentsDefault) Error() string {
	return fmt.Sprintf("[POST /baskets/{basket_id}/price_adjustments][%d] postBasketsByIDPriceAdjustments default  %+v", o._statusCode, o.Payload)
}

func (o *PostBasketsByIDPriceAdjustmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Basket)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}