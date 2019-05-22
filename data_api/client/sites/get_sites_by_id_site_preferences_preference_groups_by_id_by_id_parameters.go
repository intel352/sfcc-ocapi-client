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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams creates a new GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams object
// with the default values initialized.
func NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams() *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	var ()
	return &GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithTimeout creates a new GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithTimeout(timeout time.Duration) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	var ()
	return &GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams{

		timeout: timeout,
	}
}

// NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithContext creates a new GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithContext(ctx context.Context) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	var ()
	return &GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams{

		Context: ctx,
	}
}

// NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithHTTPClient creates a new GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParamsWithHTTPClient(client *http.Client) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	var ()
	return &GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams{
		HTTPClient: client,
	}
}

/*GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams contains all the parameters to send to the API endpoint
for the get sites by ID site preferences preference groups by ID by ID operation typically these are written to a http.Request
*/
type GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams struct {

	/*GroupID*/
	GroupID string
	/*InstanceType*/
	InstanceType string
	/*MaskPasswords*/
	MaskPasswords *bool
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithTimeout(timeout time.Duration) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithContext(ctx context.Context) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithHTTPClient(client *http.Client) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithGroupID(groupID string) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithInstanceType adds the instanceType to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithInstanceType(instanceType string) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetInstanceType(instanceType)
	return o
}

// SetInstanceType adds the instanceType to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetInstanceType(instanceType string) {
	o.InstanceType = instanceType
}

// WithMaskPasswords adds the maskPasswords to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithMaskPasswords(maskPasswords *bool) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetMaskPasswords(maskPasswords)
	return o
}

// SetMaskPasswords adds the maskPasswords to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetMaskPasswords(maskPasswords *bool) {
	o.MaskPasswords = maskPasswords
}

// WithSiteID adds the siteID to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WithSiteID(siteID string) *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get sites by ID site preferences preference groups by ID by ID params
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesByIDSitePreferencesPreferenceGroupsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	// path param instance_type
	if err := r.SetPathParam("instance_type", o.InstanceType); err != nil {
		return err
	}

	if o.MaskPasswords != nil {

		// query param mask_passwords
		var qrMaskPasswords bool
		if o.MaskPasswords != nil {
			qrMaskPasswords = *o.MaskPasswords
		}
		qMaskPasswords := swag.FormatBool(qrMaskPasswords)
		if qMaskPasswords != "" {
			if err := r.SetQueryParam("mask_passwords", qMaskPasswords); err != nil {
				return err
			}
		}

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