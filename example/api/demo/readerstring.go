package demo

import (
	"context"
	tmpio "io"
)

//go:generate gors -service ReaderString

// ReaderString
// @GORS @Path("/api")  @Path("/ReaderString")
type ReaderString interface {
	// GetReaderString
	// @GORS @GET @Path("/Get") @TextRender
	GetReaderString(context.Context, tmpio.Reader) (string, error)
	// PostReaderString
	// @GORS @POST @Path("/Post") @StringRender("text/go")
	PostReaderString(context.Context, tmpio.Reader) (string, error)
}
