// Code generated by go-swagger; DO NOT EDIT.

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

// V1CreateAuthenticatorsResult v1 create authenticators result
//
// swagger:model v1CreateAuthenticatorsResult
type V1CreateAuthenticatorsResult struct {

	// A list of Authenticator IDs.
	// Required: true
	AuthenticatorIds []string `json:"authenticatorIds"`
}

// Validate validates this v1 create authenticators result
func (m *V1CreateAuthenticatorsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticatorIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAuthenticatorsResult) validateAuthenticatorIds(formats strfmt.Registry) error {

	if err := validate.Required("authenticatorIds", "body", m.AuthenticatorIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this v1 create authenticators result based on context it is used
func (m *V1CreateAuthenticatorsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateAuthenticatorsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateAuthenticatorsResult) UnmarshalBinary(b []byte) error {
	var res V1CreateAuthenticatorsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
