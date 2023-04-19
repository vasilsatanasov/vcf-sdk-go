// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GethealthsummarytasksReader is a Reader for the Gethealthsummarytasks structure.
type GethealthsummarytasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GethealthsummarytasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGethealthsummarytasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGethealthsummarytasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGethealthsummarytasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGethealthsummarytasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGethealthsummarytasksOK creates a GethealthsummarytasksOK with default headers values
func NewGethealthsummarytasksOK() *GethealthsummarytasksOK {
	return &GethealthsummarytasksOK{}
}

/*
GethealthsummarytasksOK describes a response with status code 200, with default header values.

Ok
*/
type GethealthsummarytasksOK struct {
	Payload *models.PageOfHealthSummary
}

// IsSuccess returns true when this gethealthsummarytasks o k response has a 2xx status code
func (o *GethealthsummarytasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gethealthsummarytasks o k response has a 3xx status code
func (o *GethealthsummarytasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gethealthsummarytasks o k response has a 4xx status code
func (o *GethealthsummarytasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this gethealthsummarytasks o k response has a 5xx status code
func (o *GethealthsummarytasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this gethealthsummarytasks o k response a status code equal to that given
func (o *GethealthsummarytasksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GethealthsummarytasksOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksOK  %+v", 200, o.Payload)
}

func (o *GethealthsummarytasksOK) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksOK  %+v", 200, o.Payload)
}

func (o *GethealthsummarytasksOK) GetPayload() *models.PageOfHealthSummary {
	return o.Payload
}

func (o *GethealthsummarytasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfHealthSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGethealthsummarytasksBadRequest creates a GethealthsummarytasksBadRequest with default headers values
func NewGethealthsummarytasksBadRequest() *GethealthsummarytasksBadRequest {
	return &GethealthsummarytasksBadRequest{}
}

/*
GethealthsummarytasksBadRequest describes a response with status code 400, with default header values.

Bad request! Invalid Headers or Data. Error: {error}
*/
type GethealthsummarytasksBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this gethealthsummarytasks bad request response has a 2xx status code
func (o *GethealthsummarytasksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gethealthsummarytasks bad request response has a 3xx status code
func (o *GethealthsummarytasksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gethealthsummarytasks bad request response has a 4xx status code
func (o *GethealthsummarytasksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this gethealthsummarytasks bad request response has a 5xx status code
func (o *GethealthsummarytasksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this gethealthsummarytasks bad request response a status code equal to that given
func (o *GethealthsummarytasksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GethealthsummarytasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksBadRequest  %+v", 400, o.Payload)
}

func (o *GethealthsummarytasksBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksBadRequest  %+v", 400, o.Payload)
}

func (o *GethealthsummarytasksBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GethealthsummarytasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGethealthsummarytasksUnauthorized creates a GethealthsummarytasksUnauthorized with default headers values
func NewGethealthsummarytasksUnauthorized() *GethealthsummarytasksUnauthorized {
	return &GethealthsummarytasksUnauthorized{}
}

/*
GethealthsummarytasksUnauthorized describes a response with status code 401, with default header values.

Bad request! Authorization Header is missing or not in correct format.
*/
type GethealthsummarytasksUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this gethealthsummarytasks unauthorized response has a 2xx status code
func (o *GethealthsummarytasksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gethealthsummarytasks unauthorized response has a 3xx status code
func (o *GethealthsummarytasksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gethealthsummarytasks unauthorized response has a 4xx status code
func (o *GethealthsummarytasksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this gethealthsummarytasks unauthorized response has a 5xx status code
func (o *GethealthsummarytasksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this gethealthsummarytasks unauthorized response a status code equal to that given
func (o *GethealthsummarytasksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GethealthsummarytasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksUnauthorized  %+v", 401, o.Payload)
}

func (o *GethealthsummarytasksUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksUnauthorized  %+v", 401, o.Payload)
}

func (o *GethealthsummarytasksUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GethealthsummarytasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGethealthsummarytasksInternalServerError creates a GethealthsummarytasksInternalServerError with default headers values
func NewGethealthsummarytasksInternalServerError() *GethealthsummarytasksInternalServerError {
	return &GethealthsummarytasksInternalServerError{}
}

/*
GethealthsummarytasksInternalServerError describes a response with status code 500, with default header values.

Something went wrong. Internal server error occurred. Error {error}
*/
type GethealthsummarytasksInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this gethealthsummarytasks internal server error response has a 2xx status code
func (o *GethealthsummarytasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this gethealthsummarytasks internal server error response has a 3xx status code
func (o *GethealthsummarytasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gethealthsummarytasks internal server error response has a 4xx status code
func (o *GethealthsummarytasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this gethealthsummarytasks internal server error response has a 5xx status code
func (o *GethealthsummarytasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this gethealthsummarytasks internal server error response a status code equal to that given
func (o *GethealthsummarytasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GethealthsummarytasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GethealthsummarytasksInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/health-summary][%d] gethealthsummarytasksInternalServerError  %+v", 500, o.Payload)
}

func (o *GethealthsummarytasksInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GethealthsummarytasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
