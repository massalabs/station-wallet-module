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

// SignRequest sign request
//
// swagger:model SignRequest
type SignRequest struct {

	// A boolean property that indicates whether the sign operation is part of a batch of operations. Set to true if this operation is part of a batch, otherwise set to false.
	Batch bool `json:"batch,omitempty"`

	// correlation Id
	// Format: byte
	CorrelationID CorrelationID `json:"correlationId,omitempty"`

	// Description text of what is being signed (optional)
	// Max Length: 280
	Description string `json:"description,omitempty"`

	// Serialized attributes of the operation to be signed with the key pair corresponding to the given nickname.
	// Required: true
	// Format: byte
	Operation *strfmt.Base64 `json:"operation"`
}

// Validate validates this sign request
func (m *SignRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorrelationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignRequest) validateCorrelationID(formats strfmt.Registry) error {
	if swag.IsZero(m.CorrelationID) { // not required
		return nil
	}

	if err := m.CorrelationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("correlationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("correlationId")
		}
		return err
	}

	return nil
}

func (m *SignRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 280); err != nil {
		return err
	}

	return nil
}

func (m *SignRequest) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sign request based on the context it is used
func (m *SignRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCorrelationID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignRequest) contextValidateCorrelationID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CorrelationID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("correlationId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("correlationId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SignRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignRequest) UnmarshalBinary(b []byte) error {
	var res SignRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
