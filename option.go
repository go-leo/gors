package gors

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

type Options struct {
	Tag                     string
	DisableDefaultTag       bool
	ResponseWrapper         func(resp any) any
	ErrorHandler            func(ctx context.Context, err error) error
	IncomingHeaderMatcher   func(key string) (string, bool)
	OutgoingHeaderMatcher   func(key string) (string, bool)
	MetadataAnnotators      []func(ctx context.Context) metadata.MD
	ProtoJSONMarshalOptions protojson.MarshalOptions
}

type Option func(o *Options)

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func Tag(tag string) Option {
	return func(o *Options) {
		o.Tag = tag
	}
}

func DisableDefaultTag() Option {
	return func(o *Options) {
		o.DisableDefaultTag = true
	}
}

func ResponseWrapper(w func(resp any) any) Option {
	return func(o *Options) {
		o.ResponseWrapper = w
	}
}

func ErrorHandler(h func(ctx context.Context, err error) error) Option {
	return func(o *Options) {
		o.ErrorHandler = h
	}
}

func IncomingHeaderMatcher(m func(key string) (string, bool)) Option {
	return func(o *Options) {
		o.IncomingHeaderMatcher = m
	}
}

func OutgoingHeaderMatcher(m func(key string) (string, bool)) Option {
	return func(o *Options) {
		o.OutgoingHeaderMatcher = m
	}
}

func MetadataAnnotator(a ...func(ctx context.Context) metadata.MD) Option {
	return func(o *Options) {
		o.MetadataAnnotators = append(o.MetadataAnnotators, a...)
	}
}

func ProtoJSONMarshalOptions(mo protojson.MarshalOptions) Option {
	return func(o *Options) {
		o.ProtoJSONMarshalOptions = mo
	}
}
