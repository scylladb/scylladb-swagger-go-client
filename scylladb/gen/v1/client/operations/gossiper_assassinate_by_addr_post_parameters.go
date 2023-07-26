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
	"github.com/go-openapi/swag"
)

// NewGossiperAssassinateByAddrPostParams creates a new GossiperAssassinateByAddrPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGossiperAssassinateByAddrPostParams() *GossiperAssassinateByAddrPostParams {
	return &GossiperAssassinateByAddrPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGossiperAssassinateByAddrPostParamsWithTimeout creates a new GossiperAssassinateByAddrPostParams object
// with the ability to set a timeout on a request.
func NewGossiperAssassinateByAddrPostParamsWithTimeout(timeout time.Duration) *GossiperAssassinateByAddrPostParams {
	return &GossiperAssassinateByAddrPostParams{
		timeout: timeout,
	}
}

// NewGossiperAssassinateByAddrPostParamsWithContext creates a new GossiperAssassinateByAddrPostParams object
// with the ability to set a context for a request.
func NewGossiperAssassinateByAddrPostParamsWithContext(ctx context.Context) *GossiperAssassinateByAddrPostParams {
	return &GossiperAssassinateByAddrPostParams{
		Context: ctx,
	}
}

// NewGossiperAssassinateByAddrPostParamsWithHTTPClient creates a new GossiperAssassinateByAddrPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewGossiperAssassinateByAddrPostParamsWithHTTPClient(client *http.Client) *GossiperAssassinateByAddrPostParams {
	return &GossiperAssassinateByAddrPostParams{
		HTTPClient: client,
	}
}

/*
GossiperAssassinateByAddrPostParams contains all the parameters to send to the API endpoint

	for the gossiper assassinate by addr post operation.

	Typically these are written to a http.Request.
*/
type GossiperAssassinateByAddrPostParams struct {

	/* Addr.

	   The endpoint address
	*/
	Addr string

	/* Unsafe.

	   Set to True to perform an unsafe assassination
	*/
	Unsafe *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gossiper assassinate by addr post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GossiperAssassinateByAddrPostParams) WithDefaults() *GossiperAssassinateByAddrPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gossiper assassinate by addr post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GossiperAssassinateByAddrPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) WithTimeout(timeout time.Duration) *GossiperAssassinateByAddrPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) WithContext(ctx context.Context) *GossiperAssassinateByAddrPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) WithHTTPClient(client *http.Client) *GossiperAssassinateByAddrPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) WithAddr(addr string) *GossiperAssassinateByAddrPostParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) SetAddr(addr string) {
	o.Addr = addr
}

// WithUnsafe adds the unsafe to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) WithUnsafe(unsafe *bool) *GossiperAssassinateByAddrPostParams {
	o.SetUnsafe(unsafe)
	return o
}

// SetUnsafe adds the unsafe to the gossiper assassinate by addr post params
func (o *GossiperAssassinateByAddrPostParams) SetUnsafe(unsafe *bool) {
	o.Unsafe = unsafe
}

// WriteToRequest writes these params to a swagger request
func (o *GossiperAssassinateByAddrPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addr
	if err := r.SetPathParam("addr", o.Addr); err != nil {
		return err
	}

	if o.Unsafe != nil {

		// query param unsafe
		var qrUnsafe bool

		if o.Unsafe != nil {
			qrUnsafe = *o.Unsafe
		}
		qUnsafe := swag.FormatBool(qrUnsafe)
		if qUnsafe != "" {

			if err := r.SetQueryParam("unsafe", qUnsafe); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
