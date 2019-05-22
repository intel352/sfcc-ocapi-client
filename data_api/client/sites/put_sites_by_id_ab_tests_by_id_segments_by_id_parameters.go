// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPutSitesByIDAbTestsByIDSegmentsByIDParams creates a new PutSitesByIDAbTestsByIDSegmentsByIDParams object
// with the default values initialized.
func NewPutSitesByIDAbTestsByIDSegmentsByIDParams() *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithTimeout creates a new PutSitesByIDAbTestsByIDSegmentsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithContext creates a new PutSitesByIDAbTestsByIDSegmentsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithContext(ctx context.Context) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithHTTPClient creates a new PutSitesByIDAbTestsByIDSegmentsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDAbTestsByIDSegmentsByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID ab tests by ID segments by ID operation typically these are written to a http.Request
*/
type PutSitesByIDAbTestsByIDSegmentsByIDParams struct {

	/*Body*/
	Body *models.AbTestSegment
	/*ID
	  The id of the requested A/B Test.

	*/
	ID string
	/*SegmentID
	  The id of the segment in the test.

	*/
	SegmentID string
	/*SiteID
	  ID of the site that the A/B tests are contained within.

	*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithContext(ctx context.Context) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithBody(body *models.AbTestSegment) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetBody(body *models.AbTestSegment) {
	o.Body = body
}

// WithID adds the id to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithID(id string) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetID(id string) {
	o.ID = id
}

// WithSegmentID adds the segmentID to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithSegmentID(segmentID string) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetSegmentID(segmentID)
	return o
}

// SetSegmentID adds the segmentId to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetSegmentID(segmentID string) {
	o.SegmentID = segmentID
}

// WithSiteID adds the siteID to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WithSiteID(siteID string) *PutSitesByIDAbTestsByIDSegmentsByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID ab tests by ID segments by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDAbTestsByIDSegmentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param segment_id
	if err := r.SetPathParam("segment_id", o.SegmentID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}