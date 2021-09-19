// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorMsg error msg
//
// swagger:model errorMsg
type ErrorMsg struct {

	// message
	Message string `json:"message,omitempty"`

	// type
	// Required: true
	// Enum: [ErrInvalidParams ErrInvalidObject]
	Type *string `json:"type"`
}

// Validate validates this error msg
func (m *ErrorMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var errorMsgTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ErrInvalidParams","ErrInvalidObject"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		errorMsgTypeTypePropEnum = append(errorMsgTypeTypePropEnum, v)
	}
}

const (

	// ErrorMsgTypeErrInvalidParams captures enum value "ErrInvalidParams"
	ErrorMsgTypeErrInvalidParams string = "ErrInvalidParams"

	// ErrorMsgTypeErrInvalidObject captures enum value "ErrInvalidObject"
	ErrorMsgTypeErrInvalidObject string = "ErrInvalidObject"
)

// prop value enum
func (m *ErrorMsg) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, errorMsgTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ErrorMsg) validateType(formats strfmt.Registry) error {
	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error msg based on context it is used
func (m *ErrorMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorMsg) UnmarshalBinary(b []byte) error {
	var res ErrorMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
