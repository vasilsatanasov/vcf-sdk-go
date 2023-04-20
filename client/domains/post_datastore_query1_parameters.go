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

	"github.com/vmware/vcf-sdk-go/models"
)

// NewPOSTDatastoreQuery1Params creates a new POSTDatastoreQuery1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPOSTDatastoreQuery1Params() *POSTDatastoreQuery1Params {
	return &POSTDatastoreQuery1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTDatastoreQuery1ParamsWithTimeout creates a new POSTDatastoreQuery1Params object
// with the ability to set a timeout on a request.
func NewPOSTDatastoreQuery1ParamsWithTimeout(timeout time.Duration) *POSTDatastoreQuery1Params {
	return &POSTDatastoreQuery1Params{
		timeout: timeout,
	}
}

// NewPOSTDatastoreQuery1ParamsWithContext creates a new POSTDatastoreQuery1Params object
// with the ability to set a context for a request.
func NewPOSTDatastoreQuery1ParamsWithContext(ctx context.Context) *POSTDatastoreQuery1Params {
	return &POSTDatastoreQuery1Params{
		Context: ctx,
	}
}

// NewPOSTDatastoreQuery1ParamsWithHTTPClient creates a new POSTDatastoreQuery1Params object
// with the ability to set a custom HTTPClient for a request.
func NewPOSTDatastoreQuery1ParamsWithHTTPClient(client *http.Client) *POSTDatastoreQuery1Params {
	return &POSTDatastoreQuery1Params{
		HTTPClient: client,
	}
}

/*
POSTDatastoreQuery1Params contains all the parameters to send to the API endpoint

	for the post datastore query 1 operation.

	Typically these are written to a http.Request.
*/
type POSTDatastoreQuery1Params struct {

	/* DomainID.

	   Domain ID
	*/
	DomainID string

	/* DsCriterion.

	   dsCriterion
	*/
	DsCriterion *models.DatastoreCriterion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post datastore query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTDatastoreQuery1Params) WithDefaults() *POSTDatastoreQuery1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post datastore query 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *POSTDatastoreQuery1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) WithTimeout(timeout time.Duration) *POSTDatastoreQuery1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) WithContext(ctx context.Context) *POSTDatastoreQuery1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) WithHTTPClient(client *http.Client) *POSTDatastoreQuery1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) WithDomainID(domainID string) *POSTDatastoreQuery1Params {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithDsCriterion adds the dsCriterion to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) WithDsCriterion(dsCriterion *models.DatastoreCriterion) *POSTDatastoreQuery1Params {
	o.SetDsCriterion(dsCriterion)
	return o
}

// SetDsCriterion adds the dsCriterion to the post datastore query 1 params
func (o *POSTDatastoreQuery1Params) SetDsCriterion(dsCriterion *models.DatastoreCriterion) {
	o.DsCriterion = dsCriterion
}

// WriteToRequest writes these params to a swagger request
func (o *POSTDatastoreQuery1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}
	if o.DsCriterion != nil {
		if err := r.SetBodyParam(o.DsCriterion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}