// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v1/models"
)

// StorageServiceForceTerminatePostReader is a Reader for the StorageServiceForceTerminatePost structure.
type StorageServiceForceTerminatePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceForceTerminatePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceForceTerminatePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceForceTerminatePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceForceTerminatePostOK creates a StorageServiceForceTerminatePostOK with default headers values
func NewStorageServiceForceTerminatePostOK() *StorageServiceForceTerminatePostOK {
	return &StorageServiceForceTerminatePostOK{}
}

/*
StorageServiceForceTerminatePostOK handles this case with default header values.

StorageServiceForceTerminatePostOK storage service force terminate post o k
*/
type StorageServiceForceTerminatePostOK struct {
}

func (o *StorageServiceForceTerminatePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceForceTerminatePostDefault creates a StorageServiceForceTerminatePostDefault with default headers values
func NewStorageServiceForceTerminatePostDefault(code int) *StorageServiceForceTerminatePostDefault {
	return &StorageServiceForceTerminatePostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceForceTerminatePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceForceTerminatePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service force terminate post default response
func (o *StorageServiceForceTerminatePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceForceTerminatePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceForceTerminatePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceForceTerminatePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
