package gen

import (
	"google.golang.org/protobuf/compiler/protogen"
	"regexp"
)

var (
	UrlxPackage = protogen.GoImportPath("github.com/go-leo/gox/netx/urlx")

	ProtoJsonPackage               = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	ProtoJsonMarshalOptionsIdent   = ProtoJsonPackage.Ident("MarshalOptions")
	ProtoJsonUnmarshalOptionsIdent = ProtoJsonPackage.Ident("UnmarshalOptions")
)

var (
	ContextPackage = protogen.GoImportPath("context")
	ContextIdent   = ContextPackage.Ident("Context")

	HttpPackage          = protogen.GoImportPath("net/http")
	HttpHandlerIdent     = HttpPackage.Ident("Handler")
	HttpHandlerFuncIdent = HttpPackage.Ident("HandlerFunc")
	ResponseWriterIdent  = HttpPackage.Ident("ResponseWriter")
	RequestIdent         = HttpPackage.Ident("Request")

	FmtPackage   = protogen.GoImportPath("fmt")
	SprintfIdent = FmtPackage.Ident("Sprintf")

	MuxPackage  = protogen.GoImportPath("github.com/gorilla/mux")
	RouterIdent = MuxPackage.Ident("Router")
	VarsIdent   = MuxPackage.Ident("Vars")

	ProtoxPackage        = protogen.GoImportPath("github.com/go-leo/gox/protox")
	WrapStringSliceIdent = ProtoxPackage.Ident("WrapStringSlice")

	ProtoPackage     = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoStringIdent = ProtoPackage.Ident("String")

	GorsPackage              = protogen.GoImportPath("github.com/go-leo/gors/v2")
	ErrorEncoderIdent        = GorsPackage.Ident("ErrorEncoder")
	ResponseTransformerIdent = GorsPackage.Ident("ResponseTransformer")
	DefaultErrorEncoderIdent = GorsPackage.Ident("DefaultErrorEncoder")
	ResponseEncoderIdent     = GorsPackage.Ident("ResponseEncoder")
	HttpBodyEncoderIdent     = GorsPackage.Ident("HttpBodyEncoder")
	HttpResponseEncoderIdent = GorsPackage.Ident("HttpResponseEncoder")
	RequestDecoderIdent      = GorsPackage.Ident("RequestDecoder")
	HttpBodyDecoderIdent     = GorsPackage.Ident("HttpBodyDecoder")
	HttpRequestDecoderIdent  = GorsPackage.Ident("HttpRequestDecoder")
	FormDecoderIdent         = GorsPackage.Ident("FormDecoder")
	OptionIdent              = GorsPackage.Ident("Option")
	NewOptionsIdent          = GorsPackage.Ident("NewOptions")

	WrapperspbPackage     = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	WrapperspbStringIdent = WrapperspbPackage.Ident("String")
)

var (
	namedPathPattern = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)
