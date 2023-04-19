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

// LicenseKey Defines a license key and its attributes
//
// swagger:model LicenseKey
type LicenseKey struct {

	// Description of the license key
	// Required: true
	Description *string `json:"description"`

	// The ID of the license key
	ID string `json:"id,omitempty"`

	// Indicates if the license key has unlimited usage
	IsUnlimited bool `json:"isUnlimited,omitempty"`

	// The 29 alpha numeric character license key with hyphens
	// Example: XX0XX-XX0XX-XX0XX-XX0XX-XX0XX
	// Required: true
	Key *string `json:"key"`

	// License key usage details
	LicenseKeyUsage *LicenseKeyUsage `json:"licenseKeyUsage,omitempty"`

	// License key validity details
	LicenseKeyValidity *LicenseKeyValidity `json:"licenseKeyValidity,omitempty"`

	// The type of the product to which the license key is applicable
	// Example: One among: VCENTER, VSAN, SDDC_MANAGER, ESXI, NSXT, NSXIO, WCP, HORIZON_VIEW
	// Required: true
	ProductType *string `json:"productType"`
}

// Validate validates this license key
func (m *LicenseKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseKeyUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseKeyValidity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseKey) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *LicenseKey) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *LicenseKey) validateLicenseKeyUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseKeyUsage) { // not required
		return nil
	}

	if m.LicenseKeyUsage != nil {
		if err := m.LicenseKeyUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeyUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeyUsage")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseKey) validateLicenseKeyValidity(formats strfmt.Registry) error {
	if swag.IsZero(m.LicenseKeyValidity) { // not required
		return nil
	}

	if m.LicenseKeyValidity != nil {
		if err := m.LicenseKeyValidity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeyValidity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeyValidity")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseKey) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", m.ProductType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this license key based on the context it is used
func (m *LicenseKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicenseKeyUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicenseKeyValidity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseKey) contextValidateLicenseKeyUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.LicenseKeyUsage != nil {
		if err := m.LicenseKeyUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeyUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeyUsage")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseKey) contextValidateLicenseKeyValidity(ctx context.Context, formats strfmt.Registry) error {

	if m.LicenseKeyValidity != nil {
		if err := m.LicenseKeyValidity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeyValidity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licenseKeyValidity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseKey) UnmarshalBinary(b []byte) error {
	var res LicenseKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
