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

// NetworkDetailsSpec This specification contains network parameters required for new virtual machines being added to a workload domain
//
// swagger:model NetworkDetailsSpec
type NetworkDetailsSpec struct {

	// DNS name of the virtual machine, e.g., vc-1.domain1.vsphere.local
	DNSName string `json:"dnsName,omitempty"`

	// IPv4 gateway the VM can use to connect to the outside world
	Gateway string `json:"gateway,omitempty"`

	// IPv4 address of the virtual machine
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// Subnet mask
	SubnetMask string `json:"subnetMask,omitempty"`
}

// Validate validates this network details spec
func (m *NetworkDetailsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkDetailsSpec) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network details spec based on context it is used
func (m *NetworkDetailsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkDetailsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDetailsSpec) UnmarshalBinary(b []byte) error {
	var res NetworkDetailsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
