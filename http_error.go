package gors

// HttpError 指定http状态码和错误信息
type HttpError struct {
	// statusCode http status code
	statusCode int
	// err error message
	err error
}

func (e *HttpError) StatusCode() int {
	return e.statusCode
}

func (e *HttpError) Error() string {
	return e.err.Error()
}

// NewHttpError 创建Http错误
func NewHttpError(statusCode int, err error) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		err:        err,
	}
}
