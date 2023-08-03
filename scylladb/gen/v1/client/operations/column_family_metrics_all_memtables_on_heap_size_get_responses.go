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

// ColumnFamilyMetricsAllMemtablesOnHeapSizeGetReader is a Reader for the ColumnFamilyMetricsAllMemtablesOnHeapSizeGet structure.
type ColumnFamilyMetricsAllMemtablesOnHeapSizeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK creates a ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK with default headers values
func NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK() *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK {
	return &ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK{}
}

/*
ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK handles this case with default header values.

ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK column family metrics all memtables on heap size get o k
*/
type ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault creates a ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault with default headers values
func NewColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault(code int) *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault {
	return &ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault{
		_statusCode: code,
	}
}

/*
ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics all memtables on heap size get default response
func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsAllMemtablesOnHeapSizeGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
