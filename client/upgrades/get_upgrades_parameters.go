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

// NewGetUpgradesParams creates a new GetUpgradesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpgradesParams() *GetUpgradesParams {
	return &GetUpgradesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpgradesParamsWithTimeout creates a new GetUpgradesParams object
// with the ability to set a timeout on a request.
func NewGetUpgradesParamsWithTimeout(timeout time.Duration) *GetUpgradesParams {
	return &GetUpgradesParams{
		timeout: timeout,
	}
}

// NewGetUpgradesParamsWithContext creates a new GetUpgradesParams object
// with the ability to set a context for a request.
func NewGetUpgradesParamsWithContext(ctx context.Context) *GetUpgradesParams {
	return &GetUpgradesParams{
		Context: ctx,
	}
}

// NewGetUpgradesParamsWithHTTPClient creates a new GetUpgradesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpgradesParamsWithHTTPClient(client *http.Client) *GetUpgradesParams {
	return &GetUpgradesParams{
		HTTPClient: client,
	}
}

/*
GetUpgradesParams contains all the parameters to send to the API endpoint

	for the get upgrades operation.

	Typically these are written to a http.Request.
*/
type GetUpgradesParams struct {

	/* BundleID.

	   Bundle Id for the upgrade
	*/
	BundleID *string

	/* Status.

	   Status of the upgrades you want to retrieve
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradesParams) WithDefaults() *GetUpgradesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgrades params
func (o *GetUpgradesParams) WithTimeout(timeout time.Duration) *GetUpgradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgrades params
func (o *GetUpgradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgrades params
func (o *GetUpgradesParams) WithContext(ctx context.Context) *GetUpgradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgrades params
func (o *GetUpgradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgrades params
func (o *GetUpgradesParams) WithHTTPClient(client *http.Client) *GetUpgradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgrades params
func (o *GetUpgradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBundleID adds the bundleID to the get upgrades params
func (o *GetUpgradesParams) WithBundleID(bundleID *string) *GetUpgradesParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get upgrades params
func (o *GetUpgradesParams) SetBundleID(bundleID *string) {
	o.BundleID = bundleID
}

// WithStatus adds the status to the get upgrades params
func (o *GetUpgradesParams) WithStatus(status *string) *GetUpgradesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get upgrades params
func (o *GetUpgradesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpgradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BundleID != nil {

		// query param bundleId
		var qrBundleID string

		if o.BundleID != nil {
			qrBundleID = *o.BundleID
		}
		qBundleID := qrBundleID
		if qBundleID != "" {

			if err := r.SetQueryParam("bundleId", qBundleID); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
