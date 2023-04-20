// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

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

// NewAssignTagsToExistingClusterParams creates a new AssignTagsToExistingClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignTagsToExistingClusterParams() *AssignTagsToExistingClusterParams {
	return &AssignTagsToExistingClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignTagsToExistingClusterParamsWithTimeout creates a new AssignTagsToExistingClusterParams object
// with the ability to set a timeout on a request.
func NewAssignTagsToExistingClusterParamsWithTimeout(timeout time.Duration) *AssignTagsToExistingClusterParams {
	return &AssignTagsToExistingClusterParams{
		timeout: timeout,
	}
}

// NewAssignTagsToExistingClusterParamsWithContext creates a new AssignTagsToExistingClusterParams object
// with the ability to set a context for a request.
func NewAssignTagsToExistingClusterParamsWithContext(ctx context.Context) *AssignTagsToExistingClusterParams {
	return &AssignTagsToExistingClusterParams{
		Context: ctx,
	}
}

// NewAssignTagsToExistingClusterParamsWithHTTPClient creates a new AssignTagsToExistingClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignTagsToExistingClusterParamsWithHTTPClient(client *http.Client) *AssignTagsToExistingClusterParams {
	return &AssignTagsToExistingClusterParams{
		HTTPClient: client,
	}
}

/*
AssignTagsToExistingClusterParams contains all the parameters to send to the API endpoint

	for the assign tags to existing cluster operation.

	Typically these are written to a http.Request.
*/
type AssignTagsToExistingClusterParams struct {

	/* ID.

	   Cluster ID
	*/
	ID string

	/* TagsSpec.

	   Tags Spec
	*/
	TagsSpec *models.TagsSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign tags to existing cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingClusterParams) WithDefaults() *AssignTagsToExistingClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign tags to existing cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignTagsToExistingClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) WithTimeout(timeout time.Duration) *AssignTagsToExistingClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) WithContext(ctx context.Context) *AssignTagsToExistingClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) WithHTTPClient(client *http.Client) *AssignTagsToExistingClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) WithID(id string) *AssignTagsToExistingClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) SetID(id string) {
	o.ID = id
}

// WithTagsSpec adds the tagsSpec to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) WithTagsSpec(tagsSpec *models.TagsSpec) *AssignTagsToExistingClusterParams {
	o.SetTagsSpec(tagsSpec)
	return o
}

// SetTagsSpec adds the tagsSpec to the assign tags to existing cluster params
func (o *AssignTagsToExistingClusterParams) SetTagsSpec(tagsSpec *models.TagsSpec) {
	o.TagsSpec = tagsSpec
}

// WriteToRequest writes these params to a swagger request
func (o *AssignTagsToExistingClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.TagsSpec != nil {
		if err := r.SetBodyParam(o.TagsSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}