package gors

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin/render"
	internalrender "github.com/go-leo/gors/internal/pkg/render"
	"github.com/go-leo/gox/iox"
	"io"
	"net/http"
	"strings"
)

func ErrorRender(ctx context.Context, err error, handler func(ctx context.Context, err error)) {
	if handler != nil {
		handler(ctx, err)
		return
	}
	var e Error
	if errors.As(err, &e) {
		ResponseRender(ctx, e.StatusCode, e.Error(), "", StringRender, nil)
		return
	}
	var ePtr *Error
	if errors.As(err, &ePtr) {
		ResponseRender(ctx, e.StatusCode, e.Error(), "", StringRender, nil)
		return
	}
	ResponseRender(ctx, http.StatusInternalServerError, err.Error(), "", StringRender, nil)
}

func ResponseRender(ctx context.Context, code int, resp any, contentType string, render func(ctx context.Context, code int, resp any, contentType string), wrapper func(resp any) any) {
	c := FromContext(ctx)
	header := Header(ctx)
	for key, values := range header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	te := c.GetHeader("TE")
	var doForwardTrailers bool
	if strings.Contains(strings.ToLower(te), "trailers") {
		doForwardTrailers = true
	}

	trailer := Trailer(ctx)
	if doForwardTrailers {
		for k := range trailer {
			c.Writer.Header().Add("Trailer", k)
		}
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
	}

	if wrapper != nil {
		resp = wrapper(resp)
	}

	render(ctx, code, resp, contentType)

	if doForwardTrailers {
		for key, values := range trailer {
			for _, value := range values {
				c.Writer.Header().Add(key, value)
			}
		}
	}
}

func BytesRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: resp.([]byte)})
}

func StringRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: []byte(resp.(string))})
}

func TextRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: []byte(resp.(string))})
}

func HTMLRender(ctx context.Context, code int, resp any, contentType string) {
	FromContext(ctx).Render(code, render.Data{ContentType: contentType, Data: []byte(resp.(string))})
}

func RedirectRender(ctx context.Context, code int, resp any, contentType string) {
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
