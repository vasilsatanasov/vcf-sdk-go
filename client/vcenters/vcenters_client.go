// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v centers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v centers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETVcenter(params *GETVcenterParams, opts ...ClientOption) (*GETVcenterOK, error)

	GETVCENTERS(params *GETVCENTERSParams, opts ...ClientOption) (*GETVCENTERSOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETVcenter Gets a v center
*/
func (a *Client) GETVcenter(params *GETVcenterParams, opts ...ClientOption) (*GETVcenterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETVcenterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVcenter",
		Method:             "GET",
		PathPattern:        "/v1/vcenters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETVcenterReader{formats: a.formats},
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
	success, ok := result.(*GETVcenterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVcenter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETVCENTERS Gets v centers
*/
func (a *Client) GETVCENTERS(params *GETVCENTERSParams, opts ...ClientOption) (*GETVCENTERSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETVCENTERSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVcenters",
		Method:             "GET",
		PathPattern:        "/v1/vcenters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETVCENTERSReader{formats: a.formats},
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
	success, ok := result.(*GETVCENTERSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVcenters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
