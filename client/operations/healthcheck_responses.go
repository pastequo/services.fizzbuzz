// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HealthcheckReader is a Reader for the Healthcheck structure.
type HealthcheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthcheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthcheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthcheckOK creates a HealthcheckOK with default headers values
func NewHealthcheckOK() *HealthcheckOK {
	return &HealthcheckOK{}
}

/* HealthcheckOK describes a response with status code 200, with default header values.

Ok.
*/
type HealthcheckOK struct{}

func (o *HealthcheckOK) Error() string {
	return fmt.Sprintf("[GET /healthcheck][%d] healthcheckOK ", 200)
}

func (o *HealthcheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}