// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// GetSitesByIDAiRecommenderNamesReader is a Reader for the GetSitesByIDAiRecommenderNames structure.
type GetSitesByIDAiRecommenderNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesByIDAiRecommenderNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 400:
		result := NewGetSitesByIDAiRecommenderNamesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSitesByIDAiRecommenderNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSitesByIDAiRecommenderNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesByIDAiRecommenderNamesBadRequest creates a GetSitesByIDAiRecommenderNamesBadRequest with default headers values
func NewGetSitesByIDAiRecommenderNamesBadRequest() *GetSitesByIDAiRecommenderNamesBadRequest {
	return &GetSitesByIDAiRecommenderNamesBadRequest{}
}

/*GetSitesByIDAiRecommenderNamesBadRequest handles this case with default header values.

If CQuotient responds with a non-success status code.
*/
type GetSitesByIDAiRecommenderNamesBadRequest struct {
}

func (o *GetSitesByIDAiRecommenderNamesBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ai/recommender_names][%d] getSitesByIdAiRecommenderNamesBadRequest ", 400)
}

func (o *GetSitesByIDAiRecommenderNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDAiRecommenderNamesNotFound creates a GetSitesByIDAiRecommenderNamesNotFound with default headers values
func NewGetSitesByIDAiRecommenderNamesNotFound() *GetSitesByIDAiRecommenderNamesNotFound {
	return &GetSitesByIDAiRecommenderNamesNotFound{}
}

/*GetSitesByIDAiRecommenderNamesNotFound handles this case with default header values.

If the given site does not exist.
*/
type GetSitesByIDAiRecommenderNamesNotFound struct {
}

func (o *GetSitesByIDAiRecommenderNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ai/recommender_names][%d] getSitesByIdAiRecommenderNamesNotFound ", 404)
}

func (o *GetSitesByIDAiRecommenderNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSitesByIDAiRecommenderNamesDefault creates a GetSitesByIDAiRecommenderNamesDefault with default headers values
func NewGetSitesByIDAiRecommenderNamesDefault(code int) *GetSitesByIDAiRecommenderNamesDefault {
	return &GetSitesByIDAiRecommenderNamesDefault{
		_statusCode: code,
	}
}

/*GetSitesByIDAiRecommenderNamesDefault handles this case with default header values.

GetSitesByIDAiRecommenderNamesDefault get sites by ID ai recommender names default
*/
type GetSitesByIDAiRecommenderNamesDefault struct {
	_statusCode int

	Payload *models.RecommendersResult
}

// Code gets the status code for the get sites by ID ai recommender names default response
func (o *GetSitesByIDAiRecommenderNamesDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesByIDAiRecommenderNamesDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ai/recommender_names][%d] getSitesByIDAiRecommenderNames default  %+v", o._statusCode, o.Payload)
}

func (o *GetSitesByIDAiRecommenderNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecommendersResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
