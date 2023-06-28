package gors

import (
	"fmt"
	"github.com/go-leo/gors/internal/pkg/status"
)

// Error 指定http状态码和错误信息
type Error struct {
	// StatusCode http status code
	StatusCode int
	// Code business service code
	Code int
	// Message error message
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e Error) Proto() *status.Status {
	return &status.Status{
		Code:    int32(e.Code),
		Message: e.Message,
	}
}
