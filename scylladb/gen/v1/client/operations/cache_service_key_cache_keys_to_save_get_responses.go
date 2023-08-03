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

// CacheServiceKeyCacheKeysToSaveGetReader is a Reader for the CacheServiceKeyCacheKeysToSaveGet structure.
type CacheServiceKeyCacheKeysToSaveGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceKeyCacheKeysToSaveGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceKeyCacheKeysToSaveGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceKeyCacheKeysToSaveGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceKeyCacheKeysToSaveGetOK creates a CacheServiceKeyCacheKeysToSaveGetOK with default headers values
func NewCacheServiceKeyCacheKeysToSaveGetOK() *CacheServiceKeyCacheKeysToSaveGetOK {
	return &CacheServiceKeyCacheKeysToSaveGetOK{}
}

/*
CacheServiceKeyCacheKeysToSaveGetOK handles this case with default header values.

CacheServiceKeyCacheKeysToSaveGetOK cache service key cache keys to save get o k
*/
type CacheServiceKeyCacheKeysToSaveGetOK struct {
	Payload int32
}

func (o *CacheServiceKeyCacheKeysToSaveGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceKeyCacheKeysToSaveGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceKeyCacheKeysToSaveGetDefault creates a CacheServiceKeyCacheKeysToSaveGetDefault with default headers values
func NewCacheServiceKeyCacheKeysToSaveGetDefault(code int) *CacheServiceKeyCacheKeysToSaveGetDefault {
	return &CacheServiceKeyCacheKeysToSaveGetDefault{
		_statusCode: code,
	}
}

/*
CacheServiceKeyCacheKeysToSaveGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceKeyCacheKeysToSaveGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service key cache keys to save get default response
func (o *CacheServiceKeyCacheKeysToSaveGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceKeyCacheKeysToSaveGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceKeyCacheKeysToSaveGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceKeyCacheKeysToSaveGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
