// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewGetHostParams creates a new GetHostParams object
// with the default values initialized.
func NewGetHostParams() *GetHostParams {
	var ()
	return &GetHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostParamsWithTimeout creates a new GetHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHostParamsWithTimeout(timeout time.Duration) *GetHostParams {
	var ()
	return &GetHostParams{

		timeout: timeout,
	}
}

// NewGetHostParamsWithContext creates a new GetHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHostParamsWithContext(ctx context.Context) *GetHostParams {
	var ()
	return &GetHostParams{

		Context: ctx,
	}
}

// NewGetHostParamsWithHTTPClient creates a new GetHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHostParamsWithHTTPClient(client *http.Client) *GetHostParams {
	var ()
	return &GetHostParams{
		HTTPClient: client,
	}
}

/*GetHostParams contains all the parameters to send to the API endpoint
for the get host operation typically these are written to a http.Request
*/
type GetHostParams struct {

	/*ClusterID
	  The ID of the cluster to get hosts from

	*/
	ClusterID strfmt.UUID
	/*HostID
	  The ID of the host to retrieve

	*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get host params
func (o *GetHostParams) WithTimeout(timeout time.Duration) *GetHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host params
func (o *GetHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host params
func (o *GetHostParams) WithContext(ctx context.Context) *GetHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host params
func (o *GetHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host params
func (o *GetHostParams) WithHTTPClient(client *http.Client) *GetHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host params
func (o *GetHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get host params
func (o *GetHostParams) WithClusterID(clusterID strfmt.UUID) *GetHostParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get host params
func (o *GetHostParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the get host params
func (o *GetHostParams) WithHostID(hostID strfmt.UUID) *GetHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the get host params
func (o *GetHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
