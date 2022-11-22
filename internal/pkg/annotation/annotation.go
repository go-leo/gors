package annotation

// The list of annotation.
const (
	GORS = "@GORS"

	Path = "@Path"

	GET     = "@GET"
	HEAD    = "@HEAD"
	POST    = "@POST"
	PUT     = "@PUT"
	PATCH   = "@PATCH"
	DELETE  = "@DELETE"
	CONNECT = "@CONNECT"
	OPTIONS = "@OPTIONS"
	TRACE   = "@TRACE"

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

	BytesRender    = "@BytesRender"
	StringRender   = "@StringRender"
	TextRender     = "@TextRender"
	HTMLRender     = "@HTMLRender"
	RedirectRender = "@RedirectRender"
	ReaderRender   = "@ReaderRender"

	JSONRender         = "@JSONRender"
	IndentedJSONRender = "@IndentedJSONRender"
	SecureJSONRender   = "@SecureJSONRender"
	JsonpJSONRender    = "@JsonpJSONRender"
	PureJSONRender     = "@PureJSONRender"
	AsciiJSONRender    = "@AsciiJSONRender"
	XMLRender          = "@XMLRender"
	YAMLRender         = "@YAMLRender"
	ProtoBufRender     = "@ProtoBufRender"
	MsgPackRender      = "@MsgPackRender"
	TOMLRender         = "@TOMLRender"
)
