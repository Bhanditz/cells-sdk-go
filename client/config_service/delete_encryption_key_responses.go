// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// DeleteEncryptionKeyReader is a Reader for the DeleteEncryptionKey structure.
type DeleteEncryptionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEncryptionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEncryptionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEncryptionKeyOK creates a DeleteEncryptionKeyOK with default headers values
func NewDeleteEncryptionKeyOK() *DeleteEncryptionKeyOK {
	return &DeleteEncryptionKeyOK{}
}

/*DeleteEncryptionKeyOK handles this case with default header values.

DeleteEncryptionKeyOK delete encryption key o k
*/
type DeleteEncryptionKeyOK struct {
	Payload *models.EncryptionAdminDeleteKeyResponse
}

func (o *DeleteEncryptionKeyOK) Error() string {
	return fmt.Sprintf("[POST /config/encryption/delete][%d] deleteEncryptionKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteEncryptionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionAdminDeleteKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}