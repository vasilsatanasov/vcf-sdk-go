// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETHostCriterionReader is a Reader for the GETHostCriterion structure.
type GETHostCriterionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETHostCriterionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETHostCriterionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETHostCriterionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETHostCriterionOK creates a GETHostCriterionOK with default headers values
func NewGETHostCriterionOK() *GETHostCriterionOK {
	return &GETHostCriterionOK{}
}

/*
GETHostCriterionOK describes a response with status code 200, with default header values.

Ok
*/
type GETHostCriterionOK struct {
	Payload *models.HostCriterion
}

// IsSuccess returns true when this get host criterion o k response has a 2xx status code
func (o *GETHostCriterionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host criterion o k response has a 3xx status code
func (o *GETHostCriterionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host criterion o k response has a 4xx status code
func (o *GETHostCriterionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host criterion o k response has a 5xx status code
func (o *GETHostCriterionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host criterion o k response a status code equal to that given
func (o *GETHostCriterionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETHostCriterionOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/hosts/criteria/{name}][%d] getHostCriterionOK  %+v", 200, o.Payload)
}

func (o *GETHostCriterionOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/hosts/criteria/{name}][%d] getHostCriterionOK  %+v", 200, o.Payload)
}

func (o *GETHostCriterionOK) GetPayload() *models.HostCriterion {
	return o.Payload
}

func (o *GETHostCriterionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostCriterion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETHostCriterionNotFound creates a GETHostCriterionNotFound with default headers values
func NewGETHostCriterionNotFound() *GETHostCriterionNotFound {
	return &GETHostCriterionNotFound{}
}

/*
GETHostCriterionNotFound describes a response with status code 404, with default header values.

Criterion Not Found
*/
type GETHostCriterionNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get host criterion not found response has a 2xx status code
func (o *GETHostCriterionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host criterion not found response has a 3xx status code
func (o *GETHostCriterionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host criterion not found response has a 4xx status code
func (o *GETHostCriterionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get host criterion not found response has a 5xx status code
func (o *GETHostCriterionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get host criterion not found response a status code equal to that given
func (o *GETHostCriterionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETHostCriterionNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/hosts/criteria/{name}][%d] getHostCriterionNotFound  %+v", 404, o.Payload)
}

func (o *GETHostCriterionNotFound) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}/hosts/criteria/{name}][%d] getHostCriterionNotFound  %+v", 404, o.Payload)
}

func (o *GETHostCriterionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETHostCriterionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}