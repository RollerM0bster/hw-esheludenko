// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEventsByDayHandlerFunc turns a function with the right signature into a get events by day handler
type GetEventsByDayHandlerFunc func(GetEventsByDayParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventsByDayHandlerFunc) Handle(params GetEventsByDayParams) middleware.Responder {
	return fn(params)
}

// GetEventsByDayHandler interface for that can handle valid get events by day params
type GetEventsByDayHandler interface {
	Handle(GetEventsByDayParams) middleware.Responder
}

// NewGetEventsByDay creates a new http.Handler for the get events by day operation
func NewGetEventsByDay(ctx *middleware.Context, handler GetEventsByDayHandler) *GetEventsByDay {
	return &GetEventsByDay{Context: ctx, Handler: handler}
}

/*
	GetEventsByDay swagger:route GET /events-by-day getEventsByDay

Получить все события за день
*/
type GetEventsByDay struct {
	Context *middleware.Context
	Handler GetEventsByDayHandler
}

func (o *GetEventsByDay) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEventsByDayParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
