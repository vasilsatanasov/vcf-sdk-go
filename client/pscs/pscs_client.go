// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pscs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p s cs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p s cs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETPsc(params *GETPscParams, opts ...ClientOption) (*GETPscOK, error)

	GETPscs(params *GETPscsParams, opts ...ClientOption) (*GETPscsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETPsc Gets a p s c
*/
func (a *Client) GETPsc(params *GETPscParams, opts ...ClientOption) (*GETPscOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETPscParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPsc",
		Method:             "GET",
		PathPattern:        "/v1/pscs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETPscReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GETPscOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPsc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETPscs Gets the p s cs
*/
func (a *Client) GETPscs(params *GETPscsParams, opts ...ClientOption) (*GETPscsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETPscsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPscs",
		Method:             "GET",
		PathPattern:        "/v1/pscs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETPscsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GETPscsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPscs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}