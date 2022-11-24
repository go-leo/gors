package svc

import "context"

type StringBytesService struct {
}

func (svc *StringBytesService) GetStringBytes(ctx context.Context, s string) ([]byte, error) {
	return []byte("hello " + s), nil
}

func (svc *StringBytesService) OptionsStringBytes(ctx context.Context, s string) ([]byte, error) {
	return []byte("hello " + s), nil
}
