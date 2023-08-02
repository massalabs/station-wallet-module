// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteAssetParams creates a new DeleteAssetParams object
//
// There are no default values defined in the spec.
func NewDeleteAssetParams() DeleteAssetParams {

	return DeleteAssetParams{}
}

// DeleteAssetParams contains all the bound params for the delete asset operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteAsset
type DeleteAssetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The asset address (token address) to delete from the account. It must start with "AS" and contain only alphanumeric characters.
	  Required: true
	  Pattern: ^AS[0-9a-zA-Z]+$
	  In: query
	*/
	AssetAddress string
	/*The nickname of the account from which to delete the asset.
	  Required: true
	  In: path
	*/
	Nickname string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteAssetParams() beforehand.
func (o *DeleteAssetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAssetAddress, qhkAssetAddress, _ := qs.GetOK("assetAddress")
	if err := o.bindAssetAddress(qAssetAddress, qhkAssetAddress, route.Formats); err != nil {
		res = append(res, err)
	}

	rNickname, rhkNickname, _ := route.Params.GetOK("nickname")
	if err := o.bindNickname(rNickname, rhkNickname, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAssetAddress binds and validates parameter AssetAddress from query.
func (o *DeleteAssetParams) bindAssetAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("assetAddress", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("assetAddress", "query", raw); err != nil {
		return err
	}
	o.AssetAddress = raw

	if err := o.validateAssetAddress(formats); err != nil {
		return err
	}

	return nil
}

// validateAssetAddress carries on validations for parameter AssetAddress
func (o *DeleteAssetParams) validateAssetAddress(formats strfmt.Registry) error {

	if err := validate.Pattern("assetAddress", "query", o.AssetAddress, `^AS[0-9a-zA-Z]+$`); err != nil {
		return err
	}

	return nil
}

// bindNickname binds and validates parameter Nickname from path.
func (o *DeleteAssetParams) bindNickname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Nickname = raw

	return nil
}
