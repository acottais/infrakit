package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new lookups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for lookups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
LookupsDelByID deletes a lookup

Delete the specified lookup.
*/
func (a *Client) LookupsDelByID(params *LookupsDelByIDParams, authInfo runtime.ClientAuthInfoWriter) (*LookupsDelByIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookupsDelByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lookupsDelById",
		Method:             "DELETE",
		PathPattern:        "/lookups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LookupsDelByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LookupsDelByIDNoContent), nil

}

/*
LookupsGet gets a list of lookups

Get a list of all lookups currently stored. Lookups relate mac addresses to ip addresses.
*/
func (a *Client) LookupsGet(params *LookupsGetParams, authInfo runtime.ClientAuthInfoWriter) (*LookupsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookupsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lookupsGet",
		Method:             "GET",
		PathPattern:        "/lookups",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LookupsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LookupsGetOK), nil

}

/*
LookupsGetByID gets a lookup

Get a lookup by specifying its identifier.
*/
func (a *Client) LookupsGetByID(params *LookupsGetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*LookupsGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookupsGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lookupsGetById",
		Method:             "GET",
		PathPattern:        "/lookups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LookupsGetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LookupsGetByIDOK), nil

}

/*
LookupsPatchByID patches a lookup

Modify the properties of a lookup.
*/
func (a *Client) LookupsPatchByID(params *LookupsPatchByIDParams, authInfo runtime.ClientAuthInfoWriter) (*LookupsPatchByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookupsPatchByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lookupsPatchById",
		Method:             "PATCH",
		PathPattern:        "/lookups/{id}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LookupsPatchByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LookupsPatchByIDOK), nil

}

/*
LookupsPost posts a lookup

Create and store a new lookup.
*/
func (a *Client) LookupsPost(params *LookupsPostParams, authInfo runtime.ClientAuthInfoWriter) (*LookupsPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookupsPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "lookupsPost",
		Method:             "POST",
		PathPattern:        "/lookups",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LookupsPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LookupsPostCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}