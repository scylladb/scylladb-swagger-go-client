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

// NewCacheServiceMetricsKeyEntriesGetParams creates a new CacheServiceMetricsKeyEntriesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCacheServiceMetricsKeyEntriesGetParams() *CacheServiceMetricsKeyEntriesGetParams {
	return &CacheServiceMetricsKeyEntriesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsKeyEntriesGetParamsWithTimeout creates a new CacheServiceMetricsKeyEntriesGetParams object
// with the ability to set a timeout on a request.
func NewCacheServiceMetricsKeyEntriesGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsKeyEntriesGetParams {
	return &CacheServiceMetricsKeyEntriesGetParams{
		timeout: timeout,
	}
}

// NewCacheServiceMetricsKeyEntriesGetParamsWithContext creates a new CacheServiceMetricsKeyEntriesGetParams object
// with the ability to set a context for a request.
func NewCacheServiceMetricsKeyEntriesGetParamsWithContext(ctx context.Context) *CacheServiceMetricsKeyEntriesGetParams {
	return &CacheServiceMetricsKeyEntriesGetParams{
		Context: ctx,
	}
}

// NewCacheServiceMetricsKeyEntriesGetParamsWithHTTPClient creates a new CacheServiceMetricsKeyEntriesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCacheServiceMetricsKeyEntriesGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsKeyEntriesGetParams {
	return &CacheServiceMetricsKeyEntriesGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsKeyEntriesGetParams contains all the parameters to send to the API endpoint

	for the cache service metrics key entries get operation.

	Typically these are written to a http.Request.
*/
type CacheServiceMetricsKeyEntriesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cache service metrics key entries get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsKeyEntriesGetParams) WithDefaults() *CacheServiceMetricsKeyEntriesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cache service metrics key entries get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsKeyEntriesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsKeyEntriesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) WithContext(ctx context.Context) *CacheServiceMetricsKeyEntriesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsKeyEntriesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics key entries get params
func (o *CacheServiceMetricsKeyEntriesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsKeyEntriesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
