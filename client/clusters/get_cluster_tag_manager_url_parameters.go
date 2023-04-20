// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

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

// NewGETClusterTagManagerURLParams creates a new GETClusterTagManagerURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETClusterTagManagerURLParams() *GETClusterTagManagerURLParams {
	return &GETClusterTagManagerURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETClusterTagManagerURLParamsWithTimeout creates a new GETClusterTagManagerURLParams object
// with the ability to set a timeout on a request.
func NewGETClusterTagManagerURLParamsWithTimeout(timeout time.Duration) *GETClusterTagManagerURLParams {
	return &GETClusterTagManagerURLParams{
		timeout: timeout,
	}
}

// NewGETClusterTagManagerURLParamsWithContext creates a new GETClusterTagManagerURLParams object
// with the ability to set a context for a request.
func NewGETClusterTagManagerURLParamsWithContext(ctx context.Context) *GETClusterTagManagerURLParams {
	return &GETClusterTagManagerURLParams{
		Context: ctx,
	}
}

// NewGETClusterTagManagerURLParamsWithHTTPClient creates a new GETClusterTagManagerURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETClusterTagManagerURLParamsWithHTTPClient(client *http.Client) *GETClusterTagManagerURLParams {
	return &GETClusterTagManagerURLParams{
		HTTPClient: client,
	}
}

/*
GETClusterTagManagerURLParams contains all the parameters to send to the API endpoint

	for the get cluster tag manager Url operation.

	Typically these are written to a http.Request.
*/
type GETClusterTagManagerURLParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterTagManagerURLParams) WithDefaults() *GETClusterTagManagerURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster tag manager Url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterTagManagerURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) WithTimeout(timeout time.Duration) *GETClusterTagManagerURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) WithContext(ctx context.Context) *GETClusterTagManagerURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) WithHTTPClient(client *http.Client) *GETClusterTagManagerURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) WithID(id string) *GETClusterTagManagerURLParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cluster tag manager Url params
func (o *GETClusterTagManagerURLParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETClusterTagManagerURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}