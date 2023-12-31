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

// NewStorageServiceMovePostParams creates a new StorageServiceMovePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageServiceMovePostParams() *StorageServiceMovePostParams {
	return &StorageServiceMovePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceMovePostParamsWithTimeout creates a new StorageServiceMovePostParams object
// with the ability to set a timeout on a request.
func NewStorageServiceMovePostParamsWithTimeout(timeout time.Duration) *StorageServiceMovePostParams {
	return &StorageServiceMovePostParams{
		timeout: timeout,
	}
}

// NewStorageServiceMovePostParamsWithContext creates a new StorageServiceMovePostParams object
// with the ability to set a context for a request.
func NewStorageServiceMovePostParamsWithContext(ctx context.Context) *StorageServiceMovePostParams {
	return &StorageServiceMovePostParams{
		Context: ctx,
	}
}

// NewStorageServiceMovePostParamsWithHTTPClient creates a new StorageServiceMovePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageServiceMovePostParamsWithHTTPClient(client *http.Client) *StorageServiceMovePostParams {
	return &StorageServiceMovePostParams{
		HTTPClient: client,
	}
}

/*
StorageServiceMovePostParams contains all the parameters to send to the API endpoint

	for the storage service move post operation.

	Typically these are written to a http.Request.
*/
type StorageServiceMovePostParams struct {

	/* NewToken.

	   token to move this node to
	*/
	NewToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage service move post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceMovePostParams) WithDefaults() *StorageServiceMovePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage service move post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageServiceMovePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage service move post params
func (o *StorageServiceMovePostParams) WithTimeout(timeout time.Duration) *StorageServiceMovePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service move post params
func (o *StorageServiceMovePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service move post params
func (o *StorageServiceMovePostParams) WithContext(ctx context.Context) *StorageServiceMovePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service move post params
func (o *StorageServiceMovePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service move post params
func (o *StorageServiceMovePostParams) WithHTTPClient(client *http.Client) *StorageServiceMovePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service move post params
func (o *StorageServiceMovePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewToken adds the newToken to the storage service move post params
func (o *StorageServiceMovePostParams) WithNewToken(newToken string) *StorageServiceMovePostParams {
	o.SetNewToken(newToken)
	return o
}

// SetNewToken adds the newToken to the storage service move post params
func (o *StorageServiceMovePostParams) SetNewToken(newToken string) {
	o.NewToken = newToken
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceMovePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param new_token
	qrNewToken := o.NewToken
	qNewToken := qrNewToken
	if qNewToken != "" {

		if err := r.SetQueryParam("new_token", qNewToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
