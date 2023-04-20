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

// GETValidationOfNtpConfigurationReader is a Reader for the GETValidationOfNtpConfiguration structure.
type GETValidationOfNtpConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETValidationOfNtpConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETValidationOfNtpConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETValidationOfNtpConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETValidationOfNtpConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETValidationOfNtpConfigurationOK creates a GETValidationOfNtpConfigurationOK with default headers values
func NewGETValidationOfNtpConfigurationOK() *GETValidationOfNtpConfigurationOK {
	return &GETValidationOfNtpConfigurationOK{}
}

/*
GETValidationOfNtpConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GETValidationOfNtpConfigurationOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation of ntp configuration o k response has a 2xx status code
func (o *GETValidationOfNtpConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation of ntp configuration o k response has a 3xx status code
func (o *GETValidationOfNtpConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration o k response has a 4xx status code
func (o *GETValidationOfNtpConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of ntp configuration o k response has a 5xx status code
func (o *GETValidationOfNtpConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of ntp configuration o k response a status code equal to that given
func (o *GETValidationOfNtpConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETValidationOfNtpConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETValidationOfNtpConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationOK  %+v", 200, o.Payload)
}

func (o *GETValidationOfNtpConfigurationOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETValidationOfNtpConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationOfNtpConfigurationBadRequest creates a GETValidationOfNtpConfigurationBadRequest with default headers values
func NewGETValidationOfNtpConfigurationBadRequest() *GETValidationOfNtpConfigurationBadRequest {
	return &GETValidationOfNtpConfigurationBadRequest{}
}

/*
GETValidationOfNtpConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETValidationOfNtpConfigurationBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validation of ntp configuration bad request response has a 2xx status code
func (o *GETValidationOfNtpConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of ntp configuration bad request response has a 3xx status code
func (o *GETValidationOfNtpConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration bad request response has a 4xx status code
func (o *GETValidationOfNtpConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validation of ntp configuration bad request response has a 5xx status code
func (o *GETValidationOfNtpConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation of ntp configuration bad request response a status code equal to that given
func (o *GETValidationOfNtpConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETValidationOfNtpConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationOfNtpConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationOfNtpConfigurationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETValidationOfNtpConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationOfNtpConfigurationInternalServerError creates a GETValidationOfNtpConfigurationInternalServerError with default headers values
func NewGETValidationOfNtpConfigurationInternalServerError() *GETValidationOfNtpConfigurationInternalServerError {
	return &GETValidationOfNtpConfigurationInternalServerError{}
}

/*
GETValidationOfNtpConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETValidationOfNtpConfigurationInternalServerError struct {
}

// IsSuccess returns true when this get validation of ntp configuration internal server error response has a 2xx status code
func (o *GETValidationOfNtpConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation of ntp configuration internal server error response has a 3xx status code
func (o *GETValidationOfNtpConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation of ntp configuration internal server error response has a 4xx status code
func (o *GETValidationOfNtpConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation of ntp configuration internal server error response has a 5xx status code
func (o *GETValidationOfNtpConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validation of ntp configuration internal server error response a status code equal to that given
func (o *GETValidationOfNtpConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETValidationOfNtpConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationInternalServerError ", 500)
}

func (o *GETValidationOfNtpConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/ntp-configuration/validations/{id}][%d] getValidationOfNtpConfigurationInternalServerError ", 500)
}

func (o *GETValidationOfNtpConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}