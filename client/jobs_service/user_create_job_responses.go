// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// UserCreateJobReader is a Reader for the UserCreateJob structure.
type UserCreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCreateJobOK creates a UserCreateJobOK with default headers values
func NewUserCreateJobOK() *UserCreateJobOK {
	return &UserCreateJobOK{}
}

/*UserCreateJobOK handles this case with default header values.

UserCreateJobOK user create job o k
*/
type UserCreateJobOK struct {
	Payload *models.RestUserJobResponse
}

func (o *UserCreateJobOK) Error() string {
	return fmt.Sprintf("[PUT /jobs/user/{JobName}][%d] userCreateJobOK  %+v", 200, o.Payload)
}

func (o *UserCreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUserJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
