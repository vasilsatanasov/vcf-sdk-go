// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

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

// NewGETClusterQueryResponseParams creates a new GETClusterQueryResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETClusterQueryResponseParams() *GETClusterQueryResponseParams {
	return &GETClusterQueryResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETClusterQueryResponseParamsWithTimeout creates a new GETClusterQueryResponseParams object
// with the ability to set a timeout on a request.
func NewGETClusterQueryResponseParamsWithTimeout(timeout time.Duration) *GETClusterQueryResponseParams {
	return &GETClusterQueryResponseParams{
		timeout: timeout,
	}
}

// NewGETClusterQueryResponseParamsWithContext creates a new GETClusterQueryResponseParams object
// with the ability to set a context for a request.
func NewGETClusterQueryResponseParamsWithContext(ctx context.Context) *GETClusterQueryResponseParams {
	return &GETClusterQueryResponseParams{
		Context: ctx,
	}
}

// NewGETClusterQueryResponseParamsWithHTTPClient creates a new GETClusterQueryResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETClusterQueryResponseParamsWithHTTPClient(client *http.Client) *GETClusterQueryResponseParams {
	return &GETClusterQueryResponseParams{
		HTTPClient: client,
	}
}

/*
GETClusterQueryResponseParams contains all the parameters to send to the API endpoint

	for the get cluster query response operation.

	Typically these are written to a http.Request.
*/
type GETClusterQueryResponseParams struct {

	/* ClusterName.

	   Cluster Name
	*/
	ClusterName string

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* QueryID.

	   Query ID
	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterQueryResponseParams) WithDefaults() *GETClusterQueryResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster query response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterQueryResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithTimeout(timeout time.Duration) *GETClusterQueryResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithContext(ctx context.Context) *GETClusterQueryResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithHTTPClient(client *http.Client) *GETClusterQueryResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithClusterName(clusterName string) *GETClusterQueryResponseParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithDomainID adds the domainID to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithDomainID(domainID string) *GETClusterQueryResponseParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithQueryID adds the queryID to the get cluster query response params
func (o *GETClusterQueryResponseParams) WithQueryID(queryID string) *GETClusterQueryResponseParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the get cluster query response params
func (o *GETClusterQueryResponseParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *GETClusterQueryResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}