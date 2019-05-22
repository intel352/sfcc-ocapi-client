// Code generated by go-swagger; DO NOT EDIT.

package gift_certificate

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

// NewPostGiftCertificateParams creates a new PostGiftCertificateParams object
// with the default values initialized.
func NewPostGiftCertificateParams() *PostGiftCertificateParams {
	var ()
	return &PostGiftCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGiftCertificateParamsWithTimeout creates a new PostGiftCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGiftCertificateParamsWithTimeout(timeout time.Duration) *PostGiftCertificateParams {
	var ()
	return &PostGiftCertificateParams{

		timeout: timeout,
	}
}

// NewPostGiftCertificateParamsWithContext creates a new PostGiftCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGiftCertificateParamsWithContext(ctx context.Context) *PostGiftCertificateParams {
	var ()
	return &PostGiftCertificateParams{

		Context: ctx,
	}
}

// NewPostGiftCertificateParamsWithHTTPClient creates a new PostGiftCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGiftCertificateParamsWithHTTPClient(client *http.Client) *PostGiftCertificateParams {
	var ()
	return &PostGiftCertificateParams{
		HTTPClient: client,
	}
}

/*PostGiftCertificateParams contains all the parameters to send to the API endpoint
for the post gift certificate operation typically these are written to a http.Request
*/
type PostGiftCertificateParams struct {

	/*Body*/
	Body *models.GiftCertificateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gift certificate params
func (o *PostGiftCertificateParams) WithTimeout(timeout time.Duration) *PostGiftCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gift certificate params
func (o *PostGiftCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gift certificate params
func (o *PostGiftCertificateParams) WithContext(ctx context.Context) *PostGiftCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gift certificate params
func (o *PostGiftCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gift certificate params
func (o *PostGiftCertificateParams) WithHTTPClient(client *http.Client) *PostGiftCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gift certificate params
func (o *PostGiftCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post gift certificate params
func (o *PostGiftCertificateParams) WithBody(body *models.GiftCertificateRequest) *PostGiftCertificateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post gift certificate params
func (o *PostGiftCertificateParams) SetBody(body *models.GiftCertificateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostGiftCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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