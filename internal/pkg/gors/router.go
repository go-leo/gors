package gors

import (
	"github.com/go-leo/gors/internal/pkg/annotation"
	"github.com/go-leo/gors/internal/pkg/httpmethod"
	"github.com/go-leo/gox/stringx"
	"go/token"
	"log"
	"path"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ExtractBasePath(comments []string) string {
	var basePath string
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		if seg[0] != annotation.GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := FindPath(s)
				if !ok {
					log.Fatalf("error: %s path invalid", s)
				}
				basePath = path.Join(basePath, v)
			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	return basePath
}

type GoImportPath string

func (p GoImportPath) Ident(s string) *GoIdent {
	importPath := string(p)
	return &GoIdent{
		GoName: s,
		GoImport: &GoImport{
			PackageName: CleanPackageName(path.Base(importPath)),
			ImportPath:  importPath,
		},
	}
}

type GoIdent struct {
	GoImport *GoImport
	GoName   string
}

func (x *GoIdent) Qualify() string {
	if x.GoImport.ImportPath == "" {
		return x.GoName
	}
	return x.GoImport.PackageName + "." + x.GoName
}

type GoImport struct {
	PackageName string
	ImportPath  string
	Enable      bool
}

type ObjectArgs struct {
	Name         string
	GoImportPath GoImportPath
}

type Param struct {
	Bytes      bool
	String     bool
	ObjectArgs *ObjectArgs
	Reader     bool
}

type Result struct {
	Bytes      bool
	String     bool
	ObjectArgs *ObjectArgs
	Reader     bool
}

type RouterInfo struct {
	Method string
	Path   string

	Bindings []string

	RenderContentType string

	Render string

	BytesRender    bool
	StringRender   bool
	TextRender     bool
	HTMLRender     bool
	ReaderRender   bool
	RedirectRender bool

	JSONRender         bool
	ProtoJSONRender    bool
	IndentedJSONRender bool
	SecureJSONRender   bool
	JSONPJSONRender    bool
	PureJSONRender     bool
	AsciiJSONRender    bool
	XMLRender          bool
	YAMLRender         bool
	ProtoBufRender     bool
	MsgPackRender      bool
	TOMLRender         bool
	CustomRender       bool

	RpcMethodName string
	Param2        *Param
	Result1       *Result
}

func NewRouter(methodName string, basePath string, comments []string) *RouterInfo {
	r := &RouterInfo{}
	for _, comment := range comments {
		text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(comment), "//"))
		seg := strings.Split(text, " ")
		// 注释的开始必须以 @GORS 开头
		if seg[0] != annotation.GORS {
			continue
		}
		for _, s := range seg {
			s = strings.TrimSpace(s)
			switch {
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.Path)):
				v, ok := FindPath(s)
				if !ok {
					log.Fatalf("error: rpcmethod %s, %s path invalid", methodName, s)
				}
				r.Path = path.Join(r.Path, v)

				// method start
			case strings.ToUpper(s) == annotation.GET:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.GetMethod
			case strings.ToUpper(s) == annotation.POST:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PostMethod
			case strings.ToUpper(s) == annotation.PUT:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PutMethod
			case strings.ToUpper(s) == annotation.DELETE:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.DeleteMethod
			case strings.ToUpper(s) == annotation.PATCH:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.PatchMethod
			case strings.ToUpper(s) == annotation.HEAD:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.HeadMethod
			case strings.ToUpper(s) == annotation.CONNECT:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.ConnectMethod
			case strings.ToUpper(s) == annotation.OPTIONS:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.OptionsMethod
			case strings.ToUpper(s) == annotation.TRACE:
				if stringx.IsNotBlank(r.Method) {
					log.Fatalf("error: rpcmethod %s, there are multiple methods", methodName)
				}
				r.Method = httpmethod.TraceMethod
				// method end

				// binding start
			case strings.ToUpper(s) == strings.ToUpper(annotation.UriBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.UriBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.QueryBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.QueryBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.HeaderBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.HeaderBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.JSONBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoJSONBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.ProtoJSONBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.XMLBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.FormBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormPostBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.FormPostBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.FormMultipartBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.FormMultipartBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.ProtoBufBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.MsgPackBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.YAMLBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.TOMLBinding, "@"))
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomBinding):
				r.Bindings = append(r.Bindings, strings.TrimPrefix(annotation.CustomBinding, "@"))
				// binding end

				// render start
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.BytesRender)):
				v, _ := findBytesRender(s)
				r.RenderContentType = v
				r.BytesRender = true
				r.Render = "BytesRender"
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.StringRender)):
				v, _ := findStringRender(s)
				r.RenderContentType = v
				r.StringRender = true
				r.Render = "StringRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.TextRender):
				r.TextRender = true
				r.Render = "TextRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.HTMLRender):
				r.HTMLRender = true
				r.Render = "HTMLRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.RedirectRender):
				r.RedirectRender = true
				r.Render = "RedirectRender"
			case strings.HasPrefix(strings.ToUpper(s), strings.ToUpper(annotation.ReaderRender)):
				v, _ := findReaderRender(s)
				r.RenderContentType = v
				r.ReaderRender = true
				r.Render = "ReaderRender"

			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONRender):
				r.JSONRender = true
				r.Render = "JSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.IndentedJSONRender):
				r.IndentedJSONRender = true
				r.Render = "IndentedJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.SecureJSONRender):
				r.SecureJSONRender = true
				r.Render = "SecureJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.JSONPJSONRender):
				r.JSONPJSONRender = true
				r.Render = "JSONPJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.PureJSONRender):
				r.PureJSONRender = true
				r.Render = "PureJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.AsciiJSONRender):
				r.AsciiJSONRender = true
				r.Render = "AsciiJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoJSONRender):
				r.ProtoJSONRender = true
				r.Render = "ProtoJSONRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.XMLRender):
				r.XMLRender = true
				r.Render = "XMLRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.YAMLRender):
				r.YAMLRender = true
				r.Render = "YAMLRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.ProtoBufRender):
				r.ProtoBufRender = true
				r.Render = "ProtoBufRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.MsgPackRender):
				r.MsgPackRender = true
				r.Render = "MsgPackRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.TOMLRender):
				r.TOMLRender = true
				r.Render = "TOMLRender"
			case strings.ToUpper(s) == strings.ToUpper(annotation.CustomRender):
				r.CustomRender = true
				r.Render = "CustomRender"

				// render end

			case strings.HasPrefix(s, annotation.GORS):
			case "" == s:
			default:
				log.Printf("warning: format error: unsupport: %s", s)
			}
		}
	}
	//if stringx.IsBlank(r.Method) {
	//	r.Method = httpmethod.PostMethod
	//	log.Printf("rpcmethod %s, use default http method POST", methodName)
	//}
	r.Path = path.Join(basePath, r.Path)
	return r
}

func FindPath(s string) (string, bool) {
	reg := regexp.MustCompile(`@Path\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findStringRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@StringRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findBytesRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@BytesRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func findReaderRender(s string) (string, bool) {
	reg := regexp.MustCompile(`@ReaderRender\((.*)\)`)
	if !reg.MatchString(s) {
		return "", false
	}
	matchArr := reg.FindStringSubmatch(s)
	return matchArr[len(matchArr)-1], true
}

func CleanPackageName(name string) string {
	name = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, name)

	// Prepend '_' in the event of a Go keyword conflict or if
	// the identifier is invalid (does not start in the Unicode L category).
	r, _ := utf8.DecodeRuneInString(name)
	if token.Lookup(name).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + name
	}
	return name
}
