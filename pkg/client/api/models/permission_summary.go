// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PermissionSummary permission summary
// swagger:model PermissionSummary
type PermissionSummary struct {
	PermissionIdentifier

	// A human-readable explanation of the access provided by this permission
	Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PermissionSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PermissionIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PermissionIdentifier = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"description,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PermissionSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PermissionIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"description,omitempty"`
	}

	dataAO1.Description = m.Description

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this permission summary
func (m *PermissionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PermissionIdentifier
	if err := m.PermissionIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionSummary) UnmarshalBinary(b []byte) error {
	var res PermissionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
