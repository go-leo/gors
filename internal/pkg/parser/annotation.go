package parser

import (
	"regexp"
	"strings"
)

// The list of annotation.

type Method string

type Binding string

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

	ReaderBinding        Binding = "@ReaderBinding"
	BytesBinding         Binding = "@BytesBinding"
	StringBinding        Binding = "@StringBinding"
	UriBinding           Binding = "@UriBinding"
	QueryBinding         Binding = "@QueryBinding"
	HeaderBinding        Binding = "@HeaderBinding"
	JSONBinding          Binding = "@JSONBinding"
	XMLBinding           Binding = "@XMLBinding"
	FormBinding          Binding = "@FormBinding"
	FormPostBinding      Binding = "@FormPostBinding"
	FormMultipartBinding Binding = "@FormMultipartBinding"
	ProtoBufBinding      Binding = "@ProtoBufBinding"
	MsgPackBinding       Binding = "@MsgPackBinding"
	YAMLBinding          Binding = "@YAMLBinding"
	TOMLBinding          Binding = "@TOMLBinding"
	ProtoJSONBinding     Binding = "@ProtoJSONBinding"
	CustomBinding        Binding = "@CustomBinding"

	ReaderRender   = "@ReaderRender"
	BytesRender    = "@BytesRender"
	StringRender   = "@StringRender"
	TextRender     = "@TextRender"
	HTMLRender     = "@HTMLRender"
	RedirectRender = "@RedirectRender"

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

func ExtractValue(s string, annotation string) (string, bool) {
	reg := regexp.MustCompile(annotation + `\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}
