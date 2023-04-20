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

// PersonalityUploadSpecReferred Personality upload specification for uploading the personality from a referred vCentercluster. This mode of uplaoding personality is useful when the source vCenter cluster isinternal to the target VCF deployment.
//
// swagger:model PersonalityUploadSpecReferred
type PersonalityUploadSpecReferred struct {

	// Source cluster UUID from VCF inventory
	// Required: true
	ClusterID *string `json:"clusterId"`

	// v center Id
	// Required: true
	VCenterID *string `json:"vCenterId"`

	// vcenter Id
	VcenterID string `json:"vcenterId,omitempty"`
}

// Validate validates this personality upload spec referred
func (m *PersonalityUploadSpecReferred) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVCenterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PersonalityUploadSpecReferred) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *PersonalityUploadSpecReferred) validateVCenterID(formats strfmt.Registry) error {

	if err := validate.Required("vCenterId", "body", m.VCenterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this personality upload spec referred based on context it is used
func (m *PersonalityUploadSpecReferred) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PersonalityUploadSpecReferred) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PersonalityUploadSpecReferred) UnmarshalBinary(b []byte) error {
	var res PersonalityUploadSpecReferred
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}