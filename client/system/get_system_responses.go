// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETSystemReader is a Reader for the GETSystem structure.
type GETSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETSystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETSystemOK creates a GETSystemOK with default headers values
func NewGETSystemOK() *GETSystemOK {
	return &GETSystemOK{}
}

/*
GETSystemOK describes a response with status code 200, with default header values.

OK
*/
type GETSystemOK struct {
	Payload *models.System
}

// IsSuccess returns true when this get system o k response has a 2xx status code
func (o *GETSystemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get system o k response has a 3xx status code
func (o *GETSystemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system o k response has a 4xx status code
func (o *GETSystemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system o k response has a 5xx status code
func (o *GETSystemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get system o k response a status code equal to that given
func (o *GETSystemOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETSystemOK) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemOK  %+v", 200, o.Payload)
}

func (o *GETSystemOK) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemOK  %+v", 200, o.Payload)
}

func (o *GETSystemOK) GetPayload() *models.System {
	return o.Payload
}

func (o *GETSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.System)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSystemBadRequest creates a GETSystemBadRequest with default headers values
func NewGETSystemBadRequest() *GETSystemBadRequest {
	return &GETSystemBadRequest{}
}

/*
GETSystemBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETSystemBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get system bad request response has a 2xx status code
func (o *GETSystemBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system bad request response has a 3xx status code
func (o *GETSystemBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system bad request response has a 4xx status code
func (o *GETSystemBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get system bad request response has a 5xx status code
func (o *GETSystemBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get system bad request response a status code equal to that given
func (o *GETSystemBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETSystemBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemBadRequest  %+v", 400, o.Payload)
}

func (o *GETSystemBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemBadRequest  %+v", 400, o.Payload)
}

func (o *GETSystemBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETSystemInternalServerError creates a GETSystemInternalServerError with default headers values
func NewGETSystemInternalServerError() *GETSystemInternalServerError {
	return &GETSystemInternalServerError{}
}

/*
GETSystemInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETSystemInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get system internal server error response has a 2xx status code
func (o *GETSystemInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get system internal server error response has a 3xx status code
func (o *GETSystemInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get system internal server error response has a 4xx status code
func (o *GETSystemInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get system internal server error response has a 5xx status code
func (o *GETSystemInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get system internal server error response a status code equal to that given
func (o *GETSystemInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETSystemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSystemInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system][%d] getSystemInternalServerError  %+v", 500, o.Payload)
}

func (o *GETSystemInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETSystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}