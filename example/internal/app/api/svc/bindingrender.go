package svc

import (
	"context"
	"time"

	"github.com/go-leo/gors/example/api/demo"
	"github.com/go-leo/gors/example/api/pb"
)

type BindingRenderService struct {
}

func (svc *BindingRenderService) UriBindingIndentedJSONRender(ctx context.Context, req *demo.UriBindingReq) (*demo.IndentedJSONRenderResp, error) {
	return &demo.IndentedJSONRenderResp{
		ID:         req.ID,
		Name:       "Jax",
		Address:    "shanghai",
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) QueryBindingSecureJSONRender(ctx context.Context, req *demo.QueryBindingReq) (*demo.SecureJSONRenderResp, error) {
	return &demo.SecureJSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Address:    "shanghai",
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) HeaderBindingJsonpJSONRender(ctx context.Context, req *demo.HeaderBindingReq) (*demo.JsonpJSONRenderResp, error) {
	return &demo.JsonpJSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    "shanghai",
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) JSONBindingJSONRender(ctx context.Context, req *demo.JSONBindingReq) (*demo.JSONRenderResp, error) {
	return &demo.JSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) XMLBindingXMLRender(ctx context.Context, req *demo.XMLBindingReq) (*demo.XMLRenderResp, error) {
	return &demo.XMLRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) FormBindingJSONRender(ctx context.Context, req *demo.FormBindingReq) (*demo.JSONRenderResp, error) {
	return &demo.JSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) FormPostBindingPureJSONRender(ctx context.Context, req *demo.FormPostBindingReq) (*demo.PureJSONRenderResp, error) {
	return &demo.PureJSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) FormMultipartBindingAsciiJSONRender(ctx context.Context, req *demo.FormMultipartBindingReq) (*demo.AsciiJSONRenderResp, error) {
	return &demo.AsciiJSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) ProtoBufBindingProtoBufRender(ctx context.Context, req *pb.ProtoBufReq) (*pb.ProtoBufResp, error) {
	return &pb.ProtoBufResp{
		Id:   req.GetId(),
		Name: "jax",
	}, nil
}

func (svc *BindingRenderService) MsgPackBindingMsgPackRender(ctx context.Context, req *demo.MsgPackBindingReq) (*demo.MsgPackRenderResp, error) {
	return &demo.MsgPackRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) YAMLBindingYAMLRender(ctx context.Context, req *demo.YAMLBindingReq) (*demo.YAMLRenderResp, error) {
	return &demo.YAMLRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *BindingRenderService) TOMLBindingTOMLRender(ctx context.Context, req *demo.TOMLBindingReq) (*demo.TOMLRenderResp, error) {
	return &demo.TOMLRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Auth:       req.Auth,
		Address:    req.Address,
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}
