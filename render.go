package gors

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/render"
	internalrender "github.com/go-leo/gors/internal/pkg/render"
	"github.com/go-leo/gox/iox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	grpcstatus "google.golang.org/grpc/status"
	"io"
	"net/http"
	"strings"
)

func ErrorRender(
	ctx context.Context,
	err error,
	handler func(ctx context.Context, err error),
) {
	if handler != nil {
		handler(ctx, err)
		return
	}
	var e Error
	if errors.As(err, &e) {
		ResponseRender(ctx, e.StatusCode, e.Status(), "", PureJSONRender, nil)
		return
	}
	var ePtr *Error
	if errors.As(err, &ePtr) {
		ResponseRender(ctx, ePtr.StatusCode, ePtr.Status(), "", PureJSONRender, nil)
		return
	}
	status, ok := grpcstatus.FromError(err)
	if ok {
		ResponseRender(ctx, httpStatusFromCode(status.Code()), status.Proto(), "", ProtoJSONRender, nil)
		return
	}
	ResponseRender(ctx, http.StatusInternalServerError, err.Error(), "", TextRender, nil)
}

func ResponseRender(
	ctx context.Context,
	code int,
	resp any,
	contentType string,
	render func(ctx context.Context, code int, resp any, contentType string),
	wrapper func(resp any) any,
) {
	addHeaders(ctx)
	addTrailersHeader(ctx)
	// wrapper response
	if wrapper != nil {
		resp = wrapper(resp)
	}
	// render
	render(ctx, code, resp, contentType)
	addTrailers(ctx)
}

func addHeaders(ctx context.Context) {
	c := FromContext(ctx)
	header := Header(ctx)
	for key, values := range header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}
}

func addTrailersHeader(ctx context.Context) {
	c := FromContext(ctx)
	doForwardTrailers := strings.Contains(strings.ToLower(c.GetHeader("TE")), "trailers")
	if !doForwardTrailers {
		return
	}
	trailer := Trailer(ctx)
	if len(trailer) <= 0 {
		return
	}
	for k := range trailer {
		c.Writer.Header().Add("Trailer", k)
	}
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
}

func addTrailers(ctx context.Context) {
	c := FromContext(ctx)
	doForwardTrailers := strings.Contains(strings.ToLower(c.GetHeader("TE")), "trailers")
	if !doForwardTrailers {
		return
	}
	trailer := Trailer(ctx)
	if len(trailer) <= 0 {
		return
	}
	for key, values := range trailer {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}
}

func BytesRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: resp.([]byte)})
}

func StringRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: []byte(resp.(string))})
}

func TextRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).String(code, resp.(string))
}

func HTMLRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).Render(code, render.Data{ContentType: "text/html; charset=utf-8", Data: []byte(resp.(string))})
}

func RedirectRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).Redirect(code, resp.(string))
}

func ReaderRender(ctx context.Context, code int, resp any, contentType string) {
	r := resp.(io.Reader)
	l, ok := iox.Len(r)
	if !ok {
		l = -1
	}
	FromContext(ctx).Render(code, render.Reader{ContentType: contentType, ContentLength: int64(l), Reader: r})
}

func JSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).JSON(code, resp)
}

func IndentedJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).IndentedJSON(code, resp)
}

func SecureJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).SecureJSON(code, resp)
}

func JSONPJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).JSONP(code, resp)
}

func PureJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).PureJSON(code, resp)
}

func AsciiJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).AsciiJSON(code, resp)
}

func ProtoJSONRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).Render(code, internalrender.ProtoJSON{Data: resp})
}

func XMLRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).XML(code, resp)
}

func YAMLRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).YAML(code, resp)
}

func ProtoBufRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).ProtoBuf(code, resp)
}

func MsgPackRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).Render(code, render.MsgPack{Data: resp})
}

func TOMLRender(ctx context.Context, code int, resp any, _ string) {
	FromContext(ctx).TOML(code, resp)
}

func CustomRender(ctx context.Context, code int, resp any, _ string) {
	customRender, ok := resp.(Render)
	if !ok {
		return
	}
	customRender.Render(ctx)
}

// httpStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func httpStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		grpclog.Infof("Unknown gRPC error code: %v", code)
		return http.StatusInternalServerError
	}
}
func AddGRPCMetadata(
	ctx context.Context,
	headerMD, trailerMD metadata.MD,
	headerMatcher func(key string) (string, bool),
) {
	header := Header(ctx)
	for key, values := range headerMD {
		if h, ok := headerMatcher(key); ok {
			for _, v := range values {
				header.Add(h, v)
			}
		}
	}

	trailer := Trailer(ctx)
	for key, values := range trailerMD {
		key = fmt.Sprintf("%s%s", MetadataTrailerPrefix, key)
		for _, v := range values {
			trailer.Add(key, v)
		}
	}
}

func defaultOutgoingHeaderMatcher(key string) (string, bool) {
	return fmt.Sprintf("%s%s", MetadataHeaderPrefix, key), true
}
