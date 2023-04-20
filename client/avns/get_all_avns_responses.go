// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package avns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETAllAvnsReader is a Reader for the GETAllAvns structure.
type GETAllAvnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETAllAvnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETAllAvnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGETAllAvnsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETAllAvnsOK creates a GETAllAvnsOK with default headers values
func NewGETAllAvnsOK() *GETAllAvnsOK {
	return &GETAllAvnsOK{}
}

/*
GETAllAvnsOK describes a response with status code 200, with default header values.

Returns the list of all matching AVNs
*/
type GETAllAvnsOK struct {
	Payload []*models.Avn
}

// IsSuccess returns true when this get all avns o k response has a 2xx status code
func (o *GETAllAvnsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all avns o k response has a 3xx status code
func (o *GETAllAvnsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all avns o k response has a 4xx status code
func (o *GETAllAvnsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all avns o k response has a 5xx status code
func (o *GETAllAvnsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all avns o k response a status code equal to that given
func (o *GETAllAvnsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETAllAvnsOK) Error() string {
	return fmt.Sprintf("[GET /v1/avns][%d] getAllAvnsOK  %+v", 200, o.Payload)
}

func (o *GETAllAvnsOK) String() string {
	return fmt.Sprintf("[GET /v1/avns][%d] getAllAvnsOK  %+v", 200, o.Payload)
}

func (o *GETAllAvnsOK) GetPayload() []*models.Avn {
	return o.Payload
}

func (o *GETAllAvnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETAllAvnsInternalServerError creates a GETAllAvnsInternalServerError with default headers values
func NewGETAllAvnsInternalServerError() *GETAllAvnsInternalServerError {
	return &GETAllAvnsInternalServerError{}
}

/*
GETAllAvnsInternalServerError describes a response with status code 500, with default header values.

Unexpected error
*/
type GETAllAvnsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get all avns internal server error response has a 2xx status code
func (o *GETAllAvnsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all avns internal server error response has a 3xx status code
func (o *GETAllAvnsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all avns internal server error response has a 4xx status code
func (o *GETAllAvnsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all avns internal server error response has a 5xx status code
func (o *GETAllAvnsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all avns internal server error response a status code equal to that given
func (o *GETAllAvnsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETAllAvnsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/avns][%d] getAllAvnsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETAllAvnsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/avns][%d] getAllAvnsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETAllAvnsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETAllAvnsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}