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
)

// Domain Represents a domain
//
// swagger:model Domain
type Domain struct {

	// Capacity information for the workload domain
	Capacity *Capacity `json:"capacity,omitempty"`

	// List of clusters associated with the workload domain
	Clusters []*ClusterReference `json:"clusters"`

	// ID of the workload domain
	ID string `json:"id,omitempty"`

	// Shows whether the workload domain is joined to the Management domain SSO
	IsManagementSSODomain bool `json:"isManagementSsoDomain,omitempty"`

	// Name of the workload domain
	Name string `json:"name,omitempty"`

	// NSX-T cluster associated with the workload domain
	NSXTCluster *NsxTClusterReference `json:"nsxtCluster,omitempty"`

	// ID of the SSO domain associated with the workload domain
	SSOID string `json:"ssoId,omitempty"`

	// Name of the SSO domain associated with the workload domain
	SSOName string `json:"ssoName,omitempty"`

	// Status of the workload domain
	Status string `json:"status,omitempty"`

	// Deprecated, this list will always be returned empty
	Tags []*Tag `json:"tags"`

	// Type of the workload domain
	Type string `json:"type,omitempty"`

	// List of vCenters associated with the workload domain
	VCENTERS []*VcenterReference `json:"vcenters"`
}

// Validate validates this domain
func (m *Domain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNSXTCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVCENTERS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Domain) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *Domain) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Domain) validateNSXTCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.NSXTCluster) { // not required
		return nil
	}

	if m.NSXTCluster != nil {
		if err := m.NSXTCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtCluster")
			}
			return err
		}
	}

	return nil
}

func (m *Domain) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Domain) validateVCENTERS(formats strfmt.Registry) error {
	if swag.IsZero(m.VCENTERS) { // not required
		return nil
	}

	for i := 0; i < len(m.VCENTERS); i++ {
		if swag.IsZero(m.VCENTERS[i]) { // not required
			continue
		}

		if m.VCENTERS[i] != nil {
			if err := m.VCENTERS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcenters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcenters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain based on the context it is used
func (m *Domain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNSXTCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVCENTERS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Domain) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {
		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *Domain) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Domain) contextValidateNSXTCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.NSXTCluster != nil {
		if err := m.NSXTCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxtCluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nsxtCluster")
			}
			return err
		}
	}

	return nil
}

func (m *Domain) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Domain) contextValidateVCENTERS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VCENTERS); i++ {

		if m.VCENTERS[i] != nil {
			if err := m.VCENTERS[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcenters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcenters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Domain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Domain) UnmarshalBinary(b []byte) error {
	var res Domain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}