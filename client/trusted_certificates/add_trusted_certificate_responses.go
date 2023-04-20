// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package trusted_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// AddTrustedCertificateReader is a Reader for the AddTrustedCertificate structure.
type AddTrustedCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTrustedCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTrustedCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddTrustedCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddTrustedCertificateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTrustedCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddTrustedCertificateOK creates a AddTrustedCertificateOK with default headers values
func NewAddTrustedCertificateOK() *AddTrustedCertificateOK {
	return &AddTrustedCertificateOK{}
}

/*
AddTrustedCertificateOK describes a response with status code 200, with default header values.

OK
*/
type AddTrustedCertificateOK struct {
}

// IsSuccess returns true when this add trusted certificate o k response has a 2xx status code
func (o *AddTrustedCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add trusted certificate o k response has a 3xx status code
func (o *AddTrustedCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add trusted certificate o k response has a 4xx status code
func (o *AddTrustedCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add trusted certificate o k response has a 5xx status code
func (o *AddTrustedCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add trusted certificate o k response a status code equal to that given
func (o *AddTrustedCertificateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AddTrustedCertificateOK) Error() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateOK ", 200)
}

func (o *AddTrustedCertificateOK) String() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateOK ", 200)
}

func (o *AddTrustedCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddTrustedCertificateBadRequest creates a AddTrustedCertificateBadRequest with default headers values
func NewAddTrustedCertificateBadRequest() *AddTrustedCertificateBadRequest {
	return &AddTrustedCertificateBadRequest{}
}

/*
AddTrustedCertificateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddTrustedCertificateBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this add trusted certificate bad request response has a 2xx status code
func (o *AddTrustedCertificateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add trusted certificate bad request response has a 3xx status code
func (o *AddTrustedCertificateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add trusted certificate bad request response has a 4xx status code
func (o *AddTrustedCertificateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add trusted certificate bad request response has a 5xx status code
func (o *AddTrustedCertificateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add trusted certificate bad request response a status code equal to that given
func (o *AddTrustedCertificateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddTrustedCertificateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *AddTrustedCertificateBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *AddTrustedCertificateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddTrustedCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTrustedCertificateConflict creates a AddTrustedCertificateConflict with default headers values
func NewAddTrustedCertificateConflict() *AddTrustedCertificateConflict {
	return &AddTrustedCertificateConflict{}
}

/*
AddTrustedCertificateConflict describes a response with status code 409, with default header values.

Trusted certificate already exists
*/
type AddTrustedCertificateConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this add trusted certificate conflict response has a 2xx status code
func (o *AddTrustedCertificateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add trusted certificate conflict response has a 3xx status code
func (o *AddTrustedCertificateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add trusted certificate conflict response has a 4xx status code
func (o *AddTrustedCertificateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add trusted certificate conflict response has a 5xx status code
func (o *AddTrustedCertificateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add trusted certificate conflict response a status code equal to that given
func (o *AddTrustedCertificateConflict) IsCode(code int) bool {
	return code == 409
}

func (o *AddTrustedCertificateConflict) Error() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateConflict  %+v", 409, o.Payload)
}

func (o *AddTrustedCertificateConflict) String() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateConflict  %+v", 409, o.Payload)
}

func (o *AddTrustedCertificateConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddTrustedCertificateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTrustedCertificateInternalServerError creates a AddTrustedCertificateInternalServerError with default headers values
func NewAddTrustedCertificateInternalServerError() *AddTrustedCertificateInternalServerError {
	return &AddTrustedCertificateInternalServerError{}
}

/*
AddTrustedCertificateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddTrustedCertificateInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this add trusted certificate internal server error response has a 2xx status code
func (o *AddTrustedCertificateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add trusted certificate internal server error response has a 3xx status code
func (o *AddTrustedCertificateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add trusted certificate internal server error response has a 4xx status code
func (o *AddTrustedCertificateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add trusted certificate internal server error response has a 5xx status code
func (o *AddTrustedCertificateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add trusted certificate internal server error response a status code equal to that given
func (o *AddTrustedCertificateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddTrustedCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTrustedCertificateInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/sddc-manager/trusted-certificates][%d] addTrustedCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *AddTrustedCertificateInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddTrustedCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}