package demo

import "context"

//go:generate gors -service StringBytes

// StringBytes
// @GORS @Path("/api")  @Path("/StringBytes")
type StringBytes interface {
	// GetStringBytes
	// @GORS @GET @Path("/Get") @BytesRender
	GetStringBytes(context.Context, string) ([]byte, error)
	// OptionsStringBytes
	// @GORS @OPTIONS @Path("/Options") @BytesRender
	OptionsStringBytes(context.Context, string) ([]byte, error)
}
