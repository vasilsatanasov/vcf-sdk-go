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

// GetVasaProvidersReader is a Reader for the GetVasaProviders structure.
type GetVasaProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVasaProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVasaProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVasaProvidersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVasaProvidersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVasaProvidersOK creates a GetVasaProvidersOK with default headers values
func NewGetVasaProvidersOK() *GetVasaProvidersOK {
	return &GetVasaProvidersOK{}
}

/*
GetVasaProvidersOK describes a response with status code 200, with default header values.

Ok
*/
type GetVasaProvidersOK struct {
	Payload *models.PageOfVasaProvider
}

// IsSuccess returns true when this get vasa providers o k response has a 2xx status code
func (o *GetVasaProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vasa providers o k response has a 3xx status code
func (o *GetVasaProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers o k response has a 4xx status code
func (o *GetVasaProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa providers o k response has a 5xx status code
func (o *GetVasaProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa providers o k response a status code equal to that given
func (o *GetVasaProvidersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVasaProvidersOK) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersOK  %+v", 200, o.Payload)
}

func (o *GetVasaProvidersOK) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersOK  %+v", 200, o.Payload)
}

func (o *GetVasaProvidersOK) GetPayload() *models.PageOfVasaProvider {
	return o.Payload
}

func (o *GetVasaProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfVasaProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVasaProvidersBadRequest creates a GetVasaProvidersBadRequest with default headers values
func NewGetVasaProvidersBadRequest() *GetVasaProvidersBadRequest {
	return &GetVasaProvidersBadRequest{}
}

/*
GetVasaProvidersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetVasaProvidersBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa providers bad request response has a 2xx status code
func (o *GetVasaProvidersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa providers bad request response has a 3xx status code
func (o *GetVasaProvidersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers bad request response has a 4xx status code
func (o *GetVasaProvidersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vasa providers bad request response has a 5xx status code
func (o *GetVasaProvidersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get vasa providers bad request response a status code equal to that given
func (o *GetVasaProvidersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVasaProvidersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVasaProvidersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersBadRequest  %+v", 400, o.Payload)
}

func (o *GetVasaProvidersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVasaProvidersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVasaProvidersInternalServerError creates a GetVasaProvidersInternalServerError with default headers values
func NewGetVasaProvidersInternalServerError() *GetVasaProvidersInternalServerError {
	return &GetVasaProvidersInternalServerError{}
}

/*
GetVasaProvidersInternalServerError describes a response with status code 500, with default header values.

InternalServerError
*/
type GetVasaProvidersInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get vasa providers internal server error response has a 2xx status code
func (o *GetVasaProvidersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vasa providers internal server error response has a 3xx status code
func (o *GetVasaProvidersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vasa providers internal server error response has a 4xx status code
func (o *GetVasaProvidersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vasa providers internal server error response has a 5xx status code
func (o *GetVasaProvidersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get vasa providers internal server error response a status code equal to that given
func (o *GetVasaProvidersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVasaProvidersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVasaProvidersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/vasa-providers][%d] getVasaProvidersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVasaProvidersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVasaProvidersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
