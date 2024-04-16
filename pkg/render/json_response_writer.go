package render

import (
	"bytes"
	"net/http"
)

type JSONResponseWriter struct {
	Buffer bytes.Buffer
}

func (w *JSONResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (w *JSONResponseWriter) Write(i []byte) (int, error) {
	return w.Buffer.Write(i)
}

func (w *JSONResponseWriter) WriteHeader(statusCode int) {}
