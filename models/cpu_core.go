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

// CPUCore Represents a CPU core
//
// swagger:model CpuCore
type CPUCore struct {

	// CPU frequency in MHz
	FrequencyMHz float64 `json:"frequencyMHz,omitempty"`

	// Hardware manufacturer of the CPU
	Manufacturer string `json:"manufacturer,omitempty"`

	// Hardware model of the CPU
	Model string `json:"model,omitempty"`
}

// Validate validates this Cpu core
func (m *CPUCore) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Cpu core based on context it is used
func (m *CPUCore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CPUCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUCore) UnmarshalBinary(b []byte) error {
	var res CPUCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
