package gors

import "context"

type Options struct {
	Tag             string
	ResponseWrapper func(resp any) any
	ErrorHandler    func(ctx context.Context, err error)
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
