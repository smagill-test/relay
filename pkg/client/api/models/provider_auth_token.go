// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProviderAuthToken Metadata for authorizing a token-authenticated integration provider
// swagger:model ProviderAuthToken
type ProviderAuthToken struct {

	// type
	// Required: true
	// Enum: [token]
	Type *string `json:"type"`
}

// Validate validates this provider auth token
func (m *ProviderAuthToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var providerAuthTokenTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["token"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providerAuthTokenTypeTypePropEnum = append(providerAuthTokenTypeTypePropEnum, v)
	}
}

const (

	// ProviderAuthTokenTypeToken captures enum value "token"
	ProviderAuthTokenTypeToken string = "token"
)

// prop value enum
func (m *ProviderAuthToken) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providerAuthTokenTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProviderAuthToken) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProviderAuthToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderAuthToken) UnmarshalBinary(b []byte) error {
	var res ProviderAuthToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
