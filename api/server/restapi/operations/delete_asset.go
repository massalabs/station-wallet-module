// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAssetHandlerFunc turns a function with the right signature into a delete asset handler
type DeleteAssetHandlerFunc func(DeleteAssetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAssetHandlerFunc) Handle(params DeleteAssetParams) middleware.Responder {
	return fn(params)
}

// DeleteAssetHandler interface for that can handle valid delete asset params
type DeleteAssetHandler interface {
	Handle(DeleteAssetParams) middleware.Responder
}

// NewDeleteAsset creates a new http.Handler for the delete asset operation
func NewDeleteAsset(ctx *middleware.Context, handler DeleteAssetHandler) *DeleteAsset {
	return &DeleteAsset{Context: ctx, Handler: handler}
}

/*
	DeleteAsset swagger:route DELETE /api/accounts/{nickname}/assets deleteAsset

Delete token information from an account.
*/
type DeleteAsset struct {
	Context *middleware.Context
	Handler DeleteAssetHandler
}

func (o *DeleteAsset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteAssetParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
