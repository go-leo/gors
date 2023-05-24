// Code generated by "gors -service BindingRender"; DO NOT EDIT.

package demo

import (
	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	render "github.com/gin-gonic/gin/render"
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.UriBindingIndentedJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.IndentedJSON(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.QueryBindingSecureJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.SecureJSON(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.HeaderBindingJsonpJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.JSONP(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.JSON); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.JSONBindingJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.JSON(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.XML); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.XMLBindingXMLRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.XML(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Form); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormBindingJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.JSON(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.FormPost); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormPostBindingPureJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.PureJSON(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.FormMultipart); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.FormMultipartBindingAsciiJSONRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.AsciiJSON(statusCode, resp)
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
				if err = c.ShouldBindWith(req, binding.ProtoBuf); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.ProtoBufBindingProtoBufRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.ProtoBuf(statusCode, resp)
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
				if err = c.ShouldBindWith(req, binding.MsgPack); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.MsgPackBindingMsgPackRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.Render(statusCode, render.MsgPack{Data: resp})
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.YAML); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.YAMLBindingYAMLRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.YAML(statusCode, resp)
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
				if err = c.ShouldBindUri(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Query); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.Header); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = c.ShouldBindWith(req, binding.TOML); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				if err = gors.Validate(req); err != nil {
					c.String(http.StatusBadRequest, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypeBind)
					return
				}
				ctx := gors.NewContext(c)
				resp, err = srv.TOMLBindingTOMLRender(ctx, req)
				if gors.IsInterrupted(ctx) {
					return
				}
				if err != nil {
					if httpErr, ok := err.(*gors.HttpError); ok {
						c.String(httpErr.StatusCode(), httpErr.Error())
						_ = c.Error(err).SetType(gin.ErrorTypePublic)
						return
					}
					c.String(http.StatusInternalServerError, err.Error())
					_ = c.Error(err).SetType(gin.ErrorTypePrivate)
					return
				}
				statusCode := gors.GetCodeFromContext(ctx)
				c.TOML(statusCode, resp)
			},
		),
	}
}
