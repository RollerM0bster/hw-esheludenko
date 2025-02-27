// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteEventsIDHandlerFunc turns a function with the right signature into a delete events ID handler
type DeleteEventsIDHandlerFunc func(DeleteEventsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEventsIDHandlerFunc) Handle(params DeleteEventsIDParams) middleware.Responder {
	return fn(params)
}

// DeleteEventsIDHandler interface for that can handle valid delete events ID params
type DeleteEventsIDHandler interface {
	Handle(DeleteEventsIDParams) middleware.Responder
}

// NewDeleteEventsID creates a new http.Handler for the delete events ID operation
func NewDeleteEventsID(ctx *middleware.Context, handler DeleteEventsIDHandler) *DeleteEventsID {
	return &DeleteEventsID{Context: ctx, Handler: handler}
}

/*
	DeleteEventsID swagger:route DELETE /events/{id} deleteEventsId

Удалить событие по идентификатору
*/
type DeleteEventsID struct {
	Context *middleware.Context
	Handler DeleteEventsIDHandler
}

func (o *DeleteEventsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteEventsIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
