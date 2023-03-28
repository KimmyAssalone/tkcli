// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CreateAuthenticatorsIntent v1 create authenticators intent
//
// swagger:model v1CreateAuthenticatorsIntent
type V1CreateAuthenticatorsIntent struct {

	// @inject_tag: validate:"dive,required"
	//
	// A list of Authenticators.
	// Required: true
	Authenticators []*V1AuthenticatorParams `json:"authenticators"`

	// @inject_tag: validate:"required,uuid"
	//
	// Unique identifier for a given User.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this v1 create authenticators intent
func (m *V1CreateAuthenticatorsIntent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAuthenticatorsIntent) validateAuthenticators(formats strfmt.Registry) error {

	if err := validate.Required("authenticators", "body", m.Authenticators); err != nil {
		return err
	}

	for i := 0; i < len(m.Authenticators); i++ {
		if swag.IsZero(m.Authenticators[i]) { // not required
			continue
		}

		if m.Authenticators[i] != nil {
			if err := m.Authenticators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authenticators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authenticators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1CreateAuthenticatorsIntent) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 create authenticators intent based on the context it is used
func (m *V1CreateAuthenticatorsIntent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CreateAuthenticatorsIntent) contextValidateAuthenticators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Authenticators); i++ {

		if m.Authenticators[i] != nil {
			if err := m.Authenticators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authenticators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("authenticators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateAuthenticatorsIntent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateAuthenticatorsIntent) UnmarshalBinary(b []byte) error {
	var res V1CreateAuthenticatorsIntent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
