package annotation

import "strings"

// The list of annotation.

type Method string

const (
	GORS = "@GORS"

	Path = "@Path"

	GET     Method = "@GET"
	POST    Method = "@POST"
	PUT     Method = "@PUT"
	DELETE  Method = "@DELETE"
	PATCH   Method = "@PATCH"
	HEAD    Method = "@HEAD"
	CONNECT Method = "@CONNECT"
	OPTIONS Method = "@OPTIONS"
	TRACE   Method = "@TRACE"

	UriBinding           = "@UriBinding"
	QueryBinding         = "@QueryBinding"
	HeaderBinding        = "@HeaderBinding"
	JSONBinding          = "@JSONBinding"
	XMLBinding           = "@XMLBinding"
	FormBinding          = "@FormBinding"
	FormPostBinding      = "@FormPostBinding"
	FormMultipartBinding = "@FormMultipartBinding"
	ProtoBufBinding      = "@ProtoBufBinding"
	MsgPackBinding       = "@MsgPackBinding"
	YAMLBinding          = "@YAMLBinding"
	TOMLBinding          = "@TOMLBinding"
	CustomBinding        = "@CustomBinding"
	ProtoJSONBinding     = "@ProtoJSONBinding"

	BytesRender    = "@BytesRender"
	StringRender   = "@StringRender"
	TextRender     = "@TextRender"
	HTMLRender     = "@HTMLRender"
	RedirectRender = "@RedirectRender"
	ReaderRender   = "@ReaderRender"

	JSONRender         = "@JSONRender"
	IndentedJSONRender = "@IndentedJSONRender"
	SecureJSONRender   = "@SecureJSONRender"
	JSONPJSONRender    = "@JSONPJSONRender"
	PureJSONRender     = "@PureJSONRender"
	AsciiJSONRender    = "@AsciiJSONRender"
	XMLRender          = "@XMLRender"
	YAMLRender         = "@YAMLRender"
	ProtoBufRender     = "@ProtoBufRender"
	MsgPackRender      = "@MsgPackRender"
	TOMLRender         = "@TOMLRender"
	CustomRender       = "@CustomRender"
	ProtoJSONRender    = "@ProtoJSONRender"
)

func (m Method) String() string {
	return string(m)
}

func (m Method) EqualsIgnoreCase(str string) bool {
	return strings.ToUpper(str) == m.String()
}

func (m Method) HttpMethod() string {
	switch m {
	case GET:
		return "MethodGet"
	case POST:
		return "MethodPost"
	case PUT:
		return "MethodPut"
	case DELETE:
		return "MethodDelete"
	case PATCH:
		return "MethodPatch"
	case HEAD:
		return "MethodHead"
	case CONNECT:
		return "MethodConnect"
	case OPTIONS:
		return "MethodOptions"
	case TRACE:
		return "MethodTrace"
	default:
		return ""
	}
}
