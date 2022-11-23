package gors

import (
	"context"
	"net/http"
)

type headerKey struct{}

func InjectHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, headerKey{}, header)
}

func ExtractHeader(ctx context.Context) http.Header {
	v := ctx.Value(headerKey{})
	header, _ := v.(http.Header)
	return header
}
