// Code generated by protoc-gen-gors-gorilla. DO NOT EDIT.

package body

import (
	context "context"
	v2 "github.com/go-leo/gors/v2"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http1 "net/http"
)

type BodyGorillaService interface {
	StarBody(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBodyRequest) (*emptypb.Empty, error)
	HttpRequest(ctx context.Context, request *http.HttpRequest) (*emptypb.Empty, error)
}

func AppendBodyGorillaRoute(router *mux.Router, service BodyGorillaService, opts ...v2.Option) *mux.Router {
	options := v2.NewOptions(opts...)
	handler := bodyGorillaHandler{
		service: service,
		decoder: bodyGorillaRequestDecoder{
			unmarshalOptions: options.UnmarshalOptions(),
		},
		encoder: bodyGorillaResponseEncoder{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: v2.DefaultErrorEncoder,
	}
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/StarBody").
		Methods("POST").
		Path("/v1/star/body").
		Handler(handler.StarBody())
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/NamedBody").
		Methods("POST").
		Path("/v1/named/body").
		Handler(handler.NamedBody())
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/NonBody").
		Methods("GET").
		Path("/v1/user_body").
		Handler(handler.NonBody())
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/HttpBodyStarBody").
		Methods("PUT").
		Path("/v1/http/body/star/body").
		Handler(handler.HttpBodyStarBody())
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/HttpBodyNamedBody").
		Methods("PUT").
		Path("/v1/http/body/named/body").
		Handler(handler.HttpBodyNamedBody())
	router.NewRoute().
		Name("/leo.gors.body.v1.Body/HttpRequest").
		Methods("PUT").
		Path("/v1/http/request").
		Handler(handler.HttpRequest())
	return router
}

type bodyGorillaHandler struct {
	service      BodyGorillaService
	decoder      bodyGorillaRequestDecoder
	encoder      bodyGorillaResponseEncoder
	errorEncoder v2.ErrorEncoder
}

func (h bodyGorillaHandler) StarBody() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.StarBody(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.StarBody(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.StarBody(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h bodyGorillaHandler) NamedBody() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.NamedBody(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.NamedBody(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.NamedBody(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h bodyGorillaHandler) NonBody() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.NonBody(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.NonBody(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.NonBody(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h bodyGorillaHandler) HttpBodyStarBody() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpBodyStarBody(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpBodyStarBody(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpBodyStarBody(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h bodyGorillaHandler) HttpBodyNamedBody() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpBodyNamedBody(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpBodyNamedBody(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpBodyNamedBody(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h bodyGorillaHandler) HttpRequest() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpRequest(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpRequest(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpRequest(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type bodyGorillaRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder bodyGorillaRequestDecoder) StarBody(ctx context.Context, r *http1.Request) (*BodyRequest, error) {
	req := &BodyRequest{}
	if err := v2.RequestDecoder(ctx, r, req, decoder.unmarshalOptions); err != nil {
		return nil, err
	}
	return req, nil
}
func (decoder bodyGorillaRequestDecoder) NamedBody(ctx context.Context, r *http1.Request) (*BodyRequest, error) {
	req := &BodyRequest{}
	if req.User == nil {
		req.User = &BodyRequest_User{}
	}
	if err := v2.RequestDecoder(ctx, r, req.User, decoder.unmarshalOptions); err != nil {
		return nil, err
	}
	return req, nil
}
func (decoder bodyGorillaRequestDecoder) NonBody(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder bodyGorillaRequestDecoder) HttpBodyStarBody(ctx context.Context, r *http1.Request) (*httpbody.HttpBody, error) {
	req := &httpbody.HttpBody{}
	if err := v2.HttpBodyDecoder(ctx, r, req); err != nil {
		return nil, err
	}
	return req, nil
}
func (decoder bodyGorillaRequestDecoder) HttpBodyNamedBody(ctx context.Context, r *http1.Request) (*HttpBodyRequest, error) {
	req := &HttpBodyRequest{}
	if req.Body == nil {
		req.Body = &httpbody.HttpBody{}
	}
	if err := v2.HttpBodyDecoder(ctx, r, req.Body); err != nil {
		return nil, err
	}
	return req, nil
}
func (decoder bodyGorillaRequestDecoder) HttpRequest(ctx context.Context, r *http1.Request) (*http.HttpRequest, error) {
	req := &http.HttpRequest{}
	if err := v2.HttpRequestDecoder(ctx, r, req); err != nil {
		return nil, err
	}
	return req, nil
}

type bodyGorillaResponseEncoder struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer v2.ResponseTransformer
}

func (encoder bodyGorillaResponseEncoder) StarBody(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder bodyGorillaResponseEncoder) NamedBody(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder bodyGorillaResponseEncoder) NonBody(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder bodyGorillaResponseEncoder) HttpBodyStarBody(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder bodyGorillaResponseEncoder) HttpBodyNamedBody(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder bodyGorillaResponseEncoder) HttpRequest(ctx context.Context, w http1.ResponseWriter, resp *emptypb.Empty) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
