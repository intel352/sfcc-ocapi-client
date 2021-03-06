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
)

// NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams creates a new PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams object
// with the default values initialized.
func NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams() *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithTimeout creates a new PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithTimeout(timeout time.Duration) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams{

		timeout: timeout,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithContext creates a new PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithContext(ctx context.Context) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams{

		Context: ctx,
	}
}

// NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithHTTPClient creates a new PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParamsWithHTTPClient(client *http.Client) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	var ()
	return &PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams{
		HTTPClient: client,
	}
}

/*PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams contains all the parameters to send to the API endpoint
for the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID operation typically these are written to a http.Request
*/
type PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams struct {

	/*AbTestID*/
	AbTestID string
	/*CategoryID*/
	CategoryID string
	/*RuleContext*/
	RuleContext *string
	/*SegmentID*/
	SegmentID string
	/*SiteID*/
	SiteID string
	/*SortingRuleID*/
	SortingRuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithTimeout(timeout time.Duration) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithContext(ctx context.Context) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithHTTPClient(client *http.Client) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAbTestID adds the abTestID to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithAbTestID(abTestID string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetAbTestID(abTestID)
	return o
}

// SetAbTestID adds the abTestId to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetAbTestID(abTestID string) {
	o.AbTestID = abTestID
}

// WithCategoryID adds the categoryID to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithCategoryID(categoryID string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithRuleContext adds the ruleContext to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithRuleContext(ruleContext *string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetRuleContext(ruleContext)
	return o
}

// SetRuleContext adds the ruleContext to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetRuleContext(ruleContext *string) {
	o.RuleContext = ruleContext
}

// WithSegmentID adds the segmentID to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithSegmentID(segmentID string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetSegmentID(segmentID)
	return o
}

// SetSegmentID adds the segmentId to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetSegmentID(segmentID string) {
	o.SegmentID = segmentID
}

// WithSiteID adds the siteID to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithSiteID(siteID string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSortingRuleID adds the sortingRuleID to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WithSortingRuleID(sortingRuleID string) *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams {
	o.SetSortingRuleID(sortingRuleID)
	return o
}

// SetSortingRuleID adds the sortingRuleId to the put sites by ID ab tests by ID segments by ID sorting rules by ID by ID params
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) SetSortingRuleID(sortingRuleID string) {
	o.SortingRuleID = sortingRuleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSitesByIDAbTestsByIDSegmentsByIDSortingRulesByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ab_test_id
	if err := r.SetPathParam("ab_test_id", o.AbTestID); err != nil {
		return err
	}

	// path param category_id
	if err := r.SetPathParam("category_id", o.CategoryID); err != nil {
		return err
	}

	if o.RuleContext != nil {

		// query param rule_context
		var qrRuleContext string
		if o.RuleContext != nil {
			qrRuleContext = *o.RuleContext
		}
		qRuleContext := qrRuleContext
		if qRuleContext != "" {
			if err := r.SetQueryParam("rule_context", qRuleContext); err != nil {
				return err
			}
		}

	}

	// path param segment_id
	if err := r.SetPathParam("segment_id", o.SegmentID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param sorting_rule_id
	if err := r.SetPathParam("sorting_rule_id", o.SortingRuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
