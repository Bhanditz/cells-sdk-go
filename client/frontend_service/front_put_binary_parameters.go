// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

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

	models "github.com/pydio/cells-sdk-go/models"
)

// NewFrontPutBinaryParams creates a new FrontPutBinaryParams object
// with the default values initialized.
func NewFrontPutBinaryParams() *FrontPutBinaryParams {
	var ()
	return &FrontPutBinaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFrontPutBinaryParamsWithTimeout creates a new FrontPutBinaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFrontPutBinaryParamsWithTimeout(timeout time.Duration) *FrontPutBinaryParams {
	var ()
	return &FrontPutBinaryParams{

		timeout: timeout,
	}
}

// NewFrontPutBinaryParamsWithContext creates a new FrontPutBinaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewFrontPutBinaryParamsWithContext(ctx context.Context) *FrontPutBinaryParams {
	var ()
	return &FrontPutBinaryParams{

		Context: ctx,
	}
}

// NewFrontPutBinaryParamsWithHTTPClient creates a new FrontPutBinaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFrontPutBinaryParamsWithHTTPClient(client *http.Client) *FrontPutBinaryParams {
	var ()
	return &FrontPutBinaryParams{
		HTTPClient: client,
	}
}

/*FrontPutBinaryParams contains all the parameters to send to the API endpoint
for the front put binary operation typically these are written to a http.Request
*/
type FrontPutBinaryParams struct {

	/*BinaryType*/
	BinaryType models.RestFrontBinaryType
	/*UUID*/
	UUID string
	/*Body*/
	Body *models.RestFrontBinaryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the front put binary params
func (o *FrontPutBinaryParams) WithTimeout(timeout time.Duration) *FrontPutBinaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the front put binary params
func (o *FrontPutBinaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the front put binary params
func (o *FrontPutBinaryParams) WithContext(ctx context.Context) *FrontPutBinaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the front put binary params
func (o *FrontPutBinaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the front put binary params
func (o *FrontPutBinaryParams) WithHTTPClient(client *http.Client) *FrontPutBinaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the front put binary params
func (o *FrontPutBinaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBinaryType adds the binaryType to the front put binary params
func (o *FrontPutBinaryParams) WithBinaryType(binaryType models.RestFrontBinaryType) *FrontPutBinaryParams {
	o.SetBinaryType(binaryType)
	return o
}

// SetBinaryType adds the binaryType to the front put binary params
func (o *FrontPutBinaryParams) SetBinaryType(binaryType models.RestFrontBinaryType) {
	o.BinaryType = binaryType
}

// WithUUID adds the uuid to the front put binary params
func (o *FrontPutBinaryParams) WithUUID(uuid string) *FrontPutBinaryParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the front put binary params
func (o *FrontPutBinaryParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WithBody adds the body to the front put binary params
func (o *FrontPutBinaryParams) WithBody(body *models.RestFrontBinaryRequest) *FrontPutBinaryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the front put binary params
func (o *FrontPutBinaryParams) SetBody(body *models.RestFrontBinaryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FrontPutBinaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param BinaryType
	if err := r.SetPathParam("BinaryType", string(o.BinaryType)); err != nil {
		return err
	}

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}