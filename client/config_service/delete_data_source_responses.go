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

// DeleteDataSourceReader is a Reader for the DeleteDataSource structure.
type DeleteDataSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDataSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDataSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDataSourceOK creates a DeleteDataSourceOK with default headers values
func NewDeleteDataSourceOK() *DeleteDataSourceOK {
	return &DeleteDataSourceOK{}
}

/*DeleteDataSourceOK handles this case with default header values.

DeleteDataSourceOK delete data source o k
*/
type DeleteDataSourceOK struct {
	Payload *models.RestDeleteDataSourceResponse
}

func (o *DeleteDataSourceOK) Error() string {
	return fmt.Sprintf("[DELETE /config/datasource/{Name}][%d] deleteDataSourceOK  %+v", 200, o.Payload)
}

func (o *DeleteDataSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteDataSourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
