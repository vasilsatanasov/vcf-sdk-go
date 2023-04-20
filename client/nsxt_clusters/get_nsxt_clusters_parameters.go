// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package nsxt_clusters

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
	"github.com/go-openapi/swag"
)

// NewGETNSXTClustersParams creates a new GETNSXTClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETNSXTClustersParams() *GETNSXTClustersParams {
	return &GETNSXTClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETNSXTClustersParamsWithTimeout creates a new GETNSXTClustersParams object
// with the ability to set a timeout on a request.
func NewGETNSXTClustersParamsWithTimeout(timeout time.Duration) *GETNSXTClustersParams {
	return &GETNSXTClustersParams{
		timeout: timeout,
	}
}

// NewGETNSXTClustersParamsWithContext creates a new GETNSXTClustersParams object
// with the ability to set a context for a request.
func NewGETNSXTClustersParamsWithContext(ctx context.Context) *GETNSXTClustersParams {
	return &GETNSXTClustersParams{
		Context: ctx,
	}
}

// NewGETNSXTClustersParamsWithHTTPClient creates a new GETNSXTClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETNSXTClustersParamsWithHTTPClient(client *http.Client) *GETNSXTClustersParams {
	return &GETNSXTClustersParams{
		HTTPClient: client,
	}
}

/*
GETNSXTClustersParams contains all the parameters to send to the API endpoint

	for the get Nsxt clusters operation.

	Typically these are written to a http.Request.
*/
type GETNSXTClustersParams struct {

	/* IsShareable.

	   filter NSX-T clusters which can be shared for domain creation
	*/
	IsShareable *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Nsxt clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNSXTClustersParams) WithDefaults() *GETNSXTClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Nsxt clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETNSXTClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Nsxt clusters params
func (o *GETNSXTClustersParams) WithTimeout(timeout time.Duration) *GETNSXTClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Nsxt clusters params
func (o *GETNSXTClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Nsxt clusters params
func (o *GETNSXTClustersParams) WithContext(ctx context.Context) *GETNSXTClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Nsxt clusters params
func (o *GETNSXTClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Nsxt clusters params
func (o *GETNSXTClustersParams) WithHTTPClient(client *http.Client) *GETNSXTClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Nsxt clusters params
func (o *GETNSXTClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsShareable adds the isShareable to the get Nsxt clusters params
func (o *GETNSXTClustersParams) WithIsShareable(isShareable *bool) *GETNSXTClustersParams {
	o.SetIsShareable(isShareable)
	return o
}

// SetIsShareable adds the isShareable to the get Nsxt clusters params
func (o *GETNSXTClustersParams) SetIsShareable(isShareable *bool) {
	o.IsShareable = isShareable
}

// WriteToRequest writes these params to a swagger request
func (o *GETNSXTClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsShareable != nil {

		// query param isShareable
		var qrIsShareable bool

		if o.IsShareable != nil {
			qrIsShareable = *o.IsShareable
		}
		qIsShareable := swag.FormatBool(qrIsShareable)
		if qIsShareable != "" {

			if err := r.SetQueryParam("isShareable", qIsShareable); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}