package add

import (
	"context"
	tmpio "io"
)

//go:generate gors -service Demo

// Demo
// @GORS @Path("/api")  @Path("/add")
type Demo interface {
	// GetBytesBytes
	// @GORS @GET @Path("/GetBytesBytes") @BytesRender("text/go")
	GetBytesBytes(context.Context, []byte) ([]byte, error)
	// PostBytesBytes
	// @GORS @POST @Path("/PostBytesBytes") @StringRender("text/go")
	PostBytesBytes(context.Context, []byte) ([]byte, error)

	// GetStringString
	// @GORS @GET @Path("/GetStringString") @BytesRender("text/go")
	GetStringString(context.Context, string) (string, error)
	// PostStringString
	// @GORS @POST @Path("/PostStringString") @StringRender("text/go")
	PostStringString(context.Context, string) (string, error)

	// GetReaderReader
	// @GORS @GET @Path("/GetReaderReader") @ReaderRender
	GetReaderReader(context.Context, tmpio.Reader) (tmpio.Reader, error)
	// PostReaderReader
	// @GORS @POST @Path("/PostReaderReader") @ReaderRender
	PostReaderReader(context.Context, tmpio.Reader) (tmpio.Reader, error)

	// GetBytesString
	// @GORS @GET @Path("/GetBytesString") @HTMLRender
	GetBytesString(context.Context, []byte) (string, error)
	// PostBytesString
	// @GORS @POST @Path("/PostBytesString") @RedirectRender
	PostBytesString(context.Context, []byte) (string, error)

	// GetBytesReader
	// @GORS @GET @Path("/GetBytesReader") @ReaderRender("video/mpeg4")
	GetBytesReader(context.Context, []byte) (tmpio.Reader, error)
	// PostBytesReader
	// @GORS @POST @Path("/PostBytesReader") @ReaderRender("video/mpeg4")
	PostBytesReader(context.Context, []byte) (tmpio.Reader, error)

	// GetStringBytes
	// @GORS @GET @Path("/GetStringBytes") @HTMLRender
	GetStringBytes(context.Context, string) ([]byte, error)
	// PostStringBytes
	// @GORS @POST @Path("/PostStringBytes") @RedirectRender
	PostStringBytes(context.Context, string) ([]byte, error)

	// GetStringRender
	// @GORS @GET @Path("/GetStringRender") @ReaderRender("video/mpeg4")
	GetStringRender(context.Context, string) (tmpio.Reader, error)
	// PostStringReader
	// @GORS @POST @Path("/PostStringReader") @ReaderRender("video/mpeg4")
	PostStringReader(context.Context, string) (tmpio.Reader, error)

	// GetReaderBytes
	// @GORS @GET @Path("/GetReaderBytes") @TextRender
	GetReaderBytes(context.Context, tmpio.Reader) ([]byte, error)
	// PostReaderBytes
	// @GORS @POST @Path("/PostReaderBytes") @BytesRender("image/png")
	PostReaderBytes(context.Context, tmpio.Reader) ([]byte, error)

	// GetReaderString
	// @GORS @GET @Path("/GetReaderString") @TextRender
	GetReaderString(context.Context, tmpio.Reader) (string, error)
	// PostReaderString
	// @GORS @POST @Path("/PostReaderString") @StringRender("text/go")
	PostReaderString(context.Context, tmpio.Reader) (string, error)

	// UriBinding
	// @GORS @GET @Path("/UriBinding/:id") @UriBinding @IndentedJSONRender
	UriBinding(context.Context, *UriBindingRequest) (*UriBindingReply, error)
	// QueryBinding
	// @GORS @GET @Path("/QueryBinding") @QueryBinding @SecureJSONRender
	QueryBinding(context.Context, *QueryBindingRequest) (*QueryBindingReply, error)
	// HeaderBinding
	// @GORS @GET @Path("/HeaderBinding") @HeaderBinding @JsonpJSONRender
	HeaderBinding(context.Context, *HeaderBindingRequest) (*HeaderBindingReply, error)
	// JSONBinding
	// @GORS @GET @Path("/JSONBinding") @JSONBinding @JSONRender
	JSONBinding(context.Context, *JSONBindingRequest) (*JSONBindingReply, error)
	// XMLBinding
	// @GORS @GET @Path("/XMLBinding") @XMLBinding @XMLRender
	XMLBinding(context.Context, *XMLBindingRequest) (*XMLBindingReply, error)
	// FormBinding
	// @GORS @GET @Path("/FormBinding") @FormBinding @JsonpJSONRender
	FormBinding(context.Context, *FormBindingRequest) (*FormBindingReply, error)
	// FormPostBinding
	// @GORS @GET @Path("/FormPostBinding") @FormPostBinding @PureJSONRender
	FormPostBinding(context.Context, *FormPostBindingRequest) (*FormPostBindingReply, error)
	// FormMultipartBinding
	// @GORS @GET @Path("/FormMultipartBinding") @FormMultipartBinding @AsciiJSONRender
	FormMultipartBinding(context.Context, *FormMultipartBindingRequest) (*FormMultipartBindingReply, error)
	// ProtoBufBinding
	// @GORS @GET @Path("/ProtoBufBinding") @ProtoBufBinding @ProtoBufRender
	ProtoBufBinding(context.Context, *ProtoBufBindingRequest) (*ProtoBufBindingReply, error)
	// MsgPackBinding
	// @GORS @GET @Path("/MsgPackBinding") @MsgPackBinding @MsgPackRender
	MsgPackBinding(context.Context, *MsgPackBindingRequest) (*MsgPackBindingReply, error)
	// YAMLBinding
	// @GORS @GET @Path("/YAMLBinding") @YAMLBinding @YAMLRender
	YAMLBinding(context.Context, *YAMLBindingRequest) (*YAMLBindingReply, error)
	// TOMLBinding
	// @GORS @GET @Path("/TOMLBinding") @TOMLBinding @TOMLRender
	TOMLBinding(context.Context, *TOMLBindingRequest) (*TOMLBindingReply, error)
}

type UriBindingRequest struct {
	ID int64 `uri:"id"`
}

type UriBindingReply struct {
	V int64
}

type QueryBindingRequest struct {
	ID int64 `form:"id"`
}

type QueryBindingReply struct {
	V int64
}

type HeaderBindingRequest struct {
	ID int64 `header:"id"`
}

type HeaderBindingReply struct {
	V int64
}

type JSONBindingRequest struct {
	ID int64 `json:"id"`
}

type JSONBindingReply struct {
	V int64
}

type XMLBindingRequest struct {
	ID int64 `json:"id"`
}

type XMLBindingReply struct {
	V int64
}

type FormBindingRequest struct {
	ID int64 `form:"id"`
}

type FormBindingReply struct {
	V int64
}

type FormPostBindingRequest struct {
	ID int64 `form:"id"`
}

type FormPostBindingReply struct {
	V int64
}

type FormMultipartBindingRequest struct {
	ID int64 `form:"id"`
}

type FormMultipartBindingReply struct {
	V int64
}

type ProtoBufBindingRequest struct {
	ID int64 `form:"id"`
}

type ProtoBufBindingReply struct {
	V int64
}

type MsgPackBindingRequest struct {
	ID int64 `form:"id"`
}

type MsgPackBindingReply struct {
	V int64
}

type YAMLBindingRequest struct {
	ID int64 `yaml:"id"`
}

type YAMLBindingReply struct {
	V int64
}

type TOMLBindingRequest struct {
	ID int64 `toml:"id"`
}

type TOMLBindingReply struct {
	V int64
}
