// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductNode Spec contains information for a vRealize product node
//
// swagger:model ProductNode
type ProductNode struct {

	// The Fully Qualified Domain Name for the vRealize node (virtual appliance)
	// Example: vrops.vrack.vsphere.local
	// Required: true
	Fqdn *string `json:"fqdn"`

	// The password for a root user of vRealize appliance
	// Required: true
	Password *string `json:"password"`

	// The type of the vRealize product node
	// Example: MASTER, REPLICA, DATA, REMOTECOLLECTOR, WORKER
	// Enum: [Oneamong:MASTER REPLICA DATA REMOTECOLLECTOR WORKER PRIMARY SECONDARY]
	Type string `json:"type,omitempty"`

	// The username for a root user of vRealize appliance
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this product node
func (m *ProductNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFqdn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductNode) validateFqdn(formats strfmt.Registry) error {

	if err := validate.Required("fqdn", "body", m.Fqdn); err != nil {
		return err
	}

	return nil
}

func (m *ProductNode) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

var productNodeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Oneamong:MASTER","REPLICA","DATA","REMOTECOLLECTOR","WORKER","PRIMARY","SECONDARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productNodeTypeTypePropEnum = append(productNodeTypeTypePropEnum, v)
	}
}

const (

	// ProductNodeTypeOneamongMASTER captures enum value "Oneamong:MASTER"
	ProductNodeTypeOneamongMASTER string = "Oneamong:MASTER"

	// ProductNodeTypeREPLICA captures enum value "REPLICA"
	ProductNodeTypeREPLICA string = "REPLICA"

	// ProductNodeTypeDATA captures enum value "DATA"
	ProductNodeTypeDATA string = "DATA"

	// ProductNodeTypeREMOTECOLLECTOR captures enum value "REMOTECOLLECTOR"
	ProductNodeTypeREMOTECOLLECTOR string = "REMOTECOLLECTOR"

	// ProductNodeTypeWORKER captures enum value "WORKER"
	ProductNodeTypeWORKER string = "WORKER"

	// ProductNodeTypePRIMARY captures enum value "PRIMARY"
	ProductNodeTypePRIMARY string = "PRIMARY"

	// ProductNodeTypeSECONDARY captures enum value "SECONDARY"
	ProductNodeTypeSECONDARY string = "SECONDARY"
)

// prop value enum
func (m *ProductNode) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, productNodeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProductNode) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ProductNode) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this product node based on context it is used
func (m *ProductNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductNode) UnmarshalBinary(b []byte) error {
	var res ProductNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}