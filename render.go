package gors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	internalrender "github.com/go-leo/gors/internal/pkg/render"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/iox"
	"io"
	"net/http"
)

func MustRender(c *gin.Context, resp any, err error, contentType string, fn func(c *gin.Context, resp any, contentType string)) {
	if err == nil {
		fn(c, resp, contentType)
		return
	}

	var httpError *HttpError
	if errors.As(err, &httpError) {
		c.String(httpError.StatusCode(), httpError.Error())
		_ = c.Error(httpError).SetType(gin.ErrorTypePublic)
		return
	}

	c.String(http.StatusInternalServerError, err.Error())
	_ = c.Error(err).SetType(gin.ErrorTypePrivate)
}

func BytesRender(c *gin.Context, resp any, contentType string) {
	c.Render(HTTPStatusCode(c), render.Data{ContentType: contentType, Data: resp.([]byte)})
}

func StringRender(c *gin.Context, resp any, contentType string) {
	c.Render(HTTPStatusCode(c), render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func TextRender(c *gin.Context, resp any, contentType string) {
	c.Render(HTTPStatusCode(c), render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func HTMLRender(c *gin.Context, resp any, contentType string) {
	c.Render(HTTPStatusCode(c), render.Data{ContentType: contentType, Data: convx.StringToBytes(resp.(string))})
}

func RedirectRender(c *gin.Context, resp any, contentType string) {
	c.Redirect(HTTPStatusCode(c), resp.(string))
}

func ReaderRender(c *gin.Context, resp any, contentType string) {
	r := resp.(io.Reader)
	l, ok := iox.Len(r)
	if !ok {
		l = -1
	}
	c.Render(HTTPStatusCode(c), render.Reader{ContentType: contentType, ContentLength: int64(l), Reader: r})
}

func JSONRender(c *gin.Context, resp any, _ string) {
	c.JSON(HTTPStatusCode(c), resp)
}

func IndentedJSONRender(c *gin.Context, resp any, _ string) {
	c.IndentedJSON(HTTPStatusCode(c), resp)
}

func SecureJSONRender(c *gin.Context, resp any, _ string) {
	c.SecureJSON(HTTPStatusCode(c), resp)
}

func JSONPJSONRender(c *gin.Context, resp any, _ string) {
	c.JSONP(HTTPStatusCode(c), resp)
}

func PureJSONRender(c *gin.Context, resp any, _ string) {
	c.PureJSON(HTTPStatusCode(c), resp)
}

func AsciiJSONRender(c *gin.Context, resp any, _ string) {
	c.AsciiJSON(HTTPStatusCode(c), resp)
}

func XMLRender(c *gin.Context, resp any, _ string) {
	c.XML(HTTPStatusCode(c), resp)
}

func YAMLRender(c *gin.Context, resp any, _ string) {
	c.YAML(HTTPStatusCode(c), resp)
}

func ProtoBufRender(c *gin.Context, resp any, _ string) {
	c.ProtoBuf(HTTPStatusCode(c), resp)
}

func MsgPackRender(c *gin.Context, resp any, _ string) {
	c.Render(HTTPStatusCode(c), render.MsgPack{Data: resp})
}

func TOMLRender(c *gin.Context, resp any, _ string) {
	c.TOML(HTTPStatusCode(c), resp)
}

func CustomRender(c *gin.Context, resp any, _ string) {
	customRender, ok := resp.(Render)
	if !ok {
		return
	}
	customRender.Render(c)
}

func ProtoJSONRender(c *gin.Context, resp any, _ string) {
	c.Render(HTTPStatusCode(c), internalrender.ProtoJSON{Data: resp})
}
