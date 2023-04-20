// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vrslcm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETVRSLCMReader is a Reader for the GETVRSLCM structure.
type GETVRSLCMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETVRSLCMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETVRSLCMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGETVRSLCMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETVRSLCMOK creates a GETVRSLCMOK with default headers values
func NewGETVRSLCMOK() *GETVRSLCMOK {
	return &GETVRSLCMOK{}
}

/*
GETVRSLCMOK describes a response with status code 200, with default header values.

OK
*/
type GETVRSLCMOK struct {
	Payload *models.VRSLCM
}

// IsSuccess returns true when this get Vrslcm o k response has a 2xx status code
func (o *GETVRSLCMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Vrslcm o k response has a 3xx status code
func (o *GETVRSLCMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vrslcm o k response has a 4xx status code
func (o *GETVRSLCMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vrslcm o k response has a 5xx status code
func (o *GETVRSLCMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vrslcm o k response a status code equal to that given
func (o *GETVRSLCMOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETVRSLCMOK) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcm][%d] getVrslcmOK  %+v", 200, o.Payload)
}

func (o *GETVRSLCMOK) String() string {
	return fmt.Sprintf("[GET /v1/vrslcm][%d] getVrslcmOK  %+v", 200, o.Payload)
}

func (o *GETVRSLCMOK) GetPayload() *models.VRSLCM {
	return o.Payload
}

func (o *GETVRSLCMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRSLCM)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETVRSLCMNotFound creates a GETVRSLCMNotFound with default headers values
func NewGETVRSLCMNotFound() *GETVRSLCMNotFound {
	return &GETVRSLCMNotFound{}
}

/*
GETVRSLCMNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GETVRSLCMNotFound struct {
}

// IsSuccess returns true when this get Vrslcm not found response has a 2xx status code
func (o *GETVRSLCMNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Vrslcm not found response has a 3xx status code
func (o *GETVRSLCMNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vrslcm not found response has a 4xx status code
func (o *GETVRSLCMNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Vrslcm not found response has a 5xx status code
func (o *GETVRSLCMNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vrslcm not found response a status code equal to that given
func (o *GETVRSLCMNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GETVRSLCMNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/vrslcm][%d] getVrslcmNotFound ", 404)
}

func (o *GETVRSLCMNotFound) String() string {
	return fmt.Sprintf("[GET /v1/vrslcm][%d] getVrslcmNotFound ", 404)
}

func (o *GETVRSLCMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}