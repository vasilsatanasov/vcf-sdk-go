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

// GETClusterReader is a Reader for the GETCluster structure.
type GETClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETClusterOK creates a GETClusterOK with default headers values
func NewGETClusterOK() *GETClusterOK {
	return &GETClusterOK{}
}

/*
GETClusterOK describes a response with status code 200, with default header values.

Ok
*/
type GETClusterOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this get cluster o k response has a 2xx status code
func (o *GETClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster o k response has a 3xx status code
func (o *GETClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster o k response has a 4xx status code
func (o *GETClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster o k response has a 5xx status code
func (o *GETClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster o k response a status code equal to that given
func (o *GETClusterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETClusterOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GETClusterOK) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GETClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GETClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterNotFound creates a GETClusterNotFound with default headers values
func NewGETClusterNotFound() *GETClusterNotFound {
	return &GETClusterNotFound{}
}

/*
GETClusterNotFound describes a response with status code 404, with default header values.

Cluster not found
*/
type GETClusterNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster not found response has a 2xx status code
func (o *GETClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster not found response has a 3xx status code
func (o *GETClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster not found response has a 4xx status code
func (o *GETClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster not found response has a 5xx status code
func (o *GETClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster not found response a status code equal to that given
func (o *GETClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterNotFound  %+v", 404, o.Payload)
}

func (o *GETClusterNotFound) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterNotFound  %+v", 404, o.Payload)
}

func (o *GETClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterInternalServerError creates a GETClusterInternalServerError with default headers values
func NewGETClusterInternalServerError() *GETClusterInternalServerError {
	return &GETClusterInternalServerError{}
}

/*
GETClusterInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GETClusterInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster internal server error response has a 2xx status code
func (o *GETClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster internal server error response has a 3xx status code
func (o *GETClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster internal server error response has a 4xx status code
func (o *GETClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster internal server error response has a 5xx status code
func (o *GETClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster internal server error response a status code equal to that given
func (o *GETClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETClusterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/clusters/{id}][%d] getClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}