// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SignMessageRequest sign message request
//
// swagger:model SignMessageRequest
type SignMessageRequest struct {

	// A boolean indicating whether to display data.
	DisplayData *bool `json:"DisplayData,omitempty"`

	// The message to sign.
	Message string `json:"message,omitempty"`
}

// Validate validates this sign message request
func (m *SignMessageRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sign message request based on context it is used
func (m *SignMessageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignMessageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignMessageRequest) UnmarshalBinary(b []byte) error {
	var res SignMessageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
