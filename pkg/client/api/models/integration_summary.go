// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationSummary integration summary
// swagger:model IntegrationSummary
type IntegrationSummary struct {
	IntegrationIdentifier

	// A descriptive integration name
	// Required: true
	Name *string `json:"name"`

	// This integration's provider identifier
	// Required: true
	Provider *string `json:"provider"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IntegrationSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IntegrationIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IntegrationIdentifier = aO0

	// AO1
	var dataAO1 struct {
		Name *string `json:"name"`

		Provider *string `json:"provider"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Provider = dataAO1.Provider

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IntegrationSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.IntegrationIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Name *string `json:"name"`

		Provider *string `json:"provider"`
	}

	dataAO1.Name = m.Name

	dataAO1.Provider = m.Provider

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this integration summary
func (m *IntegrationSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IntegrationIdentifier
	if err := m.IntegrationIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationSummary) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationSummary) UnmarshalBinary(b []byte) error {
	var res IntegrationSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
