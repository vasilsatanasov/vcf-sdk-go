// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GetCredentialReader is a Reader for the GetCredential structure.
type GetCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCredentialOK creates a GetCredentialOK with default headers values
func NewGetCredentialOK() *GetCredentialOK {
	return &GetCredentialOK{}
}

/*
GetCredentialOK describes a response with status code 200, with default header values.

OK
*/
type GetCredentialOK struct {
	Payload *models.Credential
}

// IsSuccess returns true when this get credential o k response has a 2xx status code
func (o *GetCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credential o k response has a 3xx status code
func (o *GetCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential o k response has a 4xx status code
func (o *GetCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credential o k response has a 5xx status code
func (o *GetCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential o k response a status code equal to that given
func (o *GetCredentialOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCredentialOK) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialOK  %+v", 200, o.Payload)
}

func (o *GetCredentialOK) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialOK  %+v", 200, o.Payload)
}

func (o *GetCredentialOK) GetPayload() *models.Credential {
	return o.Payload
}

func (o *GetCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialBadRequest creates a GetCredentialBadRequest with default headers values
func NewGetCredentialBadRequest() *GetCredentialBadRequest {
	return &GetCredentialBadRequest{}
}

/*
GetCredentialBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCredentialBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credential bad request response has a 2xx status code
func (o *GetCredentialBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credential bad request response has a 3xx status code
func (o *GetCredentialBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential bad request response has a 4xx status code
func (o *GetCredentialBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credential bad request response has a 5xx status code
func (o *GetCredentialBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential bad request response a status code equal to that given
func (o *GetCredentialBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCredentialBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *GetCredentialBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *GetCredentialBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialUnauthorized creates a GetCredentialUnauthorized with default headers values
func NewGetCredentialUnauthorized() *GetCredentialUnauthorized {
	return &GetCredentialUnauthorized{}
}

/*
GetCredentialUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCredentialUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credential unauthorized response has a 2xx status code
func (o *GetCredentialUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credential unauthorized response has a 3xx status code
func (o *GetCredentialUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential unauthorized response has a 4xx status code
func (o *GetCredentialUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credential unauthorized response has a 5xx status code
func (o *GetCredentialUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential unauthorized response a status code equal to that given
func (o *GetCredentialUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCredentialUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCredentialUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialForbidden creates a GetCredentialForbidden with default headers values
func NewGetCredentialForbidden() *GetCredentialForbidden {
	return &GetCredentialForbidden{}
}

/*
GetCredentialForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCredentialForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credential forbidden response has a 2xx status code
func (o *GetCredentialForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credential forbidden response has a 3xx status code
func (o *GetCredentialForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential forbidden response has a 4xx status code
func (o *GetCredentialForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credential forbidden response has a 5xx status code
func (o *GetCredentialForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential forbidden response a status code equal to that given
func (o *GetCredentialForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCredentialForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialForbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialForbidden) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialForbidden  %+v", 403, o.Payload)
}

func (o *GetCredentialForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialNotFound creates a GetCredentialNotFound with default headers values
func NewGetCredentialNotFound() *GetCredentialNotFound {
	return &GetCredentialNotFound{}
}

/*
GetCredentialNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCredentialNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credential not found response has a 2xx status code
func (o *GetCredentialNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credential not found response has a 3xx status code
func (o *GetCredentialNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential not found response has a 4xx status code
func (o *GetCredentialNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get credential not found response has a 5xx status code
func (o *GetCredentialNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential not found response a status code equal to that given
func (o *GetCredentialNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCredentialNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialNotFound  %+v", 404, o.Payload)
}

func (o *GetCredentialNotFound) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialNotFound  %+v", 404, o.Payload)
}

func (o *GetCredentialNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialInternalServerError creates a GetCredentialInternalServerError with default headers values
func NewGetCredentialInternalServerError() *GetCredentialInternalServerError {
	return &GetCredentialInternalServerError{}
}

/*
GetCredentialInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCredentialInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credential internal server error response has a 2xx status code
func (o *GetCredentialInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credential internal server error response has a 3xx status code
func (o *GetCredentialInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential internal server error response has a 4xx status code
func (o *GetCredentialInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credential internal server error response has a 5xx status code
func (o *GetCredentialInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get credential internal server error response a status code equal to that given
func (o *GetCredentialInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/credentials/{id}][%d] getCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
