package svc

import (
	"context"
	"io"
	"strings"
)

type StringReaderService struct {
}

func (svc *StringReaderService) GetStringRender(ctx context.Context, s string) (io.Reader, error) {
	return strings.NewReader("hello " + s), nil
}

func (svc *StringReaderService) OptionsStringReader(ctx context.Context, s string) (io.Reader, error) {
	return strings.NewReader("hello " + s), nil
}
