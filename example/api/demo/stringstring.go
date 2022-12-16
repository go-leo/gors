package demo

import "context"

//go:generate gors -service StringString

// StringString
// @GORS @Path(/api/StringString)
type StringString interface {
	// GetStringString
	// @GORS @GET @Path(/Get) @StringRender(text/go)
	GetStringString(context.Context, string) (string, error)
	// PatchStringString
	// @GORS @PATCH @Path(/Patch) @StringRender(text/go)
	PatchStringString(context.Context, string) (string, error)
}
