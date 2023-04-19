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

// NewGETDatastoreCriterionParams creates a new GETDatastoreCriterionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETDatastoreCriterionParams() *GETDatastoreCriterionParams {
	return &GETDatastoreCriterionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETDatastoreCriterionParamsWithTimeout creates a new GETDatastoreCriterionParams object
// with the ability to set a timeout on a request.
func NewGETDatastoreCriterionParamsWithTimeout(timeout time.Duration) *GETDatastoreCriterionParams {
	return &GETDatastoreCriterionParams{
		timeout: timeout,
	}
}

// NewGETDatastoreCriterionParamsWithContext creates a new GETDatastoreCriterionParams object
// with the ability to set a context for a request.
func NewGETDatastoreCriterionParamsWithContext(ctx context.Context) *GETDatastoreCriterionParams {
	return &GETDatastoreCriterionParams{
		Context: ctx,
	}
}

// NewGETDatastoreCriterionParamsWithHTTPClient creates a new GETDatastoreCriterionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETDatastoreCriterionParamsWithHTTPClient(client *http.Client) *GETDatastoreCriterionParams {
	return &GETDatastoreCriterionParams{
		HTTPClient: client,
	}
}

/*
GETDatastoreCriterionParams contains all the parameters to send to the API endpoint

	for the get datastore criterion operation.

	Typically these are written to a http.Request.
*/
type GETDatastoreCriterionParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	/* Name.

	   Criteria Name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get datastore criterion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETDatastoreCriterionParams) WithDefaults() *GETDatastoreCriterionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get datastore criterion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETDatastoreCriterionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get datastore criterion params
func (o *GETDatastoreCriterionParams) WithTimeout(timeout time.Duration) *GETDatastoreCriterionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get datastore criterion params
func (o *GETDatastoreCriterionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get datastore criterion params
func (o *GETDatastoreCriterionParams) WithContext(ctx context.Context) *GETDatastoreCriterionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get datastore criterion params
func (o *GETDatastoreCriterionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get datastore criterion params
func (o *GETDatastoreCriterionParams) WithHTTPClient(client *http.Client) *GETDatastoreCriterionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get datastore criterion params
func (o *GETDatastoreCriterionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get datastore criterion params
func (o *GETDatastoreCriterionParams) WithID(id string) *GETDatastoreCriterionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get datastore criterion params
func (o *GETDatastoreCriterionParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the get datastore criterion params
func (o *GETDatastoreCriterionParams) WithName(name string) *GETDatastoreCriterionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get datastore criterion params
func (o *GETDatastoreCriterionParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GETDatastoreCriterionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
