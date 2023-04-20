// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package resource_warnings

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

// NewGETResourceWarningParams creates a new GETResourceWarningParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETResourceWarningParams() *GETResourceWarningParams {
	return &GETResourceWarningParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETResourceWarningParamsWithTimeout creates a new GETResourceWarningParams object
// with the ability to set a timeout on a request.
func NewGETResourceWarningParamsWithTimeout(timeout time.Duration) *GETResourceWarningParams {
	return &GETResourceWarningParams{
		timeout: timeout,
	}
}

// NewGETResourceWarningParamsWithContext creates a new GETResourceWarningParams object
// with the ability to set a context for a request.
func NewGETResourceWarningParamsWithContext(ctx context.Context) *GETResourceWarningParams {
	return &GETResourceWarningParams{
		Context: ctx,
	}
}

// NewGETResourceWarningParamsWithHTTPClient creates a new GETResourceWarningParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETResourceWarningParamsWithHTTPClient(client *http.Client) *GETResourceWarningParams {
	return &GETResourceWarningParams{
		HTTPClient: client,
	}
}

/*
GETResourceWarningParams contains all the parameters to send to the API endpoint

	for the get resource warning operation.

	Typically these are written to a http.Request.
*/
type GETResourceWarningParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource warning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETResourceWarningParams) WithDefaults() *GETResourceWarningParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource warning params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETResourceWarningParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource warning params
func (o *GETResourceWarningParams) WithTimeout(timeout time.Duration) *GETResourceWarningParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource warning params
func (o *GETResourceWarningParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource warning params
func (o *GETResourceWarningParams) WithContext(ctx context.Context) *GETResourceWarningParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource warning params
func (o *GETResourceWarningParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource warning params
func (o *GETResourceWarningParams) WithHTTPClient(client *http.Client) *GETResourceWarningParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource warning params
func (o *GETResourceWarningParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get resource warning params
func (o *GETResourceWarningParams) WithID(id string) *GETResourceWarningParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get resource warning params
func (o *GETResourceWarningParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETResourceWarningParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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