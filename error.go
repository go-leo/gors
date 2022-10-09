package gors

import "fmt"

type HttpError struct {
	err         error
	statusCode  int
	contentType string
}

type HttpErrorOption func(httpError *HttpError)

func WithContentType(contentType string) HttpErrorOption {
	return func(httpError *HttpError) {
		httpError.contentType = contentType
	}
}

func NewHttpError(err error, statusCode int, opts ...HttpErrorOption) *HttpError {
	httpErr := &HttpError{err: err, statusCode: statusCode}
	for _, opt := range opts {
		opt(httpErr)
	}
	return httpErr
}

func (h HttpError) Err() error {
	return h.err
}

func (h HttpError) StatusCode() int {
	return h.statusCode
}

func (h HttpError) ContentType() string {
	return h.contentType
}

func (h HttpError) Error() string {
	return fmt.Sprintf("http status code is %d, err: %v", h.statusCode, h.err)
}
