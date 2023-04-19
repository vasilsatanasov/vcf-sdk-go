// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system_prechecks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system prechecks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system prechecks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETPrecheckTask(params *GETPrecheckTaskParams, opts ...ClientOption) (*GETPrecheckTaskOK, error)

	PrecheckSystem(params *PrecheckSystemParams, opts ...ClientOption) (*PrecheckSystemOK, *PrecheckSystemAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETPrecheckTask Gets precheck task by ID

Monitor the progress of precheck task by the precheck task ID
*/
func (a *Client) GETPrecheckTask(params *GETPrecheckTaskParams, opts ...ClientOption) (*GETPrecheckTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETPrecheckTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrecheckTask",
		Method:             "GET",
		PathPattern:        "/v1/system/prechecks/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETPrecheckTaskReader{formats: a.formats},
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
	success, ok := result.(*GETPrecheckTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrecheckTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrecheckSystem prechecks system

Perform precheck of resource(ex: Domain, Cluster). If only resource is specified, all resources/software components under it are included. If resource(Domain, Cluster etc) and specific resources/software components are provided, only those are included in precheck
*/
func (a *Client) PrecheckSystem(params *PrecheckSystemParams, opts ...ClientOption) (*PrecheckSystemOK, *PrecheckSystemAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrecheckSystemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "precheckSystem",
		Method:             "POST",
		PathPattern:        "/v1/system/prechecks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrecheckSystemReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PrecheckSystemOK:
		return value, nil, nil
	case *PrecheckSystemAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for system_prechecks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
