package gors

import (
	"github.com/gin-gonic/gin"
)

// Validator 如果请求参数struct实现Validator接口，框架会验证数据
type Validator interface {
	Validate() error
}

func Validate(req interface{}) error {
	switch v := req.(type) {
	case Validator:
		return v.Validate()
	default:
		return nil
	}
}

// Binding 如果请求参数struct实现Binding接口，框架会用自定义绑定逻辑来绑定请求参数
type Binding interface {
	Bind(c *gin.Context) error
}

// Render 如果响应参数struct实现Render接口，框架会用自定义渲染逻辑来渲染响应参数
type Render interface {
	Render(c *gin.Context, statusCode int) error
}
