package demo

import "context"

//go:generate gors -service Service

// Service
// @GORS @Path("/api")  @Path("/v1")
type Service interface {
	// Method
	// @GORS @GET @Path("/method/:id") @UriBinding @JSONRender
	Method(context.Context, *MethodReq) (*MethodResp, error)
}

type MethodReq struct {
	ID int `uri:"id"`
}

type MethodResp struct {
	V int `json:"v,omitempty"`
}
