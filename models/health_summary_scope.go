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

// HealthSummaryScope health summary scope
//
// swagger:model HealthSummaryScope
type HealthSummaryScope struct {

	// Domains and Clusters for SOS operation.
	Domains []*Domains `json:"domains"`

	// Include all domains for SOS operation.
	IncludeAllDomains bool `json:"includeAllDomains,omitempty"`

	// Include free hosts.
	IncludeFreeHosts bool `json:"includeFreeHosts,omitempty"`
}

// Validate validates this health summary scope
func (m *HealthSummaryScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthSummaryScope) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this health summary scope based on the context it is used
func (m *HealthSummaryScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthSummaryScope) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Domains); i++ {

		if m.Domains[i] != nil {
			if err := m.Domains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthSummaryScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthSummaryScope) UnmarshalBinary(b []byte) error {
	var res HealthSummaryScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
