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

// SubTask Represents a sub-task
//
// swagger:model SubTask
type SubTask struct {

	// Sub-task completion timestamp
	CompletionTimestamp string `json:"completionTimestamp,omitempty"`

	// Sub-task creation timestamp
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Sub-task description
	Description string `json:"description,omitempty"`

	// List of errors in case of a failure
	Errors []*Error `json:"errors"`

	// Sub-task name
	Name string `json:"name,omitempty"`

	// List of resources associated with sub-task
	Resources []*Resource `json:"resources"`

	// List of stages of the sub-task
	Stages []*Stage `json:"stages"`

	// Sub-task status
	// Example: One among: PENDING, IN_PROGRESS, SUCCESSFUL, FAILED, NOT_APPLICABLE
	Status string `json:"status,omitempty"`

	// List of child subtasks associated with this subtask
	SubTasks []*SubTask `json:"subTasks"`

	// Sub-task type
	Type string `json:"type,omitempty"`
}

// Validate validates this sub task
func (m *SubTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubTask) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) validateStages(formats strfmt.Registry) error {
	if swag.IsZero(m.Stages) { // not required
		return nil
	}

	for i := 0; i < len(m.Stages); i++ {
		if swag.IsZero(m.Stages[i]) { // not required
			continue
		}

		if m.Stages[i] != nil {
			if err := m.Stages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) validateSubTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.SubTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.SubTasks); i++ {
		if swag.IsZero(m.SubTasks[i]) { // not required
			continue
		}

		if m.SubTasks[i] != nil {
			if err := m.SubTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this sub task based on the context it is used
func (m *SubTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubTask) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Resources); i++ {

		if m.Resources[i] != nil {
			if err := m.Resources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) contextValidateStages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stages); i++ {

		if m.Stages[i] != nil {
			if err := m.Stages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SubTask) contextValidateSubTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubTasks); i++ {

		if m.SubTasks[i] != nil {
			if err := m.SubTasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubTask) UnmarshalBinary(b []byte) error {
	var res SubTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}