package demo

import "context"

//go:generate gors -service BytesBytes

// BytesBytes
// @GORS @Path("/api")  @Path("/BytesBytes")
type BytesBytes interface {
	// GetBytesBytes
	// @GORS @GET @Path("/Get") @BytesRender("text/go")
	GetBytesBytes(context.Context, []byte) ([]byte, error)
	// PostBytesBytes
	// @GORS @POST @Path("/Post") @BytesRender("text/go")
	PostBytesBytes(context.Context, []byte) ([]byte, error)
}
