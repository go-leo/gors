package gors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func ShouldBind(c *gin.Context, req any, fns ...func(c *gin.Context, req any) error) error {
	for _, fn := range fns {
		if err := fn(c, req); err != nil {
			return err
		}
	}
	return nil
}

func UriBinding(c *gin.Context, req any) error {
	return c.ShouldBindUri(req)
}

func QueryBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.Query)
}

func HeaderBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.Header)
}

func FormBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.Form)
}

func FormPostBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.FormPost)
}

func FormMultipartBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.FormMultipart)
}

func JSONBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.JSON)
}

func XMLBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.XML)
}

func ProtoBufBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.ProtoBuf)
}

func MsgPackBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.MsgPack)
}

func YAMLBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.YAML)
}

func TOMLBinding(c *gin.Context, req any) error {
	return c.ShouldBindWith(req, binding.TOML)
}

func CustomBinding(c *gin.Context, req any) error {
	customBinding, ok := req.(Binding)
	if !ok {
		return nil
	}
	return customBinding.Bind(c)
}

func HandleBadRequest(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, err.Error())
	_ = c.Error(err).SetType(gin.ErrorTypeBind)
}

func UriParams(c *gin.Context) (map[string][]string, error) {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return m, nil
}

func QueryParams(c *gin.Context) (map[string][]string, error) {
	return c.Request.URL.Query(), nil
}

func HeaderParams(c *gin.Context) (map[string][]string, error) {
	return c.Request.Header, nil
}

func FormParams(c *gin.Context) (map[string][]string, error) {
	if err := c.Request.ParseForm(); err != nil {
		return nil, err
	}
	const defaultMemory = 32 << 20
	if err := c.Request.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return nil, err
	}
	return c.Request.Form, nil
}

func FormPostParams(c *gin.Context) (map[string][]string, error) {
	if err := c.Request.ParseForm(); err != nil {
		return nil, err
	}
	return c.Request.PostForm, nil
}
