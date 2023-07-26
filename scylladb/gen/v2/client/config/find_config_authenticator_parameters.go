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

// NewFindConfigAuthenticatorParams creates a new FindConfigAuthenticatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigAuthenticatorParams() *FindConfigAuthenticatorParams {
	return &FindConfigAuthenticatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAuthenticatorParamsWithTimeout creates a new FindConfigAuthenticatorParams object
// with the ability to set a timeout on a request.
func NewFindConfigAuthenticatorParamsWithTimeout(timeout time.Duration) *FindConfigAuthenticatorParams {
	return &FindConfigAuthenticatorParams{
		timeout: timeout,
	}
}

// NewFindConfigAuthenticatorParamsWithContext creates a new FindConfigAuthenticatorParams object
// with the ability to set a context for a request.
func NewFindConfigAuthenticatorParamsWithContext(ctx context.Context) *FindConfigAuthenticatorParams {
	return &FindConfigAuthenticatorParams{
		Context: ctx,
	}
}

// NewFindConfigAuthenticatorParamsWithHTTPClient creates a new FindConfigAuthenticatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigAuthenticatorParamsWithHTTPClient(client *http.Client) *FindConfigAuthenticatorParams {
	return &FindConfigAuthenticatorParams{
		HTTPClient: client,
	}
}

/*
FindConfigAuthenticatorParams contains all the parameters to send to the API endpoint

	for the find config authenticator operation.

	Typically these are written to a http.Request.
*/
type FindConfigAuthenticatorParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config authenticator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigAuthenticatorParams) WithDefaults() *FindConfigAuthenticatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config authenticator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigAuthenticatorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config authenticator params
func (o *FindConfigAuthenticatorParams) WithTimeout(timeout time.Duration) *FindConfigAuthenticatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config authenticator params
func (o *FindConfigAuthenticatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config authenticator params
func (o *FindConfigAuthenticatorParams) WithContext(ctx context.Context) *FindConfigAuthenticatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config authenticator params
func (o *FindConfigAuthenticatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config authenticator params
func (o *FindConfigAuthenticatorParams) WithHTTPClient(client *http.Client) *FindConfigAuthenticatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config authenticator params
func (o *FindConfigAuthenticatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAuthenticatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
