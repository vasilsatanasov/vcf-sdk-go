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

// NSXTIPAddressPool NSX-T IP address pool representation
//
// swagger:model NsxtIpAddressPool
type NSXTIPAddressPool struct {

	// The number of IP addresses available in the IP address pool
	AvailableIPAddresses int32 `json:"availableIpAddresses,omitempty"`

	// The list of IP address pool block subnets
	BlockSubnets []*NSXTIPAddressPoolBlockSubnet `json:"blockSubnets"`

	// Description of the IP address pool
	Description string `json:"description,omitempty"`

	// Name of the IP address pool
	Name string `json:"name,omitempty"`

	// The list of IP address pool static subnets
	StaticSubnets []*NSXTIPAddressPoolStaticSubnet `json:"staticSubnets"`

	// The total number of IP addresses in the IP address pool
	TotalIPAddresses int32 `json:"totalIpAddresses,omitempty"`
}

// Validate validates this Nsxt Ip address pool
func (m *NSXTIPAddressPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockSubnets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTIPAddressPool) validateBlockSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockSubnets) { // not required
		return nil
	}

	for i := 0; i < len(m.BlockSubnets); i++ {
		if swag.IsZero(m.BlockSubnets[i]) { // not required
			continue
		}

		if m.BlockSubnets[i] != nil {
			if err := m.BlockSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blockSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blockSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTIPAddressPool) validateStaticSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticSubnets) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticSubnets); i++ {
		if swag.IsZero(m.StaticSubnets[i]) { // not required
			continue
		}

		if m.StaticSubnets[i] != nil {
			if err := m.StaticSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Nsxt Ip address pool based on the context it is used
func (m *NSXTIPAddressPool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NSXTIPAddressPool) contextValidateBlockSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlockSubnets); i++ {

		if m.BlockSubnets[i] != nil {
			if err := m.BlockSubnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blockSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blockSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NSXTIPAddressPool) contextValidateStaticSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StaticSubnets); i++ {

		if m.StaticSubnets[i] != nil {
			if err := m.StaticSubnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NSXTIPAddressPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NSXTIPAddressPool) UnmarshalBinary(b []byte) error {
	var res NSXTIPAddressPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}