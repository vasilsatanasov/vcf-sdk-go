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

// ComputeSpec This specification contains the parameters required to add each cluster to a workload domain
//
// swagger:model ComputeSpec
type ComputeSpec struct {

	// List of clusters to be added to workload domain
	// Required: true
	ClusterSpecs []*ClusterSpec `json:"clusterSpecs"`
}

// Validate validates this compute spec
func (m *ComputeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterSpecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeSpec) validateClusterSpecs(formats strfmt.Registry) error {

	if err := validate.Required("clusterSpecs", "body", m.ClusterSpecs); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterSpecs); i++ {
		if swag.IsZero(m.ClusterSpecs[i]) { // not required
			continue
		}

		if m.ClusterSpecs[i] != nil {
			if err := m.ClusterSpecs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this compute spec based on the context it is used
func (m *ComputeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterSpecs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeSpec) contextValidateClusterSpecs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterSpecs); i++ {

		if m.ClusterSpecs[i] != nil {
			if err := m.ClusterSpecs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterSpecs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterSpecs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeSpec) UnmarshalBinary(b []byte) error {
	var res ComputeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
