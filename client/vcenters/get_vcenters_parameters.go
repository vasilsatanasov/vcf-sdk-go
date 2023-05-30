// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vcenters

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

// NewGetVCENTERSParams creates a new GetVCENTERSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVCENTERSParams() *GetVCENTERSParams {
	return &GetVCENTERSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVCENTERSParamsWithTimeout creates a new GetVCENTERSParams object
// with the ability to set a timeout on a request.
func NewGetVCENTERSParamsWithTimeout(timeout time.Duration) *GetVCENTERSParams {
	return &GetVCENTERSParams{
		timeout: timeout,
	}
}

// NewGetVCENTERSParamsWithContext creates a new GetVCENTERSParams object
// with the ability to set a context for a request.
func NewGetVCENTERSParamsWithContext(ctx context.Context) *GetVCENTERSParams {
	return &GetVCENTERSParams{
		Context: ctx,
	}
}

// NewGetVCENTERSParamsWithHTTPClient creates a new GetVCENTERSParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVCENTERSParamsWithHTTPClient(client *http.Client) *GetVCENTERSParams {
	return &GetVCENTERSParams{
		HTTPClient: client,
	}
}

/*
GetVCENTERSParams contains all the parameters to send to the API endpoint

	for the get Vcenters operation.

	Typically these are written to a http.Request.
*/
type GetVCENTERSParams struct {

	/* DomainID.

	   ID of the domain
	*/
	DomainID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vcenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVCENTERSParams) WithDefaults() *GetVCENTERSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vcenters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVCENTERSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Vcenters params
func (o *GetVCENTERSParams) WithTimeout(timeout time.Duration) *GetVCENTERSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vcenters params
func (o *GetVCENTERSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vcenters params
func (o *GetVCENTERSParams) WithContext(ctx context.Context) *GetVCENTERSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vcenters params
func (o *GetVCENTERSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vcenters params
func (o *GetVCENTERSParams) WithHTTPClient(client *http.Client) *GetVCENTERSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vcenters params
func (o *GetVCENTERSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get Vcenters params
func (o *GetVCENTERSParams) WithDomainID(domainID *string) *GetVCENTERSParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get Vcenters params
func (o *GetVCENTERSParams) SetDomainID(domainID *string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVCENTERSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DomainID != nil {

		// query param domainId
		var qrDomainID string

		if o.DomainID != nil {
			qrDomainID = *o.DomainID
		}
		qDomainID := qrDomainID
		if qDomainID != "" {

			if err := r.SetQueryParam("domainId", qDomainID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
