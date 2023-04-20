// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vsan_health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vsan health check API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vsan health check API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GETHealthCheckQuery(params *GETHealthCheckQueryParams, opts ...ClientOption) (*GETHealthCheckQueryOK, *GETHealthCheckQueryAccepted, error)

	GETHealthCheckStatus(params *GETHealthCheckStatusParams, opts ...ClientOption) (*GETHealthCheckStatusOK, error)

	GETHealthCheckStatusTask(params *GETHealthCheckStatusTaskParams, opts ...ClientOption) (*GETHealthCheckStatusTaskOK, error)

	SetHealthCheckStatus(params *SetHealthCheckStatusParams, opts ...ClientOption) (*SetHealthCheckStatusOK, *SetHealthCheckStatusAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GETHealthCheckQuery Gets v SAN health check status

Get vSAN health check status for all cluster on the domain
*/
func (a *Client) GETHealthCheckQuery(params *GETHealthCheckQueryParams, opts ...ClientOption) (*GETHealthCheckQueryOK, *GETHealthCheckQueryAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHealthCheckQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHealthCheckQuery",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/health-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHealthCheckQueryReader{formats: a.formats},
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
	case *GETHealthCheckQueryOK:
		return value, nil, nil
	case *GETHealthCheckQueryAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for vsan_health_check: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHealthCheckStatus Gets v SAN health check status by query Id

Get vSAN health check status for a given Query Id
*/
func (a *Client) GETHealthCheckStatus(params *GETHealthCheckStatusParams, opts ...ClientOption) (*GETHealthCheckStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHealthCheckStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHealthCheckStatus",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/health-checks/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHealthCheckStatusReader{formats: a.formats},
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
	success, ok := result.(*GETHealthCheckStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHealthCheckStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETHealthCheckStatusTask Gets v SAN health check update task status

Get vSAN health check update task status for a given task Id
*/
func (a *Client) GETHealthCheckStatusTask(params *GETHealthCheckStatusTaskParams, opts ...ClientOption) (*GETHealthCheckStatusTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETHealthCheckStatusTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getHealthCheckStatusTask",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainId}/health-checks/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETHealthCheckStatusTaskReader{formats: a.formats},
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
	success, ok := result.(*GETHealthCheckStatusTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHealthCheckStatusTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetHealthCheckStatus updates v SAN health check status

Update vSAN health check status for domain
*/
func (a *Client) SetHealthCheckStatus(params *SetHealthCheckStatusParams, opts ...ClientOption) (*SetHealthCheckStatusOK, *SetHealthCheckStatusAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetHealthCheckStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setHealthCheckStatus",
		Method:             "PATCH",
		PathPattern:        "/v1/domains/{domainId}/health-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetHealthCheckStatusReader{formats: a.formats},
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
	case *SetHealthCheckStatusOK:
		return value, nil, nil
	case *SetHealthCheckStatusAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for vsan_health_check: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}