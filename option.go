package gors

import (
	"google.golang.org/protobuf/encoding/protojson"
)

type Options interface {
	UnmarshalOptions() protojson.UnmarshalOptions
	MarshalOptions() protojson.MarshalOptions
	ErrorEncoder() ErrorEncoder
	ResponseTransformer() ResponseTransformer
}

type options struct {
	unmarshalOptions    protojson.UnmarshalOptions
	marshalOptions      protojson.MarshalOptions
	errorEncoder        ErrorEncoder
	responseTransformer ResponseTransformer
}

type Option func(o *options)

func (o *options) Apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

func (o *options) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

func (o *options) ErrorEncoder() ErrorEncoder {
	return o.errorEncoder
}

func (o *options) ResponseTransformer() ResponseTransformer {
	return o.responseTransformer
}

func WithUnmarshalOptions(opts protojson.UnmarshalOptions) Option {
	return func(o *options) {
		o.unmarshalOptions = opts
	}
}

func WithMarshalOptions(opts protojson.MarshalOptions) Option {
	return func(o *options) {
		o.marshalOptions = opts
	}
}

func WithErrorEncoder(encoder ErrorEncoder) Option {
	return func(o *options) {
		o.errorEncoder = encoder
	}
}

func WithResponseTransformer(transformer ResponseTransformer) Option {
	return func(o *options) {
		o.responseTransformer = transformer
	}
}

func NewOptions(opts ...Option) Options {
	o := &options{
		unmarshalOptions:    protojson.UnmarshalOptions{},
		marshalOptions:      protojson.MarshalOptions{},
		errorEncoder:        DefaultErrorEncoder,
		responseTransformer: DefaultResponseTransformer,
	}
	o = o.Apply(opts...)
	return o
}
