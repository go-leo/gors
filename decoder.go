package gors

import (
	"context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"io"
	"net/http"
)

func HttpBodyDecoder(ctx context.Context, r *http.Request, body *httpbody.HttpBody) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	body.Data = data
	body.ContentType = r.Header.Get("Content-Type")
	return nil
}
