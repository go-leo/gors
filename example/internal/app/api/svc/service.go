package svc

import (
	"context"
	"fmt"

	"github.com/go-leo/gors/example/api/demo"
)

var _ demo.Service = new(Service)

type Service struct{}

func (svc *Service) Method(ctx context.Context, req *demo.MethodReq) (*demo.MethodResp, error) {
	fmt.Println(req.ID)
	return &demo.MethodResp{V: 10}, nil
}
