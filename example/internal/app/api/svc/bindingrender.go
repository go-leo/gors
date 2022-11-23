package svc

import (
	"context"

	"github.com/go-leo/gors/example/api/demo"
	"github.com/go-leo/gors/example/api/pb"
)

type BindingRenderService struct {
}

func (svc *BindingRenderService) UriBindingIndentedJSONRender(ctx context.Context, req *demo.UriBindingReq) (*demo.IndentedJSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) QueryBindingSecureJSONRender(ctx context.Context, req *demo.QueryBindingReq) (*demo.SecureJSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) HeaderBindingJsonpJSONRender(ctx context.Context, req *demo.HeaderBindingReq) (*demo.JsonpJSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) JSONBindingJSONRender(ctx context.Context, req *demo.JSONBindingReq) (*demo.JSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) XMLBindingXMLRender(ctx context.Context, req *demo.XMLBindingReq) (*demo.XMLRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) FormBindingJSONRender(ctx context.Context, req *demo.FormBindingReq) (*demo.JSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) FormPostBindingPureJSONRender(ctx context.Context, req *demo.FormPostBindingReq) (*demo.PureJSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) FormMultipartBindingAsciiJSONRender(ctx context.Context, req *demo.FormMultipartBindingReq) (*demo.AsciiJSONRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) ProtoBufBindingProtoBufRender(ctx context.Context, req *pb.ProtoBufReq) (*pb.ProtoBufResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) MsgPackBindingMsgPackRender(ctx context.Context, req *demo.MsgPackBindingReq) (*demo.MsgPackRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) YAMLBindingYAMLRender(ctx context.Context, req *demo.YAMLBindingReq) (*demo.YAMLRenderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BindingRenderService) TOMLBindingTOMLRender(ctx context.Context, req *demo.TOMLBindingReq) (*demo.TOMLRenderResp, error) {
	//TODO implement me
	panic("implement me")
}
