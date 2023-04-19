// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new certificates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for certificates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCaConfiguration(params *DeleteCaConfigurationParams, opts ...ClientOption) (*DeleteCaConfigurationOK, *DeleteCaConfigurationNoContent, error)

	ConfigureCertificateAuthority(params *ConfigureCertificateAuthorityParams, opts ...ClientOption) (*ConfigureCertificateAuthorityOK, error)

	CreateCertificateAuthority(params *CreateCertificateAuthorityParams, opts ...ClientOption) (*CreateCertificateAuthorityOK, error)

	DownloadCSR(params *DownloadCSRParams, opts ...ClientOption) (*DownloadCSROK, error)

	GenerateCertificates(params *GenerateCertificatesParams, opts ...ClientOption) (*GenerateCertificatesOK, *GenerateCertificatesAccepted, error)

	GeneratesCSRs(params *GeneratesCSRsParams, opts ...ClientOption) (*GeneratesCSRsOK, *GeneratesCSRsAccepted, error)

	GETCSRs(params *GETCSRsParams, opts ...ClientOption) (*GETCSRsOK, error)

	GETCertificateAuthorities(params *GETCertificateAuthoritiesParams, opts ...ClientOption) (*GETCertificateAuthoritiesOK, error)

	GETCertificateAuthorityByID(params *GETCertificateAuthorityByIDParams, opts ...ClientOption) (*GETCertificateAuthorityByIDOK, error)

	GETCertificates(params *GETCertificatesParams, opts ...ClientOption) (*GETCertificatesOK, error)

	ReplaceCertificates(params *ReplaceCertificatesParams, opts ...ClientOption) (*ReplaceCertificatesOK, *ReplaceCertificatesAccepted, error)

	UploadCertificates(params *UploadCertificatesParams, opts ...ClientOption) (*UploadCertificatesOK, error)

	ViewCertificate(params *ViewCertificateParams, opts ...ClientOption) (*ViewCertificateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteCaConfiguration deletes c a configuration file

Deletes CA configuration file
*/
func (a *Client) DeleteCaConfiguration(params *DeleteCaConfigurationParams, opts ...ClientOption) (*DeleteCaConfigurationOK, *DeleteCaConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCaConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCaConfiguration",
		Method:             "DELETE",
		PathPattern:        "/v1/certificate-authorities/{caType}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCaConfigurationReader{formats: a.formats},
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
	case *DeleteCaConfigurationOK:
		return value, nil, nil
	case *DeleteCaConfigurationNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ConfigureCertificateAuthority configures existing certificate authority

Configure existing certificate authority
*/
func (a *Client) ConfigureCertificateAuthority(params *ConfigureCertificateAuthorityParams, opts ...ClientOption) (*ConfigureCertificateAuthorityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigureCertificateAuthorityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configureCertificateAuthority",
		Method:             "PATCH",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfigureCertificateAuthorityReader{formats: a.formats},
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
	success, ok := result.(*ConfigureCertificateAuthorityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for configureCertificateAuthority: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateCertificateAuthority creates a certificate authority

Creates a certificate authority. This is required to generate signed certificates by supporting CAs.
*/
func (a *Client) CreateCertificateAuthority(params *CreateCertificateAuthorityParams, opts ...ClientOption) (*CreateCertificateAuthorityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCertificateAuthorityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCertificateAuthority",
		Method:             "PUT",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCertificateAuthorityReader{formats: a.formats},
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
	success, ok := result.(*CreateCertificateAuthorityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCertificateAuthority: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadCSR downloads available CSR s in tar gz format

Download available CSR(s) in tar.gz format
*/
func (a *Client) DownloadCSR(params *DownloadCSRParams, opts ...ClientOption) (*DownloadCSROK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadCSRParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadCSR",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainName}/csrs/downloads",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadCSRReader{formats: a.formats},
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
	success, ok := result.(*DownloadCSROK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadCSR: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GenerateCertificates generates certificate s for the selected resource s in a domain

Generate certificate(s) for the selected resource(s) in a domain. CA must be configured and CSR must be generated beforehand.
*/
func (a *Client) GenerateCertificates(params *GenerateCertificatesParams, opts ...ClientOption) (*GenerateCertificatesOK, *GenerateCertificatesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generateCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{domainName}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateCertificatesReader{formats: a.formats},
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
	case *GenerateCertificatesOK:
		return value, nil, nil
	case *GenerateCertificatesAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GeneratesCSRs generates CSR s

	Generate CSR(s) for the selected resource(s) in the domain.

*Warning:*
_Avoid using wildcard certificates. Instead, use subdomain-specific certificates that are rotated often. A compromised wildcard certificate can lead to security repercussions._
*/
func (a *Client) GeneratesCSRs(params *GeneratesCSRsParams, opts ...ClientOption) (*GeneratesCSRsOK, *GeneratesCSRsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeneratesCSRsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "generatesCSRs",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{domainName}/csrs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeneratesCSRsReader{formats: a.formats},
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
	case *GeneratesCSRsOK:
		return value, nil, nil
	case *GeneratesCSRsAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCSRs Gets available CSR s in json format

Get available CSR(s) in json format
*/
func (a *Client) GETCSRs(params *GETCSRsParams, opts ...ClientOption) (*GETCSRsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCSRsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCSRs",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainName}/csrs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCSRsReader{formats: a.formats},
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
	success, ok := result.(*GETCSRsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCSRs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCertificateAuthorities Gets certificate authorities information

Get certificate authorities information
*/
func (a *Client) GETCertificateAuthorities(params *GETCertificateAuthoritiesParams, opts ...ClientOption) (*GETCertificateAuthoritiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCertificateAuthoritiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificateAuthorities",
		Method:             "GET",
		PathPattern:        "/v1/certificate-authorities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCertificateAuthoritiesReader{formats: a.formats},
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
	success, ok := result.(*GETCertificateAuthoritiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificateAuthorities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCertificateAuthorityByID Gets certificate authority information

Get certificate authority information
*/
func (a *Client) GETCertificateAuthorityByID(params *GETCertificateAuthorityByIDParams, opts ...ClientOption) (*GETCertificateAuthorityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCertificateAuthorityByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificateAuthorityById",
		Method:             "GET",
		PathPattern:        "/v1/certificate-authorities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCertificateAuthorityByIDReader{formats: a.formats},
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
	success, ok := result.(*GETCertificateAuthorityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificateAuthorityById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GETCertificates Gets latest generated certificate s in a domain

Get latest generated certificate(s) in a domain.
*/
func (a *Client) GETCertificates(params *GETCertificatesParams, opts ...ClientOption) (*GETCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCertificates",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainName}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GETCertificatesReader{formats: a.formats},
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
	success, ok := result.(*GETCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReplaceCertificates replaces certificate s for the selected resource s in a domain

Replace certificate(s) for the selected resource(s) in a domain
*/
func (a *Client) ReplaceCertificates(params *ReplaceCertificatesParams, opts ...ClientOption) (*ReplaceCertificatesOK, *ReplaceCertificatesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceCertificates",
		Method:             "PATCH",
		PathPattern:        "/v1/domains/{domainName}/certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceCertificatesReader{formats: a.formats},
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
	case *ReplaceCertificatesOK:
		return value, nil, nil
	case *ReplaceCertificatesAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadCertificates uploads certificates to the certificate store

Upload certificates to the certificate store
*/
func (a *Client) UploadCertificates(params *UploadCertificatesParams, opts ...ClientOption) (*UploadCertificatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadCertificatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadCertificates",
		Method:             "PUT",
		PathPattern:        "/v1/domains/{domainName}/certificates/uploads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadCertificatesReader{formats: a.formats},
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
	success, ok := result.(*UploadCertificatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadCertificates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ViewCertificate views certificate of all the resources in a domain

View detailed metadata about the certificate(s) of all the resources in a domain
*/
func (a *Client) ViewCertificate(params *ViewCertificateParams, opts ...ClientOption) (*ViewCertificateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewCertificateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "viewCertificate",
		Method:             "GET",
		PathPattern:        "/v1/domains/{domainName}/resource-certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ViewCertificateReader{formats: a.formats},
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
	success, ok := result.(*ViewCertificateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for viewCertificate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
