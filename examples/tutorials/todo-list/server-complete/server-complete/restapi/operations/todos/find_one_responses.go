// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/models"
)

// FindOneOKCode is the HTTP code returned for type FindOneOK
const FindOneOKCode int = 200

/*FindOneOK OK

swagger:response findOneOK
*/
type FindOneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewFindOneOK creates FindOneOK with default headers values
func NewFindOneOK() *FindOneOK {

	return &FindOneOK{}
}

// WithPayload adds the payload to the find one o k response
func (o *FindOneOK) WithPayload(payload *models.Item) *FindOneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find one o k response
func (o *FindOneOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindOneDefault error

swagger:response findOneDefault
*/
type FindOneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindOneDefault creates FindOneDefault with default headers values
func NewFindOneDefault(code int) *FindOneDefault {
	if code <= 0 {
		code = 500
	}

	return &FindOneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find one default response
func (o *FindOneDefault) WithStatusCode(code int) *FindOneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find one default response
func (o *FindOneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find one default response
func (o *FindOneDefault) WithPayload(payload *models.Error) *FindOneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find one default response
func (o *FindOneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
