// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgradables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETUpgradablesReader is a Reader for the GETUpgradables structure.
type GETUpgradablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETUpgradablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETUpgradablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGETUpgradablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETUpgradablesOK creates a GETUpgradablesOK with default headers values
func NewGETUpgradablesOK() *GETUpgradablesOK {
	return &GETUpgradablesOK{}
}

/*
GETUpgradablesOK describes a response with status code 200, with default header values.

Ok
*/
type GETUpgradablesOK struct {
	Payload *models.PageOfUpgradable
}

// IsSuccess returns true when this get upgradables o k response has a 2xx status code
func (o *GETUpgradablesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get upgradables o k response has a 3xx status code
func (o *GETUpgradablesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables o k response has a 4xx status code
func (o *GETUpgradablesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables o k response has a 5xx status code
func (o *GETUpgradablesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get upgradables o k response a status code equal to that given
func (o *GETUpgradablesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETUpgradablesOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesOK  %+v", 200, o.Payload)
}

func (o *GETUpgradablesOK) String() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesOK  %+v", 200, o.Payload)
}

func (o *GETUpgradablesOK) GetPayload() *models.PageOfUpgradable {
	return o.Payload
}

func (o *GETUpgradablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfUpgradable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETUpgradablesInternalServerError creates a GETUpgradablesInternalServerError with default headers values
func NewGETUpgradablesInternalServerError() *GETUpgradablesInternalServerError {
	return &GETUpgradablesInternalServerError{}
}

/*
GETUpgradablesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETUpgradablesInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get upgradables internal server error response has a 2xx status code
func (o *GETUpgradablesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get upgradables internal server error response has a 3xx status code
func (o *GETUpgradablesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get upgradables internal server error response has a 4xx status code
func (o *GETUpgradablesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get upgradables internal server error response has a 5xx status code
func (o *GETUpgradablesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get upgradables internal server error response a status code equal to that given
func (o *GETUpgradablesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETUpgradablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUpgradablesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/upgradables][%d] getUpgradablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETUpgradablesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETUpgradablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}