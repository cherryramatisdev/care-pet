// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/scylladb/care-pet/go/cmd/server/models"
)

// FindPetsByOwnerIDOKCode is the HTTP code returned for type FindPetsByOwnerIDOK
const FindPetsByOwnerIDOKCode int = 200

/*FindPetsByOwnerIDOK pets response

swagger:response findPetsByOwnerIdOK
*/
type FindPetsByOwnerIDOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pet `json:"body,omitempty"`
}

// NewFindPetsByOwnerIDOK creates FindPetsByOwnerIDOK with default headers values
func NewFindPetsByOwnerIDOK() *FindPetsByOwnerIDOK {

	return &FindPetsByOwnerIDOK{}
}

// WithPayload adds the payload to the find pets by owner Id o k response
func (o *FindPetsByOwnerIDOK) WithPayload(payload []*models.Pet) *FindPetsByOwnerIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets by owner Id o k response
func (o *FindPetsByOwnerIDOK) SetPayload(payload []*models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsByOwnerIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Pet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*FindPetsByOwnerIDDefault error

swagger:response findPetsByOwnerIdDefault
*/
type FindPetsByOwnerIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindPetsByOwnerIDDefault creates FindPetsByOwnerIDDefault with default headers values
func NewFindPetsByOwnerIDDefault(code int) *FindPetsByOwnerIDDefault {
	if code <= 0 {
		code = 500
	}

	return &FindPetsByOwnerIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find pets by owner id default response
func (o *FindPetsByOwnerIDDefault) WithStatusCode(code int) *FindPetsByOwnerIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find pets by owner id default response
func (o *FindPetsByOwnerIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find pets by owner id default response
func (o *FindPetsByOwnerIDDefault) WithPayload(payload *models.Error) *FindPetsByOwnerIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find pets by owner id default response
func (o *FindPetsByOwnerIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindPetsByOwnerIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
