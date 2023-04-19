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
)

// SupportBundleOption support bundle option
//
// swagger:model SupportBundleOption
type SupportBundleOption struct {

	// SupportBundle config.
	Config *SupportBundleConfig `json:"config,omitempty"`

	// Sos Additional Reports for Support Bundle.
	Include *SupportBundleIncludeItems `json:"include,omitempty"`
}

// Validate validates this support bundle option
func (m *SupportBundleOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportBundleOption) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *SupportBundleOption) validateInclude(formats strfmt.Registry) error {
	if swag.IsZero(m.Include) { // not required
		return nil
	}

	if m.Include != nil {
		if err := m.Include.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("include")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("include")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this support bundle option based on the context it is used
func (m *SupportBundleOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInclude(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportBundleOption) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *SupportBundleOption) contextValidateInclude(ctx context.Context, formats strfmt.Registry) error {

	if m.Include != nil {
		if err := m.Include.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("include")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("include")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportBundleOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportBundleOption) UnmarshalBinary(b []byte) error {
	var res SupportBundleOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
