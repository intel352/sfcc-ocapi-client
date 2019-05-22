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

// NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParams creates a new DeleteLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized.
func NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParams() *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &DeleteLibrariesByIDFolderAssignmentsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithTimeout creates a new DeleteLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithTimeout(timeout time.Duration) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &DeleteLibrariesByIDFolderAssignmentsByIDByIDParams{

		timeout: timeout,
	}
}

// NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithContext creates a new DeleteLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithContext(ctx context.Context) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &DeleteLibrariesByIDFolderAssignmentsByIDByIDParams{

		Context: ctx,
	}
}

// NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithHTTPClient creates a new DeleteLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLibrariesByIDFolderAssignmentsByIDByIDParamsWithHTTPClient(client *http.Client) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &DeleteLibrariesByIDFolderAssignmentsByIDByIDParams{
		HTTPClient: client,
	}
}

/*DeleteLibrariesByIDFolderAssignmentsByIDByIDParams contains all the parameters to send to the API endpoint
for the delete libraries by ID folder assignments by ID by ID operation typically these are written to a http.Request
*/
type DeleteLibrariesByIDFolderAssignmentsByIDByIDParams struct {

	/*ContentID
	  the ID of the content asset to retrieve.

	*/
	ContentID string
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

// WithTimeout adds the timeout to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithTimeout(timeout time.Duration) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithContext(ctx context.Context) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithHTTPClient(client *http.Client) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentID adds the contentID to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithContentID(contentID string) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithFolderID adds the folderID to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithFolderID(folderID string) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetFolderID(folderID string) {
	o.FolderID = folderID
}

// WithLibraryID adds the libraryID to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WithLibraryID(libraryID string) *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetLibraryID(libraryID)
	return o
}

// SetLibraryID adds the libraryId to the delete libraries by ID folder assignments by ID by ID params
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) SetLibraryID(libraryID string) {
	o.LibraryID = libraryID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLibrariesByIDFolderAssignmentsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param content_id
	if err := r.SetPathParam("content_id", o.ContentID); err != nil {
		return err
	}

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