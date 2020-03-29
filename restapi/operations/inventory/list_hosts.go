// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListHostsHandlerFunc turns a function with the right signature into a list hosts handler
type ListHostsHandlerFunc func(ListHostsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListHostsHandlerFunc) Handle(params ListHostsParams) middleware.Responder {
	return fn(params)
}

// ListHostsHandler interface for that can handle valid list hosts params
type ListHostsHandler interface {
	Handle(ListHostsParams) middleware.Responder
}

// NewListHosts creates a new http.Handler for the list hosts operation
func NewListHosts(ctx *middleware.Context, handler ListHostsHandler) *ListHosts {
	return &ListHosts{Context: ctx, Handler: handler}
}

/*ListHosts swagger:route GET /clusters/{cluster_id}/hosts inventory listHosts

List OpenShift bare metal hosts

*/
type ListHosts struct {
	Context *middleware.Context
	Handler ListHostsHandler
}

func (o *ListHosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListHostsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
