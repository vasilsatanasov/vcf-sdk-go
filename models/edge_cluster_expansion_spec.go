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

// EdgeClusterExpansionSpec This specification contains the parameters required to expand a NSX-T edge cluster.
//
// swagger:model EdgeClusterExpansionSpec
type EdgeClusterExpansionSpec struct {

	// List of names for the additional Tier-1(s) to be created during expansion
	AdditionalTier1Names []string `json:"additionalTier1Names"`

	// Edge Password for admin user
	// Required: true
	EdgeNodeAdminPassword *string `json:"edgeNodeAdminPassword"`

	// Edge Password for audit user
	// Required: true
	EdgeNodeAuditPassword *string `json:"edgeNodeAuditPassword"`

	// Edge Password for root user.
	// Required: true
	EdgeNodeRootPassword *string `json:"edgeNodeRootPassword"`

	// Specifications for Edge Node
	// Required: true
	EdgeNodeSpecs []*NsxTEdgeNodeSpec `json:"edgeNodeSpecs"`

	// Select whether all Tier-1(s) being created per this spec are hosted on the Edge cluster or not (default is false, meaning hosted)
	Tier1Unhosted bool `json:"tier1Unhosted,omitempty"`
}

// Validate validates this edge cluster expansion spec
func (m *EdgeClusterExpansionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeNodeAdminPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodeAuditPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodeRootPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeNodeSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterExpansionSpec) validateEdgeNodeAdminPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeAdminPassword", "body", m.EdgeNodeAdminPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterExpansionSpec) validateEdgeNodeAuditPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeAuditPassword", "body", m.EdgeNodeAuditPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterExpansionSpec) validateEdgeNodeRootPassword(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeRootPassword", "body", m.EdgeNodeRootPassword); err != nil {
		return err
	}

	return nil
}

func (m *EdgeClusterExpansionSpec) validateEdgeNodeSpecs(formats strfmt.Registry) error {

	if err := validate.Required("edgeNodeSpecs", "body", m.EdgeNodeSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.EdgeNodeSpecs); i++ {
		if swag.IsZero(m.EdgeNodeSpecs[i]) { // not required
			continue
		}

		if m.EdgeNodeSpecs[i] != nil {
			if err := m.EdgeNodeSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this edge cluster expansion spec based on the context it is used
func (m *EdgeClusterExpansionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeNodeSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterExpansionSpec) contextValidateEdgeNodeSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EdgeNodeSpecs); i++ {

		if m.EdgeNodeSpecs[i] != nil {
			if err := m.EdgeNodeSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edgeNodeSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterExpansionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterExpansionSpec) UnmarshalBinary(b []byte) error {
	var res EdgeClusterExpansionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
