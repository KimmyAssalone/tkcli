// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

	"github.com/tkhq/tkcli/api/models"
)

// NewPublicAPIServiceDeleteInvitationParams creates a new PublicAPIServiceDeleteInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublicAPIServiceDeleteInvitationParams() *PublicAPIServiceDeleteInvitationParams {
	return &PublicAPIServiceDeleteInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublicAPIServiceDeleteInvitationParamsWithTimeout creates a new PublicAPIServiceDeleteInvitationParams object
// with the ability to set a timeout on a request.
func NewPublicAPIServiceDeleteInvitationParamsWithTimeout(timeout time.Duration) *PublicAPIServiceDeleteInvitationParams {
	return &PublicAPIServiceDeleteInvitationParams{
		timeout: timeout,
	}
}

// NewPublicAPIServiceDeleteInvitationParamsWithContext creates a new PublicAPIServiceDeleteInvitationParams object
// with the ability to set a context for a request.
func NewPublicAPIServiceDeleteInvitationParamsWithContext(ctx context.Context) *PublicAPIServiceDeleteInvitationParams {
	return &PublicAPIServiceDeleteInvitationParams{
		Context: ctx,
	}
}

// NewPublicAPIServiceDeleteInvitationParamsWithHTTPClient creates a new PublicAPIServiceDeleteInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublicAPIServiceDeleteInvitationParamsWithHTTPClient(client *http.Client) *PublicAPIServiceDeleteInvitationParams {
	return &PublicAPIServiceDeleteInvitationParams{
		HTTPClient: client,
	}
}

/*
PublicAPIServiceDeleteInvitationParams contains all the parameters to send to the API endpoint

	for the public Api service delete invitation operation.

	Typically these are written to a http.Request.
*/
type PublicAPIServiceDeleteInvitationParams struct {

	// Body.
	Body *models.V1DeleteInvitationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the public Api service delete invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceDeleteInvitationParams) WithDefaults() *PublicAPIServiceDeleteInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the public Api service delete invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublicAPIServiceDeleteInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) WithTimeout(timeout time.Duration) *PublicAPIServiceDeleteInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) WithContext(ctx context.Context) *PublicAPIServiceDeleteInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) WithHTTPClient(client *http.Client) *PublicAPIServiceDeleteInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) WithBody(body *models.V1DeleteInvitationRequest) *PublicAPIServiceDeleteInvitationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public Api service delete invitation params
func (o *PublicAPIServiceDeleteInvitationParams) SetBody(body *models.V1DeleteInvitationRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PublicAPIServiceDeleteInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
