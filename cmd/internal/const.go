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
	JsonxPackage  = protogen.GoImportPath("github.com/go-leo/gox/encodingx/jsonx")

	ProtoPackage      = protogen.GoImportPath("google.golang.org/protobuf/proto")
	ProtoJsonPackage  = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	WrapperspbPackage = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	StructpbPackage   = protogen.GoImportPath("google.golang.org/protobuf/types/known/structpb")
	AnypbPackage      = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
)

var (
	namedPathPattern = regexp.MustCompile("{([^{}]+)=([^{}]+)}")
	pathPattern      = regexp.MustCompile("{([^=}]+)}")
)

var (
	ContentTypeKey  = "Content-Type"
	JsonContentType = "application/json; charset=utf-8"
)
