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

// ProductVersion Details of the product/component for the release.
//
// swagger:model ProductVersion
type ProductVersion struct {

	// any additional metadata
	AdditionalMetadata string `json:"additionalMetadata,omitempty"`

	// Name of the product. e.g ESX
	// Required: true
	Name *string `json:"name"`

	// Public name of the product, e.g VMware ESXi
	// Required: true
	PublicName *string `json:"publicName"`

	// URL for the release.
	ReleaseURL *URL `json:"releaseURL,omitempty"`

	// Version for the product, e.g 6.7.0-11675023
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this product version
func (m *ProductVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProductVersion) validatePublicName(formats strfmt.Registry) error {

	if err := validate.Required("publicName", "body", m.PublicName); err != nil {
		return err
	}

	return nil
}

func (m *ProductVersion) validateReleaseURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ReleaseURL) { // not required
		return nil
	}

	if m.ReleaseURL != nil {
		if err := m.ReleaseURL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseURL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("releaseURL")
			}
			return err
		}
	}

	return nil
}

func (m *ProductVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this product version based on the context it is used
func (m *ProductVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReleaseURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductVersion) contextValidateReleaseURL(ctx context.Context, formats strfmt.Registry) error {

	if m.ReleaseURL != nil {
		if err := m.ReleaseURL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("releaseURL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("releaseURL")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductVersion) UnmarshalBinary(b []byte) error {
	var res ProductVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
