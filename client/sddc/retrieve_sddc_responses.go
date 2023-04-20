// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package sddc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vcf-sdk-go/models"
)

// RetrieveSDDCReader is a Reader for the RetrieveSDDC structure.
type RetrieveSDDCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveSDDCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveSDDCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveSDDCBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrieveSDDCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrieveSDDCInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveSDDCOK creates a RetrieveSDDCOK with default headers values
func NewRetrieveSDDCOK() *RetrieveSDDCOK {
	return &RetrieveSDDCOK{}
}

/*
RetrieveSDDCOK describes a response with status code 200, with default header values.

OK
*/
type RetrieveSDDCOK struct {
	Payload *models.SDDCTask
}

// IsSuccess returns true when this retrieve Sddc o k response has a 2xx status code
func (o *RetrieveSDDCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retrieve Sddc o k response has a 3xx status code
func (o *RetrieveSDDCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve Sddc o k response has a 4xx status code
func (o *RetrieveSDDCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve Sddc o k response has a 5xx status code
func (o *RetrieveSDDCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve Sddc o k response a status code equal to that given
func (o *RetrieveSDDCOK) IsCode(code int) bool {
	return code == 200
}

func (o *RetrieveSDDCOK) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcOK  %+v", 200, o.Payload)
}

func (o *RetrieveSDDCOK) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcOK  %+v", 200, o.Payload)
}

func (o *RetrieveSDDCOK) GetPayload() *models.SDDCTask {
	return o.Payload
}

func (o *RetrieveSDDCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SDDCTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveSDDCBadRequest creates a RetrieveSDDCBadRequest with default headers values
func NewRetrieveSDDCBadRequest() *RetrieveSDDCBadRequest {
	return &RetrieveSDDCBadRequest{}
}

/*
RetrieveSDDCBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RetrieveSDDCBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this retrieve Sddc bad request response has a 2xx status code
func (o *RetrieveSDDCBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve Sddc bad request response has a 3xx status code
func (o *RetrieveSDDCBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve Sddc bad request response has a 4xx status code
func (o *RetrieveSDDCBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve Sddc bad request response has a 5xx status code
func (o *RetrieveSDDCBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve Sddc bad request response a status code equal to that given
func (o *RetrieveSDDCBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RetrieveSDDCBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveSDDCBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcBadRequest  %+v", 400, o.Payload)
}

func (o *RetrieveSDDCBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveSDDCBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveSDDCNotFound creates a RetrieveSDDCNotFound with default headers values
func NewRetrieveSDDCNotFound() *RetrieveSDDCNotFound {
	return &RetrieveSDDCNotFound{}
}

/*
RetrieveSDDCNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RetrieveSDDCNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this retrieve Sddc not found response has a 2xx status code
func (o *RetrieveSDDCNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve Sddc not found response has a 3xx status code
func (o *RetrieveSDDCNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve Sddc not found response has a 4xx status code
func (o *RetrieveSDDCNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this retrieve Sddc not found response has a 5xx status code
func (o *RetrieveSDDCNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this retrieve Sddc not found response a status code equal to that given
func (o *RetrieveSDDCNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RetrieveSDDCNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveSDDCNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveSDDCNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveSDDCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveSDDCInternalServerError creates a RetrieveSDDCInternalServerError with default headers values
func NewRetrieveSDDCInternalServerError() *RetrieveSDDCInternalServerError {
	return &RetrieveSDDCInternalServerError{}
}

/*
RetrieveSDDCInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RetrieveSDDCInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this retrieve Sddc internal server error response has a 2xx status code
func (o *RetrieveSDDCInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this retrieve Sddc internal server error response has a 3xx status code
func (o *RetrieveSDDCInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retrieve Sddc internal server error response has a 4xx status code
func (o *RetrieveSDDCInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this retrieve Sddc internal server error response has a 5xx status code
func (o *RetrieveSDDCInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this retrieve Sddc internal server error response a status code equal to that given
func (o *RetrieveSDDCInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RetrieveSDDCInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcInternalServerError  %+v", 500, o.Payload)
}

func (o *RetrieveSDDCInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sddcs/{id}][%d] retrieveSddcInternalServerError  %+v", 500, o.Payload)
}

func (o *RetrieveSDDCInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetrieveSDDCInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}