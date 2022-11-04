package gors

import "fmt"

type HttpError struct {
	statusCode int
	err        error
}

func NewHttpError(statusCode int, err error) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		err:        err,
	}
}

func (h HttpError) StatusCode() int {
	return h.statusCode
}

func (h HttpError) Error() string {
	return fmt.Sprintf("http status code is %d, err: %v", h.statusCode, h.err)
}
