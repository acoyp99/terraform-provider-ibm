// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ZoneReference ZoneReference
// swagger:model ZoneReference
type ZoneReference struct {

	// The URL for this zone
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`

	// The name for this zone
	// Required: true
	// Max Length: 63
	// Min Length: 1
	// Pattern: ^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$
	Name *string `json:"name"`
}

// Validate validates this zone reference
func (m *ZoneReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	if err := validate.Pattern("href", "body", string(*m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *ZoneReference) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 63); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneReference) UnmarshalBinary(b []byte) error {
	var res ZoneReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
