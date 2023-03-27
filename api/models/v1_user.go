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

// V1User v1 user
//
// swagger:model v1User
type V1User struct {

	// access type
	// Required: true
	AccessType *Externaldatav1AccessType `json:"accessType"`

	// A list of API Key parameters.
	// Required: true
	APIKeys []*V1APIKey `json:"apiKeys"`

	// A list of Authenticator parameters.
	// Required: true
	Authenticators []*V1Authenticator `json:"authenticators"`

	// created at
	// Required: true
	CreatedAt *V1Timestamp `json:"createdAt"`

	// updated at
	// Required: true
	UpdatedAt *V1Timestamp `json:"updatedAt"`

	// some users do not have emails (programmatic users)
	//
	// The user's email address.
	UserEmail string `json:"userEmail,omitempty"`

	// Unique identifier for a given User.
	// Required: true
	UserID *string `json:"userId"`

	// Human-readable name for a User.
	// Required: true
	UserName *string `json:"userName"`

	// A list of User Tag IDs.
	// Required: true
	UserTags []string `json:"userTags"`
}

// Validate validates this v1 user
func (m *V1User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPIKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1User) validateAccessType(formats strfmt.Registry) error {

	if err := validate.Required("accessType", "body", m.AccessType); err != nil {
		return err
	}

	if err := validate.Required("accessType", "body", m.AccessType); err != nil {
		return err
	}

	if m.AccessType != nil {
		if err := m.AccessType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessType")
			}
			return err
		}
	}

	return nil
}

func (m *V1User) validateAPIKeys(formats strfmt.Registry) error {

	if err := validate.Required("apiKeys", "body", m.APIKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.APIKeys); i++ {
		if swag.IsZero(m.APIKeys[i]) { // not required
			continue
		}

		if m.APIKeys[i] != nil {
			if err := m.APIKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1User) validateAuthenticators(formats strfmt.Registry) error {

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

func (m *V1User) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if m.CreatedAt != nil {
		if err := m.CreatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1User) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1User) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *V1User) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

func (m *V1User) validateUserTags(formats strfmt.Registry) error {

	if err := validate.Required("userTags", "body", m.UserTags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v1 user based on the context it is used
func (m *V1User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAPIKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1User) contextValidateAccessType(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessType != nil {
		if err := m.AccessType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessType")
			}
			return err
		}
	}

	return nil
}

func (m *V1User) contextValidateAPIKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.APIKeys); i++ {

		if m.APIKeys[i] != nil {
			if err := m.APIKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apiKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1User) contextValidateAuthenticators(ctx context.Context, formats strfmt.Registry) error {

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

func (m *V1User) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedAt != nil {
		if err := m.CreatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdAt")
			}
			return err
		}
	}

	return nil
}

func (m *V1User) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedAt != nil {
		if err := m.UpdatedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updatedAt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1User) UnmarshalBinary(b []byte) error {
	var res V1User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
