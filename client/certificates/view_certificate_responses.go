// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// ViewCertificateReader is a Reader for the ViewCertificate structure.
type ViewCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewViewCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewViewCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewViewCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewViewCertificateOK creates a ViewCertificateOK with default headers values
func NewViewCertificateOK() *ViewCertificateOK {
	return &ViewCertificateOK{}
}

/*
ViewCertificateOK describes a response with status code 200, with default header values.

OK
*/
type ViewCertificateOK struct {
	Payload *models.PageOfCertificate
}

// IsSuccess returns true when this view certificate o k response has a 2xx status code
func (o *ViewCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this view certificate o k response has a 3xx status code
func (o *ViewCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view certificate o k response has a 4xx status code
func (o *ViewCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this view certificate o k response has a 5xx status code
func (o *ViewCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this view certificate o k response a status code equal to that given
func (o *ViewCertificateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ViewCertificateOK) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateOK  %+v", 200, o.Payload)
}

func (o *ViewCertificateOK) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateOK  %+v", 200, o.Payload)
}

func (o *ViewCertificateOK) GetPayload() *models.PageOfCertificate {
	return o.Payload
}

func (o *ViewCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewCertificateNotFound creates a ViewCertificateNotFound with default headers values
func NewViewCertificateNotFound() *ViewCertificateNotFound {
	return &ViewCertificateNotFound{}
}

/*
ViewCertificateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ViewCertificateNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this view certificate not found response has a 2xx status code
func (o *ViewCertificateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view certificate not found response has a 3xx status code
func (o *ViewCertificateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view certificate not found response has a 4xx status code
func (o *ViewCertificateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this view certificate not found response has a 5xx status code
func (o *ViewCertificateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this view certificate not found response a status code equal to that given
func (o *ViewCertificateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ViewCertificateNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateNotFound  %+v", 404, o.Payload)
}

func (o *ViewCertificateNotFound) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateNotFound  %+v", 404, o.Payload)
}

func (o *ViewCertificateNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewCertificateInternalServerError creates a ViewCertificateInternalServerError with default headers values
func NewViewCertificateInternalServerError() *ViewCertificateInternalServerError {
	return &ViewCertificateInternalServerError{}
}

/*
ViewCertificateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ViewCertificateInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this view certificate internal server error response has a 2xx status code
func (o *ViewCertificateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this view certificate internal server error response has a 3xx status code
func (o *ViewCertificateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this view certificate internal server error response has a 4xx status code
func (o *ViewCertificateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this view certificate internal server error response has a 5xx status code
func (o *ViewCertificateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this view certificate internal server error response a status code equal to that given
func (o *ViewCertificateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ViewCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *ViewCertificateInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/domains/{domainName}/resource-certificates][%d] viewCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *ViewCertificateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ViewCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}