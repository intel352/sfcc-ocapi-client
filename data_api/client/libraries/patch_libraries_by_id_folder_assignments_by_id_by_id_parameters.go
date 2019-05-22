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

	models "github.com/intel352/sfcc-ocapi-client/data_api/models"
)

// NewPatchLibrariesByIDFolderAssignmentsByIDByIDParams creates a new PatchLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized.
func NewPatchLibrariesByIDFolderAssignmentsByIDByIDParams() *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &PatchLibrariesByIDFolderAssignmentsByIDByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithTimeout creates a new PatchLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithTimeout(timeout time.Duration) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &PatchLibrariesByIDFolderAssignmentsByIDByIDParams{

		timeout: timeout,
	}
}

// NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithContext creates a new PatchLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithContext(ctx context.Context) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &PatchLibrariesByIDFolderAssignmentsByIDByIDParams{

		Context: ctx,
	}
}

// NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithHTTPClient creates a new PatchLibrariesByIDFolderAssignmentsByIDByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLibrariesByIDFolderAssignmentsByIDByIDParamsWithHTTPClient(client *http.Client) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	var ()
	return &PatchLibrariesByIDFolderAssignmentsByIDByIDParams{
		HTTPClient: client,
	}
}

/*PatchLibrariesByIDFolderAssignmentsByIDByIDParams contains all the parameters to send to the API endpoint
for the patch libraries by ID folder assignments by ID by ID operation typically these are written to a http.Request
*/
type PatchLibrariesByIDFolderAssignmentsByIDByIDParams struct {

	/*Body*/
	Body *models.ContentFolderAssignment
	/*ContentID
	  the content id of the assignment

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

// WithTimeout adds the timeout to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithTimeout(timeout time.Duration) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithContext(ctx context.Context) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithHTTPClient(client *http.Client) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithBody(body *models.ContentFolderAssignment) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetBody(body *models.ContentFolderAssignment) {
	o.Body = body
}

// WithContentID adds the contentID to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithContentID(contentID string) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetContentID(contentID)
	return o
}

// SetContentID adds the contentId to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetContentID(contentID string) {
	o.ContentID = contentID
}

// WithFolderID adds the folderID to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithFolderID(folderID string) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetFolderID(folderID string) {
	o.FolderID = folderID
}

// WithLibraryID adds the libraryID to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WithLibraryID(libraryID string) *PatchLibrariesByIDFolderAssignmentsByIDByIDParams {
	o.SetLibraryID(libraryID)
	return o
}

// SetLibraryID adds the libraryId to the patch libraries by ID folder assignments by ID by ID params
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) SetLibraryID(libraryID string) {
	o.LibraryID = libraryID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLibrariesByIDFolderAssignmentsByIDByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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