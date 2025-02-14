package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
)

var (
	ContextPackage = protogen.GoImportPath("context")
	HttpPackage    = protogen.GoImportPath("net/http")
	FmtPackage     = protogen.GoImportPath("fmt")
	IOPackage      = protogen.GoImportPath("io")

	MuxPackage = protogen.GoImportPath("github.com/gorilla/mux")

	ErrorxPackage = protogen.GoImportPath("github.com/go-leo/gox/errorx")
	ProtoxPackage = protogen.GoImportPath("github.com/go-leo/gox/protox")
	UrlxPackage   = protogen.GoImportPath("github.com/go-leo/gox/netx/urlx")

	GorsPackage = protogen.GoImportPath("github.com/go-leo/gors/v2")

	ProtoPackage      = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoJsonPackage  = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	WrapperspbPackage = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	StructpbPackage   = protogen.GoImportPath("google.golang.org/protobuf/types/known/structpb")
	AnypbPackage      = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
)

var (
	ContextIdent = ContextPackage.Ident("Context")

	ErrorEncoderIdent        = GorsPackage.Ident("ErrorEncoder")
	DefaultErrorEncoderIdent = GorsPackage.Ident("DefaultErrorEncoder")
	ResponseEncoderIdent     = GorsPackage.Ident("ResponseEncoder")
	HttpBodyEncoderIdent     = GorsPackage.Ident("HttpBodyEncoder")
	HttpResponseEncoderIdent = GorsPackage.Ident("HttpResponseEncoder")

	RequestDecoderIdent     = GorsPackage.Ident("RequestDecoder")
	HttpBodyDecoderIdent    = GorsPackage.Ident("HttpBodyDecoder")
	HttpRequestDecoderIdent = GorsPackage.Ident("HttpRequestDecoder")
)

var (
	namedPathPattern = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)

var (
	ContentTypeKey  = "Content-Type"
	JsonContentType = "application/json; charset=utf-8"
)
