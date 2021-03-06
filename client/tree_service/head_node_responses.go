// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// HeadNodeReader is a Reader for the HeadNode structure.
type HeadNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHeadNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHeadNodeOK creates a HeadNodeOK with default headers values
func NewHeadNodeOK() *HeadNodeOK {
	return &HeadNodeOK{}
}

/*HeadNodeOK handles this case with default header values.

HeadNodeOK head node o k
*/
type HeadNodeOK struct {
	Payload *models.RestHeadNodeResponse
}

func (o *HeadNodeOK) Error() string {
	return fmt.Sprintf("[GET /tree/stat/{Node}][%d] headNodeOK  %+v", 200, o.Payload)
}

func (o *HeadNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestHeadNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
