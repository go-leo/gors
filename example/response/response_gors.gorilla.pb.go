// Code generated by protoc-gen-gors-gorilla. DO NOT EDIT.

package response

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

type ResponseGorillaService interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
	HttpResponse(ctx context.Context, request *emptypb.Empty) (*http.HttpResponse, error)
}

func AppendResponseGorillaRoute(router *mux.Router, service ResponseGorillaService, opts ...v2.Option) *mux.Router {
	options := v2.NewOptions(opts...)
	handler := responseGorillaHandler{
		service: service,
		decoder: responseGorillaRequestDecoder{
			unmarshalOptions: options.UnmarshalOptions(),
		},
		encoder: responseGorillaResponseEncoder{
			marshalOptions:      options.MarshalOptions(),
			unmarshalOptions:    options.UnmarshalOptions(),
			responseTransformer: options.ResponseTransformer(),
		},
		errorEncoder: v2.DefaultErrorEncoder,
	}
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/OmittedResponse").
		Methods("POST").
		Path("/v1/omitted/response").
		Handler(handler.OmittedResponse())
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/StarResponse").
		Methods("POST").
		Path("/v1/star/response").
		Handler(handler.StarResponse())
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/NamedResponse").
		Methods("POST").
		Path("/v1/named/response").
		Handler(handler.NamedResponse())
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/HttpBodyResponse").
		Methods("PUT").
		Path("/v1/http/body/omitted/response").
		Handler(handler.HttpBodyResponse())
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/HttpBodyNamedResponse").
		Methods("PUT").
		Path("/v1/http/body/named/response").
		Handler(handler.HttpBodyNamedResponse())
	router.NewRoute().
		Name("/leo.gors.example.response.v1.Response/HttpResponse").
		Methods("GET").
		Path("/v1/http/response").
		Handler(handler.HttpResponse())
	return router
}

type responseGorillaHandler struct {
	service      ResponseGorillaService
	decoder      responseGorillaRequestDecoder
	encoder      responseGorillaResponseEncoder
	errorEncoder v2.ErrorEncoder
}

func (h responseGorillaHandler) OmittedResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.OmittedResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.OmittedResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.OmittedResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h responseGorillaHandler) StarResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.StarResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.StarResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.StarResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h responseGorillaHandler) NamedResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.NamedResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.NamedResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.NamedResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h responseGorillaHandler) HttpBodyResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpBodyResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpBodyResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpBodyResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h responseGorillaHandler) HttpBodyNamedResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpBodyNamedResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpBodyNamedResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpBodyNamedResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

func (h responseGorillaHandler) HttpResponse() http1.Handler {
	return http1.HandlerFunc(func(writer http1.ResponseWriter, request *http1.Request) {
		ctx := request.Context()
		in, err := h.decoder.HttpResponse(ctx, request)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		out, err := h.service.HttpResponse(ctx, in)
		if err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
		if err := h.encoder.HttpResponse(ctx, writer, out); err != nil {
			h.errorEncoder(ctx, err, writer)
			return
		}
	})
}

type responseGorillaRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder responseGorillaRequestDecoder) OmittedResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder responseGorillaRequestDecoder) StarResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder responseGorillaRequestDecoder) NamedResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder responseGorillaRequestDecoder) HttpBodyResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder responseGorillaRequestDecoder) HttpBodyNamedResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}
func (decoder responseGorillaRequestDecoder) HttpResponse(ctx context.Context, r *http1.Request) (*emptypb.Empty, error) {
	req := &emptypb.Empty{}
	return req, nil
}

type responseGorillaResponseEncoder struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer v2.ResponseTransformer
}

func (encoder responseGorillaResponseEncoder) OmittedResponse(ctx context.Context, w http1.ResponseWriter, resp *UserResponse) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder responseGorillaResponseEncoder) StarResponse(ctx context.Context, w http1.ResponseWriter, resp *UserResponse) error {
	return v2.ResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
}
func (encoder responseGorillaResponseEncoder) NamedResponse(ctx context.Context, w http1.ResponseWriter, resp *UserResponse) error {
	return v2.ResponseEncoder(ctx, w, resp.GetUser(), encoder.marshalOptions)
}
func (encoder responseGorillaResponseEncoder) HttpBodyResponse(ctx context.Context, w http1.ResponseWriter, resp *httpbody.HttpBody) error {
	return v2.HttpBodyEncoder(ctx, w, encoder.responseTransformer(ctx, resp))
}
func (encoder responseGorillaResponseEncoder) HttpBodyNamedResponse(ctx context.Context, w http1.ResponseWriter, resp *HttpBody) error {
	return v2.HttpBodyEncoder(ctx, w, resp.GetBody())
}
func (encoder responseGorillaResponseEncoder) HttpResponse(ctx context.Context, w http1.ResponseWriter, resp *http.HttpResponse) error {
	return v2.HttpResponseEncoder(ctx, w, encoder.responseTransformer(ctx, resp))
}
