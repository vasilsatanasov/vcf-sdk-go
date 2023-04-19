// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgrades

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

// NewGETUpgradeByIDParams creates a new GETUpgradeByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETUpgradeByIDParams() *GETUpgradeByIDParams {
	return &GETUpgradeByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETUpgradeByIDParamsWithTimeout creates a new GETUpgradeByIDParams object
// with the ability to set a timeout on a request.
func NewGETUpgradeByIDParamsWithTimeout(timeout time.Duration) *GETUpgradeByIDParams {
	return &GETUpgradeByIDParams{
		timeout: timeout,
	}
}

// NewGETUpgradeByIDParamsWithContext creates a new GETUpgradeByIDParams object
// with the ability to set a context for a request.
func NewGETUpgradeByIDParamsWithContext(ctx context.Context) *GETUpgradeByIDParams {
	return &GETUpgradeByIDParams{
		Context: ctx,
	}
}

// NewGETUpgradeByIDParamsWithHTTPClient creates a new GETUpgradeByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETUpgradeByIDParamsWithHTTPClient(client *http.Client) *GETUpgradeByIDParams {
	return &GETUpgradeByIDParams{
		HTTPClient: client,
	}
}

/*
GETUpgradeByIDParams contains all the parameters to send to the API endpoint

	for the get upgrade by Id operation.

	Typically these are written to a http.Request.
*/
type GETUpgradeByIDParams struct {

	/* UpgradeID.

	   upgradeId
	*/
	UpgradeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgrade by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETUpgradeByIDParams) WithDefaults() *GETUpgradeByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgrade by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETUpgradeByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgrade by Id params
func (o *GETUpgradeByIDParams) WithTimeout(timeout time.Duration) *GETUpgradeByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgrade by Id params
func (o *GETUpgradeByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgrade by Id params
func (o *GETUpgradeByIDParams) WithContext(ctx context.Context) *GETUpgradeByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgrade by Id params
func (o *GETUpgradeByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgrade by Id params
func (o *GETUpgradeByIDParams) WithHTTPClient(client *http.Client) *GETUpgradeByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgrade by Id params
func (o *GETUpgradeByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpgradeID adds the upgradeID to the get upgrade by Id params
func (o *GETUpgradeByIDParams) WithUpgradeID(upgradeID string) *GETUpgradeByIDParams {
	o.SetUpgradeID(upgradeID)
	return o
}

// SetUpgradeID adds the upgradeId to the get upgrade by Id params
func (o *GETUpgradeByIDParams) SetUpgradeID(upgradeID string) {
	o.UpgradeID = upgradeID
}

// WriteToRequest writes these params to a swagger request
func (o *GETUpgradeByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param upgradeId
	if err := r.SetPathParam("upgradeId", o.UpgradeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
