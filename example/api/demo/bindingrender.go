package demo

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/go-leo/gors/example/api/pb"
)

//go:generate gors -service BindingRender

// BindingRender
// @GORS @Path("/api")  @Path("/BindingRender")
type BindingRender interface {
	// UriBindingIndentedJSONRender
	// @GORS @GET @Path("/UriBindingIndentedJSONRender/:id") @UriBinding @IndentedJSONRender
	UriBindingIndentedJSONRender(context.Context, *UriBindingReq) (*IndentedJSONRenderResp, error)
	// QueryBindingSecureJSONRender
	// @GORS @GET @Path("/QueryBindingSecureJSONRender/:id") @UriBinding @QueryBinding @SecureJSONRender
	QueryBindingSecureJSONRender(context.Context, *QueryBindingReq) (*SecureJSONRenderResp, error)
	// HeaderBindingJsonpJSONRender
	// @GORS @GET @Path("/HeaderBindingJsonpJSONRender/:id") @UriBinding @QueryBinding @HeaderBinding @JsonpJSONRender
	HeaderBindingJsonpJSONRender(context.Context, *HeaderBindingReq) (*JsonpJSONRenderResp, error)
	// JSONBindingJSONRender
	// @GORS @POST @Path("/JSONBindingJSONRender/:id") @UriBinding @QueryBinding @HeaderBinding @JSONBinding @JSONRender
	JSONBindingJSONRender(context.Context, *JSONBindingReq) (*JSONRenderResp, error)
	// XMLBindingXMLRender
	// @GORS @PATCH @Path("/XMLBindingXMLRender/:id") @UriBinding @QueryBinding @HeaderBinding  @XMLBinding @XMLRender
	XMLBindingXMLRender(context.Context, *XMLBindingReq) (*XMLRenderResp, error)
	// FormBindingJSONRender
	// @GORS @POST @Path("/FormBindingJSONRender/:id") @UriBinding @QueryBinding @HeaderBinding @FormBinding @JSONRender
	FormBindingJSONRender(context.Context, *FormBindingReq) (*JSONRenderResp, error)
	// FormPostBindingPureJSONRender
	// @GORS @POST @Path("/FormPostBindingPureJSONRender/:id") @UriBinding @QueryBinding @HeaderBinding @FormPostBinding @PureJSONRender
	FormPostBindingPureJSONRender(context.Context, *FormPostBindingReq) (*PureJSONRenderResp, error)
	// FormMultipartBindingAsciiJSONRender
	// @GORS @POST @Path("/FormMultipartBindingAsciiJSONRender/:id") @UriBinding @QueryBinding @HeaderBinding @FormMultipartBinding @AsciiJSONRender
	FormMultipartBindingAsciiJSONRender(context.Context, *FormMultipartBindingReq) (*AsciiJSONRenderResp, error)
	// ProtoBufBindingProtoBufRender
	// @GORS @PUT @Path("/ProtoBufBindingProtoBufRender") @ProtoBufBinding @ProtoBufRender
	ProtoBufBindingProtoBufRender(context.Context, *pb.ProtoBufReq) (*pb.ProtoBufResp, error)
	// MsgPackBindingMsgPackRender
	// @GORS @DELETE @Path("/MsgPackBindingMsgPackRender") @MsgPackBinding @MsgPackRender
	MsgPackBindingMsgPackRender(context.Context, *MsgPackBindingReq) (*MsgPackRenderResp, error)
	// YAMLBindingYAMLRender
	// @GORS @Connect @Path("/YAMLBindingYAMLRender") @YAMLBinding @YAMLRender
	YAMLBindingYAMLRender(context.Context, *YAMLBindingReq) (*YAMLRenderResp, error)
	// TOMLBindingTOMLRender
	// @GORS @Options @Path("/TOMLBindingTOMLRender") @TOMLBinding @TOMLRender
	TOMLBindingTOMLRender(context.Context, *TOMLBindingReq) (*TOMLRenderResp, error)
}

type UriBindingReq struct {
	ID int64 `uri:"id"`
}

type IndentedJSONRenderResp struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

type QueryBindingReq struct {
	ID   int64  `uri:"id"`
	Name string `form:"name"`
}

type SecureJSONRenderResp struct {
	V int64
}

type HeaderBindingReq struct {
	ID   int64  `uri:"id"`
	Name string `form:"name"`
	Auth string `header:"Authorization"`
}

type JsonpJSONRenderResp struct {
	V int64
}

type JSONBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `json:"address"`
}

type JSONRenderResp struct {
	V int64
}

type XMLBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `xml:"address"`
}

type XMLRenderResp struct {
	V int64
}

type FormBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `form:"address"`
}

type FormBindingResp struct {
	V int64
}

type FormPostBindingReq struct {
	ID      int64                 `uri:"id"`
	Name    string                `form:"name"`
	Auth    string                `header:"Authorization"`
	Address string                `form:"address"`
	Avatar  *multipart.FileHeader `form:"avatar" binding:"required"`
}

type PureJSONRenderResp struct {
	V int64
}

type FormMultipartBindingReq struct {
	ID      int64                 `uri:"id"`
	Name    string                `form:"name"`
	Auth    string                `header:"Authorization"`
	Address string                `form:"address"`
	Avatar  *multipart.FileHeader `form:"avatar" binding:"required"`
}

type AsciiJSONRenderResp struct {
	V int64
}

type ProtoBufBindingReq struct {
	ID int64 `form:"id"`
}

type ProtoBufBindingResp struct {
	V int64
}

type MsgPackBindingReq struct {
	ID int64 `form:"id"`
}

type MsgPackRenderResp struct {
	V int64
}

type YAMLBindingReq struct {
	ID int64 `yaml:"id"`
}

type YAMLRenderResp struct {
	V int64
}

type TOMLBindingReq struct {
	ID int64 `toml:"id"`
}

type TOMLRenderResp struct {
	V int64
}
