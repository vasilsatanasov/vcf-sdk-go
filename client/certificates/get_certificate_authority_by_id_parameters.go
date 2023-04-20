// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

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

// NewGETCertificateAuthorityByIDParams creates a new GETCertificateAuthorityByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETCertificateAuthorityByIDParams() *GETCertificateAuthorityByIDParams {
	return &GETCertificateAuthorityByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETCertificateAuthorityByIDParamsWithTimeout creates a new GETCertificateAuthorityByIDParams object
// with the ability to set a timeout on a request.
func NewGETCertificateAuthorityByIDParamsWithTimeout(timeout time.Duration) *GETCertificateAuthorityByIDParams {
	return &GETCertificateAuthorityByIDParams{
		timeout: timeout,
	}
}

// NewGETCertificateAuthorityByIDParamsWithContext creates a new GETCertificateAuthorityByIDParams object
// with the ability to set a context for a request.
func NewGETCertificateAuthorityByIDParamsWithContext(ctx context.Context) *GETCertificateAuthorityByIDParams {
	return &GETCertificateAuthorityByIDParams{
		Context: ctx,
	}
}

// NewGETCertificateAuthorityByIDParamsWithHTTPClient creates a new GETCertificateAuthorityByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETCertificateAuthorityByIDParamsWithHTTPClient(client *http.Client) *GETCertificateAuthorityByIDParams {
	return &GETCertificateAuthorityByIDParams{
		HTTPClient: client,
	}
}

/*
GETCertificateAuthorityByIDParams contains all the parameters to send to the API endpoint

	for the get certificate authority by Id operation.

	Typically these are written to a http.Request.
*/
type GETCertificateAuthorityByIDParams struct {

	/* ID.

	   CA type
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get certificate authority by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCertificateAuthorityByIDParams) WithDefaults() *GETCertificateAuthorityByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get certificate authority by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETCertificateAuthorityByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) WithTimeout(timeout time.Duration) *GETCertificateAuthorityByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) WithContext(ctx context.Context) *GETCertificateAuthorityByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) WithHTTPClient(client *http.Client) *GETCertificateAuthorityByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) WithID(id string) *GETCertificateAuthorityByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get certificate authority by Id params
func (o *GETCertificateAuthorityByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETCertificateAuthorityByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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