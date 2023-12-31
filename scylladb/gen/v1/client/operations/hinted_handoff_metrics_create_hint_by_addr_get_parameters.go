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

// NewHintedHandoffMetricsCreateHintByAddrGetParams creates a new HintedHandoffMetricsCreateHintByAddrGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHintedHandoffMetricsCreateHintByAddrGetParams() *HintedHandoffMetricsCreateHintByAddrGetParams {
	return &HintedHandoffMetricsCreateHintByAddrGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithTimeout creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the ability to set a timeout on a request.
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithTimeout(timeout time.Duration) *HintedHandoffMetricsCreateHintByAddrGetParams {
	return &HintedHandoffMetricsCreateHintByAddrGetParams{
		timeout: timeout,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithContext creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the ability to set a context for a request.
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithContext(ctx context.Context) *HintedHandoffMetricsCreateHintByAddrGetParams {
	return &HintedHandoffMetricsCreateHintByAddrGetParams{
		Context: ctx,
	}
}

// NewHintedHandoffMetricsCreateHintByAddrGetParamsWithHTTPClient creates a new HintedHandoffMetricsCreateHintByAddrGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewHintedHandoffMetricsCreateHintByAddrGetParamsWithHTTPClient(client *http.Client) *HintedHandoffMetricsCreateHintByAddrGetParams {
	return &HintedHandoffMetricsCreateHintByAddrGetParams{
		HTTPClient: client,
	}
}

/*
HintedHandoffMetricsCreateHintByAddrGetParams contains all the parameters to send to the API endpoint

	for the hinted handoff metrics create hint by addr get operation.

	Typically these are written to a http.Request.
*/
type HintedHandoffMetricsCreateHintByAddrGetParams struct {

	/* Addr.

	   The peer address
	*/
	Addr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hinted handoff metrics create hint by addr get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithDefaults() *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hinted handoff metrics create hint by addr get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithTimeout(timeout time.Duration) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithContext(ctx context.Context) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithHTTPClient(client *http.Client) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WithAddr(addr string) *HintedHandoffMetricsCreateHintByAddrGetParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the hinted handoff metrics create hint by addr get params
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) SetAddr(addr string) {
	o.Addr = addr
}

// WriteToRequest writes these params to a swagger request
func (o *HintedHandoffMetricsCreateHintByAddrGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
