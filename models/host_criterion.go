// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostCriterion Represents a criterion for querying the host
//
// swagger:model HostCriterion
type HostCriterion struct {

	// Arguments required for a particular criterion
	Arguments map[string]string `json:"arguments,omitempty"`

	// Description of the criterion
	Description string `json:"description,omitempty"`

	// Name of the criterion
	// Example: One among: HOST_COMPATIBLE_WITH_CLUSTER_USING_PNICS, UNMANAGED_HOSTS_IN_VCENTER, UNMANAGED_HOSTS_IN_HCIMGR
	Name string `json:"name,omitempty"`
}

// Validate validates this host criterion
func (m *HostCriterion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host criterion based on context it is used
func (m *HostCriterion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostCriterion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostCriterion) UnmarshalBinary(b []byte) error {
	var res HostCriterion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}