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

// V1Intent Intent object crafted by Turnkey based on the user request, used to assess the permissibility of an action.
//
// swagger:model v1Intent
type V1Intent struct {

	// accept invitation intent
	AcceptInvitationIntent *V1AcceptInvitationIntent `json:"acceptInvitationIntent,omitempty"`

	// approve activity intent
	ApproveActivityIntent *V1ApproveActivityIntent `json:"approveActivityIntent,omitempty"`

	// create Api keys intent
	CreateAPIKeysIntent *V1CreateAPIKeysIntent `json:"createApiKeysIntent,omitempty"`

	// create authenticators intent
	CreateAuthenticatorsIntent *V1CreateAuthenticatorsIntent `json:"createAuthenticatorsIntent,omitempty"`

	// create invitations intent
	CreateInvitationsIntent *V1CreateInvitationsIntent `json:"createInvitationsIntent,omitempty"`

	// create organization intent
	// Required: true
	CreateOrganizationIntent *V1CreateOrganizationIntent `json:"createOrganizationIntent"`

	// create policy intent
	CreatePolicyIntent *V1CreatePolicyIntent `json:"createPolicyIntent,omitempty"`

	// create policy intent v2
	CreatePolicyIntentV2 *V1CreatePolicyIntentV2 `json:"createPolicyIntentV2,omitempty"`

	// create private key tag intent
	CreatePrivateKeyTagIntent *V1CreatePrivateKeyTagIntent `json:"createPrivateKeyTagIntent,omitempty"`

	// create private keys intent
	CreatePrivateKeysIntent *V1CreatePrivateKeysIntent `json:"createPrivateKeysIntent,omitempty"`

	// create user tag intent
	CreateUserTagIntent *V1CreateUserTagIntent `json:"createUserTagIntent,omitempty"`

	// create users intent
	CreateUsersIntent *V1CreateUsersIntent `json:"createUsersIntent,omitempty"`

	// delete Api keys intent
	DeleteAPIKeysIntent *V1DeleteAPIKeysIntent `json:"deleteApiKeysIntent,omitempty"`

	// delete authenticators intent
	DeleteAuthenticatorsIntent *V1DeleteAuthenticatorsIntent `json:"deleteAuthenticatorsIntent,omitempty"`

	// delete invitation intent
	DeleteInvitationIntent *V1DeleteInvitationIntent `json:"deleteInvitationIntent,omitempty"`

	// delete organization intent
	DeleteOrganizationIntent *V1DeleteOrganizationIntent `json:"deleteOrganizationIntent,omitempty"`

	// delete policy intent
	DeletePolicyIntent *V1DeletePolicyIntent `json:"deletePolicyIntent,omitempty"`

	// delete private key tags intent
	DeletePrivateKeyTagsIntent *V1DeletePrivateKeyTagsIntent `json:"deletePrivateKeyTagsIntent,omitempty"`

	// delete user tags intent
	DeleteUserTagsIntent *V1DeleteUserTagsIntent `json:"deleteUserTagsIntent,omitempty"`

	// delete users intent
	DeleteUsersIntent *V1DeleteUsersIntent `json:"deleteUsersIntent,omitempty"`

	// disable private key intent
	DisablePrivateKeyIntent *V1DisablePrivateKeyIntent `json:"disablePrivateKeyIntent,omitempty"`

	// reject activity intent
	RejectActivityIntent *V1RejectActivityIntent `json:"rejectActivityIntent,omitempty"`

	// sign raw payload intent
	SignRawPayloadIntent *V1SignRawPayloadIntent `json:"signRawPayloadIntent,omitempty"`

	// sign transaction intent
	SignTransactionIntent *V1SignTransactionIntent `json:"signTransactionIntent,omitempty"`
}

// Validate validates this v1 intent
func (m *V1Intent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptInvitationIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApproveActivityIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateAPIKeysIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateAuthenticatorsIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateInvitationsIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateOrganizationIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatePolicyIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatePolicyIntentV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatePrivateKeyTagIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatePrivateKeysIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateUserTagIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateUsersIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteAPIKeysIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteAuthenticatorsIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteInvitationIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteOrganizationIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletePolicyIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletePrivateKeyTagsIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteUserTagsIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteUsersIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisablePrivateKeyIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectActivityIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignRawPayloadIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignTransactionIntent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Intent) validateAcceptInvitationIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptInvitationIntent) { // not required
		return nil
	}

	if m.AcceptInvitationIntent != nil {
		if err := m.AcceptInvitationIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptInvitationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptInvitationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateApproveActivityIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.ApproveActivityIntent) { // not required
		return nil
	}

	if m.ApproveActivityIntent != nil {
		if err := m.ApproveActivityIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approveActivityIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("approveActivityIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateAPIKeysIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateAPIKeysIntent) { // not required
		return nil
	}

	if m.CreateAPIKeysIntent != nil {
		if err := m.CreateAPIKeysIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createApiKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createApiKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateAuthenticatorsIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateAuthenticatorsIntent) { // not required
		return nil
	}

	if m.CreateAuthenticatorsIntent != nil {
		if err := m.CreateAuthenticatorsIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createAuthenticatorsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createAuthenticatorsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateInvitationsIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateInvitationsIntent) { // not required
		return nil
	}

	if m.CreateInvitationsIntent != nil {
		if err := m.CreateInvitationsIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createInvitationsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createInvitationsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateOrganizationIntent(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationIntent", "body", m.CreateOrganizationIntent); err != nil {
		return err
	}

	if m.CreateOrganizationIntent != nil {
		if err := m.CreateOrganizationIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreatePolicyIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatePolicyIntent) { // not required
		return nil
	}

	if m.CreatePolicyIntent != nil {
		if err := m.CreatePolicyIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPolicyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPolicyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreatePolicyIntentV2(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatePolicyIntentV2) { // not required
		return nil
	}

	if m.CreatePolicyIntentV2 != nil {
		if err := m.CreatePolicyIntentV2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPolicyIntentV2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPolicyIntentV2")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreatePrivateKeyTagIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatePrivateKeyTagIntent) { // not required
		return nil
	}

	if m.CreatePrivateKeyTagIntent != nil {
		if err := m.CreatePrivateKeyTagIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPrivateKeyTagIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPrivateKeyTagIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreatePrivateKeysIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatePrivateKeysIntent) { // not required
		return nil
	}

	if m.CreatePrivateKeysIntent != nil {
		if err := m.CreatePrivateKeysIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPrivateKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPrivateKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateUserTagIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateUserTagIntent) { // not required
		return nil
	}

	if m.CreateUserTagIntent != nil {
		if err := m.CreateUserTagIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserTagIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserTagIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateCreateUsersIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateUsersIntent) { // not required
		return nil
	}

	if m.CreateUsersIntent != nil {
		if err := m.CreateUsersIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUsersIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUsersIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteAPIKeysIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteAPIKeysIntent) { // not required
		return nil
	}

	if m.DeleteAPIKeysIntent != nil {
		if err := m.DeleteAPIKeysIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteApiKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteApiKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteAuthenticatorsIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteAuthenticatorsIntent) { // not required
		return nil
	}

	if m.DeleteAuthenticatorsIntent != nil {
		if err := m.DeleteAuthenticatorsIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAuthenticatorsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAuthenticatorsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteInvitationIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteInvitationIntent) { // not required
		return nil
	}

	if m.DeleteInvitationIntent != nil {
		if err := m.DeleteInvitationIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteInvitationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteInvitationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteOrganizationIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteOrganizationIntent) { // not required
		return nil
	}

	if m.DeleteOrganizationIntent != nil {
		if err := m.DeleteOrganizationIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOrganizationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOrganizationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeletePolicyIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletePolicyIntent) { // not required
		return nil
	}

	if m.DeletePolicyIntent != nil {
		if err := m.DeletePolicyIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletePolicyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletePolicyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeletePrivateKeyTagsIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletePrivateKeyTagsIntent) { // not required
		return nil
	}

	if m.DeletePrivateKeyTagsIntent != nil {
		if err := m.DeletePrivateKeyTagsIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletePrivateKeyTagsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletePrivateKeyTagsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteUserTagsIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteUserTagsIntent) { // not required
		return nil
	}

	if m.DeleteUserTagsIntent != nil {
		if err := m.DeleteUserTagsIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteUserTagsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteUserTagsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDeleteUsersIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteUsersIntent) { // not required
		return nil
	}

	if m.DeleteUsersIntent != nil {
		if err := m.DeleteUsersIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteUsersIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteUsersIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateDisablePrivateKeyIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.DisablePrivateKeyIntent) { // not required
		return nil
	}

	if m.DisablePrivateKeyIntent != nil {
		if err := m.DisablePrivateKeyIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disablePrivateKeyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disablePrivateKeyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateRejectActivityIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectActivityIntent) { // not required
		return nil
	}

	if m.RejectActivityIntent != nil {
		if err := m.RejectActivityIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rejectActivityIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rejectActivityIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateSignRawPayloadIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.SignRawPayloadIntent) { // not required
		return nil
	}

	if m.SignRawPayloadIntent != nil {
		if err := m.SignRawPayloadIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signRawPayloadIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signRawPayloadIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) validateSignTransactionIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.SignTransactionIntent) { // not required
		return nil
	}

	if m.SignTransactionIntent != nil {
		if err := m.SignTransactionIntent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signTransactionIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signTransactionIntent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 intent based on the context it is used
func (m *V1Intent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptInvitationIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApproveActivityIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateAPIKeysIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateAuthenticatorsIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateInvitationsIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateOrganizationIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatePolicyIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatePolicyIntentV2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatePrivateKeyTagIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatePrivateKeysIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateUserTagIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateUsersIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteAPIKeysIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteAuthenticatorsIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteInvitationIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteOrganizationIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletePolicyIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletePrivateKeyTagsIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteUserTagsIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteUsersIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisablePrivateKeyIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRejectActivityIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignRawPayloadIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignTransactionIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Intent) contextValidateAcceptInvitationIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptInvitationIntent != nil {
		if err := m.AcceptInvitationIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptInvitationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptInvitationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateApproveActivityIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.ApproveActivityIntent != nil {
		if err := m.ApproveActivityIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approveActivityIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("approveActivityIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateAPIKeysIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateAPIKeysIntent != nil {
		if err := m.CreateAPIKeysIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createApiKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createApiKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateAuthenticatorsIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateAuthenticatorsIntent != nil {
		if err := m.CreateAuthenticatorsIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createAuthenticatorsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createAuthenticatorsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateInvitationsIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateInvitationsIntent != nil {
		if err := m.CreateInvitationsIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createInvitationsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createInvitationsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateOrganizationIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateOrganizationIntent != nil {
		if err := m.CreateOrganizationIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreatePolicyIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatePolicyIntent != nil {
		if err := m.CreatePolicyIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPolicyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPolicyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreatePolicyIntentV2(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatePolicyIntentV2 != nil {
		if err := m.CreatePolicyIntentV2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPolicyIntentV2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPolicyIntentV2")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreatePrivateKeyTagIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatePrivateKeyTagIntent != nil {
		if err := m.CreatePrivateKeyTagIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPrivateKeyTagIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPrivateKeyTagIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreatePrivateKeysIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatePrivateKeysIntent != nil {
		if err := m.CreatePrivateKeysIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createPrivateKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createPrivateKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateUserTagIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateUserTagIntent != nil {
		if err := m.CreateUserTagIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUserTagIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUserTagIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateCreateUsersIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateUsersIntent != nil {
		if err := m.CreateUsersIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createUsersIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createUsersIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteAPIKeysIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteAPIKeysIntent != nil {
		if err := m.DeleteAPIKeysIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteApiKeysIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteApiKeysIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteAuthenticatorsIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteAuthenticatorsIntent != nil {
		if err := m.DeleteAuthenticatorsIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteAuthenticatorsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteAuthenticatorsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteInvitationIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteInvitationIntent != nil {
		if err := m.DeleteInvitationIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteInvitationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteInvitationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteOrganizationIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteOrganizationIntent != nil {
		if err := m.DeleteOrganizationIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOrganizationIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteOrganizationIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeletePolicyIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeletePolicyIntent != nil {
		if err := m.DeletePolicyIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletePolicyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletePolicyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeletePrivateKeyTagsIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeletePrivateKeyTagsIntent != nil {
		if err := m.DeletePrivateKeyTagsIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletePrivateKeyTagsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletePrivateKeyTagsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteUserTagsIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteUserTagsIntent != nil {
		if err := m.DeleteUserTagsIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteUserTagsIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteUserTagsIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDeleteUsersIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteUsersIntent != nil {
		if err := m.DeleteUsersIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteUsersIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteUsersIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateDisablePrivateKeyIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.DisablePrivateKeyIntent != nil {
		if err := m.DisablePrivateKeyIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disablePrivateKeyIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disablePrivateKeyIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateRejectActivityIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.RejectActivityIntent != nil {
		if err := m.RejectActivityIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rejectActivityIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rejectActivityIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateSignRawPayloadIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.SignRawPayloadIntent != nil {
		if err := m.SignRawPayloadIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signRawPayloadIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signRawPayloadIntent")
			}
			return err
		}
	}

	return nil
}

func (m *V1Intent) contextValidateSignTransactionIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.SignTransactionIntent != nil {
		if err := m.SignTransactionIntent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signTransactionIntent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signTransactionIntent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Intent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Intent) UnmarshalBinary(b []byte) error {
	var res V1Intent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
