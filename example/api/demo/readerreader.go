package demo

import (
	"context"
	"io"
)

//go:generate gors -service ReaderReader

// ReaderReader
// @GORS @Path("/api")  @Path("/ReaderReader")
type ReaderReader interface {
	// GetReaderReader
	// @GORS @GET @Path("/Get") @ReaderRender
	GetReaderReader(context.Context, io.Reader) (io.Reader, error)
	// HeadReaderReader
	// @GORS @HEAD @Path("/head") @ReaderRender
	HeadReaderReader(context.Context, io.Reader) (io.Reader, error)
}
