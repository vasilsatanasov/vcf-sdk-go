// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package network_pools

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

// NewGETNetworkPoolsParams creates a new GETNetworkPoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETNetworkPoolsParams() *GETNetworkPoolsParams {
	return &GETNetworkPoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETNetworkPoolsParamsWithTimeout creates a new GETNetworkPoolsParams object
// with the ability to set a timeout on a request.
func NewGETNetworkPoolsParamsWithTimeout(timeout time.Duration) *GETNetworkPoolsParams {
	return &GETNetworkPoolsParams{
		timeout: timeout,
	}
}

// NewGETNetworkPoolsParamsWithContext creates a new GETNetworkPoolsParams object
// with the ability to set a context for a request.
func NewGETNetworkPoolsParamsWithContext(ctx context.Context) *GETNetworkPoolsParams {
	return &GETNetworkPoolsParams{
		Context: ctx,
	}
}

// NewGETNetworkPoolsParamsWithHTTPClient creates a new GETNetworkPoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETNetworkPoolsParamsWithHTTPClient(client *http.Client) *GETNetworkPoolsParams {
	return &GETNetworkPoolsParams{
		HTTPClient: client,
	}
}

/*
GETNetworkPoolsParams contains all the parameters to send to the API endpoint

	for the get network pools operation.

	Typically these are written to a http.Request.
*/
type GETNetworkPoolsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNetworkPoolsParams) WithDefaults() *GETNetworkPoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNetworkPoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network pools params
func (o *GETNetworkPoolsParams) WithTimeout(timeout time.Duration) *GETNetworkPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network pools params
func (o *GETNetworkPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network pools params
func (o *GETNetworkPoolsParams) WithContext(ctx context.Context) *GETNetworkPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network pools params
func (o *GETNetworkPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network pools params
func (o *GETNetworkPoolsParams) WithHTTPClient(client *http.Client) *GETNetworkPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network pools params
func (o *GETNetworkPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETNetworkPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}