package svc

import (
	"context"
	"net/http"

	"github.com/go-leo/gors"
)

type BytesStringService struct {
}

func (svc *BytesStringService) GetBytesString(ctx context.Context, req []byte) (string, error) {
	return "hello " + string(req), nil
}

func (svc *BytesStringService) PutBytesString(ctx context.Context, req []byte) (string, error) {
	gors.SetStatusCode(ctx, http.StatusSeeOther)
	return "/api/BytesString/Get", nil
}
