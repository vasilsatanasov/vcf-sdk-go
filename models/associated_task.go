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

// AssociatedTask Task association for a resource warning. Indicates in which task the resource warning occurred
//
// swagger:model AssociatedTask
type AssociatedTask struct {

	// ID of the subtask where the warning for the resource occurred
	SubTaskID string `json:"subTaskId,omitempty"`

	// ID of the task where the warning for the resource occurred
	// Required: true
	TaskID *string `json:"taskId"`
}

// Validate validates this associated task
func (m *AssociatedTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssociatedTask) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("taskId", "body", m.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this associated task based on context it is used
func (m *AssociatedTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssociatedTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssociatedTask) UnmarshalBinary(b []byte) error {
	var res AssociatedTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
