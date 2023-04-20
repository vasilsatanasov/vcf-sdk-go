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

// SharesInfo Specify relative allocation between resource consumers
//
// swagger:model SharesInfo
type SharesInfo struct {

	// The allocation level
	// Example: One among: low, normal, high, custom
	Level string `json:"level,omitempty"`

	// The number of shares allocated
	Shares int32 `json:"shares,omitempty"`
}

// Validate validates this shares info
func (m *SharesInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this shares info based on context it is used
func (m *SharesInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharesInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharesInfo) UnmarshalBinary(b []byte) error {
	var res SharesInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}