// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// APIValue Value is the value of the field.
// swagger:model apiValue
type APIValue struct {

	// A double value
	DoubleValue float64 `json:"double_value,omitempty"`

	// An integer value
	IntValue string `json:"int_value,omitempty"`

	// A string value
	StringValue string `json:"string_value,omitempty"`
}

// Validate validates this api value
func (m *APIValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIValue) UnmarshalBinary(b []byte) error {
	var res APIValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}