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

// UpdateSignRule update sign rule
//
// swagger:model UpdateSignRule
type UpdateSignRule struct {

	// The contract to which the rule applies. Use wildcard (*) to apply the rule for contracts.
	// Required: true
	Contract *string `json:"contract"`

	// Description text of what is being updated (optional)
	// Max Length: 280
	Description string `json:"description,omitempty"`

	// Whether the rule is enabled or not.
	// Required: true
	Enabled *bool `json:"enabled"`

	// The name of the rule.
	// Max Length: 100
	Name string `json:"name,omitempty"`

	// rule type
	// Required: true
	RuleType RuleType `json:"ruleType"`
}

// Validate validates this update sign rule
func (m *UpdateSignRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSignRule) validateContract(formats strfmt.Registry) error {

	if err := validate.Required("contract", "body", m.Contract); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSignRule) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 280); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSignRule) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSignRule) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSignRule) validateRuleType(formats strfmt.Registry) error {

	if err := validate.Required("ruleType", "body", RuleType(m.RuleType)); err != nil {
		return err
	}

	if err := m.RuleType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ruleType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ruleType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this update sign rule based on the context it is used
func (m *UpdateSignRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSignRule) contextValidateRuleType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RuleType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ruleType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ruleType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSignRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSignRule) UnmarshalBinary(b []byte) error {
	var res UpdateSignRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
