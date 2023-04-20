// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGETCertificatesParams creates a new GETCertificatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETCertificatesParams() *GETCertificatesParams {
	return &GETCertificatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETCertificatesParamsWithTimeout creates a new GETCertificatesParams object
// with the ability to set a timeout on a request.
func NewGETCertificatesParamsWithTimeout(timeout time.Duration) *GETCertificatesParams {
	return &GETCertificatesParams{
		timeout: timeout,
	}
}

// NewGETCertificatesParamsWithContext creates a new GETCertificatesParams object
// with the ability to set a context for a request.
func NewGETCertificatesParamsWithContext(ctx context.Context) *GETCertificatesParams {
	return &GETCertificatesParams{
		Context: ctx,
	}
}

// NewGETCertificatesParamsWithHTTPClient creates a new GETCertificatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETCertificatesParamsWithHTTPClient(client *http.Client) *GETCertificatesParams {
	return &GETCertificatesParams{
		HTTPClient: client,
	}
}

/*
GETCertificatesParams contains all the parameters to send to the API endpoint

	for the get certificates operation.

	Typically these are written to a http.Request.
*/
type GETCertificatesParams struct {

	/* DomainName.

	   The domain name
	*/
	DomainName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCertificatesParams) WithDefaults() *GETCertificatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCertificatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get certificates params
func (o *GETCertificatesParams) WithTimeout(timeout time.Duration) *GETCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificates params
func (o *GETCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificates params
func (o *GETCertificatesParams) WithContext(ctx context.Context) *GETCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificates params
func (o *GETCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificates params
func (o *GETCertificatesParams) WithHTTPClient(client *http.Client) *GETCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificates params
func (o *GETCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainName adds the domainName to the get certificates params
func (o *GETCertificatesParams) WithDomainName(domainName string) *GETCertificatesParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the get certificates params
func (o *GETCertificatesParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *GETCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}