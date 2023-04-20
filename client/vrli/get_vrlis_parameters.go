// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vrli

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

// NewGETVrlisParams creates a new GETVrlisParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETVrlisParams() *GETVrlisParams {
	return &GETVrlisParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETVrlisParamsWithTimeout creates a new GETVrlisParams object
// with the ability to set a timeout on a request.
func NewGETVrlisParamsWithTimeout(timeout time.Duration) *GETVrlisParams {
	return &GETVrlisParams{
		timeout: timeout,
	}
}

// NewGETVrlisParamsWithContext creates a new GETVrlisParams object
// with the ability to set a context for a request.
func NewGETVrlisParamsWithContext(ctx context.Context) *GETVrlisParams {
	return &GETVrlisParams{
		Context: ctx,
	}
}

// NewGETVrlisParamsWithHTTPClient creates a new GETVrlisParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETVrlisParamsWithHTTPClient(client *http.Client) *GETVrlisParams {
	return &GETVrlisParams{
		HTTPClient: client,
	}
}

/*
GETVrlisParams contains all the parameters to send to the API endpoint

	for the get vrlis operation.

	Typically these are written to a http.Request.
*/
type GETVrlisParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vrlis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVrlisParams) WithDefaults() *GETVrlisParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vrlis params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETVrlisParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vrlis params
func (o *GETVrlisParams) WithTimeout(timeout time.Duration) *GETVrlisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vrlis params
func (o *GETVrlisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vrlis params
func (o *GETVrlisParams) WithContext(ctx context.Context) *GETVrlisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vrlis params
func (o *GETVrlisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vrlis params
func (o *GETVrlisParams) WithHTTPClient(client *http.Client) *GETVrlisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vrlis params
func (o *GETVrlisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GETVrlisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}