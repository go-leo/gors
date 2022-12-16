package demo

import (
	"context"
	"io"
)

//go:generate gors -service ReaderString

// ReaderString
// @GORS @Path(/api/ReaderString)
type ReaderString interface {
	// GetReaderString
	// @GORS @GET @Path(/Get) @TextRender
	GetReaderString(context.Context, io.Reader) (string, error)
	// PostReaderString
	// @GORS @POST @Path(/Post) @StringRender(text/go)
	PostReaderString(context.Context, io.Reader) (string, error)
}
