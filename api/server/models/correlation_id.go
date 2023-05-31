// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CorrelationID Correlation id of the operation batch
//
// swagger:model CorrelationId
type CorrelationID strfmt.Base64

// UnmarshalJSON sets a CorrelationID value from JSON input
func (m *CorrelationID) UnmarshalJSON(b []byte) error {
	return ((*strfmt.Base64)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a CorrelationID value as JSON output
func (m CorrelationID) MarshalJSON() ([]byte, error) {
	return (strfmt.Base64(m)).MarshalJSON()
}

// Validate validates this correlation Id
func (m CorrelationID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this correlation Id based on context it is used
func (m CorrelationID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CorrelationID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CorrelationID) UnmarshalBinary(b []byte) error {
	var res CorrelationID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
