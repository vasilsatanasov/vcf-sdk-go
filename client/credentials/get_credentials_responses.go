// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETCredentialsReader is a Reader for the GETCredentials structure.
type GETCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGETCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGETCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETCredentialsOK creates a GETCredentialsOK with default headers values
func NewGETCredentialsOK() *GETCredentialsOK {
	return &GETCredentialsOK{}
}

/*
GETCredentialsOK describes a response with status code 200, with default header values.

OK
*/
type GETCredentialsOK struct {
	Payload *models.PageOfCredential
}

// IsSuccess returns true when this get credentials o k response has a 2xx status code
func (o *GETCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credentials o k response has a 3xx status code
func (o *GETCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials o k response has a 4xx status code
func (o *GETCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials o k response has a 5xx status code
func (o *GETCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials o k response a status code equal to that given
func (o *GETCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GETCredentialsOK) String() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GETCredentialsOK) GetPayload() *models.PageOfCredential {
	return o.Payload
}

func (o *GETCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsBadRequest creates a GETCredentialsBadRequest with default headers values
func NewGETCredentialsBadRequest() *GETCredentialsBadRequest {
	return &GETCredentialsBadRequest{}
}

/*
GETCredentialsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETCredentialsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials bad request response has a 2xx status code
func (o *GETCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials bad request response has a 3xx status code
func (o *GETCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials bad request response has a 4xx status code
func (o *GETCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials bad request response has a 5xx status code
func (o *GETCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials bad request response a status code equal to that given
func (o *GETCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GETCredentialsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *GETCredentialsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsUnauthorized creates a GETCredentialsUnauthorized with default headers values
func NewGETCredentialsUnauthorized() *GETCredentialsUnauthorized {
	return &GETCredentialsUnauthorized{}
}

/*
GETCredentialsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GETCredentialsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials unauthorized response has a 2xx status code
func (o *GETCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials unauthorized response has a 3xx status code
func (o *GETCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials unauthorized response has a 4xx status code
func (o *GETCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials unauthorized response has a 5xx status code
func (o *GETCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials unauthorized response a status code equal to that given
func (o *GETCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GETCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *GETCredentialsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsForbidden creates a GETCredentialsForbidden with default headers values
func NewGETCredentialsForbidden() *GETCredentialsForbidden {
	return &GETCredentialsForbidden{}
}

/*
GETCredentialsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GETCredentialsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials forbidden response has a 2xx status code
func (o *GETCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials forbidden response has a 3xx status code
func (o *GETCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials forbidden response has a 4xx status code
func (o *GETCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credentials forbidden response has a 5xx status code
func (o *GETCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials forbidden response a status code equal to that given
func (o *GETCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GETCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *GETCredentialsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *GETCredentialsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCredentialsInternalServerError creates a GETCredentialsInternalServerError with default headers values
func NewGETCredentialsInternalServerError() *GETCredentialsInternalServerError {
	return &GETCredentialsInternalServerError{}
}

/*
GETCredentialsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETCredentialsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials internal server error response has a 2xx status code
func (o *GETCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials internal server error response has a 3xx status code
func (o *GETCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials internal server error response has a 4xx status code
func (o *GETCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials internal server error response has a 5xx status code
func (o *GETCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get credentials internal server error response a status code equal to that given
func (o *GETCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}