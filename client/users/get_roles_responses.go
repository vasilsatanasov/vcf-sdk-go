// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETRolesReader is a Reader for the GETRoles structure.
type GETRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGETRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETRolesOK creates a GETRolesOK with default headers values
func NewGETRolesOK() *GETRolesOK {
	return &GETRolesOK{}
}

/*
GETRolesOK describes a response with status code 200, with default header values.

OK
*/
type GETRolesOK struct {
	Payload *models.PageOfRole
}

// IsSuccess returns true when this get roles o k response has a 2xx status code
func (o *GETRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles o k response has a 3xx status code
func (o *GETRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles o k response has a 4xx status code
func (o *GETRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles o k response has a 5xx status code
func (o *GETRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles o k response a status code equal to that given
func (o *GETRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETRolesOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GETRolesOK) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesOK  %+v", 200, o.Payload)
}

func (o *GETRolesOK) GetPayload() *models.PageOfRole {
	return o.Payload
}

func (o *GETRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRolesBadRequest creates a GETRolesBadRequest with default headers values
func NewGETRolesBadRequest() *GETRolesBadRequest {
	return &GETRolesBadRequest{}
}

/*
GETRolesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GETRolesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles bad request response has a 2xx status code
func (o *GETRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles bad request response has a 3xx status code
func (o *GETRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles bad request response has a 4xx status code
func (o *GETRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles bad request response has a 5xx status code
func (o *GETRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles bad request response a status code equal to that given
func (o *GETRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GETRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GETRolesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRolesUnauthorized creates a GETRolesUnauthorized with default headers values
func NewGETRolesUnauthorized() *GETRolesUnauthorized {
	return &GETRolesUnauthorized{}
}

/*
GETRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GETRolesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles unauthorized response has a 2xx status code
func (o *GETRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles unauthorized response has a 3xx status code
func (o *GETRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles unauthorized response has a 4xx status code
func (o *GETRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles unauthorized response has a 5xx status code
func (o *GETRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles unauthorized response a status code equal to that given
func (o *GETRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GETRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GETRolesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETRolesInternalServerError creates a GETRolesInternalServerError with default headers values
func NewGETRolesInternalServerError() *GETRolesInternalServerError {
	return &GETRolesInternalServerError{}
}

/*
GETRolesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETRolesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles internal server error response has a 2xx status code
func (o *GETRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles internal server error response has a 3xx status code
func (o *GETRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles internal server error response has a 4xx status code
func (o *GETRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles internal server error response has a 5xx status code
func (o *GETRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get roles internal server error response a status code equal to that given
func (o *GETRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/roles][%d] getRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GETRolesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GETRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
