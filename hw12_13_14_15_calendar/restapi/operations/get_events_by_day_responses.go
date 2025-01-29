// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/models"
	"github.com/go-openapi/runtime"
)

// GetEventsByDayOKCode is the HTTP code returned for type GetEventsByDayOK
const GetEventsByDayOKCode int = 200

/*
GetEventsByDayOK Успех

swagger:response getEventsByDayOK
*/
type GetEventsByDayOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Event `json:"body,omitempty"`
}

// NewGetEventsByDayOK creates GetEventsByDayOK with default headers values
func NewGetEventsByDayOK() *GetEventsByDayOK {

	return &GetEventsByDayOK{}
}

// WithPayload adds the payload to the get events by day o k response
func (o *GetEventsByDayOK) WithPayload(payload []*models.Event) *GetEventsByDayOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events by day o k response
func (o *GetEventsByDayOK) SetPayload(payload []*models.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsByDayOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Event, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetEventsByDayInternalServerErrorCode is the HTTP code returned for type GetEventsByDayInternalServerError
const GetEventsByDayInternalServerErrorCode int = 500

/*
GetEventsByDayInternalServerError Ошибка

swagger:response getEventsByDayInternalServerError
*/
type GetEventsByDayInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventsByDayInternalServerError creates GetEventsByDayInternalServerError with default headers values
func NewGetEventsByDayInternalServerError() *GetEventsByDayInternalServerError {

	return &GetEventsByDayInternalServerError{}
}

// WithPayload adds the payload to the get events by day internal server error response
func (o *GetEventsByDayInternalServerError) WithPayload(payload *models.Error) *GetEventsByDayInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events by day internal server error response
func (o *GetEventsByDayInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsByDayInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
