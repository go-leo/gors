// Code generated by "gors -service BindingRender"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	gors "github.com/go-leo/gors"
	pb "github.com/go-leo/gors/example/api/pb"
	http "net/http"
)

func BindingRenderRoutes(srv BindingRender) []gors.Route {
	return []gors.Route{
		gors.NewRoute(
			http.MethodGet,
			"/api/BindingRender/UriBindingIndentedJSONRender/:id",
			func(c *gin.Context) {
				var req *UriBindingReq
				var resp *IndentedJSONRenderResp
				var err error
				req = new(UriBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.UriBindingIndentedJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.IndentedJSONRender)
			},
		),
		gors.NewRoute(
			http.MethodGet,
			"/api/BindingRender/QueryBindingSecureJSONRender/:id",
			func(c *gin.Context) {
				var req *QueryBindingReq
				var resp *SecureJSONRenderResp
				var err error
				req = new(QueryBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.QueryBindingSecureJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.SecureJSONRender)
			},
		),
		gors.NewRoute(
			http.MethodGet,
			"/api/BindingRender/HeaderBindingJsonpJSONRender/:id",
			func(c *gin.Context) {
				var req *HeaderBindingReq
				var resp *JsonpJSONRenderResp
				var err error
				req = new(HeaderBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.HeaderBindingJsonpJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.JSONPJSONRender)
			},
		),
		gors.NewRoute(
			http.MethodPost,
			"/api/BindingRender/JSONBindingJSONRender/:id",
			func(c *gin.Context) {
				var req *JSONBindingReq
				var resp *JSONRenderResp
				var err error
				req = new(JSONBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.JSONBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.JSONBindingJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.JSONRender)
			},
		),
		gors.NewRoute(
			http.MethodPatch,
			"/api/BindingRender/XMLBindingXMLRender/:id",
			func(c *gin.Context) {
				var req *XMLBindingReq
				var resp *XMLRenderResp
				var err error
				req = new(XMLBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.XMLBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.XMLBindingXMLRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.XMLRender)
			},
		),
		gors.NewRoute(
			http.MethodPost,
			"/api/BindingRender/FormBindingJSONRender/:id",
			func(c *gin.Context) {
				var req *FormBindingReq
				var resp *JSONRenderResp
				var err error
				req = new(FormBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.HeaderBinding,
					gors.FormBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormBindingJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.JSONRender)
			},
		),
		gors.NewRoute(
			http.MethodPost,
			"/api/BindingRender/FormPostBindingPureJSONRender/:id",
			func(c *gin.Context) {
				var req *FormPostBindingReq
				var resp *PureJSONRenderResp
				var err error
				req = new(FormPostBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.FormPostBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormPostBindingPureJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.PureJSONRender)
			},
		),
		gors.NewRoute(
			http.MethodPost,
			"/api/BindingRender/FormMultipartBindingAsciiJSONRender/:id",
			func(c *gin.Context) {
				var req *FormMultipartBindingReq
				var resp *AsciiJSONRenderResp
				var err error
				req = new(FormMultipartBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.FormMultipartBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormMultipartBindingAsciiJSONRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.AsciiJSONRender)
			},
		),
		gors.NewRoute(
			http.MethodPut,
			"/api/BindingRender/ProtoBufBindingProtoBufRender",
			func(c *gin.Context) {
				var req *pb.ProtoBufReq
				var resp *pb.ProtoBufResp
				var err error
				req = new(pb.ProtoBufReq)
				if err := gors.ShouldBind(
					c, req,
					gors.ProtoBufBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.ProtoBufBindingProtoBufRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.ProtoBufRender)
			},
		),
		gors.NewRoute(
			http.MethodDelete,
			"/api/BindingRender/MsgPackBindingMsgPackRender",
			func(c *gin.Context) {
				var req *MsgPackBindingReq
				var resp *MsgPackRenderResp
				var err error
				req = new(MsgPackBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.MsgPackBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.MsgPackBindingMsgPackRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.MsgPackRender)
			},
		),
		gors.NewRoute(
			http.MethodDelete,
			"/api/BindingRender/YAMLBindingYAMLRender/:id",
			func(c *gin.Context) {
				var req *YAMLBindingReq
				var resp *YAMLRenderResp
				var err error
				req = new(YAMLBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.YAMLBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.YAMLBindingYAMLRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.YAMLRender)
			},
		),
		gors.NewRoute(
			http.MethodPut,
			"/api/BindingRender/TOMLBindingTOMLRender/:id",
			func(c *gin.Context) {
				var req *TOMLBindingReq
				var resp *TOMLRenderResp
				var err error
				req = new(TOMLBindingReq)
				if err := gors.ShouldBind(
					c, req,
					gors.UriBinding,
					gors.QueryBinding,
					gors.HeaderBinding,
					gors.TOMLBinding,
				); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				if err = gors.Validate(req); err != nil {
					gors.HandleBadRequest(c, err)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.TOMLBindingTOMLRender(ctx, req)
				gors.MustRender(c, resp, err, "", gors.TOMLRender)
			},
		),
	}
}
