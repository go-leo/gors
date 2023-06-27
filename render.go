package gors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	internalrender "github.com/go-leo/gors/internal/pkg/render"
	internalstatus "github.com/go-leo/gors/internal/pkg/status"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/iox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
)

func HTTPErrorRender(c *gin.Context, err error) {
	var httpError *HttpError
	if errors.As(err, &httpError) {
		c.String(httpError.statusCode, httpError.err.Error())
		_ = c.Error(httpError.err).SetType(httpError.errType)
		return
	}
	c.String(http.StatusInternalServerError, err.Error())
	_ = c.Error(err).SetType(gin.ErrorTypePrivate)
}

func GRPCErrorRender(c *gin.Context, err error, headerMD metadata.MD, trailerMD metadata.MD, fn func(c *gin.Context, code int, resp any, contentType string)) {
	var httpError *HttpError
	if errors.As(err, &httpError) {
		err = httpError.err
	}

	grpcStatus := status.Convert(err)

	c.Writer.Header().Del("Trailer")
	c.Writer.Header().Del("Transfer-Encoding")

	if grpcStatus.Code() == codes.Unauthenticated {
		c.Writer.Header().Set("WWW-Authenticate", grpcStatus.Message())
	}

	handleForwardResponseGRPCMetadata(c, headerMD)

	// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
	// Unless the request includes a TE header field indicating "trailers"
	// is acceptable, as described in Section 4.3, a server SHOULD NOT
	// generate trailer fields that it believes are necessary for the user
	// agent to receive.
	doForwardTrailers := requestAcceptsTrailers(c)

	if doForwardTrailers {
		handleForwardResponseTrailerHeader(c, trailerMD)
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
	}

	st := HTTPStatusFromCode(grpcStatus.Code())
	if httpError != nil {
		st = httpError.statusCode
	}

	pb := grpcStatus.Proto()
	errResp := &internalstatus.Status{
		Code:    pb.Code,
		Message: pb.Message,
		Details: pb.Details,
	}
	fn(c, st, errResp, "")

	if doForwardTrailers {
		handleForwardResponseTrailer(c, trailerMD)
	}
}

func ResponseRender(c *gin.Context, code int, resp any, err error, contentType string, fn func(c *gin.Context, code int, resp any, contentType string)) {
	if err == nil {
		fn(c, code, resp, contentType)
		return
	}
}

func BytesRender(c *gin.Context, code int, resp any, contentType string) {
	c.Render(code, render.Data{ContentType: contentType, Data: resp.([]byte)})
}

func StringRender(c *gin.Context, code int, resp any, contentType string) {
	c.Render(code, render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func TextRender(c *gin.Context, code int, resp any, contentType string) {
	c.Render(code, render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func HTMLRender(c *gin.Context, code int, resp any, contentType string) {
	c.Render(code, render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func RedirectRender(c *gin.Context, code int, resp any, contentType string) {
	c.Redirect(code, resp.(string))
}

func ReaderRender(c *gin.Context, code int, resp any, contentType string) {
	r := resp.(io.Reader)
	l, ok := iox.Len(r)
	if !ok {
		l = -1
	}
	c.Render(code, render.Reader{ContentType: contentType, ContentLength: int64(l), Reader: r})
}

func JSONRender(c *gin.Context, code int, resp any, _ string) {
	c.JSON(code, resp)
}

func IndentedJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.IndentedJSON(code, resp)
}

func SecureJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.SecureJSON(code, resp)
}

func JSONPJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.JSONP(code, resp)
}

func PureJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.PureJSON(code, resp)
}

func AsciiJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.AsciiJSON(code, resp)
}

func XMLRender(c *gin.Context, code int, resp any, _ string) {
	c.XML(code, resp)
}

func YAMLRender(c *gin.Context, code int, resp any, _ string) {
	c.YAML(code, resp)
}

func ProtoBufRender(c *gin.Context, code int, resp any, _ string) {
	c.ProtoBuf(code, resp)
}

func MsgPackRender(c *gin.Context, code int, resp any, _ string) {
	c.Render(code, render.MsgPack{Data: resp})
}

func TOMLRender(c *gin.Context, code int, resp any, _ string) {
	c.TOML(code, resp)
}

func CustomRender(c *gin.Context, code int, resp any, _ string) {
	customRender, ok := resp.(Render)
	if !ok {
		return
	}
	customRender.Render(c)
}

func ProtoJSONRender(c *gin.Context, code int, resp any, _ string) {
	c.Render(code, internalrender.ProtoJSON{Data: resp})
}
