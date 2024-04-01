// Code generated by "gors -service ObjObj"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	pb "github.com/go-leo/gors/example/api/pb"
	http "net/http"
)

func ObjObjRoutes(srv ObjObj, opts ...gors.Option) []gors.Route {
	options := gors.NewOptions(opts...)
	return []gors.Route{
		gors.NewRoute(http.MethodGet, "/api/ObjObj/all/:id/user/:user", _ObjObj_AllRequest_Handler(srv, options)),
		gors.NewRoute(http.MethodGet, "/api/ObjObj/UriBindingIndentedJSONRender/:id", _ObjObj_UriBindingIndentedJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodGet, "/api/ObjObj/QueryBindingSecureJSONRender/:id", _ObjObj_QueryBindingSecureJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodGet, "/api/ObjObj/HeaderBindingJsonpJSONRender/:id", _ObjObj_HeaderBindingJsonpJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPost, "/api/ObjObj/JSONBindingJSONRender/:id", _ObjObj_JSONBindingJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPatch, "/api/ObjObj/XMLBindingXMLRender/:id", _ObjObj_XMLBindingXMLRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPost, "/api/ObjObj/FormBindingJSONRender/:id", _ObjObj_FormBindingJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPost, "/api/ObjObj/FormPostBindingPureJSONRender/:id", _ObjObj_FormPostBindingPureJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPost, "/api/ObjObj/FormMultipartBindingAsciiJSONRender/:id", _ObjObj_FormMultipartBindingAsciiJSONRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPut, "/api/ObjObj/ProtoBufBindingProtoBufRender", _ObjObj_ProtoBufBindingProtoBufRender_Handler(srv, options)),
		gors.NewRoute(http.MethodDelete, "/api/ObjObj/MsgPackBindingMsgPackRender", _ObjObj_MsgPackBindingMsgPackRender_Handler(srv, options)),
		gors.NewRoute(http.MethodDelete, "/api/ObjObj/YAMLBindingYAMLRender/:id", _ObjObj_YAMLBindingYAMLRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPut, "/api/ObjObj/TOMLBindingTOMLRender/:id", _ObjObj_TOMLBindingTOMLRender_Handler(srv, options)),
		gors.NewRoute(http.MethodPut, "/api/ObjObj/ProtoJSONBindingProtoJSONRender", _ObjObj_ProtoJSONBindingProtoJSONRender_Handler(srv, options)),
	}
}

func _ObjObj_AllRequest_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/AllRequest"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *AllRequestReq
		var resp *AllRequestResp
		var err error
		req = new(AllRequestReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.JSONBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.AllRequest(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.JSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_UriBindingIndentedJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/UriBindingIndentedJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *UriBindingReq
		var resp *IndentedJSONRenderResp
		var err error
		req = new(UriBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.UriBindingIndentedJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.IndentedJSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_QueryBindingSecureJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/QueryBindingSecureJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *QueryBindingReq
		var resp *SecureJSONRenderResp
		var err error
		req = new(QueryBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.QueryBindingSecureJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.SecureJSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_HeaderBindingJsonpJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/HeaderBindingJsonpJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *HeaderBindingReq
		var resp *JsonpJSONRenderResp
		var err error
		req = new(HeaderBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.HeaderBindingJsonpJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/javascript", gors.JSONPJSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_JSONBindingJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/JSONBindingJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *JSONBindingReq
		var resp *JSONRenderResp
		var err error
		req = new(JSONBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.JSONBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.JSONBindingJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.JSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_XMLBindingXMLRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/XMLBindingXMLRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *XMLBindingReq
		var resp *XMLRenderResp
		var err error
		req = new(XMLBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.XMLBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.XMLBindingXMLRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/xml", gors.XMLRender, options.ResponseWrapper)
	}
}

func _ObjObj_FormBindingJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/FormBindingJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *FormBindingReq
		var resp *JSONRenderResp
		var err error
		req = new(FormBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.HeaderBinding,
			gors.FormBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.FormBindingJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.JSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_FormPostBindingPureJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/FormPostBindingPureJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *FormPostBindingReq
		var resp *PureJSONRenderResp
		var err error
		req = new(FormPostBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.FormPostBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.FormPostBindingPureJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.PureJSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_FormMultipartBindingAsciiJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/FormMultipartBindingAsciiJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *FormMultipartBindingReq
		var resp *AsciiJSONRenderResp
		var err error
		req = new(FormMultipartBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.FormMultipartBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.FormMultipartBindingAsciiJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.AsciiJSONRender, options.ResponseWrapper)
	}
}

func _ObjObj_ProtoBufBindingProtoBufRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/ProtoBufBindingProtoBufRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *pb.ProtoBufReq
		var resp *pb.ProtoBufResp
		var err error
		req = new(pb.ProtoBufReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.ProtoBufBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.ProtoBufBindingProtoBufRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/x-protobuf", gors.ProtoBufRender, options.ResponseWrapper)
	}
}

func _ObjObj_MsgPackBindingMsgPackRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/MsgPackBindingMsgPackRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *MsgPackBindingReq
		var resp *MsgPackRenderResp
		var err error
		req = new(MsgPackBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.MsgPackBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.MsgPackBindingMsgPackRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/x-msgpack", gors.MsgPackRender, options.ResponseWrapper)
	}
}

func _ObjObj_YAMLBindingYAMLRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/YAMLBindingYAMLRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *YAMLBindingReq
		var resp *YAMLRenderResp
		var err error
		req = new(YAMLBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.YAMLBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.YAMLBindingYAMLRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/x-yaml", gors.YAMLRender, options.ResponseWrapper)
	}
}

func _ObjObj_TOMLBindingTOMLRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/TOMLBindingTOMLRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *TOMLBindingReq
		var resp *TOMLRenderResp
		var err error
		req = new(TOMLBindingReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.UriBinding,
			gors.QueryBinding,
			gors.HeaderBinding,
			gors.TOMLBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.TOMLBindingTOMLRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/toml", gors.TOMLRender, options.ResponseWrapper)
	}
}

func _ObjObj_ProtoJSONBindingProtoJSONRender_Handler(srv ObjObj, options *gors.Options) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rpcMethodName = "/demo.ObjObj/ProtoJSONBindingProtoJSONRender"
		var ctx = gors.NewContext(c, rpcMethodName)
		var req *pb.ProtoBufReq
		var resp *pb.ProtoBufResp
		var err error
		req = new(pb.ProtoBufReq)
		if err = gors.RequestBind(
			ctx, req, options.Tag,
			gors.ProtoJSONBinding,
		); err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		resp, err = srv.ProtoJSONBindingProtoJSONRender(ctx, req)
		if err != nil {
			gors.ErrorRender(ctx, err, options.ErrorHandler, options.ResponseWrapper)
			return
		}
		gors.ResponseRender(ctx, gors.StatusCode(ctx), resp, "application/json", gors.ProtoJSONRender(options.ProtoJSONMarshalOptions), options.ResponseWrapper)
	}
}
