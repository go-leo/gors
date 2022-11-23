package demo

import (
	"context"
	tmpio "io"
)

//go:generate gors -service StringReader

// StringReader
// @GORS @Path("/api")  @Path("/StringReader")
type StringReader interface {
	// GetStringRender
	// @GORS @GET @Path("/Get") @ReaderRender("video/mpeg4")
	GetStringRender(context.Context, string) (tmpio.Reader, error)
	// OptionsStringReader
	// @GORS @OPTIONS @Path("/Options") @ReaderRender("video/mpeg4")
	OptionsStringReader(context.Context, string) (tmpio.Reader, error)
}
