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

// NewColumnFamilyMetricsRowCacheHitGetParams creates a new ColumnFamilyMetricsRowCacheHitGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewColumnFamilyMetricsRowCacheHitGetParams() *ColumnFamilyMetricsRowCacheHitGetParams {
	return &ColumnFamilyMetricsRowCacheHitGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsRowCacheHitGetParamsWithTimeout creates a new ColumnFamilyMetricsRowCacheHitGetParams object
// with the ability to set a timeout on a request.
func NewColumnFamilyMetricsRowCacheHitGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsRowCacheHitGetParams {
	return &ColumnFamilyMetricsRowCacheHitGetParams{
		timeout: timeout,
	}
}

// NewColumnFamilyMetricsRowCacheHitGetParamsWithContext creates a new ColumnFamilyMetricsRowCacheHitGetParams object
// with the ability to set a context for a request.
func NewColumnFamilyMetricsRowCacheHitGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsRowCacheHitGetParams {
	return &ColumnFamilyMetricsRowCacheHitGetParams{
		Context: ctx,
	}
}

// NewColumnFamilyMetricsRowCacheHitGetParamsWithHTTPClient creates a new ColumnFamilyMetricsRowCacheHitGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewColumnFamilyMetricsRowCacheHitGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsRowCacheHitGetParams {
	return &ColumnFamilyMetricsRowCacheHitGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsRowCacheHitGetParams contains all the parameters to send to the API endpoint

	for the column family metrics row cache hit get operation.

	Typically these are written to a http.Request.
*/
type ColumnFamilyMetricsRowCacheHitGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the column family metrics row cache hit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsRowCacheHitGetParams) WithDefaults() *ColumnFamilyMetricsRowCacheHitGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the column family metrics row cache hit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsRowCacheHitGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsRowCacheHitGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsRowCacheHitGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsRowCacheHitGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics row cache hit get params
func (o *ColumnFamilyMetricsRowCacheHitGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsRowCacheHitGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
