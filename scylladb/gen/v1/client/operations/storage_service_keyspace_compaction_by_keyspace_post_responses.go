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

// StorageServiceKeyspaceCompactionByKeyspacePostReader is a Reader for the StorageServiceKeyspaceCompactionByKeyspacePost structure.
type StorageServiceKeyspaceCompactionByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceKeyspaceCompactionByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceKeyspaceCompactionByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceKeyspaceCompactionByKeyspacePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceKeyspaceCompactionByKeyspacePostOK creates a StorageServiceKeyspaceCompactionByKeyspacePostOK with default headers values
func NewStorageServiceKeyspaceCompactionByKeyspacePostOK() *StorageServiceKeyspaceCompactionByKeyspacePostOK {
	return &StorageServiceKeyspaceCompactionByKeyspacePostOK{}
}

/*
StorageServiceKeyspaceCompactionByKeyspacePostOK handles this case with default header values.

StorageServiceKeyspaceCompactionByKeyspacePostOK storage service keyspace compaction by keyspace post o k
*/
type StorageServiceKeyspaceCompactionByKeyspacePostOK struct {
}

func (o *StorageServiceKeyspaceCompactionByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStorageServiceKeyspaceCompactionByKeyspacePostDefault creates a StorageServiceKeyspaceCompactionByKeyspacePostDefault with default headers values
func NewStorageServiceKeyspaceCompactionByKeyspacePostDefault(code int) *StorageServiceKeyspaceCompactionByKeyspacePostDefault {
	return &StorageServiceKeyspaceCompactionByKeyspacePostDefault{
		_statusCode: code,
	}
}

/*
StorageServiceKeyspaceCompactionByKeyspacePostDefault handles this case with default header values.

internal server error
*/
type StorageServiceKeyspaceCompactionByKeyspacePostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service keyspace compaction by keyspace post default response
func (o *StorageServiceKeyspaceCompactionByKeyspacePostDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceKeyspaceCompactionByKeyspacePostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceKeyspaceCompactionByKeyspacePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceKeyspaceCompactionByKeyspacePostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
