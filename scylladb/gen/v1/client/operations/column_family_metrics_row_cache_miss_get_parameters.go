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

// NewColumnFamilyMetricsRowCacheMissGetParams creates a new ColumnFamilyMetricsRowCacheMissGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewColumnFamilyMetricsRowCacheMissGetParams() *ColumnFamilyMetricsRowCacheMissGetParams {
	return &ColumnFamilyMetricsRowCacheMissGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsRowCacheMissGetParamsWithTimeout creates a new ColumnFamilyMetricsRowCacheMissGetParams object
// with the ability to set a timeout on a request.
func NewColumnFamilyMetricsRowCacheMissGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsRowCacheMissGetParams {
	return &ColumnFamilyMetricsRowCacheMissGetParams{
		timeout: timeout,
	}
}

// NewColumnFamilyMetricsRowCacheMissGetParamsWithContext creates a new ColumnFamilyMetricsRowCacheMissGetParams object
// with the ability to set a context for a request.
func NewColumnFamilyMetricsRowCacheMissGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsRowCacheMissGetParams {
	return &ColumnFamilyMetricsRowCacheMissGetParams{
		Context: ctx,
	}
}

// NewColumnFamilyMetricsRowCacheMissGetParamsWithHTTPClient creates a new ColumnFamilyMetricsRowCacheMissGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewColumnFamilyMetricsRowCacheMissGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsRowCacheMissGetParams {
	return &ColumnFamilyMetricsRowCacheMissGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsRowCacheMissGetParams contains all the parameters to send to the API endpoint

	for the column family metrics row cache miss get operation.

	Typically these are written to a http.Request.
*/
type ColumnFamilyMetricsRowCacheMissGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the column family metrics row cache miss get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsRowCacheMissGetParams) WithDefaults() *ColumnFamilyMetricsRowCacheMissGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the column family metrics row cache miss get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsRowCacheMissGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsRowCacheMissGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsRowCacheMissGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsRowCacheMissGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics row cache miss get params
func (o *ColumnFamilyMetricsRowCacheMissGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsRowCacheMissGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
