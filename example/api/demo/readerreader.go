package demo

import (
	"context"
	tmpio "io"
)

//go:generate gors -service ReaderReader

// ReaderReader
// @GORS @Path("/api")  @Path("/ReaderReader")
type ReaderReader interface {
	// GetReaderReader
	// @GORS @GET @Path("/Get") @ReaderRender
	GetReaderReader(context.Context, tmpio.Reader) (tmpio.Reader, error)
	// HeadReaderReader
	// @GORS @HEAD @Path("/head") @ReaderRender
	HeadReaderReader(context.Context, tmpio.Reader) (tmpio.Reader, error)
}
