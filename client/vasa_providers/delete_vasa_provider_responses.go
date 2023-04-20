// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package vasa_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// DeleteVasaProviderReader is a Reader for the DeleteVasaProvider structure.
type DeleteVasaProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVasaProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteVasaProviderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVasaProviderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVasaProviderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVasaProviderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVasaProviderNoContent creates a DeleteVasaProviderNoContent with default headers values
func NewDeleteVasaProviderNoContent() *DeleteVasaProviderNoContent {
	return &DeleteVasaProviderNoContent{}
}

/*
DeleteVasaProviderNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteVasaProviderNoContent struct {
}

// IsSuccess returns true when this delete vasa provider no content response has a 2xx status code
func (o *DeleteVasaProviderNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete vasa provider no content response has a 3xx status code
func (o *DeleteVasaProviderNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vasa provider no content response has a 4xx status code
func (o *DeleteVasaProviderNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete vasa provider no content response has a 5xx status code
func (o *DeleteVasaProviderNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vasa provider no content response a status code equal to that given
func (o *DeleteVasaProviderNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteVasaProviderNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderNoContent ", 204)
}

func (o *DeleteVasaProviderNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderNoContent ", 204)
}

func (o *DeleteVasaProviderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVasaProviderBadRequest creates a DeleteVasaProviderBadRequest with default headers values
func NewDeleteVasaProviderBadRequest() *DeleteVasaProviderBadRequest {
	return &DeleteVasaProviderBadRequest{}
}

/*
DeleteVasaProviderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteVasaProviderBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete vasa provider bad request response has a 2xx status code
func (o *DeleteVasaProviderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete vasa provider bad request response has a 3xx status code
func (o *DeleteVasaProviderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vasa provider bad request response has a 4xx status code
func (o *DeleteVasaProviderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete vasa provider bad request response has a 5xx status code
func (o *DeleteVasaProviderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vasa provider bad request response a status code equal to that given
func (o *DeleteVasaProviderBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteVasaProviderBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVasaProviderBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVasaProviderBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVasaProviderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVasaProviderNotFound creates a DeleteVasaProviderNotFound with default headers values
func NewDeleteVasaProviderNotFound() *DeleteVasaProviderNotFound {
	return &DeleteVasaProviderNotFound{}
}

/*
DeleteVasaProviderNotFound describes a response with status code 404, with default header values.

VASA Provider not found
*/
type DeleteVasaProviderNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete vasa provider not found response has a 2xx status code
func (o *DeleteVasaProviderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete vasa provider not found response has a 3xx status code
func (o *DeleteVasaProviderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vasa provider not found response has a 4xx status code
func (o *DeleteVasaProviderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete vasa provider not found response has a 5xx status code
func (o *DeleteVasaProviderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vasa provider not found response a status code equal to that given
func (o *DeleteVasaProviderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteVasaProviderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVasaProviderNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVasaProviderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVasaProviderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVasaProviderInternalServerError creates a DeleteVasaProviderInternalServerError with default headers values
func NewDeleteVasaProviderInternalServerError() *DeleteVasaProviderInternalServerError {
	return &DeleteVasaProviderInternalServerError{}
}

/*
DeleteVasaProviderInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type DeleteVasaProviderInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete vasa provider internal server error response has a 2xx status code
func (o *DeleteVasaProviderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete vasa provider internal server error response has a 3xx status code
func (o *DeleteVasaProviderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vasa provider internal server error response has a 4xx status code
func (o *DeleteVasaProviderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete vasa provider internal server error response has a 5xx status code
func (o *DeleteVasaProviderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete vasa provider internal server error response a status code equal to that given
func (o *DeleteVasaProviderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteVasaProviderInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVasaProviderInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/vasa-providers/{id}][%d] deleteVasaProviderInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVasaProviderInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVasaProviderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}