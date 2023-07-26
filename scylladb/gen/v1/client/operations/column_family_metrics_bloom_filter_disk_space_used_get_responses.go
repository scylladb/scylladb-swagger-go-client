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

// ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader is a Reader for the ColumnFamilyMetricsBloomFilterDiskSpaceUsedGet structure.
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK creates a ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK with default headers values
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK() *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK {
	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK{}
}

/*
ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK handles this case with default header values.

ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK column family metrics bloom filter disk space used get o k
*/
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault creates a ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault with default headers values
func NewColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault(code int) *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault {
	return &ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics bloom filter disk space used get default response
func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsBloomFilterDiskSpaceUsedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
