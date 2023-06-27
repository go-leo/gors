package gors

import (
	"context"
	"google.golang.org/grpc/metadata"
)

// GRPCMetadata consists of metadata sent from gRPC server.
type GRPCMetadata struct {
	HeaderMD  metadata.MD
	TrailerMD metadata.MD
}

type grpcMetadataKey struct{}

// NewGRPCMetadataContext creates a new context with GRPCMetadata
func NewGRPCMetadataContext(ctx context.Context, md GRPCMetadata) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, grpcMetadataKey{}, md)
}

// GRPCMetadataFromContext returns the GRPCMetadata in ctx
func GRPCMetadataFromContext(ctx context.Context) (md GRPCMetadata, ok bool) {
	if ctx == nil {
		return md, false
	}
	md, ok = ctx.Value(grpcMetadataKey{}).(GRPCMetadata)
	return
}
