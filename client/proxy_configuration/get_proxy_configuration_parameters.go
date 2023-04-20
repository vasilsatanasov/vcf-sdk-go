// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package proxy_configuration

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

// NewGETProxyConfigurationParams creates a new GETProxyConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETProxyConfigurationParams() *GETProxyConfigurationParams {
	return &GETProxyConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETProxyConfigurationParamsWithTimeout creates a new GETProxyConfigurationParams object
// with the ability to set a timeout on a request.
func NewGETProxyConfigurationParamsWithTimeout(timeout time.Duration) *GETProxyConfigurationParams {
	return &GETProxyConfigurationParams{
		timeout: timeout,
	}
}

// NewGETProxyConfigurationParamsWithContext creates a new GETProxyConfigurationParams object
// with the ability to set a context for a request.
func NewGETProxyConfigurationParamsWithContext(ctx context.Context) *GETProxyConfigurationParams {
	return &GETProxyConfigurationParams{
		Context: ctx,
	}
}

// NewGETProxyConfigurationParamsWithHTTPClient creates a new GETProxyConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETProxyConfigurationParamsWithHTTPClient(client *http.Client) *GETProxyConfigurationParams {
	return &GETProxyConfigurationParams{
		HTTPClient: client,
	}
}

/*
GETProxyConfigurationParams contains all the parameters to send to the API endpoint

	for the get proxy configuration operation.

	Typically these are written to a http.Request.
*/
type GETProxyConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETProxyConfigurationParams) WithDefaults() *GETProxyConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get proxy configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETProxyConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get proxy configuration params
func (o *GETProxyConfigurationParams) WithTimeout(timeout time.Duration) *GETProxyConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxy configuration params
func (o *GETProxyConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxy configuration params
func (o *GETProxyConfigurationParams) WithContext(ctx context.Context) *GETProxyConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxy configuration params
func (o *GETProxyConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxy configuration params
func (o *GETProxyConfigurationParams) WithHTTPClient(client *http.Client) *GETProxyConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxy configuration params
func (o *GETProxyConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETProxyConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}