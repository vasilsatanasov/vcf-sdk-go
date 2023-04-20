// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vrslcm

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

// NewUpdateVRSLCMVersionInInventoryParams creates a new UpdateVRSLCMVersionInInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVRSLCMVersionInInventoryParams() *UpdateVRSLCMVersionInInventoryParams {
	return &UpdateVRSLCMVersionInInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVRSLCMVersionInInventoryParamsWithTimeout creates a new UpdateVRSLCMVersionInInventoryParams object
// with the ability to set a timeout on a request.
func NewUpdateVRSLCMVersionInInventoryParamsWithTimeout(timeout time.Duration) *UpdateVRSLCMVersionInInventoryParams {
	return &UpdateVRSLCMVersionInInventoryParams{
		timeout: timeout,
	}
}

// NewUpdateVRSLCMVersionInInventoryParamsWithContext creates a new UpdateVRSLCMVersionInInventoryParams object
// with the ability to set a context for a request.
func NewUpdateVRSLCMVersionInInventoryParamsWithContext(ctx context.Context) *UpdateVRSLCMVersionInInventoryParams {
	return &UpdateVRSLCMVersionInInventoryParams{
		Context: ctx,
	}
}

// NewUpdateVRSLCMVersionInInventoryParamsWithHTTPClient creates a new UpdateVRSLCMVersionInInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVRSLCMVersionInInventoryParamsWithHTTPClient(client *http.Client) *UpdateVRSLCMVersionInInventoryParams {
	return &UpdateVRSLCMVersionInInventoryParams{
		HTTPClient: client,
	}
}

/*
UpdateVRSLCMVersionInInventoryParams contains all the parameters to send to the API endpoint

	for the update Vrslcm version in inventory operation.

	Typically these are written to a http.Request.
*/
type UpdateVRSLCMVersionInInventoryParams struct {

	/* VRSLCMDto.

	   vrslcmDto
	*/
	VRSLCMDto *models.VRSLCM

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vrslcm version in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVRSLCMVersionInInventoryParams) WithDefaults() *UpdateVRSLCMVersionInInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vrslcm version in inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVRSLCMVersionInInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) WithTimeout(timeout time.Duration) *UpdateVRSLCMVersionInInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) WithContext(ctx context.Context) *UpdateVRSLCMVersionInInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) WithHTTPClient(client *http.Client) *UpdateVRSLCMVersionInInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVRSLCMDto adds the vRSLCMDto to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) WithVRSLCMDto(vRSLCMDto *models.VRSLCM) *UpdateVRSLCMVersionInInventoryParams {
	o.SetVRSLCMDto(vRSLCMDto)
	return o
}

// SetVRSLCMDto adds the vrslcmDto to the update Vrslcm version in inventory params
func (o *UpdateVRSLCMVersionInInventoryParams) SetVRSLCMDto(vRSLCMDto *models.VRSLCM) {
	o.VRSLCMDto = vRSLCMDto
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVRSLCMVersionInInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.VRSLCMDto != nil {
		if err := r.SetBodyParam(o.VRSLCMDto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}