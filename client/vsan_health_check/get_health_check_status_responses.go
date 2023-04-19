// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vsan_health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETHealthCheckStatusReader is a Reader for the GETHealthCheckStatus structure.
type GETHealthCheckStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETHealthCheckStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETHealthCheckStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETHealthCheckStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETHealthCheckStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETHealthCheckStatusOK creates a GETHealthCheckStatusOK with default headers values
func NewGETHealthCheckStatusOK() *GETHealthCheckStatusOK {
	return &GETHealthCheckStatusOK{}
}

/*
GETHealthCheckStatusOK describes a response with status code 200, with default header values.

Ok
*/
type GETHealthCheckStatusOK struct {
	Payload *models.HealthCheckQueryResult
}

// IsSuccess returns true when this get health check status o k response has a 2xx status code
func (o *GETHealthCheckStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get health check status o k response has a 3xx status code
func (o *GETHealthCheckStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status o k response has a 4xx status code
func (o *GETHealthCheckStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status o k response has a 5xx status code
func (o *GETHealthCheckStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status o k response a status code equal to that given
func (o *GETHealthCheckStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETHealthCheckStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusOK  %+v", 200, o.Payload)
}

func (o *GETHealthCheckStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusOK  %+v", 200, o.Payload)
}

func (o *GETHealthCheckStatusOK) GetPayload() *models.HealthCheckQueryResult {
	return o.Payload
}

func (o *GETHealthCheckStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthCheckQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETHealthCheckStatusBadRequest creates a GETHealthCheckStatusBadRequest with default headers values
func NewGETHealthCheckStatusBadRequest() *GETHealthCheckStatusBadRequest {
	return &GETHealthCheckStatusBadRequest{}
}

/*
GETHealthCheckStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETHealthCheckStatusBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status bad request response has a 2xx status code
func (o *GETHealthCheckStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status bad request response has a 3xx status code
func (o *GETHealthCheckStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status bad request response has a 4xx status code
func (o *GETHealthCheckStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get health check status bad request response has a 5xx status code
func (o *GETHealthCheckStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get health check status bad request response a status code equal to that given
func (o *GETHealthCheckStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETHealthCheckStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GETHealthCheckStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GETHealthCheckStatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETHealthCheckStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETHealthCheckStatusInternalServerError creates a GETHealthCheckStatusInternalServerError with default headers values
func NewGETHealthCheckStatusInternalServerError() *GETHealthCheckStatusInternalServerError {
	return &GETHealthCheckStatusInternalServerError{}
}

/*
GETHealthCheckStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETHealthCheckStatusInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get health check status internal server error response has a 2xx status code
func (o *GETHealthCheckStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get health check status internal server error response has a 3xx status code
func (o *GETHealthCheckStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get health check status internal server error response has a 4xx status code
func (o *GETHealthCheckStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get health check status internal server error response has a 5xx status code
func (o *GETHealthCheckStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get health check status internal server error response a status code equal to that given
func (o *GETHealthCheckStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETHealthCheckStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GETHealthCheckStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/health-checks/queries/{queryId}][%d] getHealthCheckStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GETHealthCheckStatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETHealthCheckStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
