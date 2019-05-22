// Code generated by go-swagger; DO NOT EDIT.

package baskets

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

	models "github.com/intel352/sfcc-ocapi-client/shop_api/models"
)

// NewPostBasketsReferenceParams creates a new PostBasketsReferenceParams object
// with the default values initialized.
func NewPostBasketsReferenceParams() *PostBasketsReferenceParams {
	var ()
	return &PostBasketsReferenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostBasketsReferenceParamsWithTimeout creates a new PostBasketsReferenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostBasketsReferenceParamsWithTimeout(timeout time.Duration) *PostBasketsReferenceParams {
	var ()
	return &PostBasketsReferenceParams{

		timeout: timeout,
	}
}

// NewPostBasketsReferenceParamsWithContext creates a new PostBasketsReferenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostBasketsReferenceParamsWithContext(ctx context.Context) *PostBasketsReferenceParams {
	var ()
	return &PostBasketsReferenceParams{

		Context: ctx,
	}
}

// NewPostBasketsReferenceParamsWithHTTPClient creates a new PostBasketsReferenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostBasketsReferenceParamsWithHTTPClient(client *http.Client) *PostBasketsReferenceParams {
	var ()
	return &PostBasketsReferenceParams{
		HTTPClient: client,
	}
}

/*PostBasketsReferenceParams contains all the parameters to send to the API endpoint
for the post baskets reference operation typically these are written to a http.Request
*/
type PostBasketsReferenceParams struct {

	/*Body*/
	Body *models.BasketReference

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post baskets reference params
func (o *PostBasketsReferenceParams) WithTimeout(timeout time.Duration) *PostBasketsReferenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post baskets reference params
func (o *PostBasketsReferenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post baskets reference params
func (o *PostBasketsReferenceParams) WithContext(ctx context.Context) *PostBasketsReferenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post baskets reference params
func (o *PostBasketsReferenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post baskets reference params
func (o *PostBasketsReferenceParams) WithHTTPClient(client *http.Client) *PostBasketsReferenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post baskets reference params
func (o *PostBasketsReferenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post baskets reference params
func (o *PostBasketsReferenceParams) WithBody(body *models.BasketReference) *PostBasketsReferenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post baskets reference params
func (o *PostBasketsReferenceParams) SetBody(body *models.BasketReference) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBasketsReferenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
