// Code generated by go-swagger; DO NOT EDIT.

package libraries

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

// NewGetLibrariesByIDFoldersByIDParams creates a new GetLibrariesByIDFoldersByIDParams object
// with the default values initialized.
func NewGetLibrariesByIDFoldersByIDParams() *GetLibrariesByIDFoldersByIDParams {
	var ()
	return &GetLibrariesByIDFoldersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLibrariesByIDFoldersByIDParamsWithTimeout creates a new GetLibrariesByIDFoldersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLibrariesByIDFoldersByIDParamsWithTimeout(timeout time.Duration) *GetLibrariesByIDFoldersByIDParams {
	var ()
	return &GetLibrariesByIDFoldersByIDParams{

		timeout: timeout,
	}
}

// NewGetLibrariesByIDFoldersByIDParamsWithContext creates a new GetLibrariesByIDFoldersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLibrariesByIDFoldersByIDParamsWithContext(ctx context.Context) *GetLibrariesByIDFoldersByIDParams {
	var ()
	return &GetLibrariesByIDFoldersByIDParams{

		Context: ctx,
	}
}

// NewGetLibrariesByIDFoldersByIDParamsWithHTTPClient creates a new GetLibrariesByIDFoldersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLibrariesByIDFoldersByIDParamsWithHTTPClient(client *http.Client) *GetLibrariesByIDFoldersByIDParams {
	var ()
	return &GetLibrariesByIDFoldersByIDParams{
		HTTPClient: client,
	}
}

/*GetLibrariesByIDFoldersByIDParams contains all the parameters to send to the API endpoint
for the get libraries by ID folders by ID operation typically these are written to a http.Request
*/
type GetLibrariesByIDFoldersByIDParams struct {

	/*FolderID
	  ID of a target folder.

	*/
	FolderID string
	/*LibraryID
	  ID of the shared library or the site-id in case of a private library.

	*/
	LibraryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) WithTimeout(timeout time.Duration) *GetLibrariesByIDFoldersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) WithContext(ctx context.Context) *GetLibrariesByIDFoldersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) WithHTTPClient(client *http.Client) *GetLibrariesByIDFoldersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderID adds the folderID to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) WithFolderID(folderID string) *GetLibrariesByIDFoldersByIDParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) SetFolderID(folderID string) {
	o.FolderID = folderID
}

// WithLibraryID adds the libraryID to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) WithLibraryID(libraryID string) *GetLibrariesByIDFoldersByIDParams {
	o.SetLibraryID(libraryID)
	return o
}

// SetLibraryID adds the libraryId to the get libraries by ID folders by ID params
func (o *GetLibrariesByIDFoldersByIDParams) SetLibraryID(libraryID string) {
	o.LibraryID = libraryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLibrariesByIDFoldersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param folder_id
	if err := r.SetPathParam("folder_id", o.FolderID); err != nil {
		return err
	}

	// path param library_id
	if err := r.SetPathParam("library_id", o.LibraryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}