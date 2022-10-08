package gors

import "fmt"

type HttpError struct {
	err        error
	statusCode int
}

func NewHttpError(err error, statusCode int) *HttpError {
	return &HttpError{err: err, statusCode: statusCode}
}

func (h HttpError) Err() error {
	return h.err
}

func (h HttpError) StatusCode() int {
	return h.statusCode
}

func (h HttpError) Error() string {
	return fmt.Sprintf("http status code is %d, err: %v", h.statusCode, h.err)
}
