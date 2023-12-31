// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigCrossNodeTimeoutParams creates a new FindConfigCrossNodeTimeoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigCrossNodeTimeoutParams() *FindConfigCrossNodeTimeoutParams {
	return &FindConfigCrossNodeTimeoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCrossNodeTimeoutParamsWithTimeout creates a new FindConfigCrossNodeTimeoutParams object
// with the ability to set a timeout on a request.
func NewFindConfigCrossNodeTimeoutParamsWithTimeout(timeout time.Duration) *FindConfigCrossNodeTimeoutParams {
	return &FindConfigCrossNodeTimeoutParams{
		timeout: timeout,
	}
}

// NewFindConfigCrossNodeTimeoutParamsWithContext creates a new FindConfigCrossNodeTimeoutParams object
// with the ability to set a context for a request.
func NewFindConfigCrossNodeTimeoutParamsWithContext(ctx context.Context) *FindConfigCrossNodeTimeoutParams {
	return &FindConfigCrossNodeTimeoutParams{
		Context: ctx,
	}
}

// NewFindConfigCrossNodeTimeoutParamsWithHTTPClient creates a new FindConfigCrossNodeTimeoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigCrossNodeTimeoutParamsWithHTTPClient(client *http.Client) *FindConfigCrossNodeTimeoutParams {
	return &FindConfigCrossNodeTimeoutParams{
		HTTPClient: client,
	}
}

/*
FindConfigCrossNodeTimeoutParams contains all the parameters to send to the API endpoint

	for the find config cross node timeout operation.

	Typically these are written to a http.Request.
*/
type FindConfigCrossNodeTimeoutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config cross node timeout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCrossNodeTimeoutParams) WithDefaults() *FindConfigCrossNodeTimeoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config cross node timeout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCrossNodeTimeoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) WithTimeout(timeout time.Duration) *FindConfigCrossNodeTimeoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) WithContext(ctx context.Context) *FindConfigCrossNodeTimeoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) WithHTTPClient(client *http.Client) *FindConfigCrossNodeTimeoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config cross node timeout params
func (o *FindConfigCrossNodeTimeoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCrossNodeTimeoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
