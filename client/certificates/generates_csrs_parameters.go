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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewGeneratesCSRsParams creates a new GeneratesCSRsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGeneratesCSRsParams() *GeneratesCSRsParams {
	return &GeneratesCSRsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGeneratesCSRsParamsWithTimeout creates a new GeneratesCSRsParams object
// with the ability to set a timeout on a request.
func NewGeneratesCSRsParamsWithTimeout(timeout time.Duration) *GeneratesCSRsParams {
	return &GeneratesCSRsParams{
		timeout: timeout,
	}
}

// NewGeneratesCSRsParamsWithContext creates a new GeneratesCSRsParams object
// with the ability to set a context for a request.
func NewGeneratesCSRsParamsWithContext(ctx context.Context) *GeneratesCSRsParams {
	return &GeneratesCSRsParams{
		Context: ctx,
	}
}

// NewGeneratesCSRsParamsWithHTTPClient creates a new GeneratesCSRsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGeneratesCSRsParamsWithHTTPClient(client *http.Client) *GeneratesCSRsParams {
	return &GeneratesCSRsParams{
		HTTPClient: client,
	}
}

/*
GeneratesCSRsParams contains all the parameters to send to the API endpoint

	for the generates c s rs operation.

	Typically these are written to a http.Request.
*/
type GeneratesCSRsParams struct {

	/* CSRSGenerationSpec.

	   csrsGenerationSpec
	*/
	CSRSGenerationSpec *models.CSRSGenerationSpec

	/* DomainName.

	   Domain ID or Name
	*/
	DomainName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generates c s rs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneratesCSRsParams) WithDefaults() *GeneratesCSRsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generates c s rs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GeneratesCSRsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generates c s rs params
func (o *GeneratesCSRsParams) WithTimeout(timeout time.Duration) *GeneratesCSRsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generates c s rs params
func (o *GeneratesCSRsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generates c s rs params
func (o *GeneratesCSRsParams) WithContext(ctx context.Context) *GeneratesCSRsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generates c s rs params
func (o *GeneratesCSRsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generates c s rs params
func (o *GeneratesCSRsParams) WithHTTPClient(client *http.Client) *GeneratesCSRsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generates c s rs params
func (o *GeneratesCSRsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCSRSGenerationSpec adds the cSRSGenerationSpec to the generates c s rs params
func (o *GeneratesCSRsParams) WithCSRSGenerationSpec(cSRSGenerationSpec *models.CSRSGenerationSpec) *GeneratesCSRsParams {
	o.SetCSRSGenerationSpec(cSRSGenerationSpec)
	return o
}

// SetCSRSGenerationSpec adds the csrsGenerationSpec to the generates c s rs params
func (o *GeneratesCSRsParams) SetCSRSGenerationSpec(cSRSGenerationSpec *models.CSRSGenerationSpec) {
	o.CSRSGenerationSpec = cSRSGenerationSpec
}

// WithDomainName adds the domainName to the generates c s rs params
func (o *GeneratesCSRsParams) WithDomainName(domainName string) *GeneratesCSRsParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the generates c s rs params
func (o *GeneratesCSRsParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WriteToRequest writes these params to a swagger request
func (o *GeneratesCSRsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CSRSGenerationSpec != nil {
		if err := r.SetBodyParam(o.CSRSGenerationSpec); err != nil {
			return err
		}
	}

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
