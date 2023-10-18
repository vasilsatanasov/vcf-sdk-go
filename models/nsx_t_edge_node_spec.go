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

// NsxTEdgeNodeSpec This specification contains configuration inputs required to deploy and configure an edge node
//
// swagger:model NsxTEdgeNodeSpec
type NsxTEdgeNodeSpec struct {

	// Cluster on which the edge needs to be deployed
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Edge Node Name
	// Required: true
	EdgeNodeName *string `json:"edgeNodeName"`

	// Edge TEP 1 IP
	// Required: true
	EdgeTep1IP *string `json:"edgeTep1IP"`

	// Edge TEP 2 IP
	// Required: true
	EdgeTep2IP *string `json:"edgeTep2IP"`

	// Edge TEP Gateway IP
	// Required: true
	EdgeTepGateway *string `json:"edgeTepGateway"`

	// Edge TEP VLAN
	// Required: true
	EdgeTepVlan *int32 `json:"edgeTepVlan"`

	// First NSX enabled VDS uplink for the Edge node
	// Example: One among: uplink1, uplink2, uplink3, uplink4
	FirstNsxVdsUplink string `json:"firstNsxVdsUplink,omitempty"`

	// Is inter-rack cluster(true for L2 non-uniform and L3 : At least one of management, uplink, Edge and host TEP networks is different for hosts of the cluster, false for L2 uniform :   All hosts in cluster have identical management, uplink, Edge and host TEP networks)
	// Required: true
	InterRackCluster *bool `json:"interRackCluster"`

	// Management Gateway IP
	// Required: true
	ManagementGateway *string `json:"managementGateway"`

	// Management Interface IP
	// Required: true
	ManagementIP *string `json:"managementIP"`

	// Second NSX enabled VDS uplink for the Edge node
	// Example: One among: uplink1, uplink2, uplink3, uplink4
	SecondNsxVdsUplink string `json:"secondNsxVdsUplink,omitempty"`

	// Specifications of Tier0 uplinks for the Edge Node
	UplinkNetwork []*NsxTEdgeUplinkNetwork `json:"uplinkNetwork"`
}

// Validate validates this nsx t edge node spec
func (m *NsxTEdgeNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeTep1IP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeTep2IP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeTepGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeTepVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterRackCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUplinkNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTEdgeNodeSpec) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateEdgeNodeName(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeName", "body", m.EdgeNodeName); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateEdgeTep1IP(formats strfmt.Registry) error {

	if err := validate.Required("edgeTep1IP", "body", m.EdgeTep1IP); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateEdgeTep2IP(formats strfmt.Registry) error {

	if err := validate.Required("edgeTep2IP", "body", m.EdgeTep2IP); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateEdgeTepGateway(formats strfmt.Registry) error {

	if err := validate.Required("edgeTepGateway", "body", m.EdgeTepGateway); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateEdgeTepVlan(formats strfmt.Registry) error {

	if err := validate.Required("edgeTepVlan", "body", m.EdgeTepVlan); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateInterRackCluster(formats strfmt.Registry) error {

	if err := validate.Required("interRackCluster", "body", m.InterRackCluster); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateManagementGateway(formats strfmt.Registry) error {

	if err := validate.Required("managementGateway", "body", m.ManagementGateway); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("managementIP", "body", m.ManagementIP); err != nil {
		return err
	}

	return nil
}

func (m *NsxTEdgeNodeSpec) validateUplinkNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.UplinkNetwork) { // not required
		return nil
	}

	for i := 0; i < len(m.UplinkNetwork); i++ {
		if swag.IsZero(m.UplinkNetwork[i]) { // not required
			continue
		}

		if m.UplinkNetwork[i] != nil {
			if err := m.UplinkNetwork[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uplinkNetwork" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uplinkNetwork" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nsx t edge node spec based on the context it is used
func (m *NsxTEdgeNodeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUplinkNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NsxTEdgeNodeSpec) contextValidateUplinkNetwork(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UplinkNetwork); i++ {

		if m.UplinkNetwork[i] != nil {
			if err := m.UplinkNetwork[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("uplinkNetwork" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("uplinkNetwork" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NsxTEdgeNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NsxTEdgeNodeSpec) UnmarshalBinary(b []byte) error {
	var res NsxTEdgeNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
