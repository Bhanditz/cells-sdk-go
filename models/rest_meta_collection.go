// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RestMetaCollection rest meta collection
// swagger:model restMetaCollection
type RestMetaCollection struct {

	// metadatas
	Metadatas []*RestMetadata `json:"Metadatas"`

	// node path
	NodePath string `json:"NodePath,omitempty"`
}

// Validate validates this rest meta collection
func (m *RestMetaCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadatas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestMetaCollection) validateMetadatas(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadatas) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadatas); i++ {
		if swag.IsZero(m.Metadatas[i]) { // not required
			continue
		}

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestMetaCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestMetaCollection) UnmarshalBinary(b []byte) error {
	var res RestMetaCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
