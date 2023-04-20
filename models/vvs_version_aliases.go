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

// VvsVersionAliases Vvs Mapping rest model that is located in the 2.0 manifest, these models are consumed by AP tool to determine the release Id to vcf version mapping in VVS compatibility data.
//
// swagger:model VvsVersionAliases
type VvsVersionAliases struct {

	// Product version aliases
	Aliases []string `json:"aliases"`

	// Product version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this vvs version aliases
func (m *VvsVersionAliases) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VvsVersionAliases) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vvs version aliases based on context it is used
func (m *VvsVersionAliases) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VvsVersionAliases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VvsVersionAliases) UnmarshalBinary(b []byte) error {
	var res VvsVersionAliases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}