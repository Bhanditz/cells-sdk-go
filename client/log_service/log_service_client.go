// Code generated by go-swagger; DO NOT EDIT.

package log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new log service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Syslog technicals logs in Json or c s v format
*/
func (a *Client) Syslog(params *SyslogParams) (*SyslogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyslogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Syslog",
		Method:             "POST",
		PathPattern:        "/log/sys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SyslogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SyslogOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}