// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETLocalAccountReader is a Reader for the GETLocalAccount structure.
type GETLocalAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETLocalAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETLocalAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETLocalAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGETLocalAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETLocalAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETLocalAccountOK creates a GETLocalAccountOK with default headers values
func NewGETLocalAccountOK() *GETLocalAccountOK {
	return &GETLocalAccountOK{}
}

/*
GETLocalAccountOK describes a response with status code 200, with default header values.

OK
*/
type GETLocalAccountOK struct {
	Payload *models.LocalUser
}

// IsSuccess returns true when this get local account o k response has a 2xx status code
func (o *GETLocalAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get local account o k response has a 3xx status code
func (o *GETLocalAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get local account o k response has a 4xx status code
func (o *GETLocalAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get local account o k response has a 5xx status code
func (o *GETLocalAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get local account o k response a status code equal to that given
func (o *GETLocalAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETLocalAccountOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountOK  %+v", 200, o.Payload)
}

func (o *GETLocalAccountOK) String() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountOK  %+v", 200, o.Payload)
}

func (o *GETLocalAccountOK) GetPayload() *models.LocalUser {
	return o.Payload
}

func (o *GETLocalAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETLocalAccountBadRequest creates a GETLocalAccountBadRequest with default headers values
func NewGETLocalAccountBadRequest() *GETLocalAccountBadRequest {
	return &GETLocalAccountBadRequest{}
}

/*
GETLocalAccountBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GETLocalAccountBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get local account bad request response has a 2xx status code
func (o *GETLocalAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get local account bad request response has a 3xx status code
func (o *GETLocalAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get local account bad request response has a 4xx status code
func (o *GETLocalAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get local account bad request response has a 5xx status code
func (o *GETLocalAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get local account bad request response a status code equal to that given
func (o *GETLocalAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETLocalAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GETLocalAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GETLocalAccountBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETLocalAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETLocalAccountUnauthorized creates a GETLocalAccountUnauthorized with default headers values
func NewGETLocalAccountUnauthorized() *GETLocalAccountUnauthorized {
	return &GETLocalAccountUnauthorized{}
}

/*
GETLocalAccountUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETLocalAccountUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get local account unauthorized response has a 2xx status code
func (o *GETLocalAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get local account unauthorized response has a 3xx status code
func (o *GETLocalAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get local account unauthorized response has a 4xx status code
func (o *GETLocalAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get local account unauthorized response has a 5xx status code
func (o *GETLocalAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get local account unauthorized response a status code equal to that given
func (o *GETLocalAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETLocalAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *GETLocalAccountUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *GETLocalAccountUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETLocalAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETLocalAccountInternalServerError creates a GETLocalAccountInternalServerError with default headers values
func NewGETLocalAccountInternalServerError() *GETLocalAccountInternalServerError {
	return &GETLocalAccountInternalServerError{}
}

/*
GETLocalAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETLocalAccountInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get local account internal server error response has a 2xx status code
func (o *GETLocalAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get local account internal server error response has a 3xx status code
func (o *GETLocalAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get local account internal server error response has a 4xx status code
func (o *GETLocalAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get local account internal server error response has a 5xx status code
func (o *GETLocalAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get local account internal server error response a status code equal to that given
func (o *GETLocalAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETLocalAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GETLocalAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/users/local/admin][%d] getLocalAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GETLocalAccountInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETLocalAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
