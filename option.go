package gors

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type Options struct {
	tag                   string
	disableDefaultTag     bool
	responseWrapper       func(resp any) any
	errorHandler          func(ctx context.Context, err error) error
	incomingHeaderMatcher func(key string) (string, bool)
	outgoingHeaderMatcher func(key string) (string, bool)
	metadataAnnotators    []func(ctx context.Context) metadata.MD
}

func (o *Options) Tag() string {
	return o.tag
}

func (o *Options) DefaultTag(tag string) {
	o.tag = tag
}

func (o *Options) DisableDefaultTag() bool {
	return o.disableDefaultTag
}

func (o *Options) ResponseWrapper() func(resp any) any {
	return o.responseWrapper
}

func (o *Options) ErrorHandler() func(ctx context.Context, err error) error {
	return o.errorHandler
}

func (o *Options) IncomingHeaderMatcher() func(key string) (string, bool) {
	return o.incomingHeaderMatcher
}

func (o *Options) OutgoingHeaderMatcher() func(key string) (string, bool) {
	return o.outgoingHeaderMatcher
}

func (o *Options) MetadataAnnotators() []func(ctx context.Context) metadata.MD {
	return o.metadataAnnotators
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
		o.tag = tag
	}
}

func DisableDefaultTag() Option {
	return func(o *Options) {
		o.disableDefaultTag = true
	}
}

func ResponseWrapper(w func(resp any) any) Option {
	return func(o *Options) {
		o.responseWrapper = w
	}
}

func ErrorHandler(h func(ctx context.Context, err error) error) Option {
	return func(o *Options) {
		o.errorHandler = h
	}
}

func IncomingHeaderMatcher(m func(key string) (string, bool)) Option {
	return func(o *Options) {
		o.incomingHeaderMatcher = m
	}
}

func OutgoingHeaderMatcher(m func(key string) (string, bool)) Option {
	return func(o *Options) {
		o.outgoingHeaderMatcher = m
	}
}

func MetadataAnnotator(a ...func(ctx context.Context) metadata.MD) Option {
	return func(o *Options) {
		o.metadataAnnotators = append(o.metadataAnnotators, a...)
	}
}
