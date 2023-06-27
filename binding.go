package gors

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

var HandleBadRequest = DefaultHandleBadRequest

func ShouldBind(c *gin.Context, req any, fns ...func(c *gin.Context, req any) error) error {
	for _, fn := range fns {
		if err := fn(c, req); err != nil {
			return err
		}
	}
	if err := Validate(req); err != nil {
		return err
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

func DefaultHandleBadRequest(c *gin.Context, err error) {
	c.String(http.StatusBadRequest, err.Error())
	_ = c.Error(err).SetType(gin.ErrorTypeBind)
}
