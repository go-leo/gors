package demo

import (
	"context"
	tmpio "io"
)

//go:generate gors -service BytesReader

// BytesReader
// @GORS @Path(/api/BytesReader)
type BytesReader interface {
	// GetBytesReader
	// @GORS @GET @Path(/Get) @ReaderRender(video/mpeg4)
	GetBytesReader(context.Context, []byte) (tmpio.Reader, error)
	// PatchBytesReader
	// @GORS @PATCH @Path(/Patch) @ReaderRender(video/mpeg4)
	PatchBytesReader(context.Context, []byte) (tmpio.Reader, error)
}
