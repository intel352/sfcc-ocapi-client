// Code generated by go-swagger; DO NOT EDIT.

package system_object_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system object definitions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system object definitions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByID Deletes the attribute definition by ID
*/
func (a *Client) DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByID(params *DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSystemObjectDefinitionsByIDAttributeDefinitionsByID",
		Method:             "DELETE",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definitions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSystemObjectDefinitionsByIDAttributeDefinitionsByIDNoContent), nil

}

/*
DeleteSystemObjectDefinitionsByIDAttributeGroupsByID Deletes the attribute group by ID
*/
func (a *Client) DeleteSystemObjectDefinitionsByIDAttributeGroupsByID(params *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSystemObjectDefinitionsByIDAttributeGroupsByID",
		Method:             "DELETE",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDNoContent), nil

}

/*
DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID Un-assign an attribute definition from an attribute group.
*/
func (a *Client) DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID(params *DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID",
		Method:             "DELETE",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{group_id}/attribute_definitions/{def_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent), nil

}

/*
GetSystemObjectDefinitions Action to get all the system objects with no filtering.
*/
func (a *Client) GetSystemObjectDefinitions(params *GetSystemObjectDefinitionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitions",
		Method:             "GET",
		PathPattern:        "/system_object_definitions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsReader{formats: a.formats},
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
GetSystemObjectDefinitionsByID Action to get system object information.
*/
func (a *Client) GetSystemObjectDefinitionsByID(params *GetSystemObjectDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitionsByID",
		Method:             "GET",
		PathPattern:        "/system_object_definitions/{object_type}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsByIDReader{formats: a.formats},
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
GetSystemObjectDefinitionsByIDAttributeDefinitions Action to get all the attribute definitions with no filtering.
*/
func (a *Client) GetSystemObjectDefinitionsByIDAttributeDefinitions(params *GetSystemObjectDefinitionsByIDAttributeDefinitionsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsByIDAttributeDefinitionsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitionsByIDAttributeDefinitions",
		Method:             "GET",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definitions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsByIDAttributeDefinitionsReader{formats: a.formats},
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
GetSystemObjectDefinitionsByIDAttributeDefinitionsByID Action to get attribute definition information.
*/
func (a *Client) GetSystemObjectDefinitionsByIDAttributeDefinitionsByID(params *GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitionsByIDAttributeDefinitionsByID",
		Method:             "GET",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definitions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader{formats: a.formats},
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
GetSystemObjectDefinitionsByIDAttributeGroups Action to get all the attribute groups with no filtering.
*/
func (a *Client) GetSystemObjectDefinitionsByIDAttributeGroups(params *GetSystemObjectDefinitionsByIDAttributeGroupsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsByIDAttributeGroupsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitionsByIDAttributeGroups",
		Method:             "GET",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsByIDAttributeGroupsReader{formats: a.formats},
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
GetSystemObjectDefinitionsByIDAttributeGroupsByID Action to get attribute group information.
*/
func (a *Client) GetSystemObjectDefinitionsByIDAttributeGroupsByID(params *GetSystemObjectDefinitionsByIDAttributeGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSystemObjectDefinitionsByIDAttributeGroupsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSystemObjectDefinitionsByIDAttributeGroupsByID",
		Method:             "GET",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSystemObjectDefinitionsByIDAttributeGroupsByIDReader{formats: a.formats},
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
PatchSystemObjectDefinitionsByIDAttributeDefinitionsByID Updates the attribute definition with the specified information. The request must include the If-Match header, which holds
 the last known base-point information. The value of this header is an "ETag" representing the attribute definition state. If
 the request does not contain an If-Match header with the current server customer "ETag", a 409 (IfMatchRequiredException)
 fault is returned. If the If-Match header does not match the current server attribute definition "ETag", a 412 (InvalidIfMatchException)
 fault is returned.
*/
func (a *Client) PatchSystemObjectDefinitionsByIDAttributeDefinitionsByID(params *PatchSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchSystemObjectDefinitionsByIDAttributeDefinitionsByID",
		Method:             "PATCH",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definitions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader{formats: a.formats},
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
PatchSystemObjectDefinitionsByIDAttributeGroupsByID Updates the attribute group with the specified information. The request must include the If-Match header, which holds
 the last known base-point information. The value of this header is an "ETag" representing the attribute group state. If
 the request does not contain an If-Match header with the current server customer "ETag", a 409 (IfMatchRequiredException)
 fault is returned. If the If-Match header does not match the current server attribute group "ETag", a 412 (InvalidIfMatchException)
 fault is returned.
*/
func (a *Client) PatchSystemObjectDefinitionsByIDAttributeGroupsByID(params *PatchSystemObjectDefinitionsByIDAttributeGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchSystemObjectDefinitionsByIDAttributeGroupsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchSystemObjectDefinitionsByIDAttributeGroupsByID",
		Method:             "PATCH",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchSystemObjectDefinitionsByIDAttributeGroupsByIDReader{formats: a.formats},
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
PostSystemObjectDefinitionsByIDAttributeDefinitionSearch Searches the attribute definitions of the specified system object type.

 The query attribute specifies a complex query that can be used to narrow down the search. Attributes are grouped
 into different buckets.  These are the list of searchable attributes with their corresponding buckets:

 Main:

    id - String
    display_name* - Localized String
    description* - Localized String
    key - boolean
    mandatory - boolean
    searchable - boolean
    system - boolean
    visible - boolean

 Definition version:

    localizable - boolean
    site_specific - boolean
    value_type - one of {string, int, double, text, html, date, image, boolean, money, quantity, datetime, email, password, set_of_string, set_of_int, set_of_double, enum_of_string, enum_of_int}

 Group:

    group - String


 Only attributes in the same bucket can be joined using a disjunction (OR).
 For instance, when joining localizable and description above, only a conjunction is allowed (AND), whereas display_name
 and description can be joined using a disjunction because they are in the same bucket.  If an attribute
 is used in a disjunction (OR) that violates this rule, an exception will be thrown.

 * These attributes are not searchable or sortable for built in system attributes.  They work normally for non
 system attributes.

 Note that only searchable attributes can be used in sorting.
*/
func (a *Client) PostSystemObjectDefinitionsByIDAttributeDefinitionSearch(params *PostSystemObjectDefinitionsByIDAttributeDefinitionSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSystemObjectDefinitionsByIDAttributeDefinitionSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSystemObjectDefinitionsByIDAttributeDefinitionSearch",
		Method:             "POST",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definition_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSystemObjectDefinitionsByIDAttributeDefinitionSearchReader{formats: a.formats},
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
PostSystemObjectDefinitionsByIDAttributeGroupSearch Searches for attribute groups.

 The query attribute specifies a complex query that can be used to narrow down the search. These are the list
 of searchable attributes:

 id - String
 display_name - Localized<String>
 description - Localized<String>
 position - Double
 internal - Boolean


 The output of the query can also be sorted. These are the list of sortable attributes:

 id - String
 display_name - Localized<String>
 description - Localized<String>
 position - Double
 internal - Boolean

*/
func (a *Client) PostSystemObjectDefinitionsByIDAttributeGroupSearch(params *PostSystemObjectDefinitionsByIDAttributeGroupSearchParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSystemObjectDefinitionsByIDAttributeGroupSearchParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSystemObjectDefinitionsByIDAttributeGroupSearch",
		Method:             "POST",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_group_search",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostSystemObjectDefinitionsByIDAttributeGroupSearchReader{formats: a.formats},
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
PutSystemObjectDefinitionsByIDAttributeDefinitionsByID Creates a attribute definition using the information provided.
*/
func (a *Client) PutSystemObjectDefinitionsByIDAttributeDefinitionsByID(params *PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSystemObjectDefinitionsByIDAttributeDefinitionsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putSystemObjectDefinitionsByIDAttributeDefinitionsByID",
		Method:             "PUT",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_definitions/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSystemObjectDefinitionsByIDAttributeDefinitionsByIDReader{formats: a.formats},
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
PutSystemObjectDefinitionsByIDAttributeGroupsByID Creates a attribute group using the information provided. If a attribute group with the same unique identifier, it will be deleted and a new one will be created unless the header x-dw-validate-existing=true is passed in with the request.
*/
func (a *Client) PutSystemObjectDefinitionsByIDAttributeGroupsByID(params *PutSystemObjectDefinitionsByIDAttributeGroupsByIDParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putSystemObjectDefinitionsByIDAttributeGroupsByID",
		Method:             "PUT",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSystemObjectDefinitionsByIDAttributeGroupsByIDReader{formats: a.formats},
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
PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID Assign an attribute definition to an attribute group.
*/
func (a *Client) PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID(params *PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByID",
		Method:             "PUT",
		PathPattern:        "/system_object_definitions/{object_type}/attribute_groups/{group_id}/attribute_definitions/{def_id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutSystemObjectDefinitionsByIDAttributeGroupsByIDAttributeDefinitionsByIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
