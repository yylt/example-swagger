// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEchoNameHandlerFunc turns a function with the right signature into a get echo name handler
type GetEchoNameHandlerFunc func(GetEchoNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEchoNameHandlerFunc) Handle(params GetEchoNameParams) middleware.Responder {
	return fn(params)
}

// GetEchoNameHandler interface for that can handle valid get echo name params
type GetEchoNameHandler interface {
	Handle(GetEchoNameParams) middleware.Responder
}

// NewGetEchoName creates a new http.Handler for the get echo name operation
func NewGetEchoName(ctx *middleware.Context, handler GetEchoNameHandler) *GetEchoName {
	return &GetEchoName{Context: ctx, Handler: handler}
}

/*GetEchoName swagger:route GET /echo/{name} getEchoName

GetEchoName get echo name API

*/
type GetEchoName struct {
	Context *middleware.Context
	Handler GetEchoNameHandler
}

func (o *GetEchoName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEchoNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
