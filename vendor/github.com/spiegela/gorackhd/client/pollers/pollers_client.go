package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pollers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pollers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PollersCurrentDataGet gets latest data for a poller

Get latest output data generated by the poller poller with the specified identifier.
*/
func (a *Client) PollersCurrentDataGet(params *PollersCurrentDataGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersCurrentDataGetOK, *PollersCurrentDataGetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersCurrentDataGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersCurrentDataGet",
		Method:             "GET",
		PathPattern:        "/pollers/{identifier}/data/current",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersCurrentDataGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PollersCurrentDataGetOK:
		return value, nil, nil
	case *PollersCurrentDataGetNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PollersDataGet gets output data for a poller

Get the complete history of output data generated by the poller with the specified identifier.

*/
func (a *Client) PollersDataGet(params *PollersDataGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersDataGetOK, *PollersDataGetNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersDataGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersDataGet",
		Method:             "GET",
		PathPattern:        "/pollers/{identifier}/data",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersDataGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PollersDataGetOK:
		return value, nil, nil
	case *PollersDataGetNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PollersDelete deletes the specified poller

Delete the poller with the specified identifier.
*/
func (a *Client) PollersDelete(params *PollersDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*PollersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersDelete",
		Method:             "DELETE",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersDeleteNoContent), nil

}

/*
PollersGet gets a list of all active pollers

Get list of all pollers that are currently running.
*/
func (a *Client) PollersGet(params *PollersGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersGet",
		Method:             "GET",
		PathPattern:        "/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersGetOK), nil

}

/*
PollersIDGet gets the specified poller

Get information associated with the specified poller, including type, run interval, command, and whether the poller is paused.

*/
func (a *Client) PollersIDGet(params *PollersIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersIdGet",
		Method:             "GET",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersIDGetOK), nil

}

/*
PollersLibByIDGet gets the specified poller

Get the poller definition with the specified identifier from the poller library

*/
func (a *Client) PollersLibByIDGet(params *PollersLibByIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersLibByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersLibByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersLibByIdGet",
		Method:             "GET",
		PathPattern:        "/pollers/library/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersLibByIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersLibByIDGetOK), nil

}

/*
PollersLibGet gets a list of possible pollers

Get a list of all available poller definitions in the poller library.
*/
func (a *Client) PollersLibGet(params *PollersLibGetParams, authInfo runtime.ClientAuthInfoWriter) (*PollersLibGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersLibGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersLibGet",
		Method:             "GET",
		PathPattern:        "/pollers/library",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersLibGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersLibGetOK), nil

}

/*
PollersPatch patches a poller

Modify one or more properties of the poller with the specified identifier.
*/
func (a *Client) PollersPatch(params *PollersPatchParams, authInfo runtime.ClientAuthInfoWriter) (*PollersPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersPatchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersPatch",
		Method:             "PATCH",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersPatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersPatchOK), nil

}

/*
PollersPost posts a poller

Create and start a new poller, which will run at the specified time interval.
*/
func (a *Client) PollersPost(params *PollersPostParams, authInfo runtime.ClientAuthInfoWriter) (*PollersPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPollersPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pollersPost",
		Method:             "POST",
		PathPattern:        "/pollers",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PollersPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PollersPostCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}