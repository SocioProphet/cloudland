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

// CreateFlavorParamsBodyFlavor create flavor params body flavor
// swagger:model createFlavorParamsBodyFlavor
type CreateFlavorParamsBodyFlavor struct {

	// disk
	Disk int32 `json:"disk,omitempty"`

	// name
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// raw
	Raw int32 `json:"raw,omitempty"`

	// vcpus
	Vcpus int32 `json:"vcpus,omitempty"`
}

// Validate validates this create flavor params body flavor
func (m *CreateFlavorParamsBodyFlavor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFlavorParamsBodyFlavor) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFlavorParamsBodyFlavor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFlavorParamsBodyFlavor) UnmarshalBinary(b []byte) error {
	var res CreateFlavorParamsBodyFlavor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
