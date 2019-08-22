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

// Account An organizational account
// swagger:model Account
type Account struct {

	// Timestamp when terms and conditions were accepted by an owner of this account
	// Format: date-time
	AcceptedTermsAt *strfmt.DateTime `json:"accepted_terms_at,omitempty"`

	// The unique identifier for the account
	// Required: true
	ID *string `json:"id"`

	// The name of the account
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedTermsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateAcceptedTermsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedTermsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("accepted_terms_at", "body", "date-time", m.AcceptedTermsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
