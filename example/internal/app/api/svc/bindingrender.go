package svc

import (
	"context"
	"github.com/go-leo/gors/example/api/demo"
	"github.com/go-leo/gors/example/api/pb"
	"time"
)

var _ demo.ObjObj = new(ObjObjService)

type ObjObjService struct {
}

func (svc *ObjObjService) UriBindingIndentedJSONRender(ctx context.Context, req *demo.UriBindingReq) (*demo.IndentedJSONRenderResp, error) {
	return &demo.IndentedJSONRenderResp{
		ID:         req.ID,
		Name:       "Jax",
		Address:    "shanghai",
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *ObjObjService) QueryBindingSecureJSONRender(ctx context.Context, req *demo.QueryBindingReq) (*demo.SecureJSONRenderResp, error) {
	return &demo.SecureJSONRenderResp{
		ID:         req.ID,
		Name:       req.Name,
		Address:    "shanghai",
		Birthday:   time.Now(),
		CreateTime: time.Now(),
		UnixTime:   time.Now(),
	}, nil
}

func (svc *ObjObjService) HeaderBindingJsonpJSONRender(ctx context.Context, req *demo.HeaderBindingReq) (*demo.JsonpJSONRenderResp, error) {
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

func (svc *ObjObjService) JSONBindingJSONRender(ctx context.Context, req *demo.JSONBindingReq) (*demo.JSONRenderResp, error) {
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

func (svc *ObjObjService) XMLBindingXMLRender(ctx context.Context, req *demo.XMLBindingReq) (*demo.XMLRenderResp, error) {
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

func (svc *ObjObjService) FormBindingJSONRender(ctx context.Context, req *demo.FormBindingReq) (*demo.JSONRenderResp, error) {
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

func (svc *ObjObjService) FormPostBindingPureJSONRender(ctx context.Context, req *demo.FormPostBindingReq) (*demo.PureJSONRenderResp, error) {
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

func (svc *ObjObjService) FormMultipartBindingAsciiJSONRender(ctx context.Context, req *demo.FormMultipartBindingReq) (*demo.AsciiJSONRenderResp, error) {
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

func (svc *ObjObjService) ProtoBufBindingProtoBufRender(ctx context.Context, req *pb.ProtoBufReq) (*pb.ProtoBufResp, error) {
	return &pb.ProtoBufResp{
		Id:   req.GetId(),
		Name: "jax",
	}, nil
}

func (svc *ObjObjService) MsgPackBindingMsgPackRender(ctx context.Context, req *demo.MsgPackBindingReq) (*demo.MsgPackRenderResp, error) {
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

func (svc *ObjObjService) YAMLBindingYAMLRender(ctx context.Context, req *demo.YAMLBindingReq) (*demo.YAMLRenderResp, error) {
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

func (svc *ObjObjService) TOMLBindingTOMLRender(ctx context.Context, req *demo.TOMLBindingReq) (*demo.TOMLRenderResp, error) {
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

func (svc *ObjObjService) ProtoJSONBindingProtoJSONRender(ctx context.Context, req *pb.ProtoBufReq) (*pb.ProtoBufResp, error) {
	return &pb.ProtoBufResp{
		Id:   req.Id,
		Name: "Jax",
	}, nil
}
