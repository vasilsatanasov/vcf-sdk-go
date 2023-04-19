// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sos

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

// NewHealthsummarydataParams creates a new HealthsummarydataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHealthsummarydataParams() *HealthsummarydataParams {
	return &HealthsummarydataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHealthsummarydataParamsWithTimeout creates a new HealthsummarydataParams object
// with the ability to set a timeout on a request.
func NewHealthsummarydataParamsWithTimeout(timeout time.Duration) *HealthsummarydataParams {
	return &HealthsummarydataParams{
		timeout: timeout,
	}
}

// NewHealthsummarydataParamsWithContext creates a new HealthsummarydataParams object
// with the ability to set a context for a request.
func NewHealthsummarydataParamsWithContext(ctx context.Context) *HealthsummarydataParams {
	return &HealthsummarydataParams{
		Context: ctx,
	}
}

// NewHealthsummarydataParamsWithHTTPClient creates a new HealthsummarydataParams object
// with the ability to set a custom HTTPClient for a request.
func NewHealthsummarydataParamsWithHTTPClient(client *http.Client) *HealthsummarydataParams {
	return &HealthsummarydataParams{
		HTTPClient: client,
	}
}

/*
HealthsummarydataParams contains all the parameters to send to the API endpoint

	for the healthsummarydata operation.

	Typically these are written to a http.Request.
*/
type HealthsummarydataParams struct {

	/* ID.

	   The Health Summary Id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the healthsummarydata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HealthsummarydataParams) WithDefaults() *HealthsummarydataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the healthsummarydata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HealthsummarydataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the healthsummarydata params
func (o *HealthsummarydataParams) WithTimeout(timeout time.Duration) *HealthsummarydataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the healthsummarydata params
func (o *HealthsummarydataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the healthsummarydata params
func (o *HealthsummarydataParams) WithContext(ctx context.Context) *HealthsummarydataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the healthsummarydata params
func (o *HealthsummarydataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the healthsummarydata params
func (o *HealthsummarydataParams) WithHTTPClient(client *http.Client) *HealthsummarydataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the healthsummarydata params
func (o *HealthsummarydataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the healthsummarydata params
func (o *HealthsummarydataParams) WithID(id string) *HealthsummarydataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the healthsummarydata params
func (o *HealthsummarydataParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HealthsummarydataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
