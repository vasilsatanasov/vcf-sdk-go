// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterExpansionSpec This specification contains the parameters required to add one or more vSphere hosts to an existing cluster in a workload domain
//
// swagger:model ClusterExpansionSpec
type ClusterExpansionSpec struct {

	// Use to add host to a cluster with dead host(s). Bypasses validation of disconnected hosts and vSAN cluster health. Review recovery plan VMware Support before using. False if omitted. This property is deprecated and it has no effect when using it.
	ForceHostAdditionInPresenceofDeadHosts bool `json:"forceHostAdditionInPresenceofDeadHosts,omitempty"`

	// List of vSphere host information from the free pool to consume in the workload domain
	// Required: true
	HostSpecs []*HostSpec `json:"hostSpecs"`

	// Is inter-rack cluster(true for L2 non-uniform and L3 : At least one of management, uplink, Edge and host TEP networks is different for hosts of the cluster, false for L2 uniform :  All hosts in cluster have identical management, uplink, Edge and host TEP networks) expansion. Required, only if Cluster contains NSX-T Edge Cluster
	InterRackExpansion bool `json:"interRackExpansion,omitempty"`

	// Skip thumbprint validation for ESXi hosts during add host operation.
	// This property is deprecated.
	SkipThumbprintValidation bool `json:"skipThumbprintValidation,omitempty"`

	// vSAN Network Pool Spec
	VSANNetworkSpecs []*VSANNetworkSpec `json:"vsanNetworkSpecs"`

	// Witness host Info
	WitnessSpec *WitnessSpec `json:"witnessSpec,omitempty"`

	// Witness traffic to be shared with vSAN traffic
	WitnessTrafficSharedWithVSANTraffic bool `json:"witnessTrafficSharedWithVsanTraffic,omitempty"`
}

// Validate validates this cluster expansion spec
func (m *ClusterExpansionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVSANNetworkSpecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWitnessSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterExpansionSpec) validateHostSpecs(formats strfmt.Registry) error {

	if err := validate.Required("hostSpecs", "body", m.HostSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.HostSpecs); i++ {
		if swag.IsZero(m.HostSpecs[i]) { // not required
			continue
		}

		if m.HostSpecs[i] != nil {
			if err := m.HostSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterExpansionSpec) validateVSANNetworkSpecs(formats strfmt.Registry) error {
	if swag.IsZero(m.VSANNetworkSpecs) { // not required
		return nil
	}

	for i := 0; i < len(m.VSANNetworkSpecs); i++ {
		if swag.IsZero(m.VSANNetworkSpecs[i]) { // not required
			continue
		}

		if m.VSANNetworkSpecs[i] != nil {
			if err := m.VSANNetworkSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vsanNetworkSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vsanNetworkSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterExpansionSpec) validateWitnessSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.WitnessSpec) { // not required
		return nil
	}

	if m.WitnessSpec != nil {
		if err := m.WitnessSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witnessSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witnessSpec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster expansion spec based on the context it is used
func (m *ClusterExpansionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHostSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVSANNetworkSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWitnessSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterExpansionSpec) contextValidateHostSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HostSpecs); i++ {

		if m.HostSpecs[i] != nil {
			if err := m.HostSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hostSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterExpansionSpec) contextValidateVSANNetworkSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VSANNetworkSpecs); i++ {

		if m.VSANNetworkSpecs[i] != nil {
			if err := m.VSANNetworkSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vsanNetworkSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vsanNetworkSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterExpansionSpec) contextValidateWitnessSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.WitnessSpec != nil {
		if err := m.WitnessSpec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witnessSpec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witnessSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterExpansionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterExpansionSpec) UnmarshalBinary(b []byte) error {
	var res ClusterExpansionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
