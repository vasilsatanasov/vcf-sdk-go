// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupResource Element specification to trigger backup
//
// swagger:model BackupResource
type BackupResource struct {

	// Resource type
	// Example: SDDC_MANAGER
	// Required: true
	ResourceType *string `json:"resourceType"`
}

// Validate validates this backup resource
func (m *BackupResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupResource) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup resource based on context it is used
func (m *BackupResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupResource) UnmarshalBinary(b []byte) error {
	var res BackupResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}