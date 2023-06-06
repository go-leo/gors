package svc

import (
	"context"
	"errors"
	"github.com/go-leo/gors/example/api/demo"
)

type Custom struct {
}

func (c *Custom) Custom(ctx context.Context, req *demo.CustomReq) (*demo.CustomResp, error) {
	return nil, errors.New("helloworld")
}
