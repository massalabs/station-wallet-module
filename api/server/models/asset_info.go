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

// AssetInfo Token informations
//
// swagger:model AssetInfo
type AssetInfo struct {

	// asset address
	AssetAddress string `json:"assetAddress,omitempty"`

	// decimals
	// Minimum: 0
	Decimals *int64 `json:"decimals,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`
}

// Validate validates this asset info
func (m *AssetInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecimals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetInfo) validateDecimals(formats strfmt.Registry) error {
	if swag.IsZero(m.Decimals) { // not required
		return nil
	}

	if err := validate.MinimumInt("decimals", "body", *m.Decimals, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this asset info based on context it is used
func (m *AssetInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssetInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetInfo) UnmarshalBinary(b []byte) error {
	var res AssetInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
