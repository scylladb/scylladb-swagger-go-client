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

// NewStorageServiceBatchSizeFailureThresholdGetParams creates a new StorageServiceBatchSizeFailureThresholdGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceBatchSizeFailureThresholdGetParams() *StorageServiceBatchSizeFailureThresholdGetParams {
	return &StorageServiceBatchSizeFailureThresholdGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceBatchSizeFailureThresholdGetParamsWithTimeout creates a new StorageServiceBatchSizeFailureThresholdGetParams object
// with the ability to set a timeout on a request.
func NewStorageServiceBatchSizeFailureThresholdGetParamsWithTimeout(timeout time.Duration) *StorageServiceBatchSizeFailureThresholdGetParams {
	return &StorageServiceBatchSizeFailureThresholdGetParams{
		timeout: timeout,
	}
}

// NewStorageServiceBatchSizeFailureThresholdGetParamsWithContext creates a new StorageServiceBatchSizeFailureThresholdGetParams object
// with the ability to set a context for a request.
func NewStorageServiceBatchSizeFailureThresholdGetParamsWithContext(ctx context.Context) *StorageServiceBatchSizeFailureThresholdGetParams {
	return &StorageServiceBatchSizeFailureThresholdGetParams{
		Context: ctx,
	}
}

// NewStorageServiceBatchSizeFailureThresholdGetParamsWithHTTPClient creates a new StorageServiceBatchSizeFailureThresholdGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceBatchSizeFailureThresholdGetParamsWithHTTPClient(client *http.Client) *StorageServiceBatchSizeFailureThresholdGetParams {
	return &StorageServiceBatchSizeFailureThresholdGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceBatchSizeFailureThresholdGetParams contains all the parameters to send to the API endpoint

	for the storage service batch size failure threshold get operation.

	Typically these are written to a http.Request.
*/
type StorageServiceBatchSizeFailureThresholdGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service batch size failure threshold get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceBatchSizeFailureThresholdGetParams) WithDefaults() *StorageServiceBatchSizeFailureThresholdGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service batch size failure threshold get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceBatchSizeFailureThresholdGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) WithTimeout(timeout time.Duration) *StorageServiceBatchSizeFailureThresholdGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) WithContext(ctx context.Context) *StorageServiceBatchSizeFailureThresholdGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) WithHTTPClient(client *http.Client) *StorageServiceBatchSizeFailureThresholdGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service batch size failure threshold get params
func (o *StorageServiceBatchSizeFailureThresholdGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceBatchSizeFailureThresholdGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
