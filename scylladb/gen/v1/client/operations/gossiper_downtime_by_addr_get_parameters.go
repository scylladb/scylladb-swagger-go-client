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

// NewGossiperDowntimeByAddrGetParams creates a new GossiperDowntimeByAddrGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGossiperDowntimeByAddrGetParams() *GossiperDowntimeByAddrGetParams {
	return &GossiperDowntimeByAddrGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGossiperDowntimeByAddrGetParamsWithTimeout creates a new GossiperDowntimeByAddrGetParams object
// with the ability to set a timeout on a request.
func NewGossiperDowntimeByAddrGetParamsWithTimeout(timeout time.Duration) *GossiperDowntimeByAddrGetParams {
	return &GossiperDowntimeByAddrGetParams{
		timeout: timeout,
	}
}

// NewGossiperDowntimeByAddrGetParamsWithContext creates a new GossiperDowntimeByAddrGetParams object
// with the ability to set a context for a request.
func NewGossiperDowntimeByAddrGetParamsWithContext(ctx context.Context) *GossiperDowntimeByAddrGetParams {
	return &GossiperDowntimeByAddrGetParams{
		Context: ctx,
	}
}

// NewGossiperDowntimeByAddrGetParamsWithHTTPClient creates a new GossiperDowntimeByAddrGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGossiperDowntimeByAddrGetParamsWithHTTPClient(client *http.Client) *GossiperDowntimeByAddrGetParams {
	return &GossiperDowntimeByAddrGetParams{
		HTTPClient: client,
	}
}

/*
GossiperDowntimeByAddrGetParams contains all the parameters to send to the API endpoint

	for the gossiper downtime by addr get operation.

	Typically these are written to a http.Request.
*/
type GossiperDowntimeByAddrGetParams struct {

	/* Addr.

	   The endpoint address
	*/
	Addr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gossiper downtime by addr get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GossiperDowntimeByAddrGetParams) WithDefaults() *GossiperDowntimeByAddrGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gossiper downtime by addr get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GossiperDowntimeByAddrGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) WithTimeout(timeout time.Duration) *GossiperDowntimeByAddrGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) WithContext(ctx context.Context) *GossiperDowntimeByAddrGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) WithHTTPClient(client *http.Client) *GossiperDowntimeByAddrGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) WithAddr(addr string) *GossiperDowntimeByAddrGetParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the gossiper downtime by addr get params
func (o *GossiperDowntimeByAddrGetParams) SetAddr(addr string) {
	o.Addr = addr
}

// WriteToRequest writes these params to a swagger request
func (o *GossiperDowntimeByAddrGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addr
	if err := r.SetPathParam("addr", o.Addr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
