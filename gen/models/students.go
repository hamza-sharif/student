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

// Students students
// swagger:model students
type Students struct {

	// age
	// Required: true
	Age *int32 `json:"age"`

	// class
	// Required: true
	// Min Length: 5
	Class *string `json:"class"`

	// grade
	// Required: true
	// Min Length: 2
	Grade *string `json:"grade"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`
}

// Validate validates this students
func (m *Students) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrade(formats); err != nil {
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

func (m *Students) validateAge(formats strfmt.Registry) error {

	if err := validate.Required("age", "body", m.Age); err != nil {
		return err
	}

	return nil
}

func (m *Students) validateClass(formats strfmt.Registry) error {

	if err := validate.Required("class", "body", m.Class); err != nil {
		return err
	}

	if err := validate.MinLength("class", "body", string(*m.Class), 5); err != nil {
		return err
	}

	return nil
}

func (m *Students) validateGrade(formats strfmt.Registry) error {

	if err := validate.Required("grade", "body", m.Grade); err != nil {
		return err
	}

	if err := validate.MinLength("grade", "body", string(*m.Grade), 2); err != nil {
		return err
	}

	return nil
}

func (m *Students) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Students) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Students) UnmarshalBinary(b []byte) error {
	var res Students
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}