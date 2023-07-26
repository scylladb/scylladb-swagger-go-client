// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStorageServiceNativeTransportPostParams creates a new StorageServiceNativeTransportPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceNativeTransportPostParams() *StorageServiceNativeTransportPostParams {
	return &StorageServiceNativeTransportPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceNativeTransportPostParamsWithTimeout creates a new StorageServiceNativeTransportPostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceNativeTransportPostParamsWithTimeout(timeout time.Duration) *StorageServiceNativeTransportPostParams {
	return &StorageServiceNativeTransportPostParams{
		timeout: timeout,
	}
}

// NewStorageServiceNativeTransportPostParamsWithContext creates a new StorageServiceNativeTransportPostParams object
// with the ability to set a context for a request.
func NewStorageServiceNativeTransportPostParamsWithContext(ctx context.Context) *StorageServiceNativeTransportPostParams {
	return &StorageServiceNativeTransportPostParams{
		Context: ctx,
	}
}

// NewStorageServiceNativeTransportPostParamsWithHTTPClient creates a new StorageServiceNativeTransportPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceNativeTransportPostParamsWithHTTPClient(client *http.Client) *StorageServiceNativeTransportPostParams {
	return &StorageServiceNativeTransportPostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceNativeTransportPostParams contains all the parameters to send to the API endpoint

	for the storage service native transport post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceNativeTransportPostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service native transport post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceNativeTransportPostParams) WithDefaults() *StorageServiceNativeTransportPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service native transport post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceNativeTransportPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) WithTimeout(timeout time.Duration) *StorageServiceNativeTransportPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) WithContext(ctx context.Context) *StorageServiceNativeTransportPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) WithHTTPClient(client *http.Client) *StorageServiceNativeTransportPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service native transport post params
func (o *StorageServiceNativeTransportPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceNativeTransportPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
