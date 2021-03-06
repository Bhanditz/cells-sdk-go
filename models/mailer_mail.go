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

// MailerMail mailer mail
// swagger:model mailerMail
type MailerMail struct {

	// attachments
	Attachments []string `json:"Attachments"`

	// cc
	Cc []*MailerUser `json:"Cc"`

	// content Html
	ContentHTML string `json:"ContentHtml,omitempty"`

	// content markdown
	ContentMarkdown string `json:"ContentMarkdown,omitempty"`

	// content plain
	ContentPlain string `json:"ContentPlain,omitempty"`

	// date sent
	DateSent string `json:"DateSent,omitempty"`

	// from
	From *MailerUser `json:"From,omitempty"`

	// retries
	Retries int32 `json:"Retries,omitempty"`

	// subject
	Subject string `json:"Subject,omitempty"`

	// template data
	TemplateData map[string]string `json:"TemplateData,omitempty"`

	// template Id
	TemplateID string `json:"TemplateId,omitempty"`

	// thread index
	ThreadIndex string `json:"ThreadIndex,omitempty"`

	// Could be used for Re: ... conversations
	ThreadUUID string `json:"ThreadUuid,omitempty"`

	// to
	To []*MailerUser `json:"To"`

	// send errors
	SendErrors []string `json:"sendErrors"`
}

// Validate validates this mailer mail
func (m *MailerMail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailerMail) validateCc(formats strfmt.Registry) error {

	if swag.IsZero(m.Cc) { // not required
		return nil
	}

	for i := 0; i < len(m.Cc); i++ {
		if swag.IsZero(m.Cc[i]) { // not required
			continue
		}

		if m.Cc[i] != nil {
			if err := m.Cc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Cc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MailerMail) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("From")
			}
			return err
		}
	}

	return nil
}

func (m *MailerMail) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	for i := 0; i < len(m.To); i++ {
		if swag.IsZero(m.To[i]) { // not required
			continue
		}

		if m.To[i] != nil {
			if err := m.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("To" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MailerMail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailerMail) UnmarshalBinary(b []byte) error {
	var res MailerMail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
