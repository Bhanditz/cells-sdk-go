// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// IdmPolicyConditions idm policy conditions
// swagger:model idmPolicyConditions
type IdmPolicyConditions map[string]IdmPolicyCondition

// Validate validates this idm policy conditions
func (m IdmPolicyConditions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", IdmPolicyConditions(m)); err != nil {
		return err
	}

	for k := range m {

		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}