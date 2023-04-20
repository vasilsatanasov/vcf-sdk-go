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

// NewGETClusterCriteriaParams creates a new GETClusterCriteriaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETClusterCriteriaParams() *GETClusterCriteriaParams {
	return &GETClusterCriteriaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETClusterCriteriaParamsWithTimeout creates a new GETClusterCriteriaParams object
// with the ability to set a timeout on a request.
func NewGETClusterCriteriaParamsWithTimeout(timeout time.Duration) *GETClusterCriteriaParams {
	return &GETClusterCriteriaParams{
		timeout: timeout,
	}
}

// NewGETClusterCriteriaParamsWithContext creates a new GETClusterCriteriaParams object
// with the ability to set a context for a request.
func NewGETClusterCriteriaParamsWithContext(ctx context.Context) *GETClusterCriteriaParams {
	return &GETClusterCriteriaParams{
		Context: ctx,
	}
}

// NewGETClusterCriteriaParamsWithHTTPClient creates a new GETClusterCriteriaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETClusterCriteriaParamsWithHTTPClient(client *http.Client) *GETClusterCriteriaParams {
	return &GETClusterCriteriaParams{
		HTTPClient: client,
	}
}

/*
GETClusterCriteriaParams contains all the parameters to send to the API endpoint

	for the get cluster criteria operation.

	Typically these are written to a http.Request.
*/
type GETClusterCriteriaParams struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterCriteriaParams) WithDefaults() *GETClusterCriteriaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster criteria params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETClusterCriteriaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster criteria params
func (o *GETClusterCriteriaParams) WithTimeout(timeout time.Duration) *GETClusterCriteriaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster criteria params
func (o *GETClusterCriteriaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster criteria params
func (o *GETClusterCriteriaParams) WithContext(ctx context.Context) *GETClusterCriteriaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster criteria params
func (o *GETClusterCriteriaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster criteria params
func (o *GETClusterCriteriaParams) WithHTTPClient(client *http.Client) *GETClusterCriteriaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster criteria params
func (o *GETClusterCriteriaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get cluster criteria params
func (o *GETClusterCriteriaParams) WithDomainID(domainID string) *GETClusterCriteriaParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get cluster criteria params
func (o *GETClusterCriteriaParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GETClusterCriteriaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}