package svc

import (
	"context"
)

type StringStringService struct {
}

func (svc *StringStringService) GetStringString(ctx context.Context, s string) (string, error) {
	return "hello " + s, nil
}

func (svc *StringStringService) PatchStringString(ctx context.Context, s string) (string, error) {
	return "hello " + s, nil
}
