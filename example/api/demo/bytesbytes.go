package demo

import "context"

//go:generate gors -service BytesBytes

// BytesBytes this is bytes binding and bytes render demo
// this is a interface
// @GORS @Path(/api/BytesBytes)
type BytesBytes interface {
	// GetBytesBytes get http method, receive bytes request and send bytes response
	// @GORS @GET @Path(/Get) @BytesRender(ttt.sss)
	GetBytesBytes(context.Context, []byte) ([]byte, error)
	// PostBytesBytes post http method,  receive bytes request and send bytes response
	// @GORS @POST @Path(/Post) @BytesBinding @BytesRender(text/go)
	PostBytesBytes(context.Context, []byte) ([]byte, error)
}
