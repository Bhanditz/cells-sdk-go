// Code generated by go-swagger; DO NOT EDIT.

package tree_service

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

// NewHeadNodeParams creates a new HeadNodeParams object
// with the default values initialized.
func NewHeadNodeParams() *HeadNodeParams {
	var ()
	return &HeadNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHeadNodeParamsWithTimeout creates a new HeadNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHeadNodeParamsWithTimeout(timeout time.Duration) *HeadNodeParams {
	var ()
	return &HeadNodeParams{

		timeout: timeout,
	}
}

// NewHeadNodeParamsWithContext creates a new HeadNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewHeadNodeParamsWithContext(ctx context.Context) *HeadNodeParams {
	var ()
	return &HeadNodeParams{

		Context: ctx,
	}
}

// NewHeadNodeParamsWithHTTPClient creates a new HeadNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHeadNodeParamsWithHTTPClient(client *http.Client) *HeadNodeParams {
	var ()
	return &HeadNodeParams{
		HTTPClient: client,
	}
}

/*HeadNodeParams contains all the parameters to send to the API endpoint
for the head node operation typically these are written to a http.Request
*/
type HeadNodeParams struct {

	/*Node*/
	Node string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the head node params
func (o *HeadNodeParams) WithTimeout(timeout time.Duration) *HeadNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head node params
func (o *HeadNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head node params
func (o *HeadNodeParams) WithContext(ctx context.Context) *HeadNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head node params
func (o *HeadNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head node params
func (o *HeadNodeParams) WithHTTPClient(client *http.Client) *HeadNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head node params
func (o *HeadNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNode adds the node to the head node params
func (o *HeadNodeParams) WithNode(node string) *HeadNodeParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the head node params
func (o *HeadNodeParams) SetNode(node string) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *HeadNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Node
	if err := r.SetPathParam("Node", o.Node); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}