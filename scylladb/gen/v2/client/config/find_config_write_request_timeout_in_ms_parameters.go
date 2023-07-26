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

// NewFindConfigWriteRequestTimeoutInMsParams creates a new FindConfigWriteRequestTimeoutInMsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigWriteRequestTimeoutInMsParams() *FindConfigWriteRequestTimeoutInMsParams {
	return &FindConfigWriteRequestTimeoutInMsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigWriteRequestTimeoutInMsParamsWithTimeout creates a new FindConfigWriteRequestTimeoutInMsParams object
// with the ability to set a timeout on a request.
func NewFindConfigWriteRequestTimeoutInMsParamsWithTimeout(timeout time.Duration) *FindConfigWriteRequestTimeoutInMsParams {
	return &FindConfigWriteRequestTimeoutInMsParams{
		timeout: timeout,
	}
}

// NewFindConfigWriteRequestTimeoutInMsParamsWithContext creates a new FindConfigWriteRequestTimeoutInMsParams object
// with the ability to set a context for a request.
func NewFindConfigWriteRequestTimeoutInMsParamsWithContext(ctx context.Context) *FindConfigWriteRequestTimeoutInMsParams {
	return &FindConfigWriteRequestTimeoutInMsParams{
		Context: ctx,
	}
}

// NewFindConfigWriteRequestTimeoutInMsParamsWithHTTPClient creates a new FindConfigWriteRequestTimeoutInMsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigWriteRequestTimeoutInMsParamsWithHTTPClient(client *http.Client) *FindConfigWriteRequestTimeoutInMsParams {
	return &FindConfigWriteRequestTimeoutInMsParams{
		HTTPClient: client,
	}
}

/*
FindConfigWriteRequestTimeoutInMsParams contains all the parameters to send to the API endpoint

	for the find config write request timeout in ms operation.

	Typically these are written to a http.Request.
*/
type FindConfigWriteRequestTimeoutInMsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config write request timeout in ms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigWriteRequestTimeoutInMsParams) WithDefaults() *FindConfigWriteRequestTimeoutInMsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config write request timeout in ms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigWriteRequestTimeoutInMsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) WithTimeout(timeout time.Duration) *FindConfigWriteRequestTimeoutInMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) WithContext(ctx context.Context) *FindConfigWriteRequestTimeoutInMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) WithHTTPClient(client *http.Client) *FindConfigWriteRequestTimeoutInMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config write request timeout in ms params
func (o *FindConfigWriteRequestTimeoutInMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigWriteRequestTimeoutInMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
