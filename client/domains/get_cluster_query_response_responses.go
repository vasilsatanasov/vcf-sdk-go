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

// GETClusterQueryResponseReader is a Reader for the GETClusterQueryResponse structure.
type GETClusterQueryResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETClusterQueryResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETClusterQueryResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETClusterQueryResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGETClusterQueryResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETClusterQueryResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETClusterQueryResponseOK creates a GETClusterQueryResponseOK with default headers values
func NewGETClusterQueryResponseOK() *GETClusterQueryResponseOK {
	return &GETClusterQueryResponseOK{}
}

/*
GETClusterQueryResponseOK describes a response with status code 200, with default header values.

Ok
*/
type GETClusterQueryResponseOK struct {
	Payload *models.ClusterQueryResponse
}

// IsSuccess returns true when this get cluster query response o k response has a 2xx status code
func (o *GETClusterQueryResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster query response o k response has a 3xx status code
func (o *GETClusterQueryResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster query response o k response has a 4xx status code
func (o *GETClusterQueryResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster query response o k response has a 5xx status code
func (o *GETClusterQueryResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster query response o k response a status code equal to that given
func (o *GETClusterQueryResponseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETClusterQueryResponseOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETClusterQueryResponseOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseOK  %+v", 200, o.Payload)
}

func (o *GETClusterQueryResponseOK) GetPayload() *models.ClusterQueryResponse {
	return o.Payload
}

func (o *GETClusterQueryResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterQueryResponseBadRequest creates a GETClusterQueryResponseBadRequest with default headers values
func NewGETClusterQueryResponseBadRequest() *GETClusterQueryResponseBadRequest {
	return &GETClusterQueryResponseBadRequest{}
}

/*
GETClusterQueryResponseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETClusterQueryResponseBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster query response bad request response has a 2xx status code
func (o *GETClusterQueryResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster query response bad request response has a 3xx status code
func (o *GETClusterQueryResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster query response bad request response has a 4xx status code
func (o *GETClusterQueryResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster query response bad request response has a 5xx status code
func (o *GETClusterQueryResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster query response bad request response a status code equal to that given
func (o *GETClusterQueryResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETClusterQueryResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETClusterQueryResponseBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GETClusterQueryResponseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterQueryResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterQueryResponseNotFound creates a GETClusterQueryResponseNotFound with default headers values
func NewGETClusterQueryResponseNotFound() *GETClusterQueryResponseNotFound {
	return &GETClusterQueryResponseNotFound{}
}

/*
GETClusterQueryResponseNotFound describes a response with status code 404, with default header values.

Query Not Found
*/
type GETClusterQueryResponseNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster query response not found response has a 2xx status code
func (o *GETClusterQueryResponseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster query response not found response has a 3xx status code
func (o *GETClusterQueryResponseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster query response not found response has a 4xx status code
func (o *GETClusterQueryResponseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster query response not found response has a 5xx status code
func (o *GETClusterQueryResponseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster query response not found response a status code equal to that given
func (o *GETClusterQueryResponseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETClusterQueryResponseNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETClusterQueryResponseNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseNotFound  %+v", 404, o.Payload)
}

func (o *GETClusterQueryResponseNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterQueryResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETClusterQueryResponseInternalServerError creates a GETClusterQueryResponseInternalServerError with default headers values
func NewGETClusterQueryResponseInternalServerError() *GETClusterQueryResponseInternalServerError {
	return &GETClusterQueryResponseInternalServerError{}
}

/*
GETClusterQueryResponseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETClusterQueryResponseInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get cluster query response internal server error response has a 2xx status code
func (o *GETClusterQueryResponseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster query response internal server error response has a 3xx status code
func (o *GETClusterQueryResponseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster query response internal server error response has a 4xx status code
func (o *GETClusterQueryResponseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster query response internal server error response has a 5xx status code
func (o *GETClusterQueryResponseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cluster query response internal server error response a status code equal to that given
func (o *GETClusterQueryResponseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETClusterQueryResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterQueryResponseInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainId}/clusters/{clusterName}/queries/{queryId}][%d] getClusterQueryResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GETClusterQueryResponseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETClusterQueryResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}