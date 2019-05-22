// Code generated by go-swagger; DO NOT EDIT.

package product_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProductSearchAvailabilityParams creates a new GetProductSearchAvailabilityParams object
// with the default values initialized.
func NewGetProductSearchAvailabilityParams() *GetProductSearchAvailabilityParams {
	var ()
	return &GetProductSearchAvailabilityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductSearchAvailabilityParamsWithTimeout creates a new GetProductSearchAvailabilityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductSearchAvailabilityParamsWithTimeout(timeout time.Duration) *GetProductSearchAvailabilityParams {
	var ()
	return &GetProductSearchAvailabilityParams{

		timeout: timeout,
	}
}

// NewGetProductSearchAvailabilityParamsWithContext creates a new GetProductSearchAvailabilityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProductSearchAvailabilityParamsWithContext(ctx context.Context) *GetProductSearchAvailabilityParams {
	var ()
	return &GetProductSearchAvailabilityParams{

		Context: ctx,
	}
}

// NewGetProductSearchAvailabilityParamsWithHTTPClient creates a new GetProductSearchAvailabilityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProductSearchAvailabilityParamsWithHTTPClient(client *http.Client) *GetProductSearchAvailabilityParams {
	var ()
	return &GetProductSearchAvailabilityParams{
		HTTPClient: client,
	}
}

/*GetProductSearchAvailabilityParams contains all the parameters to send to the API endpoint
for the get product search availability operation typically these are written to a http.Request
*/
type GetProductSearchAvailabilityParams struct {

	/*Count
	  The maximum number of instances per request. Default value is 25.

	*/
	Count *int32
	/*Locale
	  The locale context.

	*/
	Locale *string
	/*Q
	  The query phrase to search for.

	*/
	Q *string
	/*Refine
	  Parameter that represents a refinement attribute/value(s) pair. Refinement attribute id and
	            value(s) are separated by '='. Multiple values are supported by a sub-set of refinement attributes and
	            can be provided by separating them using a pipe (URL
	            encoded = "|"). Value ranges can be specified like this: refine=price=(100..500) Multiple refine
	            parameters can be provided by adding an underscore in combination with an integer counter right behind
	            the parameter name and a counter range 1..9. I.e. refine_1=c_refinementColor=red|green|blue. The
	            following system refinement attribute ids are supported:

	            cgid: Allows to refine per single category id. Multiple category ids are not supported.
	            price: Allows to refine per single price range. Multiple price ranges are not supported.
	            pmid: Allows to refine per promotion id(s).
	            orderable_only: Unavailable products will be excluded from the search results if true is set. Multiple
	            refinement values are not supported.


	*/
	Refine []string
	/*Sort
	  The id of the sorting option to sort the search hits.

	*/
	Sort *string
	/*Start
	  The result set index to return the first instance for. Default value is 0.

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithTimeout(timeout time.Duration) *GetProductSearchAvailabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithContext(ctx context.Context) *GetProductSearchAvailabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithHTTPClient(client *http.Client) *GetProductSearchAvailabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithCount(count *int32) *GetProductSearchAvailabilityParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetCount(count *int32) {
	o.Count = count
}

// WithLocale adds the locale to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithLocale(locale *string) *GetProductSearchAvailabilityParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WithQ adds the q to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithQ(q *string) *GetProductSearchAvailabilityParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetQ(q *string) {
	o.Q = q
}

// WithRefine adds the refine to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithRefine(refine []string) *GetProductSearchAvailabilityParams {
	o.SetRefine(refine)
	return o
}

// SetRefine adds the refine to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetRefine(refine []string) {
	o.Refine = refine
}

// WithSort adds the sort to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithSort(sort *string) *GetProductSearchAvailabilityParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStart adds the start to the get product search availability params
func (o *GetProductSearchAvailabilityParams) WithStart(start *int32) *GetProductSearchAvailabilityParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get product search availability params
func (o *GetProductSearchAvailabilityParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductSearchAvailabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Locale != nil {

		// query param locale
		var qrLocale string
		if o.Locale != nil {
			qrLocale = *o.Locale
		}
		qLocale := qrLocale
		if qLocale != "" {
			if err := r.SetQueryParam("locale", qLocale); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	valuesRefine := o.Refine

	joinedRefine := swag.JoinByFormat(valuesRefine, "")
	// query array param refine
	if err := r.SetQueryParam("refine", joinedRefine...); err != nil {
		return err
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}