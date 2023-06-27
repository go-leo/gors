package gors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	internalbinding "github.com/go-leo/gors/internal/pkg/binding"
	"github.com/go-leo/gox/stringx"
	"net/http"
)

var HandleBadRequest = DefaultHandleBadRequest

func ShouldBind(c *gin.Context, req any, tag string, fns ...func(c *gin.Context, req any, tag string) error) error {
	for _, fn := range fns {
		if err := fn(c, req, tag); err != nil {
			return err
		}
	}
	if err := Validate(req); err != nil {
		return err
	}
	return nil
}

func UriBinding(c *gin.Context, req any, tag string) error {
	if stringx.IsBlank(tag) {
		return c.ShouldBindUri(req)
	}
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}
	return binding.MapFormWithTag(req, m, tag)
}

func QueryBinding(c *gin.Context, req any, tag string) error {
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Query)
	}
	return binding.MapFormWithTag(req, c.Request.URL.Query(), tag)
}

func HeaderBinding(c *gin.Context, req any, tag string) error {
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Header)
	}
	return binding.MapFormWithTag(req, c.Request.Header, tag)
}

func FormBinding(c *gin.Context, req any, tag string) error {
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.Form)
	}
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	const defaultMemory = 32 << 20
	if err := c.Request.ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.Form, tag)
}

func FormPostBinding(c *gin.Context, req any, tag string) error {
	if stringx.IsBlank(tag) {
		return c.ShouldBindWith(req, binding.FormPost)
	}
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	return binding.MapFormWithTag(req, c.Request.PostForm, tag)
}

func FormMultipartBinding(c *gin.Context, req any, tag string) error {
	return c.ShouldBindWith(req, binding.FormMultipart)
}

func JSONBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.JSON)
}

func XMLBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.XML)
}

func ProtoBufBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.ProtoBuf)
}

func MsgPackBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.MsgPack)
}

func YAMLBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.YAML)
}

func TOMLBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, binding.TOML)
}

func CustomBinding(c *gin.Context, req any, _ string) error {
	customBinding, ok := req.(Binding)
	if !ok {
		return nil
	}
	return customBinding.Bind(c)
}

func ProtoJSONBinding(c *gin.Context, req any, _ string) error {
	return c.ShouldBindWith(req, internalbinding.ProtoJSON)
}

func DefaultHandleBadRequest(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, err.Error())
	_ = c.Error(err).SetType(gin.ErrorTypeBind)
}
