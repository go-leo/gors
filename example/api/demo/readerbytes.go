package demo

import (
	"context"
	tmpio "io"
)

//go:generate gors -service ReaderBytes

// ReaderBytes
// @GORS @Path("/api")  @Path("/ReaderBytes")
type ReaderBytes interface {
	// GetReaderBytes
	// @GORS @GET @Path("/Get") @BytesRender
	GetReaderBytes(context.Context, tmpio.Reader) ([]byte, error)
	// PostReaderBytes
	// @GORS @POST @Path("/Post") @BytesRender("text/go")
	PostReaderBytes(context.Context, tmpio.Reader) ([]byte, error)
}
