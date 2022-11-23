package demo

import "context"

//go:generate gors -service BytesString

// BytesString
// @GORS @Path("/api")  @Path("/BytesString")
type BytesString interface {
	// GetBytesString
	// @GORS @GET @Path("/Get") @HTMLRender
	GetBytesString(context.Context, []byte) (string, error)
	// PutBytesString
	// @GORS @PUT @Path("/Put") @RedirectRender
	PutBytesString(context.Context, []byte) (string, error)
}
