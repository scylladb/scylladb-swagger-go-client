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

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams() *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the ability to set a timeout on a request.
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{
		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the ability to set a context for a request.
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{
		Context: ctx,
	}
}

// NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	return &ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams contains all the parameters to send to the API endpoint

	for the column family metrics cas commit estimated recent histogram by name get operation.

	Typically these are written to a http.Request.
*/
type ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams struct {

	/* Name.

	   The column family name in keyspace:name format
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the column family metrics cas commit estimated recent histogram by name get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithDefaults() *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the column family metrics cas commit estimated recent histogram by name get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics cas commit estimated recent histogram by name get params
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCasCommitEstimatedRecentHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
