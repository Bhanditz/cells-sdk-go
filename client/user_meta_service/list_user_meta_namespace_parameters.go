// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListUserMetaNamespaceParams creates a new ListUserMetaNamespaceParams object
// with the default values initialized.
func NewListUserMetaNamespaceParams() *ListUserMetaNamespaceParams {

	return &ListUserMetaNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListUserMetaNamespaceParamsWithTimeout creates a new ListUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUserMetaNamespaceParamsWithTimeout(timeout time.Duration) *ListUserMetaNamespaceParams {

	return &ListUserMetaNamespaceParams{

		timeout: timeout,
	}
}

// NewListUserMetaNamespaceParamsWithContext creates a new ListUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUserMetaNamespaceParamsWithContext(ctx context.Context) *ListUserMetaNamespaceParams {

	return &ListUserMetaNamespaceParams{

		Context: ctx,
	}
}

// NewListUserMetaNamespaceParamsWithHTTPClient creates a new ListUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUserMetaNamespaceParamsWithHTTPClient(client *http.Client) *ListUserMetaNamespaceParams {

	return &ListUserMetaNamespaceParams{
		HTTPClient: client,
	}
}

/*ListUserMetaNamespaceParams contains all the parameters to send to the API endpoint
for the list user meta namespace operation typically these are written to a http.Request
*/
type ListUserMetaNamespaceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) WithTimeout(timeout time.Duration) *ListUserMetaNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) WithContext(ctx context.Context) *ListUserMetaNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) WithHTTPClient(client *http.Client) *ListUserMetaNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user meta namespace params
func (o *ListUserMetaNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserMetaNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
