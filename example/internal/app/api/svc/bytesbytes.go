package svc

import "context"

type BytesBytesService struct {
}

func (svc *BytesBytesService) GetBytesBytes(ctx context.Context, req []byte) ([]byte, error) {
	return append([]byte("hello "), req...), nil
}

func (svc *BytesBytesService) PostBytesBytes(ctx context.Context, req []byte) ([]byte, error) {
	return append([]byte("hello "), req...), nil
}
