package demo

import "context"

//go:generate gors -service BytesBytes

// BytesBytes
// @GORS @Path(/api/BytesBytes)
type BytesBytes interface {
	// GetBytesBytes
	// @GORS @GET @Path(/Get) @BytesRender(ttt.sss)
	GetBytesBytes(context.Context, []byte) ([]byte, error)
	// PostBytesBytes
	// @GORS @POST @Path(/Post) @BytesBinding(image/jpeg) @BytesRender(text/go)
	PostBytesBytes(context.Context, []byte) ([]byte, error)
}
