// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETDatastoreCriterion1Reader is a Reader for the GETDatastoreCriterion1 structure.
type GETDatastoreCriterion1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDatastoreCriterion1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETDatastoreCriterion1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETDatastoreCriterion1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETDatastoreCriterion1OK creates a GETDatastoreCriterion1OK with default headers values
func NewGETDatastoreCriterion1OK() *GETDatastoreCriterion1OK {
	return &GETDatastoreCriterion1OK{}
}

/*
GETDatastoreCriterion1OK describes a response with status code 200, with default header values.

Ok
*/
type GETDatastoreCriterion1OK struct {
	Payload *models.DatastoreCriterion
}

// IsSuccess returns true when this get datastore criterion1 o k response has a 2xx status code
func (o *GETDatastoreCriterion1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get datastore criterion1 o k response has a 3xx status code
func (o *GETDatastoreCriterion1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore criterion1 o k response has a 4xx status code
func (o *GETDatastoreCriterion1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get datastore criterion1 o k response has a 5xx status code
func (o *GETDatastoreCriterion1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore criterion1 o k response a status code equal to that given
func (o *GETDatastoreCriterion1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GETDatastoreCriterion1OK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/criteria/{name}][%d] getDatastoreCriterion1OK  %+v", 200, o.Payload)
}

func (o *GETDatastoreCriterion1OK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/criteria/{name}][%d] getDatastoreCriterion1OK  %+v", 200, o.Payload)
}

func (o *GETDatastoreCriterion1OK) GetPayload() *models.DatastoreCriterion {
	return o.Payload
}

func (o *GETDatastoreCriterion1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatastoreCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETDatastoreCriterion1NotFound creates a GETDatastoreCriterion1NotFound with default headers values
func NewGETDatastoreCriterion1NotFound() *GETDatastoreCriterion1NotFound {
	return &GETDatastoreCriterion1NotFound{}
}

/*
GETDatastoreCriterion1NotFound describes a response with status code 404, with default header values.

Criterion Not Found
*/
type GETDatastoreCriterion1NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get datastore criterion1 not found response has a 2xx status code
func (o *GETDatastoreCriterion1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get datastore criterion1 not found response has a 3xx status code
func (o *GETDatastoreCriterion1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get datastore criterion1 not found response has a 4xx status code
func (o *GETDatastoreCriterion1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get datastore criterion1 not found response has a 5xx status code
func (o *GETDatastoreCriterion1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get datastore criterion1 not found response a status code equal to that given
func (o *GETDatastoreCriterion1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETDatastoreCriterion1NotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/criteria/{name}][%d] getDatastoreCriterion1NotFound  %+v", 404, o.Payload)
}

func (o *GETDatastoreCriterion1NotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/datastores/criteria/{name}][%d] getDatastoreCriterion1NotFound  %+v", 404, o.Payload)
}

func (o *GETDatastoreCriterion1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETDatastoreCriterion1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}