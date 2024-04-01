package demo

import (
	"context"
	"database/sql"
	"mime/multipart"
	"time"

	"github.com/go-leo/gors/example/api/pb"
)

//go:generate gors -service ObjObj

// ObjObj
// @GORS @Path(/api/ObjObj)
type ObjObj interface {

	// AllRequest is http api
	// @GORS @GET @Path(/all/:id/user/:user) @UriBinding @UriBinding @QueryBinding @HeaderBinding @JSONBinding @JSONRender
	AllRequest(context.Context, *AllRequestReq) (*AllRequestResp, error)

	// UriBindingIndentedJSONRender
	// @GORS @GET @Path(/UriBindingIndentedJSONRender/:id) @UriBinding @IndentedJSONRender
	UriBindingIndentedJSONRender(context.Context, *UriBindingReq) (*IndentedJSONRenderResp, error)

	// QueryBindingSecureJSONRender
	// @GORS @GET @Path(/QueryBindingSecureJSONRender/:id) @UriBinding @QueryBinding @SecureJSONRender
	QueryBindingSecureJSONRender(context.Context, *QueryBindingReq) (*SecureJSONRenderResp, error)

	// HeaderBindingJsonpJSONRender
	// @GORS @GET @Path(/HeaderBindingJsonpJSONRender/:id) @UriBinding @QueryBinding @HeaderBinding @JSONPJSONRender
	HeaderBindingJsonpJSONRender(context.Context, *HeaderBindingReq) (*JsonpJSONRenderResp, error)

	// JSONBindingJSONRender
	// @GORS @POST @Path(/JSONBindingJSONRender/:id) @UriBinding @QueryBinding @HeaderBinding @JSONBinding @JSONRender
	JSONBindingJSONRender(context.Context, *JSONBindingReq) (*JSONRenderResp, error)
	// XMLBindingXMLRender
	// @GORS @PATCH @Path(/XMLBindingXMLRender/:id) @UriBinding @QueryBinding @HeaderBinding  @XMLBinding @XMLRender
	XMLBindingXMLRender(context.Context, *XMLBindingReq) (*XMLRenderResp, error)

	// FormBindingJSONRender
	// @GORS @POST @Path(/FormBindingJSONRender/:id) @UriBinding @HeaderBinding @FormBinding @JSONRender
	FormBindingJSONRender(context.Context, *FormBindingReq) (*JSONRenderResp, error)

	// FormPostBindingPureJSONRender
	// @GORS @POST @Path(/FormPostBindingPureJSONRender/:id) @UriBinding @QueryBinding @HeaderBinding @FormPostBinding @PureJSONRender
	FormPostBindingPureJSONRender(context.Context, *FormPostBindingReq) (*PureJSONRenderResp, error)

	// FormMultipartBindingAsciiJSONRender
	// @GORS @POST @Path(/FormMultipartBindingAsciiJSONRender/:id) @UriBinding @QueryBinding @HeaderBinding @FormMultipartBinding @AsciiJSONRender
	FormMultipartBindingAsciiJSONRender(context.Context, *FormMultipartBindingReq) (*AsciiJSONRenderResp, error)

	// ProtoBufBindingProtoBufRender
	// @GORS @PUT @Path(/ProtoBufBindingProtoBufRender) @ProtoBufBinding @ProtoBufRender
	ProtoBufBindingProtoBufRender(context.Context, *pb.ProtoBufReq) (*pb.ProtoBufResp, error)

	// MsgPackBindingMsgPackRender
	// @GORS @DELETE @Path(/MsgPackBindingMsgPackRender) @MsgPackBinding @MsgPackRender
	MsgPackBindingMsgPackRender(context.Context, *MsgPackBindingReq) (*MsgPackRenderResp, error)

	// YAMLBindingYAMLRender
	// @GORS @DELETE @Path(/YAMLBindingYAMLRender/:id) @UriBinding @QueryBinding @HeaderBinding @YAMLBinding @YAMLRender
	YAMLBindingYAMLRender(context.Context, *YAMLBindingReq) (*YAMLRenderResp, error)

	// TOMLBindingTOMLRender
	// @GORS @PUT @Path(/TOMLBindingTOMLRender/:id) @UriBinding @QueryBinding @HeaderBinding @TOMLBinding @TOMLRender
	TOMLBindingTOMLRender(context.Context, *TOMLBindingReq) (*TOMLRenderResp, error)

	// ProtoJSONBindingProtoJSONRender
	// @GORS @PUT @Path(/ProtoJSONBindingProtoJSONRender) @ProtoJSONBinding @ProtoJSONRender
	ProtoJSONBindingProtoJSONRender(context.Context, *pb.ProtoBufReq) (*pb.ProtoBufResp, error)
}

type Page struct {
	Num  int `json:"num"`
	Size int `json:"size"`
}

type TimeRage struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type AllRequestReq struct {
	// I comment 1
	// comment 2
	I int `uri:"i"` // comment 3
	// comment 2
	I8  int8  `form:"i8"` // comment 3
	I16 int16 `header:"i16"`
	I32 int32
	I64 int64

	U   uint
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64

	F32 float32
	F64 float64

	B  bool
	S  string
	Bs []byte `json:"bs"`
	A  [10]int

	T time.Time

	UriBindingReq    UriBindingReq
	UriBindingReqPtr *UriBindingReq
	ProtoBufReq      pb.ProtoBufReq
	ProtoBufReqPtr   *pb.ProtoBufReq

	UriBindingReqs    []UriBindingReq
	UriBindingReqPtrs []*UriBindingReq
	ProtoBufReqs      []pb.ProtoBufReq
	ProtoBufReqPtrs   []*pb.ProtoBufReq
	Page
	*TimeRage
	sql.NullByte
	*sql.NullInt32
	int
	bool
	*int64
	*string

	Avatar *multipart.FileHeader `form:"avatar"`
}

type AllRequestResp struct {
}

type UriBindingReq struct {
	ID int64 `uri:"id"`
}

type IndentedJSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type QueryBindingReq struct {
	ID   int64  `uri:"id"`
	Name string `form:"name"`
}

type SecureJSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type HeaderBindingReq struct {
	ID   int64  `uri:"id"`
	Name string `form:"name"`
	Auth string `header:"Authorization"`
}

type JsonpJSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Auth       string    `json:"auth"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type JSONBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `json:"address"`
}

type JSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Auth       string    `json:"auth"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type XMLBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `xml:"address"`
}

type XMLRenderResp struct {
	ID         int64     `xml:"id"`
	Name       string    `xml:"name"`
	Auth       string    `xml:"auth"`
	Address    string    `xml:"address"`
	Birthday   time.Time `xml:"birthday"`
	CreateTime time.Time `xml:"createTime"`
	UnixTime   time.Time `xml:"unixTime"`
}

type FormBindingReq struct {
	ID      int64                 `uri:"id"`
	Name    string                `form:"name"`
	Auth    string                `header:"Authorization"`
	Address string                `form:"address"`
	Avatar  *multipart.FileHeader `form:"avatar"`
}

type FormBindingResp struct {
	V int64
}

type FormPostBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `form:"address"`
}

type PureJSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Auth       string    `json:"auth"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type FormMultipartBindingReq struct {
	ID      int64                 `uri:"id"`
	Name    string                `form:"name"`
	Auth    string                `header:"Authorization"`
	Address string                `form:"address"`
	Avatar  *multipart.FileHeader `form:"avatar"`
}

type AsciiJSONRenderResp struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Auth       string    `json:"auth"`
	Address    string    `json:"address"`
	Birthday   time.Time `json:"birthday"`
	CreateTime time.Time `json:"createTime"`
	UnixTime   time.Time `json:"unixTime"`
}

type MsgPackBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string
}

type MsgPackRenderResp struct {
	ID         int64
	Name       string
	Auth       string
	Address    string
	Birthday   time.Time
	CreateTime time.Time
	UnixTime   time.Time
}

type YAMLBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `yaml:"address"`
}

type YAMLRenderResp struct {
	ID         int64     `yaml:"id"`
	Name       string    `yaml:"name"`
	Auth       string    `yaml:"auth"`
	Address    string    `yaml:"address"`
	Birthday   time.Time `yaml:"birthday"`
	CreateTime time.Time `yaml:"createTime"`
	UnixTime   time.Time `yaml:"unixTime"`
}

type TOMLBindingReq struct {
	ID      int64  `uri:"id"`
	Name    string `form:"name"`
	Auth    string `header:"Authorization"`
	Address string `toml:"address"`
}

type TOMLRenderResp struct {
	ID         int64     `toml:"id"`
	Name       string    `toml:"name"`
	Auth       string    `toml:"auth"`
	Address    string    `toml:"address"`
	Birthday   time.Time `toml:"birthday"`
	CreateTime time.Time `toml:"createTime"`
	UnixTime   time.Time `toml:"unixTime"`
}
