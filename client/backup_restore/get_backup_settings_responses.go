// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package backup_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETBackupSettingsReader is a Reader for the GETBackupSettings structure.
type GETBackupSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETBackupSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETBackupSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETBackupSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETBackupSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETBackupSettingsOK creates a GETBackupSettingsOK with default headers values
func NewGETBackupSettingsOK() *GETBackupSettingsOK {
	return &GETBackupSettingsOK{}
}

/*
GETBackupSettingsOK describes a response with status code 200, with default header values.

Ok
*/
type GETBackupSettingsOK struct {
	Payload *models.BackupConfiguration
}

// IsSuccess returns true when this get backup settings o k response has a 2xx status code
func (o *GETBackupSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup settings o k response has a 3xx status code
func (o *GETBackupSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings o k response has a 4xx status code
func (o *GETBackupSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup settings o k response has a 5xx status code
func (o *GETBackupSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup settings o k response a status code equal to that given
func (o *GETBackupSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETBackupSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *GETBackupSettingsOK) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsOK  %+v", 200, o.Payload)
}

func (o *GETBackupSettingsOK) GetPayload() *models.BackupConfiguration {
	return o.Payload
}

func (o *GETBackupSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETBackupSettingsBadRequest creates a GETBackupSettingsBadRequest with default headers values
func NewGETBackupSettingsBadRequest() *GETBackupSettingsBadRequest {
	return &GETBackupSettingsBadRequest{}
}

/*
GETBackupSettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETBackupSettingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup settings bad request response has a 2xx status code
func (o *GETBackupSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup settings bad request response has a 3xx status code
func (o *GETBackupSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings bad request response has a 4xx status code
func (o *GETBackupSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get backup settings bad request response has a 5xx status code
func (o *GETBackupSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup settings bad request response a status code equal to that given
func (o *GETBackupSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETBackupSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GETBackupSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GETBackupSettingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETBackupSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETBackupSettingsInternalServerError creates a GETBackupSettingsInternalServerError with default headers values
func NewGETBackupSettingsInternalServerError() *GETBackupSettingsInternalServerError {
	return &GETBackupSettingsInternalServerError{}
}

/*
GETBackupSettingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GETBackupSettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get backup settings internal server error response has a 2xx status code
func (o *GETBackupSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get backup settings internal server error response has a 3xx status code
func (o *GETBackupSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup settings internal server error response has a 4xx status code
func (o *GETBackupSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup settings internal server error response has a 5xx status code
func (o *GETBackupSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get backup settings internal server error response a status code equal to that given
func (o *GETBackupSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETBackupSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETBackupSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/system/backup-configuration][%d] getBackupSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETBackupSettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETBackupSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
