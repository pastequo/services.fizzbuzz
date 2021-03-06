// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FizzbuzzParams fizzbuzz params
//
// swagger:model fizzbuzzParams
type FizzbuzzParams struct {

	// limit
	// Required: true
	// Minimum: 1
	Limit *int32 `json:"limit"`

	// word1
	// Required: true
	Word1 *FizzbuzzWord `json:"word1"`

	// word2
	// Required: true
	Word2 *FizzbuzzWord `json:"word2"`
}

// Validate validates this fizzbuzz params
func (m *FizzbuzzParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWord1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWord2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FizzbuzzParams) validateLimit(formats strfmt.Registry) error {
	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	if err := validate.MinimumInt("limit", "body", int64(*m.Limit), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *FizzbuzzParams) validateWord1(formats strfmt.Registry) error {
	if err := validate.Required("word1", "body", m.Word1); err != nil {
		return err
	}

	if m.Word1 != nil {
		if err := m.Word1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("word1")
			}
			return err
		}
	}

	return nil
}

func (m *FizzbuzzParams) validateWord2(formats strfmt.Registry) error {
	if err := validate.Required("word2", "body", m.Word2); err != nil {
		return err
	}

	if m.Word2 != nil {
		if err := m.Word2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("word2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fizzbuzz params based on the context it is used
func (m *FizzbuzzParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWord1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWord2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FizzbuzzParams) contextValidateWord1(ctx context.Context, formats strfmt.Registry) error {
	if m.Word1 != nil {
		if err := m.Word1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("word1")
			}
			return err
		}
	}

	return nil
}

func (m *FizzbuzzParams) contextValidateWord2(ctx context.Context, formats strfmt.Registry) error {
	if m.Word2 != nil {
		if err := m.Word2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("word2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FizzbuzzParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FizzbuzzParams) UnmarshalBinary(b []byte) error {
	var res FizzbuzzParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
