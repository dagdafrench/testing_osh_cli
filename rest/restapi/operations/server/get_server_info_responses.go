package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dagdafrench/testing_osh_cli/rest/models"
)

/*GetServerInfoOK Server info response

swagger:response getServerInfoOK
*/
type GetServerInfoOK struct {

	// In: body
	Payload GetServerInfoOKBodyBody `json:"body,omitempty"`
}

// NewGetServerInfoOK creates GetServerInfoOK with default headers values
func NewGetServerInfoOK() *GetServerInfoOK {
	return &GetServerInfoOK{}
}

// WithPayload adds the payload to the get server info o k response
func (o *GetServerInfoOK) WithPayload(payload GetServerInfoOKBodyBody) *GetServerInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server info o k response
func (o *GetServerInfoOK) SetPayload(payload GetServerInfoOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetServerInfoDefault Unexpected error

swagger:response getServerInfoDefault
*/
type GetServerInfoDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetServerInfoDefault creates GetServerInfoDefault with default headers values
func NewGetServerInfoDefault(code int) *GetServerInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServerInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get server info default response
func (o *GetServerInfoDefault) WithStatusCode(code int) *GetServerInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get server info default response
func (o *GetServerInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get server info default response
func (o *GetServerInfoDefault) WithPayload(payload *models.ErrorResponse) *GetServerInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get server info default response
func (o *GetServerInfoDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServerInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
