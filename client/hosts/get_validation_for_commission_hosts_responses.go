// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETValidationForCommissionHostsReader is a Reader for the GETValidationForCommissionHosts structure.
type GETValidationForCommissionHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETValidationForCommissionHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETValidationForCommissionHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETValidationForCommissionHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETValidationForCommissionHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETValidationForCommissionHostsOK creates a GETValidationForCommissionHostsOK with default headers values
func NewGETValidationForCommissionHostsOK() *GETValidationForCommissionHostsOK {
	return &GETValidationForCommissionHostsOK{}
}

/*
GETValidationForCommissionHostsOK describes a response with status code 200, with default header values.

OK
*/
type GETValidationForCommissionHostsOK struct {
	Payload *models.Validation
}

// IsSuccess returns true when this get validation for commission hosts o k response has a 2xx status code
func (o *GETValidationForCommissionHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get validation for commission hosts o k response has a 3xx status code
func (o *GETValidationForCommissionHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for commission hosts o k response has a 4xx status code
func (o *GETValidationForCommissionHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation for commission hosts o k response has a 5xx status code
func (o *GETValidationForCommissionHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation for commission hosts o k response a status code equal to that given
func (o *GETValidationForCommissionHostsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETValidationForCommissionHostsOK) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsOK  %+v", 200, o.Payload)
}

func (o *GETValidationForCommissionHostsOK) String() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsOK  %+v", 200, o.Payload)
}

func (o *GETValidationForCommissionHostsOK) GetPayload() *models.Validation {
	return o.Payload
}

func (o *GETValidationForCommissionHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Validation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationForCommissionHostsBadRequest creates a GETValidationForCommissionHostsBadRequest with default headers values
func NewGETValidationForCommissionHostsBadRequest() *GETValidationForCommissionHostsBadRequest {
	return &GETValidationForCommissionHostsBadRequest{}
}

/*
GETValidationForCommissionHostsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETValidationForCommissionHostsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get validation for commission hosts bad request response has a 2xx status code
func (o *GETValidationForCommissionHostsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation for commission hosts bad request response has a 3xx status code
func (o *GETValidationForCommissionHostsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for commission hosts bad request response has a 4xx status code
func (o *GETValidationForCommissionHostsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get validation for commission hosts bad request response has a 5xx status code
func (o *GETValidationForCommissionHostsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get validation for commission hosts bad request response a status code equal to that given
func (o *GETValidationForCommissionHostsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETValidationForCommissionHostsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationForCommissionHostsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsBadRequest  %+v", 400, o.Payload)
}

func (o *GETValidationForCommissionHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETValidationForCommissionHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETValidationForCommissionHostsInternalServerError creates a GETValidationForCommissionHostsInternalServerError with default headers values
func NewGETValidationForCommissionHostsInternalServerError() *GETValidationForCommissionHostsInternalServerError {
	return &GETValidationForCommissionHostsInternalServerError{}
}

/*
GETValidationForCommissionHostsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETValidationForCommissionHostsInternalServerError struct {
}

// IsSuccess returns true when this get validation for commission hosts internal server error response has a 2xx status code
func (o *GETValidationForCommissionHostsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get validation for commission hosts internal server error response has a 3xx status code
func (o *GETValidationForCommissionHostsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get validation for commission hosts internal server error response has a 4xx status code
func (o *GETValidationForCommissionHostsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get validation for commission hosts internal server error response has a 5xx status code
func (o *GETValidationForCommissionHostsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get validation for commission hosts internal server error response a status code equal to that given
func (o *GETValidationForCommissionHostsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETValidationForCommissionHostsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsInternalServerError ", 500)
}

func (o *GETValidationForCommissionHostsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/hosts/validations/{id}][%d] getValidationForCommissionHostsInternalServerError ", 500)
}

func (o *GETValidationForCommissionHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
