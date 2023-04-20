// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package resource_functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// GETResourcesFunctionalitiesAllowedGlobalReader is a Reader for the GETResourcesFunctionalitiesAllowedGlobal structure.
type GETResourcesFunctionalitiesAllowedGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETResourcesFunctionalitiesAllowedGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGETResourcesFunctionalitiesAllowedGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGETResourcesFunctionalitiesAllowedGlobalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGETResourcesFunctionalitiesAllowedGlobalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGETResourcesFunctionalitiesAllowedGlobalOK creates a GETResourcesFunctionalitiesAllowedGlobalOK with default headers values
func NewGETResourcesFunctionalitiesAllowedGlobalOK() *GETResourcesFunctionalitiesAllowedGlobalOK {
	return &GETResourcesFunctionalitiesAllowedGlobalOK{}
}

/*
GETResourcesFunctionalitiesAllowedGlobalOK describes a response with status code 200, with default header values.

OK
*/
type GETResourcesFunctionalitiesAllowedGlobalOK struct {
	Payload *models.ResourceFunctionalitiesGlobalConfiguration
}

// IsSuccess returns true when this get resources functionalities allowed global o k response has a 2xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources functionalities allowed global o k response has a 3xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global o k response has a 4xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources functionalities allowed global o k response has a 5xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources functionalities allowed global o k response a status code equal to that given
func (o *GETResourcesFunctionalitiesAllowedGlobalOK) IsCode(code int) bool {
	return code == 200
}

func (o *GETResourcesFunctionalitiesAllowedGlobalOK) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalOK  %+v", 200, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalOK) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalOK  %+v", 200, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalOK) GetPayload() *models.ResourceFunctionalitiesGlobalConfiguration {
	return o.Payload
}

func (o *GETResourcesFunctionalitiesAllowedGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceFunctionalitiesGlobalConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourcesFunctionalitiesAllowedGlobalBadRequest creates a GETResourcesFunctionalitiesAllowedGlobalBadRequest with default headers values
func NewGETResourcesFunctionalitiesAllowedGlobalBadRequest() *GETResourcesFunctionalitiesAllowedGlobalBadRequest {
	return &GETResourcesFunctionalitiesAllowedGlobalBadRequest{}
}

/*
GETResourcesFunctionalitiesAllowedGlobalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GETResourcesFunctionalitiesAllowedGlobalBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resources functionalities allowed global bad request response has a 2xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources functionalities allowed global bad request response has a 3xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global bad request response has a 4xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources functionalities allowed global bad request response has a 5xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources functionalities allowed global bad request response a status code equal to that given
func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalBadRequest  %+v", 400, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourcesFunctionalitiesAllowedGlobalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETResourcesFunctionalitiesAllowedGlobalInternalServerError creates a GETResourcesFunctionalitiesAllowedGlobalInternalServerError with default headers values
func NewGETResourcesFunctionalitiesAllowedGlobalInternalServerError() *GETResourcesFunctionalitiesAllowedGlobalInternalServerError {
	return &GETResourcesFunctionalitiesAllowedGlobalInternalServerError{}
}

/*
GETResourcesFunctionalitiesAllowedGlobalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GETResourcesFunctionalitiesAllowedGlobalInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resources functionalities allowed global internal server error response has a 2xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources functionalities allowed global internal server error response has a 3xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources functionalities allowed global internal server error response has a 4xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources functionalities allowed global internal server error response has a 5xx status code
func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resources functionalities allowed global internal server error response a status code equal to that given
func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/resource-functionalities/global][%d] getResourcesFunctionalitiesAllowedGlobalInternalServerError  %+v", 500, o.Payload)
}

func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GETResourcesFunctionalitiesAllowedGlobalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}