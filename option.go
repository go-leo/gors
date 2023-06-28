package gors

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type Options struct {
	Tag                string
	ResponseWrapper    func(resp any) any
	ErrorHandler       func(ctx context.Context, err error)
	HeaderMatcher      func(key string) (string, bool)
	MetadataAnnotators []func(ctx context.Context) metadata.MD
}

type Option func(o *Options)

func New(opts ...Option) *Options {
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

func ResponseWrapper(w func(resp any) any) Option {
	return func(o *Options) {
		o.ResponseWrapper = w
	}
}

func ErrorHandler(h func(ctx context.Context, err error)) Option {
	return func(o *Options) {
		o.ErrorHandler = h
	}
}

func HeaderMatcher(m func(key string) (string, bool)) Option {
	return func(o *Options) {
		o.HeaderMatcher = m
	}
}

func MetadataAnnotator(a ...func(ctx context.Context) metadata.MD) Option {
	return func(o *Options) {
		o.MetadataAnnotators = append(o.MetadataAnnotators, a...)
	}
}
