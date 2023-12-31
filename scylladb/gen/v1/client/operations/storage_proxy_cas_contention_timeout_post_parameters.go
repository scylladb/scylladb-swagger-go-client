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

// NewStorageProxyCasContentionTimeoutPostParams creates a new StorageProxyCasContentionTimeoutPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageProxyCasContentionTimeoutPostParams() *StorageProxyCasContentionTimeoutPostParams {
	return &StorageProxyCasContentionTimeoutPostParams{
		requestTimeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyCasContentionTimeoutPostParamsWithTimeout creates a new StorageProxyCasContentionTimeoutPostParams object
// with the ability to set a timeout on a request.
func NewStorageProxyCasContentionTimeoutPostParamsWithTimeout(timeout time.Duration) *StorageProxyCasContentionTimeoutPostParams {
	return &StorageProxyCasContentionTimeoutPostParams{
		requestTimeout: timeout,
	}
}

// NewStorageProxyCasContentionTimeoutPostParamsWithContext creates a new StorageProxyCasContentionTimeoutPostParams object
// with the ability to set a context for a request.
func NewStorageProxyCasContentionTimeoutPostParamsWithContext(ctx context.Context) *StorageProxyCasContentionTimeoutPostParams {
	return &StorageProxyCasContentionTimeoutPostParams{
		Context: ctx,
	}
}

// NewStorageProxyCasContentionTimeoutPostParamsWithHTTPClient creates a new StorageProxyCasContentionTimeoutPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageProxyCasContentionTimeoutPostParamsWithHTTPClient(client *http.Client) *StorageProxyCasContentionTimeoutPostParams {
	return &StorageProxyCasContentionTimeoutPostParams{
		HTTPClient: client,
	}
}

/*
StorageProxyCasContentionTimeoutPostParams contains all the parameters to send to the API endpoint

	for the storage proxy cas contention timeout post operation.

	Typically these are written to a http.Request.
*/
type StorageProxyCasContentionTimeoutPostParams struct {

	/* Timeout.

	   timeout in second
	*/
	Timeout string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithDefaults hydrates default values in the storage proxy cas contention timeout post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyCasContentionTimeoutPostParams) WithDefaults() *StorageProxyCasContentionTimeoutPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage proxy cas contention timeout post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageProxyCasContentionTimeoutPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithRequestTimeout adds the timeout to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) WithRequestTimeout(timeout time.Duration) *StorageProxyCasContentionTimeoutPostParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) WithContext(ctx context.Context) *StorageProxyCasContentionTimeoutPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) WithHTTPClient(client *http.Client) *StorageProxyCasContentionTimeoutPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimeout adds the timeout to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) WithTimeout(timeout string) *StorageProxyCasContentionTimeoutPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy cas contention timeout post params
func (o *StorageProxyCasContentionTimeoutPostParams) SetTimeout(timeout string) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyCasContentionTimeoutPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param timeout
	qrTimeout := o.Timeout
	qTimeout := qrTimeout
	if qTimeout != "" {

		if err := r.SetQueryParam("timeout", qTimeout); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
