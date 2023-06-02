// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModifyNicknameRequest modify nickname request
//
// swagger:model ModifyNicknameRequest
type ModifyNicknameRequest struct {

	// new nickname
	NewNickname Nickname `json:"newNickname,omitempty"`
}

// Validate validates this modify nickname request
func (m *ModifyNicknameRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewNickname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyNicknameRequest) validateNewNickname(formats strfmt.Registry) error {
	if swag.IsZero(m.NewNickname) { // not required
		return nil
	}

	if err := m.NewNickname.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("newNickname")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("newNickname")
		}
		return err
	}

	return nil
}

// ContextValidate validate this modify nickname request based on the context it is used
func (m *ModifyNicknameRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewNickname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyNicknameRequest) contextValidateNewNickname(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NewNickname.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("newNickname")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("newNickname")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModifyNicknameRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyNicknameRequest) UnmarshalBinary(b []byte) error {
	var res ModifyNicknameRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
