// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

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

// NewGETBringupValidationReportParams creates a new GETBringupValidationReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGETBringupValidationReportParams() *GETBringupValidationReportParams {
	return &GETBringupValidationReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGETBringupValidationReportParamsWithTimeout creates a new GETBringupValidationReportParams object
// with the ability to set a timeout on a request.
func NewGETBringupValidationReportParamsWithTimeout(timeout time.Duration) *GETBringupValidationReportParams {
	return &GETBringupValidationReportParams{
		timeout: timeout,
	}
}

// NewGETBringupValidationReportParamsWithContext creates a new GETBringupValidationReportParams object
// with the ability to set a context for a request.
func NewGETBringupValidationReportParamsWithContext(ctx context.Context) *GETBringupValidationReportParams {
	return &GETBringupValidationReportParams{
		Context: ctx,
	}
}

// NewGETBringupValidationReportParamsWithHTTPClient creates a new GETBringupValidationReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGETBringupValidationReportParamsWithHTTPClient(client *http.Client) *GETBringupValidationReportParams {
	return &GETBringupValidationReportParams{
		HTTPClient: client,
	}
}

/*
GETBringupValidationReportParams contains all the parameters to send to the API endpoint

	for the get bringup validation report operation.

	Typically these are written to a http.Request.
*/
type GETBringupValidationReportParams struct {

	/* CurClientTime.

	   Current client local time of the the report generation
	*/
	CurClientTime *string

	/* StartTime.

	   Start time of validation to be put in the report
	*/
	StartTime *string

	/* ValidationID.

	   Bringup validation ID
	*/
	ValidationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bringup validation report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBringupValidationReportParams) WithDefaults() *GETBringupValidationReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bringup validation report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GETBringupValidationReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithTimeout(timeout time.Duration) *GETBringupValidationReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithContext(ctx context.Context) *GETBringupValidationReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithHTTPClient(client *http.Client) *GETBringupValidationReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCurClientTime adds the curClientTime to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithCurClientTime(curClientTime *string) *GETBringupValidationReportParams {
	o.SetCurClientTime(curClientTime)
	return o
}

// SetCurClientTime adds the curClientTime to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetCurClientTime(curClientTime *string) {
	o.CurClientTime = curClientTime
}

// WithStartTime adds the startTime to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithStartTime(startTime *string) *GETBringupValidationReportParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WithValidationID adds the validationID to the get bringup validation report params
func (o *GETBringupValidationReportParams) WithValidationID(validationID string) *GETBringupValidationReportParams {
	o.SetValidationID(validationID)
	return o
}

// SetValidationID adds the validationId to the get bringup validation report params
func (o *GETBringupValidationReportParams) SetValidationID(validationID string) {
	o.ValidationID = validationID
}

// WriteToRequest writes these params to a swagger request
func (o *GETBringupValidationReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CurClientTime != nil {

		// query param curClientTime
		var qrCurClientTime string

		if o.CurClientTime != nil {
			qrCurClientTime = *o.CurClientTime
		}
		qCurClientTime := qrCurClientTime
		if qCurClientTime != "" {

			if err := r.SetQueryParam("curClientTime", qCurClientTime); err != nil {
				return err
			}
		}
	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime string

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}
	}

	// path param validationId
	if err := r.SetPathParam("validationId", o.ValidationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}