// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package upgrades

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// CommitRescheduleUpgradeReader is a Reader for the CommitRescheduleUpgrade structure.
type CommitRescheduleUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitRescheduleUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitRescheduleUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCommitRescheduleUpgradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCommitRescheduleUpgradeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCommitRescheduleUpgradeOK creates a CommitRescheduleUpgradeOK with default headers values
func NewCommitRescheduleUpgradeOK() *CommitRescheduleUpgradeOK {
	return &CommitRescheduleUpgradeOK{}
}

/*
CommitRescheduleUpgradeOK describes a response with status code 200, with default header values.

Ok
*/
type CommitRescheduleUpgradeOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this commit reschedule upgrade o k response has a 2xx status code
func (o *CommitRescheduleUpgradeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this commit reschedule upgrade o k response has a 3xx status code
func (o *CommitRescheduleUpgradeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit reschedule upgrade o k response has a 4xx status code
func (o *CommitRescheduleUpgradeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit reschedule upgrade o k response has a 5xx status code
func (o *CommitRescheduleUpgradeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this commit reschedule upgrade o k response a status code equal to that given
func (o *CommitRescheduleUpgradeOK) IsCode(code int) bool {
	return code == 200
}

func (o *CommitRescheduleUpgradeOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeOK  %+v", 200, o.Payload)
}

func (o *CommitRescheduleUpgradeOK) String() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeOK  %+v", 200, o.Payload)
}

func (o *CommitRescheduleUpgradeOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *CommitRescheduleUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitRescheduleUpgradeNotFound creates a CommitRescheduleUpgradeNotFound with default headers values
func NewCommitRescheduleUpgradeNotFound() *CommitRescheduleUpgradeNotFound {
	return &CommitRescheduleUpgradeNotFound{}
}

/*
CommitRescheduleUpgradeNotFound describes a response with status code 404, with default header values.

Upgrade not found
*/
type CommitRescheduleUpgradeNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this commit reschedule upgrade not found response has a 2xx status code
func (o *CommitRescheduleUpgradeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit reschedule upgrade not found response has a 3xx status code
func (o *CommitRescheduleUpgradeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit reschedule upgrade not found response has a 4xx status code
func (o *CommitRescheduleUpgradeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this commit reschedule upgrade not found response has a 5xx status code
func (o *CommitRescheduleUpgradeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this commit reschedule upgrade not found response a status code equal to that given
func (o *CommitRescheduleUpgradeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CommitRescheduleUpgradeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeNotFound  %+v", 404, o.Payload)
}

func (o *CommitRescheduleUpgradeNotFound) String() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeNotFound  %+v", 404, o.Payload)
}

func (o *CommitRescheduleUpgradeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitRescheduleUpgradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitRescheduleUpgradeInternalServerError creates a CommitRescheduleUpgradeInternalServerError with default headers values
func NewCommitRescheduleUpgradeInternalServerError() *CommitRescheduleUpgradeInternalServerError {
	return &CommitRescheduleUpgradeInternalServerError{}
}

/*
CommitRescheduleUpgradeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CommitRescheduleUpgradeInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this commit reschedule upgrade internal server error response has a 2xx status code
func (o *CommitRescheduleUpgradeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this commit reschedule upgrade internal server error response has a 3xx status code
func (o *CommitRescheduleUpgradeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this commit reschedule upgrade internal server error response has a 4xx status code
func (o *CommitRescheduleUpgradeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this commit reschedule upgrade internal server error response has a 5xx status code
func (o *CommitRescheduleUpgradeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this commit reschedule upgrade internal server error response a status code equal to that given
func (o *CommitRescheduleUpgradeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CommitRescheduleUpgradeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeInternalServerError  %+v", 500, o.Payload)
}

func (o *CommitRescheduleUpgradeInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /v1/upgrades/{upgradeId}][%d] commitRescheduleUpgradeInternalServerError  %+v", 500, o.Payload)
}

func (o *CommitRescheduleUpgradeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CommitRescheduleUpgradeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}