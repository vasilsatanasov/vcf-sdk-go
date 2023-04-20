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

// POSTClustersQueryReader is a Reader for the POSTClustersQuery structure.
type POSTClustersQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *POSTClustersQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPOSTClustersQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPOSTClustersQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPOSTClustersQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPOSTClustersQueryOK creates a POSTClustersQueryOK with default headers values
func NewPOSTClustersQueryOK() *POSTClustersQueryOK {
	return &POSTClustersQueryOK{}
}

/*
POSTClustersQueryOK describes a response with status code 200, with default header values.

Ok
*/
type POSTClustersQueryOK struct {
	Payload *models.ClusterQueryResponse
}

// IsSuccess returns true when this post clusters query o k response has a 2xx status code
func (o *POSTClustersQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post clusters query o k response has a 3xx status code
func (o *POSTClustersQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query o k response has a 4xx status code
func (o *POSTClustersQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post clusters query o k response has a 5xx status code
func (o *POSTClustersQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post clusters query o k response a status code equal to that given
func (o *POSTClustersQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *POSTClustersQueryOK) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryOK  %+v", 200, o.Payload)
}

func (o *POSTClustersQueryOK) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryOK  %+v", 200, o.Payload)
}

func (o *POSTClustersQueryOK) GetPayload() *models.ClusterQueryResponse {
	return o.Payload
}

func (o *POSTClustersQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTClustersQueryBadRequest creates a POSTClustersQueryBadRequest with default headers values
func NewPOSTClustersQueryBadRequest() *POSTClustersQueryBadRequest {
	return &POSTClustersQueryBadRequest{}
}

/*
POSTClustersQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type POSTClustersQueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post clusters query bad request response has a 2xx status code
func (o *POSTClustersQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post clusters query bad request response has a 3xx status code
func (o *POSTClustersQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query bad request response has a 4xx status code
func (o *POSTClustersQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post clusters query bad request response has a 5xx status code
func (o *POSTClustersQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post clusters query bad request response a status code equal to that given
func (o *POSTClustersQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *POSTClustersQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryBadRequest  %+v", 400, o.Payload)
}

func (o *POSTClustersQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryBadRequest  %+v", 400, o.Payload)
}

func (o *POSTClustersQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *POSTClustersQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPOSTClustersQueryInternalServerError creates a POSTClustersQueryInternalServerError with default headers values
func NewPOSTClustersQueryInternalServerError() *POSTClustersQueryInternalServerError {
	return &POSTClustersQueryInternalServerError{}
}

/*
POSTClustersQueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type POSTClustersQueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this post clusters query internal server error response has a 2xx status code
func (o *POSTClustersQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post clusters query internal server error response has a 3xx status code
func (o *POSTClustersQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post clusters query internal server error response has a 4xx status code
func (o *POSTClustersQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post clusters query internal server error response has a 5xx status code
func (o *POSTClustersQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post clusters query internal server error response a status code equal to that given
func (o *POSTClustersQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *POSTClustersQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *POSTClustersQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/domains/{domainId}/clusters/queries][%d] postClustersQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *POSTClustersQueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *POSTClustersQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}