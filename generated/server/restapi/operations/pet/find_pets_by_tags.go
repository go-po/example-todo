// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindPetsByTagsHandlerFunc turns a function with the right signature into a find pets by tags handler
type FindPetsByTagsHandlerFunc func(FindPetsByTagsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn FindPetsByTagsHandlerFunc) Handle(params FindPetsByTagsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// FindPetsByTagsHandler interface for that can handle valid find pets by tags params
type FindPetsByTagsHandler interface {
	Handle(FindPetsByTagsParams, interface{}) middleware.Responder
}

// NewFindPetsByTags creates a new http.Handler for the find pets by tags operation
func NewFindPetsByTags(ctx *middleware.Context, handler FindPetsByTagsHandler) *FindPetsByTags {
	return &FindPetsByTags{Context: ctx, Handler: handler}
}

/*FindPetsByTags swagger:route GET /pet/findByTags pet findPetsByTags

Finds Pets by tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

*/
type FindPetsByTags struct {
	Context *middleware.Context
	Handler FindPetsByTagsHandler
}

func (o *FindPetsByTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindPetsByTagsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
