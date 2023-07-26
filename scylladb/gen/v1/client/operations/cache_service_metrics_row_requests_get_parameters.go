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

// NewCacheServiceMetricsRowRequestsGetParams creates a new CacheServiceMetricsRowRequestsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCacheServiceMetricsRowRequestsGetParams() *CacheServiceMetricsRowRequestsGetParams {
	return &CacheServiceMetricsRowRequestsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowRequestsGetParamsWithTimeout creates a new CacheServiceMetricsRowRequestsGetParams object
// with the ability to set a timeout on a request.
func NewCacheServiceMetricsRowRequestsGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowRequestsGetParams {
	return &CacheServiceMetricsRowRequestsGetParams{
		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowRequestsGetParamsWithContext creates a new CacheServiceMetricsRowRequestsGetParams object
// with the ability to set a context for a request.
func NewCacheServiceMetricsRowRequestsGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowRequestsGetParams {
	return &CacheServiceMetricsRowRequestsGetParams{
		Context: ctx,
	}
}

// NewCacheServiceMetricsRowRequestsGetParamsWithHTTPClient creates a new CacheServiceMetricsRowRequestsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCacheServiceMetricsRowRequestsGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowRequestsGetParams {
	return &CacheServiceMetricsRowRequestsGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsRowRequestsGetParams contains all the parameters to send to the API endpoint

	for the cache service metrics row requests get operation.

	Typically these are written to a http.Request.
*/
type CacheServiceMetricsRowRequestsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cache service metrics row requests get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsRowRequestsGetParams) WithDefaults() *CacheServiceMetricsRowRequestsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cache service metrics row requests get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CacheServiceMetricsRowRequestsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowRequestsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowRequestsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowRequestsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row requests get params
func (o *CacheServiceMetricsRowRequestsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowRequestsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
