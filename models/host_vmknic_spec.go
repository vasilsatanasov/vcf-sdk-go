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

// HostVmknicSpec Spec contains parameters for ESXi Host Vmknic
//
// swagger:model HostVmknicSpec
type HostVmknicSpec struct {

	// Vmknic IP address
	// Max Length: 15
	// Min Length: 7
	IPAddress string `json:"ipAddress,omitempty"`

	// Vmknic mac address
	// Max Length: 18
	// Min Length: 18
	MacAddress string `json:"macAddress,omitempty"`

	// Portgroup
	// Example: One among:VSAN, VMOTION, MANAGEMENT
	// Required: true
	Portgroup *string `json:"portgroup"`
}

// Validate validates this host vmknic spec
func (m *HostVmknicSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortgroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostVmknicSpec) validateIPAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddress) { // not required
		return nil
	}

	if err := validate.MinLength("ipAddress", "body", m.IPAddress, 7); err != nil {
		return err
	}

	if err := validate.MaxLength("ipAddress", "body", m.IPAddress, 15); err != nil {
		return err
	}

	return nil
}

func (m *HostVmknicSpec) validateMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.MacAddress) { // not required
		return nil
	}

	if err := validate.MinLength("macAddress", "body", m.MacAddress, 18); err != nil {
		return err
	}

	if err := validate.MaxLength("macAddress", "body", m.MacAddress, 18); err != nil {
		return err
	}

	return nil
}

func (m *HostVmknicSpec) validatePortgroup(formats strfmt.Registry) error {

	if err := validate.Required("portgroup", "body", m.Portgroup); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host vmknic spec based on context it is used
func (m *HostVmknicSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostVmknicSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostVmknicSpec) UnmarshalBinary(b []byte) error {
	var res HostVmknicSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}