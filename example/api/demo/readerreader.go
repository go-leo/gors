package demo

import (
	"context"
	"github.com/go-leo/gors"
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
	// GetJSONReader
	// @GORS @Get @Path(/get/json) @JSONBinding @ReaderRender(video/mp4)
	GetJSONReader(context.Context, *gors.Empty) (io.Reader, error)
}
