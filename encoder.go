package gors

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

func DefaultErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(interface{ Headers() http.Header }); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(interface{ StatusCode() int }); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Default().Println("gors: response write error: ", err)
	}
}
