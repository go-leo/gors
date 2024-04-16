package gors

import (
	"context"
	"google.golang.org/protobuf/encoding/protojson"

	// "errors"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin/render"
	renderPkg "github.com/go-leo/gors/pkg/render"
	"github.com/go-leo/gox/iox"
	"google.golang.org/grpc/metadata"
)

func ErrorRender(
	ctx context.Context,
	err error,
	handler func(ctx context.Context, err error) error,
	wrapper func(resp any) any,
) {
	_ = FromContext(ctx).Error(err)
	if handler != nil {
		err = handler(ctx, err)
		if err == nil {
			return
		}
	}
	gorserr := FromError(err)
	ResponseRender(ctx, gorserr.StatusCode, gorserr.Status(), "", JSONRender, wrapper)
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
	length := int64(-1)
	l, ok := iox.Len(r)
	if ok {
		length = int64(l)
	}
	FromContext(ctx).Render(code, render.Reader{ContentType: contentType, ContentLength: length, Reader: r})
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

func ProtoJSONRender(mo protojson.MarshalOptions) func(ctx context.Context, code int, resp any, _ string) {
	return func(ctx context.Context, code int, resp any, _ string) {
		FromContext(ctx).Render(code, renderPkg.ProtoJSON{Data: resp, MarshalOptions: mo})
	}
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

func AddGRPCMetadata(
	ctx context.Context,
	headerMD, trailerMD metadata.MD,
	headerMatcher func(key string) (string, bool),
) {
	if headerMatcher == nil {
		headerMatcher = defaultOutgoingHeaderMatcher
	}
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

var DefaultMarshalOptions = protojson.MarshalOptions{}
