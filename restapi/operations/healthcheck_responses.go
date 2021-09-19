// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// HealthcheckOKCode is the HTTP code returned for type HealthcheckOK
const HealthcheckOKCode int = 200

/*HealthcheckOK Ok.

swagger:response healthcheckOK
*/
type HealthcheckOK struct{}

// NewHealthcheckOK creates HealthcheckOK with default headers values
func NewHealthcheckOK() *HealthcheckOK {
	return &HealthcheckOK{}
}

// WriteResponse to the client
func (o *HealthcheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
