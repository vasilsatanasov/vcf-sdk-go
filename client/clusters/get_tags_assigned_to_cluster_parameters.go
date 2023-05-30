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

// NewGetTagsAssignedToClusterParams creates a new GetTagsAssignedToClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTagsAssignedToClusterParams() *GetTagsAssignedToClusterParams {
	return &GetTagsAssignedToClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagsAssignedToClusterParamsWithTimeout creates a new GetTagsAssignedToClusterParams object
// with the ability to set a timeout on a request.
func NewGetTagsAssignedToClusterParamsWithTimeout(timeout time.Duration) *GetTagsAssignedToClusterParams {
	return &GetTagsAssignedToClusterParams{
		timeout: timeout,
	}
}

// NewGetTagsAssignedToClusterParamsWithContext creates a new GetTagsAssignedToClusterParams object
// with the ability to set a context for a request.
func NewGetTagsAssignedToClusterParamsWithContext(ctx context.Context) *GetTagsAssignedToClusterParams {
	return &GetTagsAssignedToClusterParams{
		Context: ctx,
	}
}

// NewGetTagsAssignedToClusterParamsWithHTTPClient creates a new GetTagsAssignedToClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTagsAssignedToClusterParamsWithHTTPClient(client *http.Client) *GetTagsAssignedToClusterParams {
	return &GetTagsAssignedToClusterParams{
		HTTPClient: client,
	}
}

/*
GetTagsAssignedToClusterParams contains all the parameters to send to the API endpoint

	for the get tags assigned to cluster operation.

	Typically these are written to a http.Request.
*/
type GetTagsAssignedToClusterParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tags assigned to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagsAssignedToClusterParams) WithDefaults() *GetTagsAssignedToClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tags assigned to cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagsAssignedToClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) WithTimeout(timeout time.Duration) *GetTagsAssignedToClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) WithContext(ctx context.Context) *GetTagsAssignedToClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) WithHTTPClient(client *http.Client) *GetTagsAssignedToClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) WithID(id string) *GetTagsAssignedToClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get tags assigned to cluster params
func (o *GetTagsAssignedToClusterParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagsAssignedToClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
