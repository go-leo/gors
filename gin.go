package gors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
