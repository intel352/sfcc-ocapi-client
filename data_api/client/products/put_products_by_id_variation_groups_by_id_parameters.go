// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewPutProductsByIDVariationGroupsByIDParams creates a new PutProductsByIDVariationGroupsByIDParams object
// with the default values initialized.
func NewPutProductsByIDVariationGroupsByIDParams() *PutProductsByIDVariationGroupsByIDParams {
	var ()
	return &PutProductsByIDVariationGroupsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutProductsByIDVariationGroupsByIDParamsWithTimeout creates a new PutProductsByIDVariationGroupsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutProductsByIDVariationGroupsByIDParamsWithTimeout(timeout time.Duration) *PutProductsByIDVariationGroupsByIDParams {
	var ()
	return &PutProductsByIDVariationGroupsByIDParams{

		timeout: timeout,
	}
}

// NewPutProductsByIDVariationGroupsByIDParamsWithContext creates a new PutProductsByIDVariationGroupsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutProductsByIDVariationGroupsByIDParamsWithContext(ctx context.Context) *PutProductsByIDVariationGroupsByIDParams {
	var ()
	return &PutProductsByIDVariationGroupsByIDParams{

		Context: ctx,
	}
}

// NewPutProductsByIDVariationGroupsByIDParamsWithHTTPClient creates a new PutProductsByIDVariationGroupsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutProductsByIDVariationGroupsByIDParamsWithHTTPClient(client *http.Client) *PutProductsByIDVariationGroupsByIDParams {
	var ()
	return &PutProductsByIDVariationGroupsByIDParams{
		HTTPClient: client,
	}
}

/*PutProductsByIDVariationGroupsByIDParams contains all the parameters to send to the API endpoint
for the put products by ID variation groups by ID operation typically these are written to a http.Request
*/
type PutProductsByIDVariationGroupsByIDParams struct {

	/*Body*/
	Body *models.VariationGroup
	/*ID
	  The id of the variation group product.

	*/
	ID string
	/*MasterProductID
	  The id of the master product that contains the variation group.

	*/
	MasterProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithTimeout(timeout time.Duration) *PutProductsByIDVariationGroupsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithContext(ctx context.Context) *PutProductsByIDVariationGroupsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithHTTPClient(client *http.Client) *PutProductsByIDVariationGroupsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithBody(body *models.VariationGroup) *PutProductsByIDVariationGroupsByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetBody(body *models.VariationGroup) {
	o.Body = body
}

// WithID adds the id to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithID(id string) *PutProductsByIDVariationGroupsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetID(id string) {
	o.ID = id
}

// WithMasterProductID adds the masterProductID to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) WithMasterProductID(masterProductID string) *PutProductsByIDVariationGroupsByIDParams {
	o.SetMasterProductID(masterProductID)
	return o
}

// SetMasterProductID adds the masterProductId to the put products by ID variation groups by ID params
func (o *PutProductsByIDVariationGroupsByIDParams) SetMasterProductID(masterProductID string) {
	o.MasterProductID = masterProductID
}

// WriteToRequest writes these params to a swagger request
func (o *PutProductsByIDVariationGroupsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param master_product_id
	if err := r.SetPathParam("master_product_id", o.MasterProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
