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

// NewStorageProxyMetricsCasReadConditionNotMetGetParams creates a new StorageProxyMetricsCasReadConditionNotMetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageProxyMetricsCasReadConditionNotMetGetParams() *StorageProxyMetricsCasReadConditionNotMetGetParams {
	return &StorageProxyMetricsCasReadConditionNotMetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithTimeout creates a new StorageProxyMetricsCasReadConditionNotMetGetParams object
// with the ability to set a timeout on a request.
func NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	return &StorageProxyMetricsCasReadConditionNotMetGetParams{
		timeout: timeout,
	}
}

// NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithContext creates a new StorageProxyMetricsCasReadConditionNotMetGetParams object
// with the ability to set a context for a request.
func NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithContext(ctx context.Context) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	return &StorageProxyMetricsCasReadConditionNotMetGetParams{
		Context: ctx,
	}
}

// NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithHTTPClient creates a new StorageProxyMetricsCasReadConditionNotMetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageProxyMetricsCasReadConditionNotMetGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	return &StorageProxyMetricsCasReadConditionNotMetGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyMetricsCasReadConditionNotMetGetParams contains all the parameters to send to the API endpoint

	for the storage proxy metrics cas read condition not met get operation.

	Typically these are written to a http.Request.
*/
type StorageProxyMetricsCasReadConditionNotMetGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage proxy metrics cas read condition not met get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) WithDefaults() *StorageProxyMetricsCasReadConditionNotMetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage proxy metrics cas read condition not met get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) WithContext(ctx context.Context) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadConditionNotMetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics cas read condition not met get params
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsCasReadConditionNotMetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
