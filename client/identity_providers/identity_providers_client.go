// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package identity_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identity providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddEmbeddedIdentitySource(params *AddEmbeddedIdentitySourceParams, opts ...ClientOption) (*AddEmbeddedIdentitySourceOK, *AddEmbeddedIdentitySourceNoContent, error)

	AddExternalIdentityProvider(params *AddExternalIdentityProviderParams, opts ...ClientOption) (*AddExternalIdentityProviderOK, *AddExternalIdentityProviderCreated, error)

	DeleteExternalIdentityProvider(params *DeleteExternalIdentityProviderParams, opts ...ClientOption) (*DeleteExternalIdentityProviderOK, *DeleteExternalIdentityProviderNoContent, error)

	DeleteIdentitySource(params *DeleteIdentitySourceParams, opts ...ClientOption) (*DeleteIdentitySourceOK, *DeleteIdentitySourceNoContent, error)

	GetAllIdps(params *GetAllIdpsParams, opts ...ClientOption) (*GetAllIdpsOK, error)

	GetIdentityProviderByID(params *GetIdentityProviderByIDParams, opts ...ClientOption) (*GetIdentityProviderByIDOK, error)

	UpdateEmbeddedIdentitySource(params *UpdateEmbeddedIdentitySourceParams, opts ...ClientOption) (*UpdateEmbeddedIdentitySourceOK, *UpdateEmbeddedIdentitySourceNoContent, error)

	UpdateExternalIdentityProvider(params *UpdateExternalIdentityProviderParams, opts ...ClientOption) (*UpdateExternalIdentityProviderOK, *UpdateExternalIdentityProviderNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddEmbeddedIdentitySource adds an embedded identity source

Add an identity source to an embedded IDP by id, if it exists
*/
func (a *Client) AddEmbeddedIdentitySource(params *AddEmbeddedIdentitySourceParams, opts ...ClientOption) (*AddEmbeddedIdentitySourceOK, *AddEmbeddedIdentitySourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEmbeddedIdentitySourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addEmbeddedIdentitySource",
		Method:             "POST",
		PathPattern:        "/v1/identity-providers/{id}/identity-sources",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddEmbeddedIdentitySourceReader{formats: a.formats},
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
	case *AddEmbeddedIdentitySourceOK:
		return value, nil, nil
	case *AddEmbeddedIdentitySourceNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddExternalIdentityProvider adds an external identity provider

Add an External Identity Provider
*/
func (a *Client) AddExternalIdentityProvider(params *AddExternalIdentityProviderParams, opts ...ClientOption) (*AddExternalIdentityProviderOK, *AddExternalIdentityProviderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddExternalIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addExternalIdentityProvider",
		Method:             "POST",
		PathPattern:        "/v1/identity-providers",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddExternalIdentityProviderReader{formats: a.formats},
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
	case *AddExternalIdentityProviderOK:
		return value, nil, nil
	case *AddExternalIdentityProviderCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteExternalIdentityProvider deletes an external identity provider

Delete an Identity Provider by its identifier, if it exists
*/
func (a *Client) DeleteExternalIdentityProvider(params *DeleteExternalIdentityProviderParams, opts ...ClientOption) (*DeleteExternalIdentityProviderOK, *DeleteExternalIdentityProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExternalIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteExternalIdentityProvider",
		Method:             "DELETE",
		PathPattern:        "/v1/identity-providers/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExternalIdentityProviderReader{formats: a.formats},
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
	case *DeleteExternalIdentityProviderOK:
		return value, nil, nil
	case *DeleteExternalIdentityProviderNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteIdentitySource deletes an identity source

Delete an Identity Source by domain name, if it exists
*/
func (a *Client) DeleteIdentitySource(params *DeleteIdentitySourceParams, opts ...ClientOption) (*DeleteIdentitySourceOK, *DeleteIdentitySourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteIdentitySourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteIdentitySource",
		Method:             "DELETE",
		PathPattern:        "/v1/identity-providers/{id}/identity-sources/{domainName}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteIdentitySourceReader{formats: a.formats},
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
	case *DeleteIdentitySourceOK:
		return value, nil, nil
	case *DeleteIdentitySourceNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllIdps gets all identity providers

Get a list of all Identity Providers
*/
func (a *Client) GetAllIdps(params *GetAllIdpsParams, opts ...ClientOption) (*GetAllIdpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllIdpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllIdps",
		Method:             "GET",
		PathPattern:        "/v1/identity-providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllIdpsReader{formats: a.formats},
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
	success, ok := result.(*GetAllIdpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllIdps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetIdentityProviderByID gets identity provider by Id

Get a specific Identity Provider using it's Id
*/
func (a *Client) GetIdentityProviderByID(params *GetIdentityProviderByIDParams, opts ...ClientOption) (*GetIdentityProviderByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdentityProviderByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIdentityProviderById",
		Method:             "GET",
		PathPattern:        "/v1/identity-providers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIdentityProviderByIDReader{formats: a.formats},
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
	success, ok := result.(*GetIdentityProviderByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIdentityProviderById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateEmbeddedIdentitySource updates an embedded identity source

Update the identity source associated with the embedded IDP by name, if it exists
*/
func (a *Client) UpdateEmbeddedIdentitySource(params *UpdateEmbeddedIdentitySourceParams, opts ...ClientOption) (*UpdateEmbeddedIdentitySourceOK, *UpdateEmbeddedIdentitySourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEmbeddedIdentitySourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateEmbeddedIdentitySource",
		Method:             "PATCH",
		PathPattern:        "/v1/identity-providers/{id}/identity-sources/{domainName}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateEmbeddedIdentitySourceReader{formats: a.formats},
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
	case *UpdateEmbeddedIdentitySourceOK:
		return value, nil, nil
	case *UpdateEmbeddedIdentitySourceNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateExternalIdentityProvider updates an external identity provider

Update the identity provider by its identifier, if it exists
*/
func (a *Client) UpdateExternalIdentityProvider(params *UpdateExternalIdentityProviderParams, opts ...ClientOption) (*UpdateExternalIdentityProviderOK, *UpdateExternalIdentityProviderNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalIdentityProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateExternalIdentityProvider",
		Method:             "PATCH",
		PathPattern:        "/v1/identity-providers/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalIdentityProviderReader{formats: a.formats},
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
	case *UpdateExternalIdentityProviderOK:
		return value, nil, nil
	case *UpdateExternalIdentityProviderNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for identity_providers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
