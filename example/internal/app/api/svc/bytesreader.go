package svc

import (
	"bytes"
	"context"
	"io"
)

type BytesReaderService struct {
}

func (svc *BytesReaderService) GetBytesReader(ctx context.Context, req []byte) (io.Reader, error) {
	return bytes.NewReader(append([]byte("hello "), req...)), nil
}

func (svc *BytesReaderService) PatchBytesReader(ctx context.Context, req []byte) (io.Reader, error) {
	return bytes.NewReader(append([]byte("hello "), req...)), nil
}
