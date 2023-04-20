// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewDecommissionHostsParams creates a new DecommissionHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDecommissionHostsParams() *DecommissionHostsParams {
	return &DecommissionHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDecommissionHostsParamsWithTimeout creates a new DecommissionHostsParams object
// with the ability to set a timeout on a request.
func NewDecommissionHostsParamsWithTimeout(timeout time.Duration) *DecommissionHostsParams {
	return &DecommissionHostsParams{
		timeout: timeout,
	}
}

// NewDecommissionHostsParamsWithContext creates a new DecommissionHostsParams object
// with the ability to set a context for a request.
func NewDecommissionHostsParamsWithContext(ctx context.Context) *DecommissionHostsParams {
	return &DecommissionHostsParams{
		Context: ctx,
	}
}

// NewDecommissionHostsParamsWithHTTPClient creates a new DecommissionHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDecommissionHostsParamsWithHTTPClient(client *http.Client) *DecommissionHostsParams {
	return &DecommissionHostsParams{
		HTTPClient: client,
	}
}

/*
DecommissionHostsParams contains all the parameters to send to the API endpoint

	for the decommission hosts operation.

	Typically these are written to a http.Request.
*/
type DecommissionHostsParams struct {

	/* HostDecommissionSpecs.

	   hostDecommissionSpecs
	*/
	HostDecommissionSpecs []*models.HostDecommissionSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the decommission hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecommissionHostsParams) WithDefaults() *DecommissionHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the decommission hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DecommissionHostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the decommission hosts params
func (o *DecommissionHostsParams) WithTimeout(timeout time.Duration) *DecommissionHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the decommission hosts params
func (o *DecommissionHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the decommission hosts params
func (o *DecommissionHostsParams) WithContext(ctx context.Context) *DecommissionHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the decommission hosts params
func (o *DecommissionHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the decommission hosts params
func (o *DecommissionHostsParams) WithHTTPClient(client *http.Client) *DecommissionHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the decommission hosts params
func (o *DecommissionHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostDecommissionSpecs adds the hostDecommissionSpecs to the decommission hosts params
func (o *DecommissionHostsParams) WithHostDecommissionSpecs(hostDecommissionSpecs []*models.HostDecommissionSpec) *DecommissionHostsParams {
	o.SetHostDecommissionSpecs(hostDecommissionSpecs)
	return o
}

// SetHostDecommissionSpecs adds the hostDecommissionSpecs to the decommission hosts params
func (o *DecommissionHostsParams) SetHostDecommissionSpecs(hostDecommissionSpecs []*models.HostDecommissionSpec) {
	o.HostDecommissionSpecs = hostDecommissionSpecs
}

// WriteToRequest writes these params to a swagger request
func (o *DecommissionHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.HostDecommissionSpecs != nil {
		if err := r.SetBodyParam(o.HostDecommissionSpecs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}