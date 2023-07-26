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

// NewFindConfigCommitlogDirectoryParams creates a new FindConfigCommitlogDirectoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigCommitlogDirectoryParams() *FindConfigCommitlogDirectoryParams {
	return &FindConfigCommitlogDirectoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogDirectoryParamsWithTimeout creates a new FindConfigCommitlogDirectoryParams object
// with the ability to set a timeout on a request.
func NewFindConfigCommitlogDirectoryParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogDirectoryParams {
	return &FindConfigCommitlogDirectoryParams{
		timeout: timeout,
	}
}

// NewFindConfigCommitlogDirectoryParamsWithContext creates a new FindConfigCommitlogDirectoryParams object
// with the ability to set a context for a request.
func NewFindConfigCommitlogDirectoryParamsWithContext(ctx context.Context) *FindConfigCommitlogDirectoryParams {
	return &FindConfigCommitlogDirectoryParams{
		Context: ctx,
	}
}

// NewFindConfigCommitlogDirectoryParamsWithHTTPClient creates a new FindConfigCommitlogDirectoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigCommitlogDirectoryParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogDirectoryParams {
	return &FindConfigCommitlogDirectoryParams{
		HTTPClient: client,
	}
}

/*
FindConfigCommitlogDirectoryParams contains all the parameters to send to the API endpoint

	for the find config commitlog directory operation.

	Typically these are written to a http.Request.
*/
type FindConfigCommitlogDirectoryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config commitlog directory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCommitlogDirectoryParams) WithDefaults() *FindConfigCommitlogDirectoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config commitlog directory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigCommitlogDirectoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) WithContext(ctx context.Context) *FindConfigCommitlogDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog directory params
func (o *FindConfigCommitlogDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
