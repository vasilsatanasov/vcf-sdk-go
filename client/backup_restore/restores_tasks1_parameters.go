// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

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

// NewRestoresTasks1Params creates a new RestoresTasks1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoresTasks1Params() *RestoresTasks1Params {
	return &RestoresTasks1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoresTasks1ParamsWithTimeout creates a new RestoresTasks1Params object
// with the ability to set a timeout on a request.
func NewRestoresTasks1ParamsWithTimeout(timeout time.Duration) *RestoresTasks1Params {
	return &RestoresTasks1Params{
		timeout: timeout,
	}
}

// NewRestoresTasks1ParamsWithContext creates a new RestoresTasks1Params object
// with the ability to set a context for a request.
func NewRestoresTasks1ParamsWithContext(ctx context.Context) *RestoresTasks1Params {
	return &RestoresTasks1Params{
		Context: ctx,
	}
}

// NewRestoresTasks1ParamsWithHTTPClient creates a new RestoresTasks1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRestoresTasks1ParamsWithHTTPClient(client *http.Client) *RestoresTasks1Params {
	return &RestoresTasks1Params{
		HTTPClient: client,
	}
}

/*
RestoresTasks1Params contains all the parameters to send to the API endpoint

	for the restores tasks 1 operation.

	Typically these are written to a http.Request.
*/
type RestoresTasks1Params struct {

	/* RestoreSpec.

	   restoreSpec
	*/
	RestoreSpec *models.RestoreSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restores tasks 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoresTasks1Params) WithDefaults() *RestoresTasks1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restores tasks 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoresTasks1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restores tasks 1 params
func (o *RestoresTasks1Params) WithTimeout(timeout time.Duration) *RestoresTasks1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restores tasks 1 params
func (o *RestoresTasks1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restores tasks 1 params
func (o *RestoresTasks1Params) WithContext(ctx context.Context) *RestoresTasks1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restores tasks 1 params
func (o *RestoresTasks1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restores tasks 1 params
func (o *RestoresTasks1Params) WithHTTPClient(client *http.Client) *RestoresTasks1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restores tasks 1 params
func (o *RestoresTasks1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRestoreSpec adds the restoreSpec to the restores tasks 1 params
func (o *RestoresTasks1Params) WithRestoreSpec(restoreSpec *models.RestoreSpec) *RestoresTasks1Params {
	o.SetRestoreSpec(restoreSpec)
	return o
}

// SetRestoreSpec adds the restoreSpec to the restores tasks 1 params
func (o *RestoresTasks1Params) SetRestoreSpec(restoreSpec *models.RestoreSpec) {
	o.RestoreSpec = restoreSpec
}

// WriteToRequest writes these params to a swagger request
func (o *RestoresTasks1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RestoreSpec != nil {
		if err := r.SetBodyParam(o.RestoreSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
