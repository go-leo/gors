package gors

import "fmt"

// HttpError http 错误，
type HttpError struct {
	// statusCode http status code
	statusCode int
	// err error message
	err error
}

func (h HttpError) StatusCode() int {
	return h.statusCode
}

func (h HttpError) Error() string {
	return fmt.Sprintf("http status code is %d, err: %v", h.statusCode, h.err)
}

// NewHttpError 创建一个HttpError
func NewHttpError(statusCode int, err error) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		err:        err,
	}
}
