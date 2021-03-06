// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteUsersByID Action to delete a single user.
*/
func (a *Client) DeleteUsersByID(params *DeleteUsersByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUsersByID",
		Method:             "DELETE",
		PathPattern:        "/users/{login}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteUsersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUsersByIDNoContent), nil

}

/*
GetUsers Action to get all users with no filtering.
*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetUsersByID Action to get a user.
*/
func (a *Client) GetUsersByID(params *GetUsersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersByID",
		Method:             "GET",
		PathPattern:        "/users/{login}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetUsersThis Action to get the user password expiration information.
*/
func (a *Client) GetUsersThis(params *GetUsersThisParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersThisParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersThis",
		Method:             "GET",
		PathPattern:        "/users/this",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUsersThisReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PatchUsersByID Action to update a user.

 Note: The locked flag and the user password cannot be updated with this resource.
*/
func (a *Client) PatchUsersByID(params *PatchUsersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUsersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchUsersByID",
		Method:             "PATCH",
		PathPattern:        "/users/{login}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchUsersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PatchUsersThisPassword Action to change a user password.
*/
func (a *Client) PatchUsersThisPassword(params *PatchUsersThisPasswordParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUsersThisPasswordParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchUsersThisPassword",
		Method:             "PATCH",
		PathPattern:        "/users/this/password",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchUsersThisPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutUsersByID Action to create or overwrite a user.

 If a user with the given login already exists, the existing user will be overwritten.
 If no such login exists, a new user is created.
*/
func (a *Client) PutUsersByID(params *PutUsersByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUsersByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUsersByID",
		Method:             "PUT",
		PathPattern:        "/users/{login}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUsersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
