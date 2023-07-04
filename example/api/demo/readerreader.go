package demo

import (
	"context"
	"io"
)

//go:generate gors -service ReaderReader

// ReaderReader
// @GORS @Path(/api/ReaderReader)
type ReaderReader interface {
	// GetReaderReader
	// @GORS @GET @Path(/Get) @ReaderRender
	GetReaderReader(context.Context, io.Reader) (io.Reader, error)
	// HeadReaderReader
	// @GORS @HEAD @Path(/head) @ReaderBinding(video/mp3) @ReaderRender(video/mp4)
	HeadReaderReader(context.Context, io.Reader) (io.Reader, error)
}
