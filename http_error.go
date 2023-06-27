package gors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpError 指定http状态码和错误信息
type HttpError struct {
	// statusCode http status code
	statusCode int
	// err error message
	err error
	// errType gin.ErrorType
	errType gin.ErrorType
}

func (e *HttpError) Error() string {
	return e.err.Error()
}

// NewHttpError 创建Http错误
func NewHttpError(statusCode int, err error) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		err:        err,
		errType:    gin.ErrorTypePublic,
	}
}

func BindError(err error) *HttpError {
	return &HttpError{
		statusCode: http.StatusBadGateway,
		err:        err,
		errType:    gin.ErrorTypeBind,
	}
}
