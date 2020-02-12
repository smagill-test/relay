// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowRunState workflow run state
// swagger:model WorkflowRunState
type WorkflowRunState struct {
	WorkflowRunStateSummary

	// Status of individual run steps, indexed by name
	Steps map[string]WorkflowRunStepState `json:"steps,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowRunState) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowRunStateSummary
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowRunStateSummary = aO0

	// AO1
	var dataAO1 struct {
		Steps map[string]WorkflowRunStepState `json:"steps,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Steps = dataAO1.Steps

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowRunState) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowRunStateSummary)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Steps map[string]WorkflowRunStepState `json:"steps,omitempty"`
	}

	dataAO1.Steps = m.Steps

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow run state
func (m *WorkflowRunState) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowRunStateSummary
	if err := m.WorkflowRunStateSummary.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRunState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRunState) UnmarshalBinary(b []byte) error {
	var res WorkflowRunState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}