// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1TagType v1 tag type
//
// swagger:model v1TagType
type V1TagType string

func NewV1TagType(value V1TagType) *V1TagType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V1TagType.
func (m V1TagType) Pointer() *V1TagType {
	return &m
}

const (

	// V1TagTypeTAGTYPEUSER captures enum value "TAG_TYPE_USER"
	V1TagTypeTAGTYPEUSER V1TagType = "TAG_TYPE_USER"

	// V1TagTypeTAGTYPEPRIVATEKEY captures enum value "TAG_TYPE_PRIVATE_KEY"
	V1TagTypeTAGTYPEPRIVATEKEY V1TagType = "TAG_TYPE_PRIVATE_KEY"
)

// for schema
var v1TagTypeEnum []interface{}

func init() {
	var res []V1TagType
	if err := json.Unmarshal([]byte(`["TAG_TYPE_USER","TAG_TYPE_PRIVATE_KEY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1TagTypeEnum = append(v1TagTypeEnum, v)
	}
}

func (m V1TagType) validateV1TagTypeEnum(path, location string, value V1TagType) error {
	if err := validate.EnumCase(path, location, value, v1TagTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 tag type
func (m V1TagType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1TagTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 tag type based on context it is used
func (m V1TagType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
