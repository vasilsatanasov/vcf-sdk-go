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

// PortgroupSpec This specification contains vCenter port group configurations
//
// swagger:model PortgroupSpec
type PortgroupSpec struct {

	// List of active uplinks associated with portgroup. This is only supported for VxRail.
	ActiveUplinks []string `json:"activeUplinks"`

	// Port group name
	// Required: true
	Name *string `json:"name"`

	// Port group transport type
	// Example: One among: VSAN, VMOTION, MANAGEMENT, PUBLIC, NFS, VREALIZE, ISCSI, EDGE_INFRA_OVERLAY_UPLINK
	// Required: true
	TransportType *string `json:"transportType"`
}

// Validate validates this portgroup spec
func (m *PortgroupSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortgroupSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PortgroupSpec) validateTransportType(formats strfmt.Registry) error {

	if err := validate.Required("transportType", "body", m.TransportType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this portgroup spec based on context it is used
func (m *PortgroupSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortgroupSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortgroupSpec) UnmarshalBinary(b []byte) error {
	var res PortgroupSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}